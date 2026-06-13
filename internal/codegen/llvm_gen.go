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
	module              *ir.Module
	moduleName          string
	curBlock            *ir.Block
	curFunc             *ir.Func
	symStack            []map[string]value.Value
	fieldNameToIndex    map[*types.StructType]map[string]int
	recordTypes         map[string]*ast.RecordType
	blockCounter        int
	currentProc         *ast.ProcedureDecl
	framePtr            value.Value
	varInfo             map[string]VarInfo
	procFrames          map[*ast.ProcedureDecl]*types.StructType
	currentResultAlloca value.Value // alloca для возвращаемого значения текущей функции
	currentRetType      types.Type  // тип возврата текущей функции
}

type VarInfo struct {
	NestingLevel int
	FrameOffset  int
	FrameType    *types.StructType
	IsGlobal     bool
	GlobalValue  value.Value
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
		varInfo:          make(map[string]VarInfo),
		procFrames:       make(map[*ast.ProcedureDecl]*types.StructType),
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

func (g *Generator) buildFrameType(proc *ast.ProcedureDecl) *types.StructType {
	if t, ok := g.procFrames[proc]; ok {
		return t
	}
	var fields []types.Type
	// static link only for nesting level >= 2
	if proc.NestingLevel > 1 {
		fields = append(fields, types.I8Ptr)
	}
	if proc.Declarations != nil {
		for _, v := range proc.Declarations.Vars {
			typ := g.oberonTypeToLLVM(v.Type)
			for range v.Names {
				fields = append(fields, typ)
			}
		}
	}
	if proc.Signature != nil {
		for _, ps := range proc.Signature.Params {
			if ps.ByRef {
				continue
			}
			typ := g.oberonTypeToLLVM(ps.Type)
			for range ps.Names {
				fields = append(fields, typ)
			}
		}
	}
	frameType := types.NewStruct(fields...)
	g.procFrames[proc] = frameType
	return frameType
}

func (g *Generator) registerVar(name string, nestingLevel int, frameOffset int, frameType *types.StructType) {
	g.varInfo[name] = VarInfo{
		NestingLevel: nestingLevel,
		FrameOffset:  frameOffset,
		FrameType:    frameType,
	}
}

func (g *Generator) registerVarParam(name string, nestingLevel int, val value.Value) {
	g.varInfo[name] = VarInfo{
		NestingLevel: nestingLevel,
		FrameOffset:  -1,
		GlobalValue:  val,
		IsGlobal:     false,
	}
}

func (g *Generator) getVarPtr(name string) value.Value {
	info, ok := g.varInfo[name]
	if !ok {
		if val, ok := g.lookup(name); ok {
			if _, ok := val.Type().(*types.PointerType); ok {
				return val
			}
			panic("variable " + name + " is not a pointer")
		}
		panic("unknown variable: " + name)
	}
	if info.IsGlobal {
		return info.GlobalValue
	}
	if info.FrameOffset == -1 {
		// VAR-параметр или переданный указатель
		return info.GlobalValue
	}
	if info.NestingLevel == g.currentProc.NestingLevel {
		frameType := g.buildFrameType(g.currentProc)
		ptr := g.curBlock.NewGetElementPtr(frameType, g.framePtr,
			constant.NewInt(types.I32, 0),
			constant.NewInt(types.I32, int64(info.FrameOffset)))
		return ptr
	}
	// Нелокальная переменная: поднимаемся по статическим ссылкам
	curProc := g.currentProc
	frame := g.framePtr
	for curProc.NestingLevel > info.NestingLevel {
		frameType := g.buildFrameType(curProc)
		staticLinkPtr := g.curBlock.NewGetElementPtr(frameType, frame,
			constant.NewInt(types.I32, 0),
			constant.NewInt(types.I32, 0))
		frame = g.curBlock.NewLoad(types.I8Ptr, staticLinkPtr)
		curProc = curProc.Parent
		if curProc == nil {
			panic("broken parent chain for variable " + name)
		}
	}
	frameTyped := g.curBlock.NewBitCast(frame, types.NewPointer(info.FrameType))
	ptr := g.curBlock.NewGetElementPtr(info.FrameType, frameTyped,
		constant.NewInt(types.I32, 0),
		constant.NewInt(types.I32, int64(info.FrameOffset)))
	return ptr
}

func (g *Generator) oberonTypeToLLVM(t ast.TypeExpr) types.Type {
	switch tt := t.(type) {
	case *ast.NamedType:
		switch tt.Name {
		case "INTEGER":
			return types.I32
		case "BOOLEAN":
			return types.I1
		case "REAL":
			return types.Float
		case "SET":
			return types.I32
		default:
			if rec, ok := g.recordTypes[tt.Name]; ok {
				return g.oberonTypeToLLVM(rec)
			}
			return types.I32
		}
	case *ast.ArrayType:
		elemType := g.oberonTypeToLLVM(tt.Elem)
		if len(tt.Lengths) == 0 {
			return types.NewPointer(elemType)
		}
		for i := len(tt.Lengths) - 1; i >= 0; i-- {
			length := g.constExprToInt(tt.Lengths[i])
			elemType = types.NewArray(uint64(length), elemType)
		}
		return elemType
	case *ast.RecordType:
		var fields []types.Type
		var fieldNames []string
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
		if ex.IsReal {
			panic("constExprToInt called on REAL constant")
		}
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

func (g *Generator) constEval(e ast.Expr) (isReal bool, intVal int64, realVal float64) {
	switch ex := e.(type) {
	case *ast.NumberExpr:
		if ex.IsReal {
			val, err := strconv.ParseFloat(ex.Text, 64)
			if err != nil {
				panic("invalid real number: " + ex.Text)
			}
			return true, 0, val
		}
		val, err := strconv.ParseInt(ex.Text, 10, 64)
		if err != nil {
			panic("invalid integer number: " + ex.Text)
		}
		return false, val, 0
	case *ast.BinaryExpr:
		leftIsReal, leftInt, leftReal := g.constEval(ex.Left)
		rightIsReal, rightInt, rightReal := g.constEval(ex.Right)
		if leftIsReal || rightIsReal {
			l := leftReal
			if !leftIsReal {
				l = float64(leftInt)
			}
			r := rightReal
			if !rightIsReal {
				r = float64(rightInt)
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
		}
		var res int64
		switch ex.Op {
		case "+":
			res = leftInt + rightInt
		case "-":
			res = leftInt - rightInt
		case "*":
			res = leftInt * rightInt
		case "/":
			res = leftInt / rightInt
		default:
			panic("unsupported integer const op")
		}
		return false, res, 0
	default:
		panic("non-constant expression in const declaration")
	}
}

func (g *Generator) constExprToIntOrNil(e ast.Expr) *int64 {
	switch ex := e.(type) {
	case *ast.NumberExpr:
		if ex.IsReal {
			return nil
		}
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

func (g *Generator) isRealType(typ types.Type) bool {
	return typ == types.Float
}

func (g *Generator) promoteToReal(val value.Value) value.Value {
	if val.Type() == types.I32 {
		return g.curBlock.NewSIToFP(val, types.Float)
	}
	return val
}

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
		if isReal {
			g.addSymbol(c.Name.Name, constant.NewFloat(types.Float, realVal))
		} else {
			g.addSymbol(c.Name.Name, constant.NewInt(types.I32, intVal))
		}
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
			if isGlobal {
				global := g.module.NewGlobal(name, llvmType)
				global.Init = constant.NewZeroInitializer(llvmType)
				if nameDef.Exported {
					global.Linkage = enum.LinkageExternal
				}
				g.addSymbol(name, global)
				g.varInfo[name] = VarInfo{IsGlobal: true, GlobalValue: global}
			}
		}
	}
	for _, proc := range decls.Procedures {
		g.generateProcedure(proc, isGlobal)
	}
}

func (g *Generator) generateProcedure(proc *ast.ProcedureDecl, isGlobal bool) {
	oldProc := g.currentProc
	g.currentProc = proc

	frameType := g.buildFrameType(proc)

	var llvmParams []*ir.Param
	if proc.NestingLevel > 1 {
		llvmParams = append(llvmParams, ir.NewParam("__static_link", types.I8Ptr))
	}
	if proc.Signature != nil {
		for _, paramSec := range proc.Signature.Params {
			paramType := g.oberonTypeToLLVM(paramSec.Type)
			if paramSec.ByRef {
				paramType = types.NewPointer(paramType)
			}
			for range paramSec.Names {
				llvmParams = append(llvmParams, ir.NewParam("", paramType))
			}
		}
	}
	var retType types.Type = types.Void
	if proc.Signature != nil && proc.Signature.ReturnType != "" {
		retType = g.namedTypeToLLVM(proc.Signature.ReturnType)
	}
	funcName := proc.Name.Name
	f := g.module.NewFunc(funcName, retType, llvmParams...)
	if proc.Name.Exported {
		f.Linkage = enum.LinkageExternal
	}
	g.addSymbol(funcName, f)
	g.varInfo[funcName] = VarInfo{NestingLevel: proc.NestingLevel}

	oldFunc := g.curFunc
	oldBlock := g.curBlock
	oldFramePtr := g.framePtr
	oldResultAlloca := g.currentResultAlloca // NEW
	oldRetType := g.currentRetType           // NEW
	defer func() {
		g.curFunc = oldFunc
		g.curBlock = oldBlock
		g.framePtr = oldFramePtr
		g.currentProc = oldProc
		g.currentResultAlloca = oldResultAlloca // NEW
		g.currentRetType = oldRetType           // NEW
	}()
	g.curFunc = f
	entry := f.NewBlock("entry")
	g.curBlock = entry
	g.pushScope()
	defer g.popScope()

	frameAlloca := entry.NewAlloca(frameType)
	g.framePtr = frameAlloca

	fieldOffset := 0
	if proc.NestingLevel > 1 {
		staticLink := f.Params[0]
		linkField := entry.NewGetElementPtr(frameType, frameAlloca,
			constant.NewInt(types.I32, 0),
			constant.NewInt(types.I32, 0))
		entry.NewStore(staticLink, linkField)
		fieldOffset = 1
	}

	// Создаём alloca для возвращаемого значения (если функция)
	var resultAlloca value.Value = nil
	if retType != types.Void {
		resultAlloca = entry.NewAlloca(retType)
		// Не добавляем в symStack, чтобы не мешать вызовам функции
	}
	g.currentResultAlloca = resultAlloca // NEW
	g.currentRetType = retType           // NEW

	// Локальные переменные
	if proc.Declarations != nil {
		for _, v := range proc.Declarations.Vars {
			varType := g.oberonTypeToLLVM(v.Type)
			for _, nameDef := range v.Names {
				ptr := entry.NewGetElementPtr(frameType, frameAlloca,
					constant.NewInt(types.I32, 0),
					constant.NewInt(types.I32, int64(fieldOffset)))
				entry.NewStore(constant.NewZeroInitializer(varType), ptr)
				g.registerVar(nameDef.Name, proc.NestingLevel, fieldOffset, frameType)
				fieldOffset++
			}
		}
	}

	// Параметры
	paramIdx := 0
	if proc.NestingLevel > 1 {
		paramIdx = 1
	}
	if proc.Signature != nil {
		for _, paramSec := range proc.Signature.Params {
			for _, paramName := range paramSec.Names {
				paramVal := f.Params[paramIdx]
				if paramSec.ByRef {
					// VAR-параметр – сохраняем как есть
					g.registerVarParam(paramName, proc.NestingLevel, paramVal)
				} else {
					ptr := entry.NewGetElementPtr(frameType, frameAlloca,
						constant.NewInt(types.I32, 0),
						constant.NewInt(types.I32, int64(fieldOffset)))
					entry.NewStore(paramVal, ptr)
					g.registerVar(paramName, proc.NestingLevel, fieldOffset, frameType)
					fieldOffset++
				}
				paramIdx++
			}
		}
	}

	// Генерация вложенных процедур
	if proc.Declarations != nil {
		for _, nestedProc := range proc.Declarations.Procedures {
			g.generateProcedure(nestedProc, false)
		}
	}

	// Генерация тела
	for _, stmt := range proc.Body {
		g.genStmt(stmt)
	}

	// Возврат значения
	if proc.ReturnExpr != nil {
		retVal := g.genExpr(proc.ReturnExpr)
		g.curBlock.NewRet(retVal)
	} else if retType != types.Void {
		if resultAlloca == nil {
			panic("missing return alloca")
		}
		retVal := g.curBlock.NewLoad(retType, resultAlloca)
		g.curBlock.NewRet(retVal)
	} else {
		g.curBlock.NewRet(nil)
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
			val, _ := strconv.ParseFloat(ex.Text, 64)
			return constant.NewFloat(types.Float, val)
		}
		val, _ := strconv.ParseInt(ex.Text, 10, 32)
		return constant.NewInt(types.I32, val)
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
			if sub.Type() == types.Float {
				zero := constant.NewFloat(types.Float, 0)
				return g.curBlock.NewFSub(zero, sub)
			}
			intType, ok := sub.Type().(*types.IntType)
			if !ok {
				panic("unary minus on non-integer/non-float")
			}
			zero := constant.NewInt(intType, 0)
			return g.curBlock.NewSub(zero, sub)
		} else if ex.Op == "~" {
			one := constant.NewInt(types.I1, 1)
			return g.curBlock.NewXor(one, sub)
		}
		return sub
	case *ast.DesignatorExpr:
		if len(ex.Selectors) == 0 {
			name := ex.Base.Name
			if val, ok := g.lookup(name); ok {
				if _, ok := val.(*constant.Int); ok {
					return val
				}
				if _, ok := val.(*constant.Float); ok {
					return val
				}
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
	var basePtr value.Value
	if len(d.Selectors) == 0 && d.Base.Module == "" {
		// Простая переменная: всегда получаем через getVarPtr
		name := d.Base.Name
		basePtr = g.getVarPtr(name)
	} else {
		// Сложный дизайнатор: начинаем с lookup (возможно, глобальный или функция)
		name := d.Base.Name
		baseVal, ok := g.lookup(name)
		if !ok {
			baseVal = g.getVarPtr(name)
		}
		basePtr = baseVal
	}
	cur := basePtr
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
	// Built-in functions
	if des, ok := c.Callee.(*ast.DesignatorExpr); ok {
		name := des.Base.Name
		if des.Base.Module != "" {
			name = des.Base.Module + "." + name
		}
		switch name {
		case "INC":
			if len(c.Args) == 1 {
				if des, ok := c.Args[0].(*ast.DesignatorExpr); ok {
					ptr := g.genDesignatorPtr(des)
					if ptrType, ok := ptr.Type().(*types.PointerType); ok {
						load := g.curBlock.NewLoad(ptrType.ElemType, ptr)
						one := constant.NewInt(types.I32, 1)
						added := g.curBlock.NewAdd(load, one)
						g.curBlock.NewStore(added, ptr)
						return nil
					}
				}
				panic("INC argument must be a variable")
			}
		case "DEC":
			if len(c.Args) == 1 {
				if des, ok := c.Args[0].(*ast.DesignatorExpr); ok {
					ptr := g.genDesignatorPtr(des)
					if ptrType, ok := ptr.Type().(*types.PointerType); ok {
						load := g.curBlock.NewLoad(ptrType.ElemType, ptr)
						one := constant.NewInt(types.I32, 1)
						sub := g.curBlock.NewSub(load, one)
						g.curBlock.NewStore(sub, ptr)
						return nil
					}
				}
				panic("DEC argument must be a variable")
			}
		case "ABS":
			if len(c.Args) == 1 {
				arg := g.genExpr(c.Args[0])
				if arg.Type() == types.I32 {
					zero := constant.NewInt(types.I32, 0)
					cmp := g.curBlock.NewICmp(enum.IPredSLT, arg, zero)
					neg := g.curBlock.NewSub(zero, arg)
					return g.curBlock.NewSelect(cmp, neg, arg)
				} else if arg.Type() == types.Float {
					zero := constant.NewFloat(types.Float, 0)
					cmp := g.curBlock.NewFCmp(enum.FPredOLT, arg, zero)
					neg := g.curBlock.NewFSub(zero, arg)
					return g.curBlock.NewSelect(cmp, neg, arg)
				}
				panic("ABS: invalid argument type")
			}
		case "ODD":
			if len(c.Args) == 1 {
				arg := g.genExpr(c.Args[0])
				if arg.Type() != types.I32 {
					panic("ODD argument must be INTEGER")
				}
				one := constant.NewInt(types.I32, 1)
				and := g.curBlock.NewAnd(arg, one)
				zero := constant.NewInt(types.I32, 0)
				return g.curBlock.NewICmp(enum.IPredNE, and, zero)
			}
		case "REAL":
			if len(c.Args) == 1 {
				arg := g.genExpr(c.Args[0])
				if arg.Type() == types.I32 {
					return g.curBlock.NewSIToFP(arg, types.Float)
				}
				panic("REAL argument must be INTEGER")
			}
		}
	}

	// Normal function call
	var calleeVal value.Value
	var calleeName string
	if des, ok := c.Callee.(*ast.DesignatorExpr); ok {
		calleeName = des.Base.Name
		if des.Base.Module != "" {
			calleeName = des.Base.Module + "." + calleeName
		}
		var found bool
		calleeVal, found = g.lookup(calleeName)
		if !found {
			for _, f := range g.module.Funcs {
				if f.Name() == calleeName {
					calleeVal = f
					found = true
					break
				}
			}
			if !found {
				panic("unknown function: " + calleeName)
			}
		}
	} else {
		calleeVal = g.genExpr(c.Callee)
	}
	ptrType, ok := calleeVal.Type().(*types.PointerType)
	if !ok {
		panic(fmt.Sprintf("callee is not a pointer, type = %T", calleeVal.Type()))
	}
	funcType, ok := ptrType.ElemType.(*types.FuncType)
	if !ok {
		panic(fmt.Sprintf("pointer does not point to function, elem type = %T", ptrType.ElemType))
	}
	args := []value.Value{}

	if info, ok := g.varInfo[calleeName]; ok && info.NestingLevel > 1 {
		if g.framePtr == nil {
			panic("no frame pointer for nested call")
		}
		framePtrAsI8 := g.curBlock.NewBitCast(g.framePtr, types.I8Ptr)
		args = append(args, framePtrAsI8)
	}

	for i, a := range c.Args {
		expectedType := funcType.Params[i+len(args)]
		var argVal value.Value
		if _, isPtr := expectedType.(*types.PointerType); isPtr {
			if des, ok := a.(*ast.DesignatorExpr); ok {
				argVal = g.genDesignatorPtr(des)
			} else {
				argVal = g.genExpr(a)
			}
		} else {
			argVal = g.genExpr(a)
		}
		if argVal.Type() != expectedType {
			if argVal.Type() == types.I1 && expectedType == types.I32 {
				argVal = g.curBlock.NewZExt(argVal, types.I32)
			} else if argVal.Type() == types.I32 && expectedType == types.I1 {
				argVal = g.curBlock.NewTrunc(argVal, types.I1)
			}
		}
		args = append(args, argVal)
	}
	return g.curBlock.NewCall(calleeVal, args...)
}

func (g *Generator) genStmt(s ast.Stmt) {
	switch st := s.(type) {
	case *ast.AssignmentStmt:
		// Присваивание имени функции (возвращаемого значения)
		if len(st.Target.Selectors) == 0 && st.Target.Base.Module == "" {
			name := st.Target.Base.Name
			if g.currentProc != nil && name == g.currentProc.Name.Name && g.currentRetType != types.Void {
				val := g.genExpr(st.Value)
				if g.currentResultAlloca == nil {
					panic("no result alloca for function")
				}
				g.curBlock.NewStore(val, g.currentResultAlloca)
				return
			}
		}
		// Старое обычное присваивание (через указатель)
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
	varPtr := g.getVarPtr(st.Var)
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
	case "REAL":
		return types.Float
	case "SET":
		return types.I32
	default:
		return types.I32
	}
}
