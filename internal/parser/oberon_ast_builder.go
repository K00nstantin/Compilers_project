package parser

import (
	"github.com/K00nstantin/Compilers_project/internal/ast"
	"github.com/antlr4-go/antlr/v4"
)

type ASTBuilder struct{ *BaseoberonVisitor }

func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{BaseoberonVisitor: &BaseoberonVisitor{BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{}}}
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

	return &ast.ConstDecl{
		Name:  name,
		Value: value,
	}
}

func (v *ASTBuilder) VisitConstExpression(ctx *ConstExpressionContext) interface{} {
	return v.Visit(ctx.Expression())
}

func (v *ASTBuilder) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	name := v.Visit(ctx.Identdef()).(ast.IdentDef)
	tp := v.Visit(ctx.Type_()).(ast.TypeExpr)
	return &ast.TypeDecl{
		Name: name,
		Type: tp,
	}
}

func (v *ASTBuilder) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	names := v.Visit(ctx.IdentList()).([]ast.IdentDef)
	tp := v.Visit(ctx.Type_()).(ast.TypeExpr)
	return &ast.VarDecl{
		Names: names,
		Type:  tp,
	}
}

func (v *ASTBuilder) VisitProcedureDeclaration(ctx *ProcedureDeclarationContext) interface{} {
	head := v.Visit(ctx.ProcedureHeading()).(*ast.ProcedureDecl)
	body := v.Visit(ctx.ProcedureBody()).(*ast.ProcedureDecl)
	if body != nil {
		head.Declarations = body.Declarations
		head.Body = body.Body
		head.ReturnExpr = body.ReturnExpr
	}
	if endName := ctx.Ident(); endName != nil {
		head.EndName = endName.GetText()
	}
	return head
}

