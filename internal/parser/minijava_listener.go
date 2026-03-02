// Code generated from MiniJava.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MiniJava
import "github.com/antlr4-go/antlr/v4"

// MiniJavaListener is a complete listener for a parse tree produced by MiniJavaParser.
type MiniJavaListener interface {
	antlr.ParseTreeListener

	// EnterGoal is called when entering the goal production.
	EnterGoal(c *GoalContext)

	// EnterMainClassDeclaration is called when entering the mainClassDeclaration production.
	EnterMainClassDeclaration(c *MainClassDeclarationContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterMainClassBody is called when entering the mainClassBody production.
	EnterMainClassBody(c *MainClassBodyContext)

	// EnterMainMethod is called when entering the mainMethod production.
	EnterMainMethod(c *MainMethodContext)

	// EnterMainMethodDeclaration is called when entering the mainMethodDeclaration production.
	EnterMainMethodDeclaration(c *MainMethodDeclarationContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterFieldDeclaration is called when entering the fieldDeclaration production.
	EnterFieldDeclaration(c *FieldDeclarationContext)

	// EnterVarDeclaration is called when entering the varDeclaration production.
	EnterVarDeclaration(c *VarDeclarationContext)

	// EnterMethodDeclaration is called when entering the methodDeclaration production.
	EnterMethodDeclaration(c *MethodDeclarationContext)

	// EnterMethodBody is called when entering the methodBody production.
	EnterMethodBody(c *MethodBodyContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterNestedStatement is called when entering the nestedStatement production.
	EnterNestedStatement(c *NestedStatementContext)

	// EnterIfElseStatement is called when entering the ifElseStatement production.
	EnterIfElseStatement(c *IfElseStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterAssignStatement is called when entering the assignStatement production.
	EnterAssignStatement(c *AssignStatementContext)

	// EnterArrayAssignStatement is called when entering the arrayAssignStatement production.
	EnterArrayAssignStatement(c *ArrayAssignStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterRecurStatement is called when entering the recurStatement production.
	EnterRecurStatement(c *RecurStatementContext)

	// EnterLtExpression is called when entering the ltExpression production.
	EnterLtExpression(c *LtExpressionContext)

	// EnterObjectInstantiationExpression is called when entering the objectInstantiationExpression production.
	EnterObjectInstantiationExpression(c *ObjectInstantiationExpressionContext)

	// EnterArrayInstantiationExpression is called when entering the arrayInstantiationExpression production.
	EnterArrayInstantiationExpression(c *ArrayInstantiationExpressionContext)

	// EnterIdentifierExpression is called when entering the identifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// EnterMethodCallExpression is called when entering the methodCallExpression production.
	EnterMethodCallExpression(c *MethodCallExpressionContext)

	// EnterNotExpression is called when entering the notExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterBooleanLitExpression is called when entering the booleanLitExpression production.
	EnterBooleanLitExpression(c *BooleanLitExpressionContext)

	// EnterParenExpression is called when entering the parenExpression production.
	EnterParenExpression(c *ParenExpressionContext)

	// EnterIntLitExpression is called when entering the intLitExpression production.
	EnterIntLitExpression(c *IntLitExpressionContext)

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterArrayAccessExpression is called when entering the arrayAccessExpression production.
	EnterArrayAccessExpression(c *ArrayAccessExpressionContext)

	// EnterAddExpression is called when entering the addExpression production.
	EnterAddExpression(c *AddExpressionContext)

	// EnterThisExpression is called when entering the thisExpression production.
	EnterThisExpression(c *ThisExpressionContext)

	// EnterArrayLengthExpression is called when entering the arrayLengthExpression production.
	EnterArrayLengthExpression(c *ArrayLengthExpressionContext)

	// EnterNegExpression is called when entering the negExpression production.
	EnterNegExpression(c *NegExpressionContext)

	// EnterSubExpression is called when entering the subExpression production.
	EnterSubExpression(c *SubExpressionContext)

	// EnterMulExpression is called when entering the mulExpression production.
	EnterMulExpression(c *MulExpressionContext)

	// EnterMethodArgumentList is called when entering the methodArgumentList production.
	EnterMethodArgumentList(c *MethodArgumentListContext)

	// EnterIntArrayType is called when entering the intArrayType production.
	EnterIntArrayType(c *IntArrayTypeContext)

	// EnterBooleanType is called when entering the booleanType production.
	EnterBooleanType(c *BooleanTypeContext)

	// EnterIntType is called when entering the intType production.
	EnterIntType(c *IntTypeContext)

	// ExitGoal is called when exiting the goal production.
	ExitGoal(c *GoalContext)

	// ExitMainClassDeclaration is called when exiting the mainClassDeclaration production.
	ExitMainClassDeclaration(c *MainClassDeclarationContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitMainClassBody is called when exiting the mainClassBody production.
	ExitMainClassBody(c *MainClassBodyContext)

	// ExitMainMethod is called when exiting the mainMethod production.
	ExitMainMethod(c *MainMethodContext)

	// ExitMainMethodDeclaration is called when exiting the mainMethodDeclaration production.
	ExitMainMethodDeclaration(c *MainMethodDeclarationContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitFieldDeclaration is called when exiting the fieldDeclaration production.
	ExitFieldDeclaration(c *FieldDeclarationContext)

	// ExitVarDeclaration is called when exiting the varDeclaration production.
	ExitVarDeclaration(c *VarDeclarationContext)

	// ExitMethodDeclaration is called when exiting the methodDeclaration production.
	ExitMethodDeclaration(c *MethodDeclarationContext)

	// ExitMethodBody is called when exiting the methodBody production.
	ExitMethodBody(c *MethodBodyContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitNestedStatement is called when exiting the nestedStatement production.
	ExitNestedStatement(c *NestedStatementContext)

	// ExitIfElseStatement is called when exiting the ifElseStatement production.
	ExitIfElseStatement(c *IfElseStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitAssignStatement is called when exiting the assignStatement production.
	ExitAssignStatement(c *AssignStatementContext)

	// ExitArrayAssignStatement is called when exiting the arrayAssignStatement production.
	ExitArrayAssignStatement(c *ArrayAssignStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitRecurStatement is called when exiting the recurStatement production.
	ExitRecurStatement(c *RecurStatementContext)

	// ExitLtExpression is called when exiting the ltExpression production.
	ExitLtExpression(c *LtExpressionContext)

	// ExitObjectInstantiationExpression is called when exiting the objectInstantiationExpression production.
	ExitObjectInstantiationExpression(c *ObjectInstantiationExpressionContext)

	// ExitArrayInstantiationExpression is called when exiting the arrayInstantiationExpression production.
	ExitArrayInstantiationExpression(c *ArrayInstantiationExpressionContext)

	// ExitIdentifierExpression is called when exiting the identifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)

	// ExitMethodCallExpression is called when exiting the methodCallExpression production.
	ExitMethodCallExpression(c *MethodCallExpressionContext)

	// ExitNotExpression is called when exiting the notExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitBooleanLitExpression is called when exiting the booleanLitExpression production.
	ExitBooleanLitExpression(c *BooleanLitExpressionContext)

	// ExitParenExpression is called when exiting the parenExpression production.
	ExitParenExpression(c *ParenExpressionContext)

	// ExitIntLitExpression is called when exiting the intLitExpression production.
	ExitIntLitExpression(c *IntLitExpressionContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitArrayAccessExpression is called when exiting the arrayAccessExpression production.
	ExitArrayAccessExpression(c *ArrayAccessExpressionContext)

	// ExitAddExpression is called when exiting the addExpression production.
	ExitAddExpression(c *AddExpressionContext)

	// ExitThisExpression is called when exiting the thisExpression production.
	ExitThisExpression(c *ThisExpressionContext)

	// ExitArrayLengthExpression is called when exiting the arrayLengthExpression production.
	ExitArrayLengthExpression(c *ArrayLengthExpressionContext)

	// ExitNegExpression is called when exiting the negExpression production.
	ExitNegExpression(c *NegExpressionContext)

	// ExitSubExpression is called when exiting the subExpression production.
	ExitSubExpression(c *SubExpressionContext)

	// ExitMulExpression is called when exiting the mulExpression production.
	ExitMulExpression(c *MulExpressionContext)

	// ExitMethodArgumentList is called when exiting the methodArgumentList production.
	ExitMethodArgumentList(c *MethodArgumentListContext)

	// ExitIntArrayType is called when exiting the intArrayType production.
	ExitIntArrayType(c *IntArrayTypeContext)

	// ExitBooleanType is called when exiting the booleanType production.
	ExitBooleanType(c *BooleanTypeContext)

	// ExitIntType is called when exiting the intType production.
	ExitIntType(c *IntTypeContext)
}
