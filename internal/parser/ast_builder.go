package parser

import (
	"strconv"

	"github.com/K00nstantin/Compilers_project/internal/ast"
	"github.com/antlr4-go/antlr/v4"
)

type ASTBuilder struct {
	*BaseMiniJavaVisitor
}

func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{
		BaseMiniJavaVisitor: &BaseMiniJavaVisitor{BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{}},
	}
}

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

func (v *ASTBuilder) VisitGoal(ctx *GoalContext) interface{} {
	mainClassCtx := ctx.MainClassDeclaration()
	mainClass := v.Visit(mainClassCtx).(*ast.MainClass)
	classes := []*ast.ClassDecl{}
	for _, classCtx := range ctx.AllClassDeclaration() {
		class := v.Visit(classCtx).(*ast.ClassDecl)
		classes = append(classes, class)
	}
	return &ast.Program{
		MainClass: mainClass,
		Classes:   classes,
	}
}

func (v *ASTBuilder) VisitMainClassDeclaration(ctx *MainClassDeclarationContext) interface{} {
	name := ctx.ID().GetText()
	mainMethodCtx := ctx.MainClassBody().MainMethod()
	mainMethod := v.Visit(mainMethodCtx).(*ast.MainMethod)
	return &ast.MainClass{
		Name:       name,
		MainMethod: mainMethod,
	}
}

func (v *ASTBuilder) VisitMainMethod(ctx *MainMethodContext) interface{} {
	stmtCtx := ctx.Statement()
	stmt := v.Visit(stmtCtx).(ast.Stmt)

	return &ast.MainMethod{
		Body: []ast.Stmt{stmt},
	}
}

func (v *ASTBuilder) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	name := ctx.ID().GetText()
	parent := ""
	if ctx.Type_() != nil {
		parent = ctx.Type_().GetText()
	}
	fields := []*ast.VarDecl{}
	methods := []*ast.MethodDecl{}

	for _, field := range ctx.ClassBody().AllFieldDeclaration() {
		f := v.Visit(field).(*ast.VarDecl)
		fields = append(fields, f)
	}

	for _, method := range ctx.ClassBody().AllMethodDeclaration() {
		m := v.Visit(method).(*ast.MethodDecl)
		methods = append(methods, m)
	}
	return &ast.ClassDecl{
		Name:    name,
		Parent:  parent,
		Fields:  fields,
		Methods: methods,
	}
}

func (v *ASTBuilder) VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{} {
	tp := ctx.Type_().GetText()
	name := ctx.ID().GetText()
	return &ast.VarDecl{
		Type: tp,
		Name: name,
	}
}

func (v *ASTBuilder) VisitMethodDeclaration(ctx *MethodDeclarationContext) interface{} {
	pub := ctx.GetToken(MiniJavaParserT__4, 0) != nil
	rettyppe := ""
	if ctx.Type_() != nil {
		rettyppe = ctx.Type_().GetText()
	}
	name := ctx.ID().GetText()
	params := []*ast.VarDecl{}
	if ctx.FormalParameters() != nil {
		if plist := ctx.FormalParameters().FormalParameterList(); plist != nil {
			for _, param := range plist.AllFormalParameter() {
				p := v.Visit(param).(*ast.VarDecl)
				params = append(params, p)
			}
		}
	}

	vars := []*ast.VarDecl{}
	for _, vr := range ctx.MethodBody().AllVarDeclaration() {
		vs := v.Visit(vr).(*ast.VarDecl)
		vars = append(vars, vs)
	}

	body := []ast.Stmt{}
	for _, stmt := range ctx.MethodBody().AllStatement() {
		s := v.Visit(stmt).(ast.Stmt)
		body = append(body, s)
	}

	return &ast.MethodDecl{
		Public:     pub,
		ReturnType: rettyppe,
		Name:       name,
		Params:     params,
		Vars:       vars,
		Body:       body,
	}
}

