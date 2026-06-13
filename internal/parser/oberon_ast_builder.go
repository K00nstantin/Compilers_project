package parser

import (
	"fmt"

	"github.com/K00nstantin/Compilers_project/internal/ast"
	"github.com/antlr4-go/antlr/v4"
)

type ASTBuilder struct {
	*BaseoberonVisitor
	nestingLevel int
	currentProc  *ast.ProcedureDecl
}

func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{
		BaseoberonVisitor: &BaseoberonVisitor{BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{}},
		nestingLevel:      0,
		currentProc:       nil,
	}
}

var _ oberonVisitor = (*ASTBuilder)(nil)

func (v *ASTBuilder) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}
	return tree.Accept(v)
}

func (v *ASTBuilder) VisitChildren(node antlr.RuleNode) interface{} {
	if node == nil {
		return nil
	}
	var result interface{}
	for i := 0; i < node.GetChildCount(); i++ {
		if child, ok := node.GetChild(i).(antlr.ParseTree); ok && child != nil {
			result = child.Accept(v)
		}
	}
	return result
}

func (v *ASTBuilder) VisitModule(ctx *ModuleContext) interface{} {
	names := ctx.AllIdent()
	mod := &ast.Module{Declarations: &ast.DeclarationBlock{}}
	if len(names) > 0 {
		mod.Name = names[0].GetText()
		mod.EndName = names[len(names)-1].GetText()
	}
	if il := ctx.ImportList(); il != nil {
		mod.Imports = v.Visit(il).([]*ast.Import)
	}
	if ds := ctx.DeclarationSequence(); ds != nil {
		if decls := v.Visit(ds).(*ast.DeclarationBlock); decls != nil {
			mod.Declarations = decls
		}
	}
	if ss := ctx.StatementSequence(); ss != nil {
		mod.Body = v.Visit(ss).([]ast.Stmt)
	}
	return mod
}

func (v *ASTBuilder) VisitImportList(ctx *ImportListContext) interface{} {
	imports := make([]*ast.Import, 0, len(ctx.AllImport_()))
	for _, ic := range ctx.AllImport_() {
		imports = append(imports, v.Visit(ic).(*ast.Import))
	}
	return imports
}

func (v *ASTBuilder) VisitImport_(ctx *Import_Context) interface{} {
	ids := ctx.AllIdent()
	if len(ids) == 1 {
		name := ids[0].GetText()
		return &ast.Import{Alias: name, Module: name}
	}
	return &ast.Import{Alias: ids[0].GetText(), Module: ids[1].GetText()}
}

func (v *ASTBuilder) VisitDeclarationSequence(ctx *DeclarationSequenceContext) interface{} {
	block := &ast.DeclarationBlock{}
	for _, c := range ctx.AllConstDeclaration() {
		decl := v.Visit(c).(*ast.ConstDecl)
		block.Consts = append(block.Consts, decl)
	}
	for _, t := range ctx.AllTypeDeclaration() {
		tp := v.Visit(t).(*ast.TypeDecl)
		block.Types = append(block.Types, tp)
	}
	for _, vrs := range ctx.AllVariableDeclaration() {
		vs := v.Visit(vrs).(*ast.VarDecl)
		block.Vars = append(block.Vars, vs)
	}
	for _, prs := range ctx.AllProcedureDeclaration() {
		ps := v.Visit(prs).(*ast.ProcedureDecl)
		block.Procedures = append(block.Procedures, ps)
	}
	return block
}

func (v *ASTBuilder) VisitConstDeclaration(ctx *ConstDeclarationContext) interface{} {
	name := v.Visit(ctx.Identdef()).(ast.IdentDef)
	value := v.Visit(ctx.ConstExpression()).(ast.Expr)
	return &ast.ConstDecl{Name: name, Value: value}
}

func (v *ASTBuilder) VisitConstExpression(ctx *ConstExpressionContext) interface{} {
	expr := v.Visit(ctx.Expression())
	if expr == nil {
		panic("const expression is nil")
	}
	return expr
}

