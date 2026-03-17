// Code generated from /home/konstantin/go/Compilers_project/grammar/MiniJava.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MiniJava
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MiniJavaParser.
type MiniJavaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MiniJavaParser#goal.
	VisitGoal(ctx *GoalContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#mainClassDeclaration.
	VisitMainClassDeclaration(ctx *MainClassDeclarationContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#mainClassBody.
	VisitMainClassBody(ctx *MainClassBodyContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#mainMethod.
	VisitMainMethod(ctx *MainMethodContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#mainMethodDeclaration.
	VisitMainMethodDeclaration(ctx *MainMethodDeclarationContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#classBody.
	VisitClassBody(ctx *ClassBodyContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#fieldDeclaration.
	VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#varDeclaration.
	VisitVarDeclaration(ctx *VarDeclarationContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#methodDeclaration.
	VisitMethodDeclaration(ctx *MethodDeclarationContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#methodBody.
	VisitMethodBody(ctx *MethodBodyContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#formalParameters.
	VisitFormalParameters(ctx *FormalParametersContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#formalParameterList.
	VisitFormalParameterList(ctx *FormalParameterListContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#formalParameter.
	VisitFormalParameter(ctx *FormalParameterContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#nestedStatement.
	VisitNestedStatement(ctx *NestedStatementContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#ifElseStatement.
	VisitIfElseStatement(ctx *IfElseStatementContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#printStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#assignStatement.
	VisitAssignStatement(ctx *AssignStatementContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#arrayAssignStatement.
	VisitArrayAssignStatement(ctx *ArrayAssignStatementContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#recurStatement.
	VisitRecurStatement(ctx *RecurStatementContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#ltExpression.
	VisitLtExpression(ctx *LtExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#objectInstantiationExpression.
	VisitObjectInstantiationExpression(ctx *ObjectInstantiationExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#arrayInstantiationExpression.
	VisitArrayInstantiationExpression(ctx *ArrayInstantiationExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#identifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#methodCallExpression.
	VisitMethodCallExpression(ctx *MethodCallExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#notExpression.
	VisitNotExpression(ctx *NotExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#booleanLitExpression.
	VisitBooleanLitExpression(ctx *BooleanLitExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#parenExpression.
	VisitParenExpression(ctx *ParenExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#intLitExpression.
	VisitIntLitExpression(ctx *IntLitExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#arrayAccessExpression.
	VisitArrayAccessExpression(ctx *ArrayAccessExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#addExpression.
	VisitAddExpression(ctx *AddExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#thisExpression.
	VisitThisExpression(ctx *ThisExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#arrayLengthExpression.
	VisitArrayLengthExpression(ctx *ArrayLengthExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#negExpression.
	VisitNegExpression(ctx *NegExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#subExpression.
	VisitSubExpression(ctx *SubExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#mulExpression.
	VisitMulExpression(ctx *MulExpressionContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#methodArgumentList.
	VisitMethodArgumentList(ctx *MethodArgumentListContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#intArrayType.
	VisitIntArrayType(ctx *IntArrayTypeContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#booleanType.
	VisitBooleanType(ctx *BooleanTypeContext) interface{}

	// Visit a parse tree produced by MiniJavaParser#intType.
	VisitIntType(ctx *IntTypeContext) interface{}
}
