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

/////////////////////////////////////////////////////////////////////////

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
