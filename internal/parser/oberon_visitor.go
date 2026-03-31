// Code generated from /home/konstantin/go/Compilers_project/grammar/oberon.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // oberon
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by oberonParser.
type oberonVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by oberonParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by oberonParser#qualident.
	VisitQualident(ctx *QualidentContext) interface{}

	// Visit a parse tree produced by oberonParser#identdef.
	VisitIdentdef(ctx *IdentdefContext) interface{}

	// Visit a parse tree produced by oberonParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by oberonParser#real.
	VisitReal(ctx *RealContext) interface{}

	// Visit a parse tree produced by oberonParser#scaleFactor.
	VisitScaleFactor(ctx *ScaleFactorContext) interface{}

	// Visit a parse tree produced by oberonParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by oberonParser#constDeclaration.
	VisitConstDeclaration(ctx *ConstDeclarationContext) interface{}

	// Visit a parse tree produced by oberonParser#constExpression.
	VisitConstExpression(ctx *ConstExpressionContext) interface{}

	// Visit a parse tree produced by oberonParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by oberonParser#type_.
	VisitType_(ctx *Type_Context) interface{}

	// Visit a parse tree produced by oberonParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by oberonParser#length.
	VisitLength(ctx *LengthContext) interface{}

	// Visit a parse tree produced by oberonParser#recordType.
	VisitRecordType(ctx *RecordTypeContext) interface{}

	// Visit a parse tree produced by oberonParser#baseType.
	VisitBaseType(ctx *BaseTypeContext) interface{}

	// Visit a parse tree produced by oberonParser#fieldListSequence.
	VisitFieldListSequence(ctx *FieldListSequenceContext) interface{}

	// Visit a parse tree produced by oberonParser#fieldList.
	VisitFieldList(ctx *FieldListContext) interface{}

	// Visit a parse tree produced by oberonParser#identList.
	VisitIdentList(ctx *IdentListContext) interface{}

	// Visit a parse tree produced by oberonParser#pointerType.
	VisitPointerType(ctx *PointerTypeContext) interface{}

	// Visit a parse tree produced by oberonParser#procedureType.
	VisitProcedureType(ctx *ProcedureTypeContext) interface{}

	// Visit a parse tree produced by oberonParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by oberonParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by oberonParser#relation.
	VisitRelation(ctx *RelationContext) interface{}

	// Visit a parse tree produced by oberonParser#simpleExpression.
	VisitSimpleExpression(ctx *SimpleExpressionContext) interface{}

	// Visit a parse tree produced by oberonParser#addOperator.
	VisitAddOperator(ctx *AddOperatorContext) interface{}

	// Visit a parse tree produced by oberonParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by oberonParser#mulOperator.
	VisitMulOperator(ctx *MulOperatorContext) interface{}

	// Visit a parse tree produced by oberonParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by oberonParser#designator.
	VisitDesignator(ctx *DesignatorContext) interface{}

	// Visit a parse tree produced by oberonParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by oberonParser#set_.
	VisitSet_(ctx *Set_Context) interface{}

	// Visit a parse tree produced by oberonParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Visit a parse tree produced by oberonParser#expList.
	VisitExpList(ctx *ExpListContext) interface{}

	// Visit a parse tree produced by oberonParser#actualParameters.
	VisitActualParameters(ctx *ActualParametersContext) interface{}

	// Visit a parse tree produced by oberonParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by oberonParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by oberonParser#procedureCall.
	VisitProcedureCall(ctx *ProcedureCallContext) interface{}

	// Visit a parse tree produced by oberonParser#statementSequence.
	VisitStatementSequence(ctx *StatementSequenceContext) interface{}

	// Visit a parse tree produced by oberonParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by oberonParser#caseStatement.
	VisitCaseStatement(ctx *CaseStatementContext) interface{}

	// Visit a parse tree produced by oberonParser#case_.
	VisitCase_(ctx *Case_Context) interface{}

	// Visit a parse tree produced by oberonParser#caseLabelList.
	VisitCaseLabelList(ctx *CaseLabelListContext) interface{}

	// Visit a parse tree produced by oberonParser#labelRange.
	VisitLabelRange(ctx *LabelRangeContext) interface{}

	// Visit a parse tree produced by oberonParser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by oberonParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by oberonParser#repeatStatement.
	VisitRepeatStatement(ctx *RepeatStatementContext) interface{}

	// Visit a parse tree produced by oberonParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by oberonParser#procedureDeclaration.
	VisitProcedureDeclaration(ctx *ProcedureDeclarationContext) interface{}

	// Visit a parse tree produced by oberonParser#procedureHeading.
	VisitProcedureHeading(ctx *ProcedureHeadingContext) interface{}

	// Visit a parse tree produced by oberonParser#procedureBody.
	VisitProcedureBody(ctx *ProcedureBodyContext) interface{}

	// Visit a parse tree produced by oberonParser#declarationSequence.
	VisitDeclarationSequence(ctx *DeclarationSequenceContext) interface{}

	// Visit a parse tree produced by oberonParser#formalParameters.
	VisitFormalParameters(ctx *FormalParametersContext) interface{}

	// Visit a parse tree produced by oberonParser#fPSection.
	VisitFPSection(ctx *FPSectionContext) interface{}

	// Visit a parse tree produced by oberonParser#formalType.
	VisitFormalType(ctx *FormalTypeContext) interface{}

	// Visit a parse tree produced by oberonParser#module.
	VisitModule(ctx *ModuleContext) interface{}

	// Visit a parse tree produced by oberonParser#importList.
	VisitImportList(ctx *ImportListContext) interface{}

	// Visit a parse tree produced by oberonParser#import_.
	VisitImport_(ctx *Import_Context) interface{}
}
