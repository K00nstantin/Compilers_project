// Code generated from MiniJava.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MiniJava
import "github.com/antlr4-go/antlr/v4"

// BaseMiniJavaListener is a complete listener for a parse tree produced by MiniJavaParser.
type BaseMiniJavaListener struct{}

var _ MiniJavaListener = &BaseMiniJavaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMiniJavaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMiniJavaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMiniJavaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMiniJavaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGoal is called when production goal is entered.
func (s *BaseMiniJavaListener) EnterGoal(ctx *GoalContext) {}

// ExitGoal is called when production goal is exited.
func (s *BaseMiniJavaListener) ExitGoal(ctx *GoalContext) {}

// EnterMainClassDeclaration is called when production mainClassDeclaration is entered.
func (s *BaseMiniJavaListener) EnterMainClassDeclaration(ctx *MainClassDeclarationContext) {}

// ExitMainClassDeclaration is called when production mainClassDeclaration is exited.
func (s *BaseMiniJavaListener) ExitMainClassDeclaration(ctx *MainClassDeclarationContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseMiniJavaListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseMiniJavaListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterMainClassBody is called when production mainClassBody is entered.
func (s *BaseMiniJavaListener) EnterMainClassBody(ctx *MainClassBodyContext) {}

// ExitMainClassBody is called when production mainClassBody is exited.
func (s *BaseMiniJavaListener) ExitMainClassBody(ctx *MainClassBodyContext) {}

// EnterMainMethod is called when production mainMethod is entered.
func (s *BaseMiniJavaListener) EnterMainMethod(ctx *MainMethodContext) {}

// ExitMainMethod is called when production mainMethod is exited.
func (s *BaseMiniJavaListener) ExitMainMethod(ctx *MainMethodContext) {}

// EnterMainMethodDeclaration is called when production mainMethodDeclaration is entered.
func (s *BaseMiniJavaListener) EnterMainMethodDeclaration(ctx *MainMethodDeclarationContext) {}

// ExitMainMethodDeclaration is called when production mainMethodDeclaration is exited.
func (s *BaseMiniJavaListener) ExitMainMethodDeclaration(ctx *MainMethodDeclarationContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseMiniJavaListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseMiniJavaListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BaseMiniJavaListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BaseMiniJavaListener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterVarDeclaration is called when production varDeclaration is entered.
func (s *BaseMiniJavaListener) EnterVarDeclaration(ctx *VarDeclarationContext) {}

// ExitVarDeclaration is called when production varDeclaration is exited.
func (s *BaseMiniJavaListener) ExitVarDeclaration(ctx *VarDeclarationContext) {}

// EnterMethodDeclaration is called when production methodDeclaration is entered.
func (s *BaseMiniJavaListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {}

// ExitMethodDeclaration is called when production methodDeclaration is exited.
func (s *BaseMiniJavaListener) ExitMethodDeclaration(ctx *MethodDeclarationContext) {}

// EnterMethodBody is called when production methodBody is entered.
func (s *BaseMiniJavaListener) EnterMethodBody(ctx *MethodBodyContext) {}

// ExitMethodBody is called when production methodBody is exited.
func (s *BaseMiniJavaListener) ExitMethodBody(ctx *MethodBodyContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BaseMiniJavaListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BaseMiniJavaListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseMiniJavaListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseMiniJavaListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BaseMiniJavaListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BaseMiniJavaListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterType is called when production type is entered.
func (s *BaseMiniJavaListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseMiniJavaListener) ExitType(ctx *TypeContext) {}

// EnterNestedStatement is called when production nestedStatement is entered.
func (s *BaseMiniJavaListener) EnterNestedStatement(ctx *NestedStatementContext) {}

// ExitNestedStatement is called when production nestedStatement is exited.
func (s *BaseMiniJavaListener) ExitNestedStatement(ctx *NestedStatementContext) {}

// EnterIfElseStatement is called when production ifElseStatement is entered.
func (s *BaseMiniJavaListener) EnterIfElseStatement(ctx *IfElseStatementContext) {}

// ExitIfElseStatement is called when production ifElseStatement is exited.
func (s *BaseMiniJavaListener) ExitIfElseStatement(ctx *IfElseStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseMiniJavaListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseMiniJavaListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseMiniJavaListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseMiniJavaListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterAssignStatement is called when production assignStatement is entered.
func (s *BaseMiniJavaListener) EnterAssignStatement(ctx *AssignStatementContext) {}

// ExitAssignStatement is called when production assignStatement is exited.
func (s *BaseMiniJavaListener) ExitAssignStatement(ctx *AssignStatementContext) {}

// EnterArrayAssignStatement is called when production arrayAssignStatement is entered.
func (s *BaseMiniJavaListener) EnterArrayAssignStatement(ctx *ArrayAssignStatementContext) {}

// ExitArrayAssignStatement is called when production arrayAssignStatement is exited.
func (s *BaseMiniJavaListener) ExitArrayAssignStatement(ctx *ArrayAssignStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseMiniJavaListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseMiniJavaListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterRecurStatement is called when production recurStatement is entered.
func (s *BaseMiniJavaListener) EnterRecurStatement(ctx *RecurStatementContext) {}

// ExitRecurStatement is called when production recurStatement is exited.
func (s *BaseMiniJavaListener) ExitRecurStatement(ctx *RecurStatementContext) {}

// EnterLtExpression is called when production ltExpression is entered.
func (s *BaseMiniJavaListener) EnterLtExpression(ctx *LtExpressionContext) {}

// ExitLtExpression is called when production ltExpression is exited.
func (s *BaseMiniJavaListener) ExitLtExpression(ctx *LtExpressionContext) {}

// EnterObjectInstantiationExpression is called when production objectInstantiationExpression is entered.
func (s *BaseMiniJavaListener) EnterObjectInstantiationExpression(ctx *ObjectInstantiationExpressionContext) {
}

// ExitObjectInstantiationExpression is called when production objectInstantiationExpression is exited.
func (s *BaseMiniJavaListener) ExitObjectInstantiationExpression(ctx *ObjectInstantiationExpressionContext) {
}

// EnterArrayInstantiationExpression is called when production arrayInstantiationExpression is entered.
func (s *BaseMiniJavaListener) EnterArrayInstantiationExpression(ctx *ArrayInstantiationExpressionContext) {
}

// ExitArrayInstantiationExpression is called when production arrayInstantiationExpression is exited.
func (s *BaseMiniJavaListener) ExitArrayInstantiationExpression(ctx *ArrayInstantiationExpressionContext) {
}

// EnterIdentifierExpression is called when production identifierExpression is entered.
func (s *BaseMiniJavaListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production identifierExpression is exited.
func (s *BaseMiniJavaListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterMethodCallExpression is called when production methodCallExpression is entered.
func (s *BaseMiniJavaListener) EnterMethodCallExpression(ctx *MethodCallExpressionContext) {}

// ExitMethodCallExpression is called when production methodCallExpression is exited.
func (s *BaseMiniJavaListener) ExitMethodCallExpression(ctx *MethodCallExpressionContext) {}

// EnterNotExpression is called when production notExpression is entered.
func (s *BaseMiniJavaListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production notExpression is exited.
func (s *BaseMiniJavaListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterBooleanLitExpression is called when production booleanLitExpression is entered.
func (s *BaseMiniJavaListener) EnterBooleanLitExpression(ctx *BooleanLitExpressionContext) {}

// ExitBooleanLitExpression is called when production booleanLitExpression is exited.
func (s *BaseMiniJavaListener) ExitBooleanLitExpression(ctx *BooleanLitExpressionContext) {}

// EnterParenExpression is called when production parenExpression is entered.
func (s *BaseMiniJavaListener) EnterParenExpression(ctx *ParenExpressionContext) {}

// ExitParenExpression is called when production parenExpression is exited.
func (s *BaseMiniJavaListener) ExitParenExpression(ctx *ParenExpressionContext) {}

// EnterIntLitExpression is called when production intLitExpression is entered.
func (s *BaseMiniJavaListener) EnterIntLitExpression(ctx *IntLitExpressionContext) {}

// ExitIntLitExpression is called when production intLitExpression is exited.
func (s *BaseMiniJavaListener) ExitIntLitExpression(ctx *IntLitExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseMiniJavaListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseMiniJavaListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterArrayAccessExpression is called when production arrayAccessExpression is entered.
func (s *BaseMiniJavaListener) EnterArrayAccessExpression(ctx *ArrayAccessExpressionContext) {}

// ExitArrayAccessExpression is called when production arrayAccessExpression is exited.
func (s *BaseMiniJavaListener) ExitArrayAccessExpression(ctx *ArrayAccessExpressionContext) {}

// EnterAddExpression is called when production addExpression is entered.
func (s *BaseMiniJavaListener) EnterAddExpression(ctx *AddExpressionContext) {}

// ExitAddExpression is called when production addExpression is exited.
func (s *BaseMiniJavaListener) ExitAddExpression(ctx *AddExpressionContext) {}

// EnterThisExpression is called when production thisExpression is entered.
func (s *BaseMiniJavaListener) EnterThisExpression(ctx *ThisExpressionContext) {}

// ExitThisExpression is called when production thisExpression is exited.
func (s *BaseMiniJavaListener) ExitThisExpression(ctx *ThisExpressionContext) {}

// EnterArrayLengthExpression is called when production arrayLengthExpression is entered.
func (s *BaseMiniJavaListener) EnterArrayLengthExpression(ctx *ArrayLengthExpressionContext) {}

// ExitArrayLengthExpression is called when production arrayLengthExpression is exited.
func (s *BaseMiniJavaListener) ExitArrayLengthExpression(ctx *ArrayLengthExpressionContext) {}

// EnterNegExpression is called when production negExpression is entered.
func (s *BaseMiniJavaListener) EnterNegExpression(ctx *NegExpressionContext) {}

// ExitNegExpression is called when production negExpression is exited.
func (s *BaseMiniJavaListener) ExitNegExpression(ctx *NegExpressionContext) {}

// EnterSubExpression is called when production subExpression is entered.
func (s *BaseMiniJavaListener) EnterSubExpression(ctx *SubExpressionContext) {}

// ExitSubExpression is called when production subExpression is exited.
func (s *BaseMiniJavaListener) ExitSubExpression(ctx *SubExpressionContext) {}

// EnterMulExpression is called when production mulExpression is entered.
func (s *BaseMiniJavaListener) EnterMulExpression(ctx *MulExpressionContext) {}

// ExitMulExpression is called when production mulExpression is exited.
func (s *BaseMiniJavaListener) ExitMulExpression(ctx *MulExpressionContext) {}

// EnterMethodArgumentList is called when production methodArgumentList is entered.
func (s *BaseMiniJavaListener) EnterMethodArgumentList(ctx *MethodArgumentListContext) {}

// ExitMethodArgumentList is called when production methodArgumentList is exited.
func (s *BaseMiniJavaListener) ExitMethodArgumentList(ctx *MethodArgumentListContext) {}

// EnterIntArrayType is called when production intArrayType is entered.
func (s *BaseMiniJavaListener) EnterIntArrayType(ctx *IntArrayTypeContext) {}

// ExitIntArrayType is called when production intArrayType is exited.
func (s *BaseMiniJavaListener) ExitIntArrayType(ctx *IntArrayTypeContext) {}

// EnterBooleanType is called when production booleanType is entered.
func (s *BaseMiniJavaListener) EnterBooleanType(ctx *BooleanTypeContext) {}

// ExitBooleanType is called when production booleanType is exited.
func (s *BaseMiniJavaListener) ExitBooleanType(ctx *BooleanTypeContext) {}

// EnterIntType is called when production intType is entered.
func (s *BaseMiniJavaListener) EnterIntType(ctx *IntTypeContext) {}

// ExitIntType is called when production intType is exited.
func (s *BaseMiniJavaListener) ExitIntType(ctx *IntTypeContext) {}