func (v *ASTBuilder) VisitFormalParameter(ctx *FormalParameterContext) interface{} {
	tp := ctx.Type_().GetText()
	name := ctx.ID().GetText()
	return &ast.VarDecl{
		Type: tp,
		Name: name,
	}
}

func (v *ASTBuilder) VisitVarDeclaration(ctx *VarDeclarationContext) interface{} {
	tp := ctx.Type_().GetText()
	name := ctx.ID().GetText()
	return &ast.VarDecl{
		Type: tp,
		Name: name,
	}
}

func (v *ASTBuilder) VisitNestedStatement(ctx *NestedStatementContext) interface{} {
	stmts := []ast.Stmt{}
	for _, stmt := range ctx.AllStatement() {
		s := v.Visit(stmt).(ast.Stmt)
		stmts = append(stmts, s)
	}
	return &ast.BlockStmt{
		Stmts: stmts,
	}
}

func (v *ASTBuilder) VisitIfElseStatement(ctx *IfElseStatementContext) interface{} {
	cond := v.Visit(ctx.Expression()).(ast.Expression)
	then := v.Visit(ctx.Statement(0)).(ast.Stmt)
	else_ := v.Visit(ctx.Statement(1)).(ast.Stmt)
	return &ast.IfStmt{
		Cond: cond,
		Then: then,
		Else: else_,
	}
}

func (v *ASTBuilder) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	cond := v.Visit(ctx.Expression()).(ast.Expression)
	body := v.Visit(ctx.Statement()).(ast.Stmt)
	return &ast.WhileStmt{
		Cond: cond,
		Body: body,
	}
}

func (v *ASTBuilder) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	expr := v.Visit(ctx.Expression()).(ast.Expression)
	return &ast.PrintStmt{
		Expr: expr,
	}
}

func (v *ASTBuilder) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	name := ctx.ID().GetText()
	expr := v.Visit(ctx.Expression()).(ast.Expression)
	return &ast.AssignStmt{
		Name: name,
		Expr: expr,
	}
}

func (v *ASTBuilder) VisitArrayAssignStatement(ctx *ArrayAssignStatementContext) interface{} {
	name := ctx.ID().GetText()
	ind := v.Visit(ctx.Expression(0)).(ast.Expression)
	expr := v.Visit(ctx.Expression(1)).(ast.Expression)
	return &ast.ArrayAssignStmt{
		Name:  name,
		Index: ind,
		Expr:  expr,
	}
}

func (v *ASTBuilder) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	expr := v.Visit(ctx.Expression()).(ast.Expression)
	return &ast.ReturnStmt{
		Expr: expr,
	}
}

func (v *ASTBuilder) VisitRecurStatement(ctx *RecurStatementContext) interface{} {
	cond := v.Visit(ctx.Expression(0)).(ast.Expression)
	args := []ast.Expression{}
	if mal := ctx.MethodArgumentList(); mal != nil {
		for _, arg := range mal.AllExpression() {
			a := v.Visit(arg).(ast.Expression)
			args = append(args, a)
		}
	}

	elseexpr := v.Visit(ctx.Expression(1)).(ast.Expression)
	return &ast.RecurStmt{
		Cond: cond,
		Args: args,
		Else: elseexpr,
	}
}

func (v *ASTBuilder) VisitArrayAccessExpression(ctx *ArrayAccessExpressionContext) interface{} {
	arr := v.Visit(ctx.Expression(0)).(ast.Expression)
	ind := v.Visit(ctx.Expression(1)).(ast.Expression)
	return &ast.ArrayAccessExpr{
		Array: arr,
		Index: ind,
	}
}

func (v *ASTBuilder) VisitArrayLengthExpression(ctx *ArrayLengthExpressionContext) interface{} {
	array := v.Visit(ctx.Expression()).(ast.Expression)
	return &ast.ArrayLengthExpr{Array: array}
}