func (v *ASTBuilder) VisitProcedureHeading(ctx *ProcedureHeadingContext) interface{} {
	proc := &ast.ProcedureDecl{}
	proc.Name = v.Visit(ctx.Identdef()).(ast.IdentDef)
	fp, ok := v.Visit(ctx.FormalParameters()).(*ast.ProcedureSignature)
	if ok && fp != nil {
		proc.Signature = fp
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

func (v *ASTBuilder) VisitFPSection(ctx *FPSectionContext) interface{} {
	section := &ast.ParamSection{ByRef: ctx.VAR() != nil}
	for _, id := range ctx.AllIdent() {
		section.Names = append(section.Names, id.GetText())
	}
	if ft := ctx.FormalType(); ft != nil {
		section.Type = ft.GetText()
	}
	return section
}

func (v *ASTBuilder) VisitType_(ctx *Type_Context) interface{} {
	switch {
	case ctx.Qualident() != nil:
		return &ast.NamedType{Name: v.qualidentText(ctx.Qualident())}
	case ctx.ArrayType() != nil:
		return v.Visit(ctx.ArrayType())
	case ctx.RecordType() != nil:
		return v.Visit(ctx.RecordType())
	case ctx.PointerType() != nil:
		return v.Visit(ctx.PointerType())
	case ctx.ProcedureType() != nil:
		return v.Visit(ctx.ProcedureType())
	default:
		panic("type")
	}
}

func (v *ASTBuilder) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	a := &ast.ArrayType{}
	for _, l := range ctx.AllLength() {
		a.Lengths = append(a.Lengths, v.Visit(l).(ast.Expr))
	}
	a.Elem = v.Visit(ctx.Type_()).(ast.TypeExpr)
	return a
}

func (v *ASTBuilder) VisitLength(ctx *LengthContext) interface{} {
	return v.Visit(ctx.ConstExpression())
}

func (v *ASTBuilder) VisitRecordType(ctx *RecordTypeContext) interface{} {
	r := &ast.RecordType{}
	base := ctx.BaseType()
	if base != nil && base.Qualident() != nil {
		r.Base = v.qualidentText(base.Qualident())
	}
	fls := ctx.FieldListSequence()
	if fls != nil {
		for _, f := range fls.AllFieldList() {
			r.Fields = append(r.Fields, v.Visit(f).(*ast.FieldDecl))
		}
	}
	return r
}

func (v *ASTBuilder) VisitFieldList(ctx *FieldListContext) interface{} {
	return &ast.FieldDecl{Names: v.Visit(ctx.IdentList()).([]ast.IdentDef), Type: v.Visit(ctx.Type_()).(ast.TypeExpr)}
}

func (v *ASTBuilder) VisitPointerType(ctx *PointerTypeContext) interface{} {
	return &ast.PointerType{Target: v.Visit(ctx.Type_()).(ast.TypeExpr)}
}

func (v *ASTBuilder) VisitProcedureType(ctx *ProcedureTypeContext) interface{} {
	proc := &ast.ProcedureType{}
	sig := ctx.FormalParameters()
	if sig != nil {
		proc.Signature = v.Visit(sig).(*ast.ProcedureSignature)
	}

	return proc
}

func (v *ASTBuilder) VisitIdentList(ctx *IdentListContext) interface{} {
	list := make([]ast.IdentDef, 0, len(ctx.AllIdentdef()))
	for _, i := range ctx.AllIdentdef() {
		list = append(list, v.Visit(i).(ast.IdentDef))
	}
	return list
}

func (v *ASTBuilder) VisitIdentDef(ctx *IdentdefContext) interface{} {
	return &ast.IdentDef{Name: ctx.Ident().GetText(), Exported: ctx.GetChildCount() > 1}
}

func (v *ASTBuilder) VisitStatementSequence(ctx *StatementSequenceContext) interface{} {
	ss := make([]ast.Stmt, 0, len(ctx.AllStatement()))
	for _, s := range ctx.AllStatement() {
		ts := v.Visit(s)
		if ts != nil {
			ss = append(ss, ts.(ast.Stmt))
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
	return &ast.AssignmentStmt{Target: v.Visit(ctx.Designator()).(*ast.DesignatorExpr), Value: v.Visit(ctx.Expression()).(ast.Expr)}
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
		stmt.Branches = append(stmt.Branches, &ast.IfBranch{Cond: v.Visit(exprs[i]).(ast.Expr), Body: v.Visit(seqs[i]).([]ast.Stmt)})
	}
	if len(seqs) > len(exprs) {
		stmt.ElseBody = v.Visit(seqs[len(seqs)-1]).([]ast.Stmt)
	}
	return stmt
}

func (v *ASTBuilder) VisitCaseStatement(ctx *CaseStatementContext) interface{} {
	c := &ast.CaseStmt{Expr: v.Visit(ctx.Expression()).(ast.Expr)}
	for _, cc := range ctx.AllCase_() {
		c.Branches = append(c.Branches, v.Visit(cc).(*ast.CaseBranch))
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
		return nil
	}
	lr.From = v.Visit(all[0]).(ast.Expr)
	lr.To = lr.From
	if len(all) > 1 {
		lr.To = v.Visit(all[1]).(ast.Expr)
	}
	return lr
}

func (v *ASTBuilder) VisitLabel(ctx *LabelContext) interface{} {
	return &ast.NumberExpr{Text: ctx.GetText()}
}

func (v *ASTBuilder) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	w := &ast.WhileStmt{}
	exprs := ctx.AllExpression()
	seqs := ctx.AllStatementSequence()
	for i := 0; i < len(exprs) && i < len(seqs); i++ {
		w.Branches = append(w.Branches, &ast.IfBranch{Cond: v.Visit(exprs[i]).(ast.Expr), Body: v.Visit(seqs[i]).([]ast.Stmt)})
	}
	return w
}

func (v *ASTBuilder) VisitRepeatStatement(ctx *RepeatStatementContext) interface{} {
	return &ast.RepeatStmt{Body: v.Visit(ctx.StatementSequence()).([]ast.Stmt), Until: v.Visit(ctx.Expression()).(ast.Expr)}
}

func (v *ASTBuilder) VisitForStatement(ctx *ForStatementContext) interface{} {
	f := &ast.ForStmt{}
	f.Var = ctx.Ident().GetText()
	exprs := ctx.AllExpression()
	f.From = v.Visit(exprs[0]).(ast.Expr)
	f.To = v.Visit(exprs[1]).(ast.Expr)
	f.Body = v.Visit(ctx.StatementSequence()).([]ast.Stmt)
	if by := ctx.ConstExpression(); by != nil {
		f.By = v.Visit(ctx.ConstExpression()).(ast.Expr)
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
		list = append(list, v.Visit(e).(ast.Expr))
	}
	return list
}

func (v *ASTBuilder) VisitExpression(ctx *ExpressionContext) interface{} {
	all := ctx.AllSimpleExpression()
	if len(all) == 0 {
		return nil
	}
	left := v.Visit(all[0]).(ast.Expr)
	if ctx.Relation() == nil && len(all) < 2 {
		return left
	}
	right := v.Visit(all[1]).(ast.Expr)
	rel := ctx.Relation().GetText()
	return &ast.BinaryExpr{Left: left, Op: rel, Right: right}
}

func (v *ASTBuilder) VisitSimpleExpression(ctx *SimpleExpressionContext) interface{} {
	allterm := ctx.AllTerm()
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
		op := op.GetText()
		right := v.Visit(allterm[i+1]).(ast.Expr)
		expr = &ast.BinaryExpr{Left: expr, Op: op, Right: right}
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
		return &ast.NumberExpr{Text: ctx.Number().GetText()}
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
		base := v.Visit(ctx.Designator()).(ast.Expr)
		if ap := ctx.ActualParameters(); ap != nil {
			return &ast.CallExpr{Callee: base, Args: v.Visit(ap).([]ast.Expr)}
		}
		return base
	case ctx.Expression() != nil:
		return v.Visit(ctx.Expression())
	case ctx.Factor() != nil:
		return &ast.UnaryExpr{Op: "~", Expr: v.Visit(ctx.Factor()).(ast.Expr)}
	default:
		panic("factor")
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
	des := &ast.DesignatorExpr{Base: v.qualident(ctx.Qualident())}
	for _, s := range ctx.AllSelector() {
		des.Selectors = append(des.Selectors, v.Visit(s).(ast.Selector))
	}
	return des
}

func (v *ASTBuilder) VisitSelector(ctx *SelectorContext) interface{} {
	switch {
	case ctx.Ident() != nil:
		return ast.Selector{Field: ctx.Ident().GetText()}
	case ctx.ExpList() != nil:
		return ast.Selector{Index: v.Visit(ctx.ExpList()).([]ast.Expr)}
	case ctx.Qualident() != nil:
		return ast.Selector{Type: v.qualidentText(ctx.Qualident())}
	default:
		return ast.Selector{Deref: true}
	}
}

func (v *ASTBuilder) qualident(ctx IQualidentContext) ast.QualIdent {
	ids := ctx.AllIdent()
	if len(ids) == 1 {
		return ast.QualIdent{Name: ids[0].GetText()}
	}
	if len(ids) >= 2 {
		return ast.QualIdent{Module: ids[0].GetText(), Name: ids[1].GetText()}
	}
	return ast.QualIdent{}
}

func (v *ASTBuilder) qualidentText(ctx IQualidentContext) string {
	q := v.qualident(ctx)
	if q.Module == "" {
		return q.Name
	}
	return q.Module + "." + q.Name
}