func (v *ASTBuilder) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	name := v.Visit(ctx.Identdef()).(ast.IdentDef)
	tp := v.Visit(ctx.Type_()).(ast.TypeExpr)
	return &ast.TypeDecl{Name: name, Type: tp}
}

func (v *ASTBuilder) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	names := v.Visit(ctx.IdentList()).([]ast.IdentDef)
	tp := v.Visit(ctx.Type_()).(ast.TypeExpr)
	return &ast.VarDecl{Names: names, Type: tp}
}

func (v *ASTBuilder) VisitProcedureDeclaration(ctx *ProcedureDeclarationContext) interface{} {
	head := v.Visit(ctx.ProcedureHeading()).(*ast.ProcedureDecl)
	// Уровень = текущий уровень + 1
	head.NestingLevel = v.nestingLevel + 1
	head.Parent = v.currentProc

	v.nestingLevel++
	oldProc := v.currentProc
	v.currentProc = head

	body := v.Visit(ctx.ProcedureBody()).(*ast.ProcedureDecl)
	head.Declarations = body.Declarations
	head.Body = body.Body
	head.ReturnExpr = body.ReturnExpr
	if endName := ctx.Ident(); endName != nil {
		head.EndName = endName.GetText()
	}

	v.nestingLevel--
	v.currentProc = oldProc
	return head
}

func (v *ASTBuilder) VisitProcedureHeading(ctx *ProcedureHeadingContext) interface{} {
	proc := &ast.ProcedureDecl{}
	if identdef := ctx.Identdef(); identdef != nil {
		if ident := identdef.Ident(); ident != nil {
			proc.Name = ast.IdentDef{
				Name:     ident.GetText(),
				Exported: identdef.GetChildCount() > 1,
			}
		} else {
			proc.Name = ast.IdentDef{Name: "__error", Exported: false}
		}
	} else {
		proc.Name = ast.IdentDef{Name: "__error", Exported: false}
	}
	if fp := ctx.FormalParameters(); fp != nil {
		if val := v.Visit(fp); val != nil {
			if sig, ok := val.(*ast.ProcedureSignature); ok {
				proc.Signature = sig
			}
		}
	}
	return proc
}

func (v *ASTBuilder) VisitProcedureBody(ctx *ProcedureBodyContext) interface{} {
	proc := &ast.ProcedureDecl{}
	proc.Declarations = v.Visit(ctx.DeclarationSequence()).(*ast.DeclarationBlock)
	bd, ok := v.Visit(ctx.StatementSequence()).([]ast.Stmt)
	if ok && bd != nil {
		proc.Body = bd
	}
	re, ok := v.Visit(ctx.Expression()).(ast.Expr)
	if ok && re != nil {
		proc.ReturnExpr = re
	}
	return proc
}

func (v *ASTBuilder) VisitFormalParameters(ctx *FormalParametersContext) interface{} {
	sign := &ast.ProcedureSignature{}
	for _, ps := range ctx.AllFPSection() {
		fps := v.Visit(ps).(*ast.ParamSection)
		sign.Params = append(sign.Params, fps)
	}
	if rt := ctx.Qualident(); rt != nil {
		sign.ReturnType = v.qualidentText(rt)
	}
	return sign
}

// VisitFPSection обрабатывает секцию формальных параметров
func (v *ASTBuilder) VisitFPSection(ctx *FPSectionContext) interface{} {
	section := &ast.ParamSection{ByRef: ctx.VAR() != nil}
	for _, id := range ctx.AllIdent() {
		section.Names = append(section.Names, id.GetText())
	}
	if ft := ctx.FormalType(); ft != nil {
		// ft имеет интерфейсный тип IFormalTypeContext, приводим к конкретному
		if formalTypeCtx, ok := ft.(*FormalTypeContext); ok {
			section.Type = v.VisitFormalType(formalTypeCtx).(ast.TypeExpr)
		} else {
			panic(fmt.Sprintf("unexpected formal type context type: %T", ft))
		}
	}
	return section
}