func (v *ASTBuilder) VisitMethodCallExpression(ctx *MethodCallExpressionContext) interface{} {
	object := v.Visit(ctx.Expression()).(ast.Expression)
	method := ctx.ID().GetText()
	args := []ast.Expression{}
	for _, arg := range ctx.MethodArgumentList().AllExpression() {
		a := v.Visit(arg).(ast.Expression)
		args = append(args, a)
	}

	return &ast.MethodCallExpr{
		Object: object,
		Method: method,
		Args:   args,
	}
}

func (v *ASTBuilder) VisitNegExpression(ctx *NegExpressionContext) interface{} {
	expr := v.Visit(ctx.Expression()).(ast.Expression)
	return &ast.NegExpr{
		Expr: expr,
	}
}

func (v *ASTBuilder) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	expr := v.Visit(ctx.Expression()).(ast.Expression)
	return &ast.NotExpr{
		Expr: expr,
	}
}

func (v *ASTBuilder) VisitArrayInstantiationExpression(ctx *ArrayInstantiationExpressionContext) interface{} {
	size := v.Visit(ctx.Expression()).(ast.Expression)
	return &ast.ArrayInstantiationExpr{
		Size: size,
	}
}

func (v *ASTBuilder) VisitObjectInstantiationExpression(ctx *ObjectInstantiationExpressionContext) interface{} {
	className := ctx.ID().GetText()
	return &ast.ObjectInstantiationExpr{
		ClassName: className,
	}
}

func (v *ASTBuilder) VisitAddExpression(ctx *AddExpressionContext) interface{} {
	left := v.Visit(ctx.Expression(0)).(ast.Expression)
	right := v.Visit(ctx.Expression(1)).(ast.Expression)
	return &ast.BinaryOpExpr{
		Left:  left,
		Right: right,
		Op:    "+",
	}
}

func (v *ASTBuilder) VisitSubExpression(ctx *SubExpressionContext) interface{} {
	left := v.Visit(ctx.Expression(0)).(ast.Expression)
	right := v.Visit(ctx.Expression(1)).(ast.Expression)
	return &ast.BinaryOpExpr{
		Left:  left,
		Right: right,
		Op:    "-",
	}
}

func (v *ASTBuilder) VisitMulExpression(ctx *MulExpressionContext) interface{} {
	left := v.Visit(ctx.Expression(0)).(ast.Expression)
	right := v.Visit(ctx.Expression(1)).(ast.Expression)
	return &ast.BinaryOpExpr{
		Left:  left,
		Right: right,
		Op:    "*",
	}
}

func (v *ASTBuilder) VisitLtExpression(ctx *LtExpressionContext) interface{} {
	left := v.Visit(ctx.Expression(0)).(ast.Expression)
	right := v.Visit(ctx.Expression(1)).(ast.Expression)
	return &ast.BinaryOpExpr{
		Left:  left,
		Right: right,
		Op:    "<",
	}
}

func (v *ASTBuilder) VisitAndExpression(ctx *AndExpressionContext) interface{} {
	left := v.Visit(ctx.Expression(0)).(ast.Expression)
	right := v.Visit(ctx.Expression(1)).(ast.Expression)
	return &ast.BinaryOpExpr{
		Left:  left,
		Right: right,
		Op:    "&&",
	}
}

func (v *ASTBuilder) VisitIntLitExpression(ctx *IntLitExpressionContext) interface{} {
	val, _ := strconv.Atoi(ctx.INT().GetText())
	return &ast.IntLitExpr{
		Value: val,
	}
}

func (v *ASTBuilder) VisitBooleanLitExpression(ctx *BooleanLitExpressionContext) interface{} {
	val := ctx.BOOL().GetText() == "true"
	return &ast.BoolLitExpr{
		Value: val,
	}
}

func (v *ASTBuilder) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return &ast.IdentifierExpr{Name: ctx.ID().GetText()}
}

func (v *ASTBuilder) VisitThisExpression(ctx *ThisExpressionContext) interface{} {
	return &ast.ThisExpr{}
}

func (v *ASTBuilder) VisitParenExpression(ctx *ParenExpressionContext) interface{} {
	expr := v.Visit(ctx.Expression()).(ast.Expression)
	return &ast.ParenExpr{Expr: expr}
}
