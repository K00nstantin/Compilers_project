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