// VisitFormalType обрабатывает формальный тип (ARRAY OF ...)
func (v *ASTBuilder) VisitFormalType(ctx *FormalTypeContext) interface{} {
	// Если есть ARRAY, значит это открытый массив
	if len(ctx.AllARRAY()) > 0 {
		// Базовый тип - последний qualident
		var result ast.TypeExpr = &ast.NamedType{Name: v.qualidentText(ctx.Qualident())}
		// Создаём цепочку открытых массивов (каждый ARRAY OF даёт один уровень)
		for i := 0; i < len(ctx.AllARRAY()); i++ {
			result = &ast.ArrayType{Lengths: []ast.Expr{}, Elem: result}
		}
		return result
	}
	// Иначе это просто qualident
	return &ast.NamedType{Name: v.qualidentText(ctx.Qualident())}
}

func (v *ASTBuilder) VisitType_(ctx *Type_Context) interface{} {
	switch {
	case ctx.Qualident() != nil:
		return &ast.NamedType{Name: v.qualidentText(ctx.Qualident())}
	case ctx.ArrayType() != nil:
		arr := v.Visit(ctx.ArrayType())
		if arr == nil {
			panic("arrayType returned nil")
		}
		return arr
	case ctx.RecordType() != nil:
		rec := v.Visit(ctx.RecordType())
		if rec == nil {
			panic("recordType returned nil")
		}
		return rec
	case ctx.PointerType() != nil:
		ptr := v.Visit(ctx.PointerType())
		if ptr == nil {
			panic("pointerType returned nil")
		}
		return ptr
	case ctx.ProcedureType() != nil:
		proc := v.Visit(ctx.ProcedureType())
		if proc == nil {
			panic("procedureType returned nil")
		}
		return proc
	default:
		panic(fmt.Sprintf("unknown type: %s", ctx.GetText()))
	}
}

func (v *ASTBuilder) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	a := &ast.ArrayType{}
	lengths := ctx.AllLength()
	if len(lengths) == 0 {
		panic("array type has no lengths")
	}
	for _, l := range lengths {
		lengthExpr := v.Visit(l)
		if lengthExpr == nil {
			panic("length expression is nil")
		}
		expr, ok := lengthExpr.(ast.Expr)
		if !ok {
			panic(fmt.Sprintf("length expression not Expr: %T", lengthExpr))
		}
		a.Lengths = append(a.Lengths, expr)
	}
	elemType := v.Visit(ctx.Type_())
	if elemType == nil {
		panic("array element type is nil")
	}
	typeExpr, ok := elemType.(ast.TypeExpr)
	if !ok {
		panic(fmt.Sprintf("array element type not TypeExpr: %T", elemType))
	}
	a.Elem = typeExpr
	return a
}

func (v *ASTBuilder) VisitLength(ctx *LengthContext) interface{} {
	expr := v.Visit(ctx.ConstExpression())
	if expr == nil {
		panic("length constant expression is nil")
	}
	return expr
}

func (v *ASTBuilder) VisitRecordType(ctx *RecordTypeContext) interface{} {
	r := &ast.RecordType{}
	if base := ctx.BaseType(); base != nil {
		if q := base.Qualident(); q != nil {
			r.Base = v.qualidentText(q)
		}
	}
	if fls := ctx.FieldListSequence(); fls != nil {
		for _, f := range fls.AllFieldList() {
			fd := v.Visit(f).(*ast.FieldDecl)
			if fd != nil && fd.Type != nil {
				r.Fields = append(r.Fields, fd)
				for _, nameDef := range fd.Names {
					r.FieldOrder = append(r.FieldOrder, nameDef.Name)
				}
			}
		}
	}
	return r
}

func (v *ASTBuilder) VisitFieldList(ctx *FieldListContext) interface{} {
	namesRaw := v.Visit(ctx.IdentList())
	if namesRaw == nil {
		panic("field list: identList is nil")
	}
	names, ok := namesRaw.([]ast.IdentDef)
	if !ok {
		panic(fmt.Sprintf("field list: identList not []IdentDef: %T", namesRaw))
	}
	if len(names) == 0 {
		return &ast.FieldDecl{Names: []ast.IdentDef{}, Type: nil}
	}
	typRaw := v.Visit(ctx.Type_())
	if typRaw == nil {
		fmt.Println("Warning: field list has no type, skipping")
		return &ast.FieldDecl{Names: names, Type: nil}
	}
	typ, ok := typRaw.(ast.TypeExpr)
	if !ok {
		panic(fmt.Sprintf("field list: type not TypeExpr: %T", typRaw))
	}
	return &ast.FieldDecl{Names: names, Type: typ}
}

