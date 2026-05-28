package codegen

import (
	"github.com/K00nstantin/Compilers_project/internal/ast"
	"github.com/llir/llvm/ir"
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
