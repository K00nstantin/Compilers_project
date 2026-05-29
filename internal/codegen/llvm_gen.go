package codegen

import (
	"fmt"
	"strconv"

	"github.com/K00nstantin/Compilers_project/internal/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Generator struct {
	module   *ir.Module
	curBlock *ir.Block
	curFunc  *ir.Func
	symStack []map[string]value.Value
}

func NewGenerator(moduleName string) *Generator {
	m := ir.NewModule()
	m.SourceFilename = moduleName + ".mod"
	g := &Generator{
		module:   m,
		symStack: []map[string]value.Value{},
	}
	g.pushScope()
	return g
}

func (g *Generator) pushScope() {
	g.symStack = append(g.symStack, make(map[string]value.Value))
}

func (g *Generator) Generate(astMod *ast.Module) *ir.Module {
	return g.module
}

func (g *Generator) popScope() {
	if len(g.symStack) > 0 {
		g.symStack = g.symStack[:len(g.symStack)-1]
	}
}

func (g *Generator) addSymbol(name string, val value.Value) {
	if len(g.symStack) == 0 {
		g.pushScope()
	}
	g.symStack[len(g.symStack)-1][name] = val
}

func (g *Generator) lookup(name string) (value.Value, bool) {
	for i := len(g.symStack) - 1; i > 0; i-- {
		if val, ok := g.symStack[i][name]; ok {
			return val, true
		}
	}
	return nil, false
}

func (g *Generator) oberonTypeToLLVM(t ast.TypeExpr) types.Type {
	switch tt := t.(type) {
	case *ast.NamedType:
		switch tt.Name {
		case "INTEGER":
			return types.I32
		case "BOOLEAN":
			return types.I1
		default:
			return types.I32
		}
	case *ast.ArrayType:
		if len(tt.Lengths) != 1 {
			panic("multi dismentional arrays are not supported yet")
		}
		length := g.constExprToInt(tt.Lengths[0])
		elemType := g.oberonTypeToLLVM(tt.Elem)
		return types.NewArray(uint64(length), elemType)
	case *ast.RecordType:
		fields := make([]types.Type, len(tt.Fields))
		for i, f := range tt.Fields {
			fields[i] = g.oberonTypeToLLVM(f.Type)
		}
		return types.NewStruct(fields...)
	case *ast.PointerType:
		return types.NewPointer(g.oberonTypeToLLVM(tt.Target))
	default:
		panic(fmt.Sprintf("unknown type %T", t))
	}
}

func (g *Generator) constExprToInt(e ast.Expr) int {
	switch exp := e.(type) {
	case *ast.NumberExpr:
		val, _ := strconv.Atoi(exp.Text)
		return val
	case *ast.BinaryExpr:
		left := g.constExprToInt(exp.Left)
		right := g.constExprToInt(exp.Right)
		switch exp.Op {
		case "+":
			return left + right
		case "-":
			return left - right
		case "*":
			return left * right
		case "/":
			return left / right
		}
	}
	return 0
}