func (v *ASTBuilder) VisitPointerType(ctx *PointerTypeContext) interface{} {
	return &ast.PointerType{Target: v.Visit(ctx.Type_()).(ast.TypeExpr)}
}

func (v *ASTBuilder) VisitProcedureType(ctx *ProcedureTypeContext) interface{} {
	proc := &ast.ProcedureType{}
	if sig := ctx.FormalParameters(); sig != nil {
		proc.Signature = v.Visit(sig).(*ast.ProcedureSignature)
	}
	return proc
}

func (v *ASTBuilder) VisitIdentList(ctx *IdentListContext) interface{} {
	list := make([]ast.IdentDef, 0, len(ctx.AllIdentdef()))
	for _, i := range ctx.AllIdentdef() {
		if i == nil {
			continue
		}
		val := v.Visit(i)
		if val == nil {
			continue
		}
		if id, ok := val.(ast.IdentDef); ok {
			list = append(list, id)
		}
	}
	return list
}

func (v *ASTBuilder) VisitIdentdef(ctx *IdentdefContext) interface{} {
	if ident := ctx.Ident(); ident != nil {
		return ast.IdentDef{
			Name:     ident.GetText(),
			Exported: ctx.GetChildCount() > 1,
		}
	}
	return ast.IdentDef{Name: "", Exported: false}
}

func (v *ASTBuilder) VisitStatementSequence(ctx *StatementSequenceContext) interface{} {
	allStmts := ctx.AllStatement()
	ss := make([]ast.Stmt, 0, len(allStmts))
	for _, stmtCtx := range allStmts {
		if stmt := v.Visit(stmtCtx); stmt != nil {
			ss = append(ss, stmt.(ast.Stmt))
		}
	}
	return ss
}

func (v *ASTBuilder) VisitStatement(ctx *StatementContext) interface{} {
	switch {
	case ctx.Assignment() != nil:
		return v.Visit(ctx.Assignment())
	case ctx.ProcedureCall() != nil:
		return v.Visit(ctx.ProcedureCall())
	case ctx.IfStatement() != nil:
		return v.Visit(ctx.IfStatement())
	case ctx.CaseStatement() != nil:
		return v.Visit(ctx.CaseStatement())
	case ctx.WhileStatement() != nil:
		return v.Visit(ctx.WhileStatement())
	case ctx.RepeatStatement() != nil:
		return v.Visit(ctx.RepeatStatement())
	case ctx.ForStatement() != nil:
		return v.Visit(ctx.ForStatement())
	default:
		return nil
	}
}

func (v *ASTBuilder) VisitAssignment(ctx *AssignmentContext) interface{} {
	targetRaw := v.Visit(ctx.Designator())
	if targetRaw == nil {
		panic("assignment target is nil")
	}
	target, ok := targetRaw.(*ast.DesignatorExpr)
	if !ok {
		panic(fmt.Sprintf("assignment target is not DesignatorExpr, got %T", targetRaw))
	}
	exprCtx := ctx.Expression()
	if exprCtx == nil {
		panic(fmt.Sprintf("Expression context is nil for assignment target %v", target))
	}
	valueRaw := v.Visit(exprCtx)
	if valueRaw == nil {
		panic(fmt.Sprintf("Visit returned nil for expression: %s", exprCtx.GetText()))
	}
	value, ok := valueRaw.(ast.Expr)
	if !ok {
		panic(fmt.Sprintf("assignment value is not Expr, got %T", valueRaw))
	}
	return &ast.AssignmentStmt{Target: target, Value: value}
}

