// Code generated from /home/konstantin/go/Compilers_project/grammar/oberon.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // oberon
import "github.com/antlr4-go/antlr/v4"

// BaseoberonListener is a complete listener for a parse tree produced by oberonParser.
type BaseoberonListener struct{}

var _ oberonListener = &BaseoberonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseoberonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseoberonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseoberonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseoberonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseoberonListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseoberonListener) ExitIdent(ctx *IdentContext) {}

// EnterQualident is called when production qualident is entered.
func (s *BaseoberonListener) EnterQualident(ctx *QualidentContext) {}

// ExitQualident is called when production qualident is exited.
func (s *BaseoberonListener) ExitQualident(ctx *QualidentContext) {}

// EnterIdentdef is called when production identdef is entered.
func (s *BaseoberonListener) EnterIdentdef(ctx *IdentdefContext) {}

// ExitIdentdef is called when production identdef is exited.
func (s *BaseoberonListener) ExitIdentdef(ctx *IdentdefContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseoberonListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseoberonListener) ExitInteger(ctx *IntegerContext) {}

// EnterReal is called when production real is entered.
func (s *BaseoberonListener) EnterReal(ctx *RealContext) {}

// ExitReal is called when production real is exited.
func (s *BaseoberonListener) ExitReal(ctx *RealContext) {}

// EnterScaleFactor is called when production scaleFactor is entered.
func (s *BaseoberonListener) EnterScaleFactor(ctx *ScaleFactorContext) {}

// ExitScaleFactor is called when production scaleFactor is exited.
func (s *BaseoberonListener) ExitScaleFactor(ctx *ScaleFactorContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseoberonListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseoberonListener) ExitNumber(ctx *NumberContext) {}

// EnterConstDeclaration is called when production constDeclaration is entered.
func (s *BaseoberonListener) EnterConstDeclaration(ctx *ConstDeclarationContext) {}

// ExitConstDeclaration is called when production constDeclaration is exited.
func (s *BaseoberonListener) ExitConstDeclaration(ctx *ConstDeclarationContext) {}

// EnterConstExpression is called when production constExpression is entered.
func (s *BaseoberonListener) EnterConstExpression(ctx *ConstExpressionContext) {}

// ExitConstExpression is called when production constExpression is exited.
func (s *BaseoberonListener) ExitConstExpression(ctx *ConstExpressionContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseoberonListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseoberonListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseoberonListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseoberonListener) ExitType_(ctx *Type_Context) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseoberonListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseoberonListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterLength is called when production length is entered.
func (s *BaseoberonListener) EnterLength(ctx *LengthContext) {}

// ExitLength is called when production length is exited.
func (s *BaseoberonListener) ExitLength(ctx *LengthContext) {}

// EnterRecordType is called when production recordType is entered.
func (s *BaseoberonListener) EnterRecordType(ctx *RecordTypeContext) {}

