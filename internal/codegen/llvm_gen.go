package codegen

import (
	"fmt"
	"strconv"

	"github.com/K00nstantin/Compilers_project/internal/ast"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Generator struct {
	module           *ir.Module
	moduleName       string
	curBlock         *ir.Block
	curFunc          *ir.Func
	symStack         []map[string]value.Value
	fieldNameToIndex map[*types.StructType]map[string]int
	recordTypes      map[string]*ast.RecordType
	blockCounter     int
}

func NewGenerator(moduleName string) *Generator {
	m := ir.NewModule()
	m.SourceFilename = moduleName + ".mod"
	g := &Generator{
		module:           m,
		moduleName:       moduleName,
		curBlock:         nil,
		curFunc:          nil,
		symStack:         []map[string]value.Value{},
		fieldNameToIndex: make(map[*types.StructType]map[string]int),
		recordTypes:      make(map[string]*ast.RecordType),
		blockCounter:     0,
	}
	g.pushScope()
	return g
}

func (g *Generator) pushScope() {
	g.symStack = append(g.symStack, make(map[string]value.Value))
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
	for i := len(g.symStack) - 1; i >= 0; i-- {
		if val, ok := g.symStack[i][name]; ok {
			return val, true
		}
	}
	return nil, false
}

func (g *Generator) nextBlockName(prefix string) string {
	g.blockCounter++
	return fmt.Sprintf("%s_%d", prefix, g.blockCounter)
}

func (g *Generator) oberonTypeToLLVM(t ast.TypeExpr) types.Type {
	switch tt := t.(type) {
	case *ast.NamedType:
		switch tt.Name {
		case "INTEGER":
			return types.I32
		case "BOOLEAN":
			return types.I1
		case "SET":
			return types.I32
		case "REAL":
			return types.Float
		default:
			if rec, ok := g.recordTypes[tt.Name]; ok {
				return g.oberonTypeToLLVM(rec)
			}
			return types.I32
		}
	case *ast.ArrayType:
		elemType := g.oberonTypeToLLVM(tt.Elem)
		if len(tt.Lengths) == 0 {
			// Открытый массив: представляем как указатель на элемент
			return types.NewPointer(elemType)
		}
		// Фиксированный массив: вложенные массивы
		for i := len(tt.Lengths) - 1; i >= 0; i-- {
			length := g.constExprToInt(tt.Lengths[i])
			elemType = types.NewArray(uint64(length), elemType)
		}
		return elemType
	case *ast.RecordType:
		var fields []types.Type
		var fieldNames []string

		// Рекурсивный сбор полей (сначала базовые, потом текущие)
		var collectFields func(rec *ast.RecordType)
		collectFields = func(rec *ast.RecordType) {
			if rec.Base != "" {
				baseRec, ok := g.recordTypes[rec.Base]
				if !ok {
					panic("base record not found: " + rec.Base)
				}
				collectFields(baseRec)
			}
			for _, fd := range rec.Fields {
				fieldType := g.oberonTypeToLLVM(fd.Type)
				for _, nameDef := range fd.Names {
					fields = append(fields, fieldType)
					fieldNames = append(fieldNames, nameDef.Name)
				}
			}
		}
		collectFields(tt)

		structType := types.NewStruct(fields...)
		nameMap := make(map[string]int)
		for i, name := range fieldNames {
			nameMap[name] = i
		}
		g.fieldNameToIndex[structType] = nameMap
		return structType
	case *ast.PointerType:
		return types.NewPointer(g.oberonTypeToLLVM(tt.Target))

	default:
		panic(fmt.Sprintf("unknown type: %T", t))
	}
}

func (g *Generator) constExprToInt(e ast.Expr) int {
	switch ex := e.(type) {
	case *ast.NumberExpr:
		val, _ := strconv.Atoi(ex.Text)
		return val
	case *ast.BinaryExpr:
		left := g.constExprToInt(ex.Left)
		right := g.constExprToInt(ex.Right)
		switch ex.Op {
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

// constExprToIntOrNil возвращает *int64, если выражение константно, иначе nil
func (g *Generator) constExprToIntOrNil(e ast.Expr) *int64 {
	switch ex := e.(type) {
	case *ast.NumberExpr:
		val, err := strconv.ParseInt(ex.Text, 10, 64)
		if err != nil {
			return nil
		}
		return &val
	case *ast.BinaryExpr:
		left := g.constExprToIntOrNil(ex.Left)
		right := g.constExprToIntOrNil(ex.Right)
		if left == nil || right == nil {
			return nil
		}
		var res int64
		switch ex.Op {
		case "+":
			res = *left + *right
		case "-":
			res = *left - *right
		case "*":
			res = *left * *right
		case "/":
			res = *left / *right
		default:
			return nil
		}
		return &res
	default:
		return nil
	}
}

// isSetValue определяет, является ли LLVM значение множеством (i32 с пометкой, что это множество).
func (g *Generator) isSetValue(v value.Value) bool {
	return v.Type() == types.I32
}

// genSetExpr генерирует битовую маску для литерала множества.
func (g *Generator) genSetExpr(s *ast.SetExpr) value.Value {
	mask := int64(0)
	for _, elem := range s.Elements {
		fromVal := g.constExprToIntOrNil(elem.From)
		if fromVal == nil {
			panic("non-constant set element")
		}
		from := *fromVal
		to := from
		if elem.To != nil && elem.To != elem.From {
			toVal := g.constExprToIntOrNil(elem.To)
			if toVal == nil {
				panic("non-constant set range end")
			}
			to = *toVal
		}
		if from > to {
			from, to = to, from
		}
		for i := from; i <= to; i++ {
			if i >= 0 && i < 32 {
				mask |= 1 << i
			}
		}
	}
	return constant.NewInt(types.I32, mask)
}

func (g *Generator) Generate(astMod *ast.Module) *ir.Module {
	g.processDeclarations(astMod.Declarations, true)
	g.generateModuleInit(astMod.Body)
	g.generateMain()
	return g.module
}

func (g *Generator) processDeclarations(decls *ast.DeclarationBlock, isGlobal bool) {
	if decls == nil {
		return
	}
	for _, c := range decls.Consts {
		isReal, intVal, realVal := g.constEval(c.Value)
		var llvmConst value.Value
		if isReal {
			llvmConst = constant.NewFloat(types.Float, realVal)
		} else {
			llvmConst = constant.NewInt(types.I32, intVal)
		}
		g.addSymbol(c.Name.Name, llvmConst)
	}
	for _, t := range decls.Types {
		if rec, ok := t.Type.(*ast.RecordType); ok {
			g.recordTypes[t.Name.Name] = rec
		}
	}
	for _, v := range decls.Vars {
		llvmType := g.oberonTypeToLLVM(v.Type)
		for _, nameDef := range v.Names {
			name := nameDef.Name
			var val value.Value
			if isGlobal {
				global := g.module.NewGlobal(name, llvmType)
				global.Init = constant.NewZeroInitializer(llvmType)
				if nameDef.Exported {
					global.Linkage = enum.LinkageExternal
				}
				val = global
			} else {
				val = nil
			}
			if val != nil {
				g.addSymbol(name, val)
			}
		}
	}
	for _, proc := range decls.Procedures {
		g.generateProcedure(proc, isGlobal)
	}
}

func (g *Generator) generateProcedure(proc *ast.ProcedureDecl, isGlobal bool) {
	var retType types.Type = types.Void
	if proc.Signature != nil && proc.Signature.ReturnType != "" {
		retType = g.namedTypeToLLVM(proc.Signature.ReturnType)
	}
	var params []*ir.Param
	if proc.Signature != nil {
		for _, paramSec := range proc.Signature.Params {
			typ := g.oberonTypeToLLVM(paramSec.Type)
			if paramSec.ByRef {
				typ = types.NewPointer(typ)
			}
			for _, paramName := range paramSec.Names {
				params = append(params, ir.NewParam(paramName, typ))
			}
		}
	}
	funcName := proc.Name.Name
	f := g.module.NewFunc(funcName, retType, params...)
	if proc.Name.Exported {
		f.Linkage = enum.LinkageExternal
	}
	g.addSymbol(funcName, f)
	oldFunc := g.curFunc
	oldBlock := g.curBlock
	defer func() {
		g.curFunc = oldFunc
		g.curBlock = oldBlock
	}()
	g.curFunc = f
	entryBlock := f.NewBlock("entry")
	g.curBlock = entryBlock
	g.pushScope()
	defer g.popScope()

	paramIdx := 0
	for _, paramSec := range proc.Signature.Params {
		for _, paramName := range paramSec.Names {
			param := f.Params[paramIdx]
			if paramSec.ByRef {
				// VAR параметр: уже указатель, сохраняем как есть
				g.addSymbol(paramName, param)
			} else {
				// Обычный параметр: копируем на стек
				alloca := g.curBlock.NewAlloca(param.Typ)
				g.curBlock.NewStore(param, alloca)
				g.addSymbol(paramName, alloca)
			}
			paramIdx++
		}
	}

	if proc.Declarations != nil {
		for _, v := range proc.Declarations.Vars {
			llvmType := g.oberonTypeToLLVM(v.Type)
			for _, nameDef := range v.Names {
				alloca := g.curBlock.NewAlloca(llvmType)
				g.curBlock.NewStore(constant.NewZeroInitializer(llvmType), alloca)
				g.addSymbol(nameDef.Name, alloca)
			}
		}
	}
	for _, stmt := range proc.Body {
		g.genStmt(stmt)
	}
	if proc.ReturnExpr != nil {
		retVal := g.genExpr(proc.ReturnExpr)
		g.curBlock.NewRet(retVal)
	} else if retType == types.Void {
		g.curBlock.NewRet(nil)
	} else {
		panic("missing return value")
	}
}

func (g *Generator) generateModuleInit(body []ast.Stmt) {
	if len(body) == 0 {
		return
	}
	initFunc := g.module.NewFunc("__init_"+g.moduleName, types.Void)
	initBlock := initFunc.NewBlock("entry")
	oldFunc := g.curFunc
	oldBlock := g.curBlock
	defer func() {
		g.curFunc = oldFunc
		g.curBlock = oldBlock
	}()
	g.curFunc = initFunc
	g.curBlock = initBlock
	for _, stmt := range body {
		g.genStmt(stmt)
	}
	g.curBlock.NewRet(nil)
}

func (g *Generator) generateMain() {
	mainFunc := g.module.NewFunc("main", types.I32)
	block := mainFunc.NewBlock("entry")
	oldBlock := g.curBlock
	defer func() { g.curBlock = oldBlock }()
	g.curBlock = block

	initName := "__init_" + g.moduleName
	var initFunc *ir.Func
	for _, f := range g.module.Funcs {
		if f.Name() == initName {
			initFunc = f
			break
		}
	}
	if initFunc != nil {
		g.curBlock.NewCall(initFunc)
	}

	if resVal, ok := g.lookup("result"); ok {
		if ptrType, ok := resVal.Type().(*types.PointerType); ok {
			val := g.curBlock.NewLoad(ptrType.ElemType, resVal)
			g.curBlock.NewRet(val)
			return
		}
	}
	g.curBlock.NewRet(constant.NewInt(types.I32, 0))
}

func (g *Generator) genExpr(e ast.Expr) value.Value {
	switch ex := e.(type) {
	case *ast.NumberExpr:
		if ex.IsReal {
			val, err := strconv.ParseFloat(ex.Text, 64)
			if err != nil {
				panic("invalid real number: " + ex.Text)
			}
			return constant.NewFloat(types.Float, val)
		} else {
			val, err := strconv.ParseInt(ex.Text, 10, 32)
			if err != nil {
				panic("invalid integer number: " + ex.Text)
			}
			return constant.NewInt(types.I32, val)
		}
	case *ast.BoolExpr:
		if ex.Value {
			return constant.NewInt(types.I1, 1)
		}
		return constant.NewInt(types.I1, 0)
	case *ast.StringExpr:
		strConst := constant.NewCharArrayFromString(ex.Value)
		globalStr := g.module.NewGlobalDef("str", strConst)
		return globalStr
	case *ast.SetExpr:
		return g.genSetExpr(ex)
	case *ast.BinaryExpr:
		left := g.genExpr(ex.Left)
		right := g.genExpr(ex.Right)
		isReal := g.isRealType(left.Type()) || g.isRealType(right.Type())
		if isReal {
			left = g.promoteToReal(left)
			right = g.promoteToReal(right)
		}
		switch ex.Op {
		case "+":
			if isReal {
				return g.curBlock.NewFAdd(left, right)
			}
			return g.curBlock.NewAdd(left, right)
		case "-":
			if isReal {
				return g.curBlock.NewFSub(left, right)
			}
			return g.curBlock.NewSub(left, right)
		case "*":
			if isReal {
				return g.curBlock.NewFMul(left, right)
			}
			return g.curBlock.NewMul(left, right)
		case "/":
			if isReal {
				return g.curBlock.NewFDiv(left, right)
			}
			return g.curBlock.NewSDiv(left, right)
		case "DIV":
			if isReal {
				panic("DIV not allowed on REAL")
			}
			return g.curBlock.NewSDiv(left, right)
		case "MOD":
			if isReal {
				panic("MOD not allowed on REAL")
			}
			return g.curBlock.NewSRem(left, right)
		case "=":
			if isReal {
				return g.curBlock.NewFCmp(enum.FPredOEQ, left, right)
			}
			return g.curBlock.NewICmp(enum.IPredEQ, left, right)
		case "#":
			if isReal {
				return g.curBlock.NewFCmp(enum.FPredONE, left, right)
			}
			return g.curBlock.NewICmp(enum.IPredNE, left, right)
		case "<":
			if isReal {
				return g.curBlock.NewFCmp(enum.FPredOLT, left, right)
			}
			return g.curBlock.NewICmp(enum.IPredSLT, left, right)
		case "<=":
			if isReal {
				return g.curBlock.NewFCmp(enum.FPredOLE, left, right)
			}
			return g.curBlock.NewICmp(enum.IPredSLE, left, right)
		case ">":
			if isReal {
				return g.curBlock.NewFCmp(enum.FPredOGT, left, right)
			}
			return g.curBlock.NewICmp(enum.IPredSGT, left, right)
		case ">=":
			if isReal {
				return g.curBlock.NewFCmp(enum.FPredOGE, left, right)
			}
			return g.curBlock.NewICmp(enum.IPredSGE, left, right)
		case "&":
			if isReal {
				panic("& not allowed on REAL")
			}
			return g.curBlock.NewAnd(left, right)
		case "OR":
			if isReal {
				panic("OR not allowed on REAL")
			}
			return g.curBlock.NewOr(left, right)
		case "IN":
			if isReal {
				panic("IN not allowed on REAL")
			}
			// обработка IN (как раньше)
			elem := left
			set := right
			if elem.Type() == types.I1 {
				elem = g.curBlock.NewZExt(elem, types.I32)
			}
			one := constant.NewInt(types.I32, 1)
			shifted := g.curBlock.NewShl(one, elem)
			and := g.curBlock.NewAnd(set, shifted)
			zero := constant.NewInt(types.I32, 0)
			return g.curBlock.NewICmp(enum.IPredNE, and, zero)
		default:
			panic("unknown binary operator: " + ex.Op)
		}
	case *ast.UnaryExpr:
		sub := g.genExpr(ex.Expr)
		if ex.Op == "-" {
			intType, ok := sub.Type().(*types.IntType)
			if !ok {
				panic("unary minus on non-integer")
			}
			zero := constant.NewInt(intType, 0)
			return g.curBlock.NewSub(zero, sub)
		} else if ex.Op == "~" {
			one := constant.NewInt(types.I1, 1)
			return g.curBlock.NewXor(one, sub)
		}
		return sub
	case *ast.DesignatorExpr:
		name := ex.Base.Name
		if val, ok := g.lookup(name); ok {
			// Если это константа (целая или вещественная), возвращаем её напрямую
			switch val.(type) {
			case *constant.Int, *constant.Float:
				return val
			}
		}
		ptr := g.genDesignatorPtr(ex)
		ptrType, ok := ptr.Type().(*types.PointerType)
		if !ok {
			panic("designator is not a pointer")
		}
		return g.curBlock.NewLoad(ptrType.ElemType, ptr)
	case *ast.CallExpr:
		return g.genCall(ex)
	default:
		panic(fmt.Sprintf("unhandled expression type: %T", e))
	}
}

func (g *Generator) fieldIndex(structType *types.StructType, fieldName string) int {
	if m, ok := g.fieldNameToIndex[structType]; ok {
		if idx, ok := m[fieldName]; ok {
			return idx
		}
	}
	panic(fmt.Sprintf("field %s not found in struct type %v", fieldName, structType))
}

func (g *Generator) genDesignatorPtr(d *ast.DesignatorExpr) value.Value {
	name := d.Base.Name
	if d.Base.Module != "" {
		name = d.Base.Module + "." + d.Base.Name
	}
	baseVal, ok := g.lookup(name)
	if !ok && d.Base.Module != "" {
		moduleName := d.Base.Module
		baseVal, ok = g.lookup(moduleName)
		if ok {
			newSelectors := make([]ast.Selector, 0, len(d.Selectors)+1)
			newSelectors = append(newSelectors, ast.Selector{Field: d.Base.Name})
			newSelectors = append(newSelectors, d.Selectors...)
			newDes := &ast.DesignatorExpr{
				Base:      ast.QualIdent{Name: moduleName},
				Selectors: newSelectors,
			}
			return g.genDesignatorPtr(newDes)
		}
	}
	if !ok {
		panic("unknown identifier: " + name)
	}
	cur := baseVal

	for _, sel := range d.Selectors {
		if sel.Field != "" {
			ptrType, ok := cur.Type().(*types.PointerType)
			if !ok {
				panic("field access on non-pointer")
			}
			structType, ok := ptrType.ElemType.(*types.StructType)
			if !ok {
				panic("field access on non-struct")
			}
			fieldIdx := g.fieldIndex(structType, sel.Field)
			gep := g.curBlock.NewGetElementPtr(structType, cur,
				constant.NewInt(types.I32, 0),
				constant.NewInt(types.I32, int64(fieldIdx)))
			cur = gep
		} else if len(sel.Index) > 0 {
			ptrType, ok := cur.Type().(*types.PointerType)
			if !ok {
				panic("index access on non-pointer")
			}
			elemType := ptrType.ElemType
			indices := []value.Value{}
			if _, isArray := elemType.(*types.ArrayType); isArray {
				indices = append(indices, constant.NewInt(types.I32, 0))
			}
			for _, idxExpr := range sel.Index {
				idxVal := g.genExpr(idxExpr)
				if idxVal.Type() != types.I32 {
					if idxVal.Type() == types.I1 {
						idxVal = g.curBlock.NewZExt(idxVal, types.I32)
					} else if intType, ok := idxVal.Type().(*types.IntType); ok && intType.BitSize < 32 {
						idxVal = g.curBlock.NewZExt(idxVal, types.I32)
					} else {
						idxVal = g.curBlock.NewTrunc(idxVal, types.I32)
					}
				}
				indices = append(indices, idxVal)
				if arrType, ok := elemType.(*types.ArrayType); ok {
					elemType = arrType.ElemType
				} else {
					if len(sel.Index) > 1 {
						panic("too many indices for open array or non-array type")
					}
					break
				}
			}
			gep := g.curBlock.NewGetElementPtr(ptrType.ElemType, cur, indices...)
			cur = gep
		} else if sel.Deref {
			ptrType, ok := cur.Type().(*types.PointerType)
			if !ok {
				panic("deref on non-pointer")
			}
			cur = g.curBlock.NewLoad(ptrType.ElemType, cur)
		}
	}
	return cur
}

func (g *Generator) genCall(c *ast.CallExpr) value.Value {
	// Проверка на встроенные вызовы (по имени designator)
	if des, ok := c.Callee.(*ast.DesignatorExpr); ok {
		name := des.Base.Name
		if des.Base.Module != "" {
			name = des.Base.Module + "." + name
		}
		// Встроенные процедуры (не возвращают значение, возвращаем nil)
		if name == "INC" && len(c.Args) == 1 {
			arg := c.Args[0]
			if des, ok := arg.(*ast.DesignatorExpr); ok {
				ptr := g.genDesignatorPtr(des)
				if ptrType, ok := ptr.Type().(*types.PointerType); ok {
					load := g.curBlock.NewLoad(ptrType.ElemType, ptr)
					if load.Type() != types.I32 {
						panic("INC requires INTEGER variable")
					}
					one := constant.NewInt(types.I32, 1)
					added := g.curBlock.NewAdd(load, one)
					g.curBlock.NewStore(added, ptr)
					return nil
				}
			}
			panic("INC argument must be an INTEGER variable")
		}
		if name == "DEC" && len(c.Args) == 1 {
			arg := c.Args[0]
			if des, ok := arg.(*ast.DesignatorExpr); ok {
				ptr := g.genDesignatorPtr(des)
				if ptrType, ok := ptr.Type().(*types.PointerType); ok {
					load := g.curBlock.NewLoad(ptrType.ElemType, ptr)
					if load.Type() != types.I32 {
						panic("DEC requires INTEGER variable")
					}
					one := constant.NewInt(types.I32, 1)
					sub := g.curBlock.NewSub(load, one)
					g.curBlock.NewStore(sub, ptr)
					return nil
				}
			}
			panic("DEC argument must be an INTEGER variable")
		}
		// Встроенные функции (возвращают значение)
		if name == "ABS" && len(c.Args) == 1 {
			arg := g.genExpr(c.Args[0])
			if arg.Type() == types.I32 {
				zero := constant.NewInt(types.I32, 0)
				cmp := g.curBlock.NewICmp(enum.IPredSLT, arg, zero)
				neg := g.curBlock.NewSub(zero, arg)
				return g.curBlock.NewSelect(cmp, neg, arg)
			} else if arg.Type() == types.Float {
				// вызов llvm.fabs.f64
				fabsFunc := g.module.NewFunc("llvm.fabs.f64", types.Float, ir.NewParam("x", types.Float))
				return g.curBlock.NewCall(fabsFunc, arg)
			}
			panic("ABS: invalid argument type")
		}
		if name == "ODD" && len(c.Args) == 1 {
			arg := g.genExpr(c.Args[0])
			if arg.Type() != types.I32 {
				panic("ODD argument must be INTEGER")
			}
			one := constant.NewInt(types.I32, 1)
			and := g.curBlock.NewAnd(arg, one)
			zero := constant.NewInt(types.I32, 0)
			return g.curBlock.NewICmp(enum.IPredNE, and, zero)
		}
		if name == "REAL" && len(c.Args) == 1 {
			arg := g.genExpr(c.Args[0])
			if arg.Type() == types.I32 {
				return g.curBlock.NewSIToFP(arg, types.Float)
			}
			panic("REAL argument must be INTEGER")
		}
	}

	// Обычный вызов: поиск callee (пользовательская функция)
	var callee value.Value
	if des, ok := c.Callee.(*ast.DesignatorExpr); ok {
		name := des.Base.Name
		if des.Base.Module != "" {
			name = des.Base.Module + "." + name
		}
		var found bool
		callee, found = g.lookup(name)
		if !found {
			for _, f := range g.module.Funcs {
				if f.Name() == name {
					callee = f
					found = true
					break
				}
			}
			if !found {
				panic("unknown function: " + name)
			}
		}
	} else {
		callee = g.genExpr(c.Callee)
	}
	ptrType, ok := callee.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Sprintf("callee is not a pointer, type = %T", callee.Type()))
	}
	funcType, ok := ptrType.ElemType.(*types.FuncType)
	if !ok {
		panic(fmt.Sprintf("pointer does not point to function, elem type = %T", ptrType.ElemType))
	}
	args := make([]value.Value, len(c.Args))
	for i, a := range c.Args {
		expectedType := funcType.Params[i]
		var argVal value.Value

		// Если параметр – указатель (VAR параметр или открытый массив)
		if _, isPtr := expectedType.(*types.PointerType); isPtr {
			if des, ok := a.(*ast.DesignatorExpr); ok {
				// Передаём указатель на аргумент
				argVal = g.genDesignatorPtr(des)
			} else {
				argVal = g.genExpr(a)
			}
		} else {
			argVal = g.genExpr(a)
		}

		// Приведение типа
		if argVal.Type() != expectedType {
			if argVal.Type() == types.I1 && expectedType == types.I32 {
				argVal = g.curBlock.NewZExt(argVal, types.I32)
			} else if argVal.Type() == types.I32 && expectedType == types.I1 {
				argVal = g.curBlock.NewTrunc(argVal, types.I1)
			}
		}
		args[i] = argVal
	}
	return g.curBlock.NewCall(callee, args...)
}

func (g *Generator) genStmt(s ast.Stmt) {
	switch st := s.(type) {
	case *ast.AssignmentStmt:
		ptr := g.genDesignatorPtr(st.Target)
		val := g.genExpr(st.Value)
		g.curBlock.NewStore(val, ptr)
	case *ast.ProcedureCallStmt:
		g.genCall(st.Call)
	case *ast.IfStmt:
		g.genIfStmt(st)
	case *ast.WhileStmt:
		g.genWhileStmt(st)
	case *ast.RepeatStmt:
		g.genRepeatStmt(st)
	case *ast.ForStmt:
		g.genForStmt(st)
	case *ast.CaseStmt:
		g.genCaseStmt(st)
	default:
		panic(fmt.Sprintf("statement not implemented: %T", s))
	}
}

func (g *Generator) genIfStmt(st *ast.IfStmt) {
	mergeBlock := g.curFunc.NewBlock(g.nextBlockName("ifmerge"))
	var elseBlock *ir.Block
	if len(st.ElseBody) > 0 {
		elseBlock = g.curFunc.NewBlock(g.nextBlockName("ifelse"))
	}
	for i, branch := range st.Branches {
		cond := g.genExpr(branch.Cond)
		thenBlock := g.curFunc.NewBlock(g.nextBlockName(fmt.Sprintf("ifthen_%d", i)))
		var nextBlock *ir.Block
		if i == len(st.Branches)-1 && elseBlock != nil {
			nextBlock = elseBlock
		} else {
			nextBlock = g.curFunc.NewBlock(g.nextBlockName(fmt.Sprintf("ifcond_%d", i+1)))
		}
		g.curBlock.NewCondBr(cond, thenBlock, nextBlock)
		g.curBlock = thenBlock
		for _, s := range branch.Body {
			g.genStmt(s)
		}
		g.curBlock.NewBr(mergeBlock)
		if i < len(st.Branches)-1 || elseBlock != nil {
			g.curBlock = nextBlock
		}
	}
	if elseBlock != nil {
		g.curBlock = elseBlock
		for _, s := range st.ElseBody {
			g.genStmt(s)
		}
		g.curBlock.NewBr(mergeBlock)
	}
	g.curBlock = mergeBlock
}

func (g *Generator) genWhileStmt(st *ast.WhileStmt) {
	if len(st.Branches) == 0 {
		return
	}
	condBlock := g.curFunc.NewBlock(g.nextBlockName("whilecond"))
	bodyBlock := g.curFunc.NewBlock(g.nextBlockName("whilebody"))
	exitBlock := g.curFunc.NewBlock(g.nextBlockName("whileexit"))
	g.curBlock.NewBr(condBlock)
	g.curBlock = condBlock
	cond := g.genExpr(st.Branches[0].Cond)
	g.curBlock.NewCondBr(cond, bodyBlock, exitBlock)
	g.curBlock = bodyBlock
	for _, s := range st.Branches[0].Body {
		g.genStmt(s)
	}
	g.curBlock.NewBr(condBlock)
	g.curBlock = exitBlock
}

func (g *Generator) genRepeatStmt(st *ast.RepeatStmt) {
	bodyBlock := g.curFunc.NewBlock(g.nextBlockName("repeatbody"))
	condBlock := g.curFunc.NewBlock(g.nextBlockName("repeatcond"))
	exitBlock := g.curFunc.NewBlock(g.nextBlockName("repeatexit"))
	g.curBlock.NewBr(bodyBlock)
	g.curBlock = bodyBlock
	for _, s := range st.Body {
		g.genStmt(s)
	}
	g.curBlock.NewBr(condBlock)
	g.curBlock = condBlock
	cond := g.genExpr(st.Until)
	g.curBlock.NewCondBr(cond, exitBlock, bodyBlock)
	g.curBlock = exitBlock
}

func (g *Generator) genForStmt(st *ast.ForStmt) {
	varPtr, ok := g.lookup(st.Var)
	if !ok {
		panic("for loop variable not found: " + st.Var)
	}
	fromVal := g.genExpr(st.From)
	g.curBlock.NewStore(fromVal, varPtr)
	condBlock := g.curFunc.NewBlock(g.nextBlockName("forcond"))
	bodyBlock := g.curFunc.NewBlock(g.nextBlockName("forbody"))
	exitBlock := g.curFunc.NewBlock(g.nextBlockName("forexit"))
	g.curBlock.NewBr(condBlock)
	g.curBlock = condBlock
	ptrType, ok := varPtr.Type().(*types.PointerType)
	if !ok {
		panic("loop variable is not a pointer")
	}
	curVal := g.curBlock.NewLoad(ptrType.ElemType, varPtr)
	toVal := g.genExpr(st.To)
	cmp := g.curBlock.NewICmp(enum.IPredSLE, curVal, toVal)
	g.curBlock.NewCondBr(cmp, bodyBlock, exitBlock)
	g.curBlock = bodyBlock
	for _, s := range st.Body {
		g.genStmt(s)
	}
	step := value.Value(constant.NewInt(types.I32, 1))
	if st.HasBy {
		step = g.genExpr(st.By)
	}
	newVal := g.curBlock.NewAdd(curVal, step)
	g.curBlock.NewStore(newVal, varPtr)
	g.curBlock.NewBr(condBlock)
	g.curBlock = exitBlock
}

func (g *Generator) genCaseStmt(stmt *ast.CaseStmt) {
	selector := g.genExpr(stmt.Expr)
	exitBlock := g.curFunc.NewBlock(g.nextBlockName("case_exit"))
	var elseBlock *ir.Block = exitBlock

	for _, branch := range stmt.Branches {
		if len(branch.Labels) == 0 {
			elseBlock = g.curFunc.NewBlock(g.nextBlockName("case_else"))
			break
		}
	}
	for i, branch := range stmt.Branches {
		if len(branch.Labels) == 0 {
			continue
		}
		var cond value.Value
		for j, label := range branch.Labels {
			fromVal := g.constExprToInt(label.From)
			var cmp value.Value
			if label.To != nil {
				toVal := g.constExprToInt(label.To)
				cmpFrom := g.curBlock.NewICmp(enum.IPredSGE, selector, constant.NewInt(types.I32, int64(fromVal)))
				cmpTo := g.curBlock.NewICmp(enum.IPredSLE, selector, constant.NewInt(types.I32, int64(toVal)))
				cmp = g.curBlock.NewAnd(cmpFrom, cmpTo)
			} else {
				cmp = g.curBlock.NewICmp(enum.IPredEQ, selector, constant.NewInt(types.I32, int64(fromVal)))
			}
			if j == 0 {
				cond = cmp
			} else {
				cond = g.curBlock.NewOr(cond, cmp)
			}
		}
		thenBlock := g.curFunc.NewBlock(g.nextBlockName(fmt.Sprintf("case_body_%d", i)))
		nextBlock := g.curFunc.NewBlock(g.nextBlockName(fmt.Sprintf("case_next_%d", i)))
		g.curBlock.NewCondBr(cond, thenBlock, nextBlock)

		g.curBlock = thenBlock
		for _, s := range branch.Body {
			g.genStmt(s)
		}
		g.curBlock.NewBr(exitBlock)

		g.curBlock = nextBlock
	}
	if elseBlock != exitBlock {
		g.curBlock.NewBr(elseBlock)
		g.curBlock = elseBlock
		for _, branch := range stmt.Branches {
			if len(branch.Labels) == 0 {
				for _, s := range branch.Body {
					g.genStmt(s)
				}
				break
			}
		}
		g.curBlock.NewBr(exitBlock)
	} else {
		g.curBlock.NewBr(exitBlock)
	}
	g.curBlock = exitBlock
}

func (g *Generator) namedTypeToLLVM(typeName string) types.Type {
	switch typeName {
	case "INTEGER":
		return types.I32
	case "BOOLEAN":
		return types.I1
	case "SET":
		return types.I32
	default:
		return types.I32
	}
}

func (g *Generator) isRealType(typ types.Type) bool {
	return typ == types.Float
}

func (g *Generator) promoteToReal(val value.Value) value.Value {
	if val.Type() == types.I32 {
		return g.curBlock.NewSIToFP(val, types.Float)
	}
	return val
}

// Константное вычисление (без LLVM IR)
func (g *Generator) constEval(e ast.Expr) (isReal bool, intVal int64, realVal float64) {
	switch ex := e.(type) {
	case *ast.NumberExpr:
		if ex.IsReal {
			val, _ := strconv.ParseFloat(ex.Text, 64)
			return true, 0, val
		} else {
			val, _ := strconv.ParseInt(ex.Text, 10, 64)
			return false, val, 0
		}
	case *ast.BinaryExpr:
		isRealL, intL, realL := g.constEval(ex.Left)
		isRealR, intR, realR := g.constEval(ex.Right)
		if isRealL || isRealR {
			l := realL
			if !isRealL {
				l = float64(intL)
			}
			r := realR
			if !isRealR {
				r = float64(intR)
			}
			var res float64
			switch ex.Op {
			case "+":
				res = l + r
			case "-":
				res = l - r
			case "*":
				res = l * r
			case "/":
				res = l / r
			default:
				panic("unsupported real const op")
			}
			return true, 0, res
		} else {
			var res int64
			switch ex.Op {
			case "+":
				res = intL + intR
			case "-":
				res = intL - intR
			case "*":
				res = intL * intR
			case "/":
				res = intL / intR
			case "DIV":
				res = intL / intR
			case "MOD":
				res = intL % intR
			default:
				panic("unsupported integer const op")
			}
			return false, res, 0
		}
	default:
		panic("non-constant expression in const declaration")
	}
}