func (v *ASTBuilder) VisitExpression(ctx *ExpressionContext) interface{} {
	all := ctx.AllSimpleExpression()
	if len(all) == 0 {
		panic("expression has no simple expression")
	}
	leftVal := v.Visit(all[0])
	if leftVal == nil {
		panic("left simple expression returned nil")
	}
	left, ok := leftVal.(ast.Expr)
	if !ok {
		panic(fmt.Sprintf("left simple expression is not Expr: %T", leftVal))
	}
	if ctx.Relation() == nil && len(all) < 2 {
		return left
	}
	if len(all) < 2 {
		panic("relation present but no right simple expression")
	}
	rightVal := v.Visit(all[1])
	if rightVal == nil {
		panic("right simple expression returned nil")
	}
	right, ok := rightVal.(ast.Expr)
	if !ok {
		panic(fmt.Sprintf("right simple expression is not Expr: %T", rightVal))
	}
	rel := ctx.Relation().GetText()
	return &ast.BinaryExpr{Left: left, Op: rel, Right: right}
}

func (v *ASTBuilder) VisitProcedureCall(ctx *ProcedureCallContext) interface{} {
	call := &ast.CallExpr{}
	call.Callee = v.Visit(ctx.Designator()).(ast.Expr)
	if params := ctx.ActualParameters(); params != nil {
		call.Args = v.Visit(params).([]ast.Expr)
	}
	return &ast.ProcedureCallStmt{Call: call}
}

func (v *ASTBuilder) VisitIfStatement(ctx *IfStatementContext) interface{} {
	exprs := ctx.AllExpression()
	seqs := ctx.AllStatementSequence()
	stmt := &ast.IfStmt{}
	for i := 0; i < len(exprs) && i < len(seqs); i++ {
		stmt.Branches = append(stmt.Branches, &ast.IfBranch{
			Cond: v.Visit(exprs[i]).(ast.Expr),
			Body: v.Visit(seqs[i]).([]ast.Stmt),
		})
	}
	if len(seqs) > len(exprs) {
		stmt.ElseBody = v.Visit(seqs[len(seqs)-1]).([]ast.Stmt)
	}
	return stmt
}

func (v *ASTBuilder) VisitCaseStatement(ctx *CaseStatementContext) interface{} {
	c := &ast.CaseStmt{Expr: v.Visit(ctx.Expression()).(ast.Expr)}
	for _, cc := range ctx.AllCase_() {
		branch := v.Visit(cc).(*ast.CaseBranch)
		c.Branches = append(c.Branches, branch)
	}
	return c
}

func (v *ASTBuilder) VisitCase_(ctx *Case_Context) interface{} {
	c := &ast.CaseBranch{}
	if ll := ctx.CaseLabelList(); ll != nil {
		c.Labels = v.Visit(ll).([]*ast.CaseLabel)
	}
	if ss := ctx.StatementSequence(); ss != nil {
		c.Body = v.Visit(ss).([]ast.Stmt)
	}
	return c
}

func (v *ASTBuilder) VisitCaseLabelList(ctx *CaseLabelListContext) interface{} {
	out := make([]*ast.CaseLabel, 0, len(ctx.AllLabelRange()))
	for _, lr := range ctx.AllLabelRange() {
		out = append(out, v.Visit(lr).(*ast.CaseLabel))
	}
	return out
}

func (v *ASTBuilder) VisitLabelRange(ctx *LabelRangeContext) interface{} {
	lr := &ast.CaseLabel{}
	all := ctx.AllLabel()
	if len(all) == 0 {
		return lr
	}
	lr.From = v.Visit(all[0]).(ast.Expr)
	lr.To = lr.From
	if len(all) > 1 {
		lr.To = v.Visit(all[1]).(ast.Expr)
	}
	return lr
}

func (v *ASTBuilder) VisitLabel(ctx *LabelContext) interface{} {
	if tok := ctx.INTEGER(); tok != nil {
		return &ast.NumberExpr{Text: tok.GetText(), IsReal: false}
	}
	return &ast.NumberExpr{Text: ctx.GetText()}
}

func (v *ASTBuilder) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	w := &ast.WhileStmt{}
	exprs := ctx.AllExpression()
	seqs := ctx.AllStatementSequence()
	for i := 0; i < len(exprs) && i < len(seqs); i++ {
		w.Branches = append(w.Branches, &ast.IfBranch{
			Cond: v.Visit(exprs[i]).(ast.Expr),
			Body: v.Visit(seqs[i]).([]ast.Stmt),
		})
	}
	return w
}