// ExitRecordType is called when production recordType is exited.
func (s *BaseoberonListener) ExitRecordType(ctx *RecordTypeContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseoberonListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseoberonListener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterFieldListSequence is called when production fieldListSequence is entered.
func (s *BaseoberonListener) EnterFieldListSequence(ctx *FieldListSequenceContext) {}

// ExitFieldListSequence is called when production fieldListSequence is exited.
func (s *BaseoberonListener) ExitFieldListSequence(ctx *FieldListSequenceContext) {}

// EnterFieldList is called when production fieldList is entered.
func (s *BaseoberonListener) EnterFieldList(ctx *FieldListContext) {}

// ExitFieldList is called when production fieldList is exited.
func (s *BaseoberonListener) ExitFieldList(ctx *FieldListContext) {}

// EnterIdentList is called when production identList is entered.
func (s *BaseoberonListener) EnterIdentList(ctx *IdentListContext) {}

// ExitIdentList is called when production identList is exited.
func (s *BaseoberonListener) ExitIdentList(ctx *IdentListContext) {}

// EnterPointerType is called when production pointerType is entered.
func (s *BaseoberonListener) EnterPointerType(ctx *PointerTypeContext) {}

// ExitPointerType is called when production pointerType is exited.
func (s *BaseoberonListener) ExitPointerType(ctx *PointerTypeContext) {}

// EnterProcedureType is called when production procedureType is entered.
func (s *BaseoberonListener) EnterProcedureType(ctx *ProcedureTypeContext) {}

// ExitProcedureType is called when production procedureType is exited.
func (s *BaseoberonListener) ExitProcedureType(ctx *ProcedureTypeContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseoberonListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseoberonListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseoberonListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseoberonListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseoberonListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseoberonListener) ExitRelation(ctx *RelationContext) {}

// EnterSimpleExpression is called when production simpleExpression is entered.
func (s *BaseoberonListener) EnterSimpleExpression(ctx *SimpleExpressionContext) {}

// ExitSimpleExpression is called when production simpleExpression is exited.
func (s *BaseoberonListener) ExitSimpleExpression(ctx *SimpleExpressionContext) {}

// EnterAddOperator is called when production addOperator is entered.
func (s *BaseoberonListener) EnterAddOperator(ctx *AddOperatorContext) {}

// ExitAddOperator is called when production addOperator is exited.
func (s *BaseoberonListener) ExitAddOperator(ctx *AddOperatorContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseoberonListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseoberonListener) ExitTerm(ctx *TermContext) {}

// EnterMulOperator is called when production mulOperator is entered.
func (s *BaseoberonListener) EnterMulOperator(ctx *MulOperatorContext) {}

// ExitMulOperator is called when production mulOperator is exited.
func (s *BaseoberonListener) ExitMulOperator(ctx *MulOperatorContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseoberonListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseoberonListener) ExitFactor(ctx *FactorContext) {}

// EnterDesignator is called when production designator is entered.
func (s *BaseoberonListener) EnterDesignator(ctx *DesignatorContext) {}

// ExitDesignator is called when production designator is exited.
func (s *BaseoberonListener) ExitDesignator(ctx *DesignatorContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseoberonListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseoberonListener) ExitSelector(ctx *SelectorContext) {}

// EnterSet_ is called when production set_ is entered.
func (s *BaseoberonListener) EnterSet_(ctx *Set_Context) {}

// ExitSet_ is called when production set_ is exited.
func (s *BaseoberonListener) ExitSet_(ctx *Set_Context) {}

// EnterElement is called when production element is entered.
func (s *BaseoberonListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseoberonListener) ExitElement(ctx *ElementContext) {}

// EnterExpList is called when production expList is entered.
func (s *BaseoberonListener) EnterExpList(ctx *ExpListContext) {}

// ExitExpList is called when production expList is exited.
func (s *BaseoberonListener) ExitExpList(ctx *ExpListContext) {}

// EnterActualParameters is called when production actualParameters is entered.
func (s *BaseoberonListener) EnterActualParameters(ctx *ActualParametersContext) {}

// ExitActualParameters is called when production actualParameters is exited.
func (s *BaseoberonListener) ExitActualParameters(ctx *ActualParametersContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseoberonListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseoberonListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseoberonListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseoberonListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterProcedureCall is called when production procedureCall is entered.
func (s *BaseoberonListener) EnterProcedureCall(ctx *ProcedureCallContext) {}

// ExitProcedureCall is called when production procedureCall is exited.
func (s *BaseoberonListener) ExitProcedureCall(ctx *ProcedureCallContext) {}

// EnterStatementSequence is called when production statementSequence is entered.
func (s *BaseoberonListener) EnterStatementSequence(ctx *StatementSequenceContext) {}

// ExitStatementSequence is called when production statementSequence is exited.
func (s *BaseoberonListener) ExitStatementSequence(ctx *StatementSequenceContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseoberonListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseoberonListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterCaseStatement is called when production caseStatement is entered.
func (s *BaseoberonListener) EnterCaseStatement(ctx *CaseStatementContext) {}

// ExitCaseStatement is called when production caseStatement is exited.
func (s *BaseoberonListener) ExitCaseStatement(ctx *CaseStatementContext) {}

// EnterCase_ is called when production case_ is entered.
func (s *BaseoberonListener) EnterCase_(ctx *Case_Context) {}

// ExitCase_ is called when production case_ is exited.
func (s *BaseoberonListener) ExitCase_(ctx *Case_Context) {}

// EnterCaseLabelList is called when production caseLabelList is entered.
func (s *BaseoberonListener) EnterCaseLabelList(ctx *CaseLabelListContext) {}

// ExitCaseLabelList is called when production caseLabelList is exited.
func (s *BaseoberonListener) ExitCaseLabelList(ctx *CaseLabelListContext) {}

// EnterLabelRange is called when production labelRange is entered.
func (s *BaseoberonListener) EnterLabelRange(ctx *LabelRangeContext) {}

// ExitLabelRange is called when production labelRange is exited.
func (s *BaseoberonListener) ExitLabelRange(ctx *LabelRangeContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseoberonListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseoberonListener) ExitLabel(ctx *LabelContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseoberonListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseoberonListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterRepeatStatement is called when production repeatStatement is entered.
func (s *BaseoberonListener) EnterRepeatStatement(ctx *RepeatStatementContext) {}

// ExitRepeatStatement is called when production repeatStatement is exited.
func (s *BaseoberonListener) ExitRepeatStatement(ctx *RepeatStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseoberonListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseoberonListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterProcedureDeclaration is called when production procedureDeclaration is entered.
func (s *BaseoberonListener) EnterProcedureDeclaration(ctx *ProcedureDeclarationContext) {}

// ExitProcedureDeclaration is called when production procedureDeclaration is exited.
func (s *BaseoberonListener) ExitProcedureDeclaration(ctx *ProcedureDeclarationContext) {}

// EnterProcedureHeading is called when production procedureHeading is entered.
func (s *BaseoberonListener) EnterProcedureHeading(ctx *ProcedureHeadingContext) {}

// ExitProcedureHeading is called when production procedureHeading is exited.
func (s *BaseoberonListener) ExitProcedureHeading(ctx *ProcedureHeadingContext) {}

// EnterProcedureBody is called when production procedureBody is entered.
func (s *BaseoberonListener) EnterProcedureBody(ctx *ProcedureBodyContext) {}

// ExitProcedureBody is called when production procedureBody is exited.
func (s *BaseoberonListener) ExitProcedureBody(ctx *ProcedureBodyContext) {}

// EnterDeclarationSequence is called when production declarationSequence is entered.
func (s *BaseoberonListener) EnterDeclarationSequence(ctx *DeclarationSequenceContext) {}

// ExitDeclarationSequence is called when production declarationSequence is exited.
func (s *BaseoberonListener) ExitDeclarationSequence(ctx *DeclarationSequenceContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BaseoberonListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BaseoberonListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFPSection is called when production fPSection is entered.
func (s *BaseoberonListener) EnterFPSection(ctx *FPSectionContext) {}

// ExitFPSection is called when production fPSection is exited.
func (s *BaseoberonListener) ExitFPSection(ctx *FPSectionContext) {}

// EnterFormalType is called when production formalType is entered.
func (s *BaseoberonListener) EnterFormalType(ctx *FormalTypeContext) {}

// ExitFormalType is called when production formalType is exited.
func (s *BaseoberonListener) ExitFormalType(ctx *FormalTypeContext) {}

// EnterModule is called when production module is entered.
func (s *BaseoberonListener) EnterModule(ctx *ModuleContext) {}

// ExitModule is called when production module is exited.
func (s *BaseoberonListener) ExitModule(ctx *ModuleContext) {}

// EnterImportList is called when production importList is entered.
func (s *BaseoberonListener) EnterImportList(ctx *ImportListContext) {}

// ExitImportList is called when production importList is exited.
func (s *BaseoberonListener) ExitImportList(ctx *ImportListContext) {}

// EnterImport_ is called when production import_ is entered.
func (s *BaseoberonListener) EnterImport_(ctx *Import_Context) {}

// ExitImport_ is called when production import_ is exited.
func (s *BaseoberonListener) ExitImport_(ctx *Import_Context) {}
