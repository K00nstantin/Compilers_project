// Code generated from MiniJava.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MiniJava
import "github.com/antlr4-go/antlr/v4"

type BaseMiniJavaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMiniJavaVisitor) VisitGoal(ctx *GoalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitMainClassDeclaration(ctx *MainClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitMainClassBody(ctx *MainClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitMainMethod(ctx *MainMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitMainMethodDeclaration(ctx *MainMethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitClassBody(ctx *ClassBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitFieldDeclaration(ctx *FieldDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitVarDeclaration(ctx *VarDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitMethodDeclaration(ctx *MethodDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitMethodBody(ctx *MethodBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitFormalParameters(ctx *FormalParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitFormalParameterList(ctx *FormalParameterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitFormalParameter(ctx *FormalParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitNestedStatement(ctx *NestedStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitIfElseStatement(ctx *IfElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitArrayAssignStatement(ctx *ArrayAssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitRecurStatement(ctx *RecurStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitLtExpression(ctx *LtExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitObjectInstantiationExpression(ctx *ObjectInstantiationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitArrayInstantiationExpression(ctx *ArrayInstantiationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitMethodCallExpression(ctx *MethodCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitNotExpression(ctx *NotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitBooleanLitExpression(ctx *BooleanLitExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitParenExpression(ctx *ParenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitIntLitExpression(ctx *IntLitExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitAndExpression(ctx *AndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitArrayAccessExpression(ctx *ArrayAccessExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitAddExpression(ctx *AddExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitThisExpression(ctx *ThisExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitArrayLengthExpression(ctx *ArrayLengthExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitNegExpression(ctx *NegExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitSubExpression(ctx *SubExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitMulExpression(ctx *MulExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitMethodArgumentList(ctx *MethodArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitIntArrayType(ctx *IntArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitBooleanType(ctx *BooleanTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniJavaVisitor) VisitIntType(ctx *IntTypeContext) interface{} {
	return v.VisitChildren(ctx)
}