func (v *ASTBuilder) VisitRepeatStatement(ctx *RepeatStatementContext) interface{} {
	return &ast.RepeatStmt{
		Body:  v.Visit(ctx.StatementSequence()).([]ast.Stmt),
		Until: v.Visit(ctx.Expression()).(ast.Expr),
	}
}

func (v *ASTBuilder) VisitForStatement(ctx *ForStatementContext) interface{} {
	f := &ast.ForStmt{}
	f.Var = ctx.Ident().GetText()
	exprs := ctx.AllExpression()
	if len(exprs) < 2 {
		panic("for statement missing from or to expression")
	}
	f.From = v.Visit(exprs[0]).(ast.Expr)
	f.To = v.Visit(exprs[1]).(ast.Expr)
	f.Body = v.Visit(ctx.StatementSequence()).([]ast.Stmt)
	if by := ctx.ConstExpression(); by != nil {
		f.By = v.Visit(by).(ast.Expr)
		f.HasBy = true
	}
	return f
}

func (v *ASTBuilder) VisitActualParameters(ctx *ActualParametersContext) interface{} {
	if ctx.ExpList() != nil {
		return v.Visit(ctx.ExpList())
	}
	return []ast.Expr{}
}

func (v *ASTBuilder) VisitExpList(ctx *ExpListContext) interface{} {
	list := make([]ast.Expr, 0, len(ctx.AllExpression()))
	for _, e := range ctx.AllExpression() {
		expr := v.Visit(e)
		if expr == nil {
			panic("expression in expList is nil")
		}
		list = append(list, expr.(ast.Expr))
	}
	return list
}

func (v *ASTBuilder) VisitSimpleExpression(ctx *SimpleExpressionContext) interface{} {
	allterm := ctx.AllTerm()
	if len(allterm) == 0 {
		for i := 0; i < ctx.GetChildCount(); i++ {
			child := ctx.GetChild(i)
			switch c := child.(type) {
			case *NumberContext:
				return &ast.NumberExpr{Text: c.GetText()}
			case *DesignatorContext:
				return v.Visit(c)
			case *FactorContext:
				return v.Visit(c)
			case *antlr.ErrorNodeImpl:
				text := c.GetText()
				if text != "" {
					return &ast.NumberExpr{Text: text}
				}
			case antlr.TerminalNode:
				return &ast.NumberExpr{Text: c.GetText()}
			}
		}
		return nil
	}
	expr := v.Visit(allterm[0]).(ast.Expr)
	if fc, ok := allterm[0].GetChild(0).(antlr.TerminalNode); ok {
		if op := fc.GetText(); op == "+" || op == "-" {
			expr = &ast.UnaryExpr{Op: op, Expr: expr}
		}
	}
	for i, op := range ctx.AllAddOperator() {
		if i+1 >= len(allterm) {
			break
		}
		opText := op.GetText()
		right := v.Visit(allterm[i+1]).(ast.Expr)
		expr = &ast.BinaryExpr{Left: expr, Op: opText, Right: right}
	}
	return expr
}

func (v *ASTBuilder) VisitTerm(ctx *TermContext) interface{} {
	af := ctx.AllFactor()
	if len(af) == 0 {
		return nil
	}
	expr := v.Visit(af[0]).(ast.Expr)
	for i, op := range ctx.AllMulOperator() {
		if i+1 >= len(af) {
			break
		}
		expr = &ast.BinaryExpr{Left: expr, Op: op.GetText(), Right: v.Visit(af[i+1]).(ast.Expr)}
	}
	return expr
}

func (v *ASTBuilder) VisitFactor(ctx *FactorContext) interface{} {
	switch {
	case ctx.Number() != nil:
		numCtx := ctx.Number()
		if numCtx.INTEGER() != nil {
			return &ast.NumberExpr{Text: numCtx.GetText(), IsReal: false}
		} else if numCtx.REAL() != nil {
			return &ast.NumberExpr{Text: numCtx.GetText(), IsReal: true}
		} else {
			return &ast.NumberExpr{Text: numCtx.GetText(), IsReal: false}
		}
	case ctx.STRING() != nil:
		return &ast.StringExpr{Value: ctx.STRING().GetText()}
	case ctx.NIL() != nil:
		return &ast.NilExpr{}
	case ctx.TRUE() != nil:
		return &ast.BoolExpr{Value: true}
	case ctx.FALSE() != nil:
		return &ast.BoolExpr{Value: false}
	case ctx.Set_() != nil:
		return v.Visit(ctx.Set_())
	case ctx.Designator() != nil:
		baseRaw := v.Visit(ctx.Designator())
		if baseRaw == nil {
			panic("designator in factor returned nil")
		}
		base, ok := baseRaw.(ast.Expr)
		if !ok {
			panic(fmt.Sprintf("designator did not return Expr: %T", baseRaw))
		}
		if ap := ctx.ActualParameters(); ap != nil {
			argsRaw := v.Visit(ap)
			args, ok := argsRaw.([]ast.Expr)
			if !ok {
				panic(fmt.Sprintf("actualParameters did not return []Expr: %T", argsRaw))
			}
			return &ast.CallExpr{Callee: base, Args: args}
		}
		return base
	case ctx.Expression() != nil:
		return v.Visit(ctx.Expression())
	case ctx.Factor() != nil:
		return &ast.UnaryExpr{Op: "~", Expr: v.Visit(ctx.Factor()).(ast.Expr)}
	default:
		panic("unknown factor")
	}
}

func (v *ASTBuilder) VisitSet_(ctx *Set_Context) interface{} {
	elems := &ast.SetExpr{}
	for _, elem := range ctx.AllElement() {
		elems.Elements = append(elems.Elements, v.Visit(elem).(*ast.SetElement))
	}
	return elems
}

func (v *ASTBuilder) VisitElement(ctx *ElementContext) interface{} {
	exprs := ctx.AllExpression()
	if len(exprs) == 0 {
		return nil
	}
	from := v.Visit(exprs[0]).(ast.Expr)
	to := from
	if len(exprs) > 1 {
		to = v.Visit(exprs[1]).(ast.Expr)
	}
	return &ast.SetElement{From: from, To: to}
}

func (v *ASTBuilder) VisitDesignator(ctx *DesignatorContext) interface{} {
	q := ctx.Qualident()
	if q == nil {
		panic("designator without qualident")
	}
	base := v.qualident(q)
	des := &ast.DesignatorExpr{Base: base}
	for _, selCtx := range ctx.AllSelector() {
		sel := v.Visit(selCtx).(ast.Selector)
		des.Selectors = append(des.Selectors, sel)
	}
	return des
}

func (v *ASTBuilder) VisitSelector(ctx *SelectorContext) interface{} {
	switch {
	case ctx.Ident() != nil:
		return ast.Selector{Field: ctx.Ident().GetText()}
	case ctx.ExpList() != nil:
		exprsRaw := v.Visit(ctx.ExpList())
		if exprsRaw == nil {
			return ast.Selector{Index: []ast.Expr{}}
		}
		exprs, ok := exprsRaw.([]ast.Expr)
		if !ok {
			panic(fmt.Sprintf("expList returned %T, expected []ast.Expr", exprsRaw))
		}
		return ast.Selector{Index: exprs}
	// case ctx.Qualident() != nil: // Удалено, так как этой альтернативы больше нет в грамматике
	//     return ast.Selector{Type: v.qualidentText(ctx.Qualident())}
	default:
		return ast.Selector{Deref: true}
	}
}

func (v *ASTBuilder) qualident(ctx IQualidentContext) ast.QualIdent {
	ids := ctx.AllIdent()
	if len(ids) == 0 {
		panic("qualident without ident")
	}
	if len(ids) == 1 {
		return ast.QualIdent{Name: ids[0].GetText()}
	}
	return ast.QualIdent{Module: ids[0].GetText(), Name: ids[1].GetText()}
}

func (v *ASTBuilder) qualidentText(ctx IQualidentContext) string {
	q := v.qualident(ctx)
	if q.Module == "" {
		return q.Name
	}
	return q.Module + "." + q.Name
}
