// Code generated from /home/konstantin/go/Compilers_project/grammar/oberon.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // oberon
import "github.com/antlr4-go/antlr/v4"

// oberonListener is a complete listener for a parse tree produced by oberonParser.
type oberonListener interface {
	antlr.ParseTreeListener

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterQualident is called when entering the qualident production.
	EnterQualident(c *QualidentContext)

	// EnterIdentdef is called when entering the identdef production.
	EnterIdentdef(c *IdentdefContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterConstDeclaration is called when entering the constDeclaration production.
	EnterConstDeclaration(c *ConstDeclarationContext)

	// EnterConstExpression is called when entering the constExpression production.
	EnterConstExpression(c *ConstExpressionContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterLength is called when entering the length production.
	EnterLength(c *LengthContext)

	// EnterRecordType is called when entering the recordType production.
	EnterRecordType(c *RecordTypeContext)

	// EnterBaseType is called when entering the baseType production.
	EnterBaseType(c *BaseTypeContext)

	// EnterFieldListSequence is called when entering the fieldListSequence production.
	EnterFieldListSequence(c *FieldListSequenceContext)

	// EnterFieldList is called when entering the fieldList production.
	EnterFieldList(c *FieldListContext)

	// EnterIdentList is called when entering the identList production.
	EnterIdentList(c *IdentListContext)

	// EnterPointerType is called when entering the pointerType production.
	EnterPointerType(c *PointerTypeContext)

	// EnterProcedureType is called when entering the procedureType production.
	EnterProcedureType(c *ProcedureTypeContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterSimpleExpression is called when entering the simpleExpression production.
	EnterSimpleExpression(c *SimpleExpressionContext)

	// EnterAddOperator is called when entering the addOperator production.
	EnterAddOperator(c *AddOperatorContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterMulOperator is called when entering the mulOperator production.
	EnterMulOperator(c *MulOperatorContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterDesignator is called when entering the designator production.
	EnterDesignator(c *DesignatorContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterSet_ is called when entering the set_ production.
	EnterSet_(c *Set_Context)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterExpList is called when entering the expList production.
	EnterExpList(c *ExpListContext)

	// EnterActualParameters is called when entering the actualParameters production.
	EnterActualParameters(c *ActualParametersContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterProcedureCall is called when entering the procedureCall production.
	EnterProcedureCall(c *ProcedureCallContext)

	// EnterStatementSequence is called when entering the statementSequence production.
	EnterStatementSequence(c *StatementSequenceContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterCaseStatement is called when entering the caseStatement production.
	EnterCaseStatement(c *CaseStatementContext)

	// EnterCase_ is called when entering the case_ production.
	EnterCase_(c *Case_Context)

	// EnterCaseLabelList is called when entering the caseLabelList production.
	EnterCaseLabelList(c *CaseLabelListContext)

	// EnterLabelRange is called when entering the labelRange production.
	EnterLabelRange(c *LabelRangeContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterRepeatStatement is called when entering the repeatStatement production.
	EnterRepeatStatement(c *RepeatStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterProcedureDeclaration is called when entering the procedureDeclaration production.
	EnterProcedureDeclaration(c *ProcedureDeclarationContext)

	// EnterProcedureHeading is called when entering the procedureHeading production.
	EnterProcedureHeading(c *ProcedureHeadingContext)

	// EnterProcedureBody is called when entering the procedureBody production.
	EnterProcedureBody(c *ProcedureBodyContext)

	// EnterDeclarationSequence is called when entering the declarationSequence production.
	EnterDeclarationSequence(c *DeclarationSequenceContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterFPSection is called when entering the fPSection production.
	EnterFPSection(c *FPSectionContext)

	// EnterFormalType is called when entering the formalType production.
	EnterFormalType(c *FormalTypeContext)

	// EnterModule is called when entering the module production.
	EnterModule(c *ModuleContext)

	// EnterImportList is called when entering the importList production.
	EnterImportList(c *ImportListContext)

	// EnterImport_ is called when entering the import_ production.
	EnterImport_(c *Import_Context)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitQualident is called when exiting the qualident production.
	ExitQualident(c *QualidentContext)

	// ExitIdentdef is called when exiting the identdef production.
	ExitIdentdef(c *IdentdefContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitConstDeclaration is called when exiting the constDeclaration production.
	ExitConstDeclaration(c *ConstDeclarationContext)

	// ExitConstExpression is called when exiting the constExpression production.
	ExitConstExpression(c *ConstExpressionContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitLength is called when exiting the length production.
	ExitLength(c *LengthContext)

	// ExitRecordType is called when exiting the recordType production.
	ExitRecordType(c *RecordTypeContext)

	// ExitBaseType is called when exiting the baseType production.
	ExitBaseType(c *BaseTypeContext)

	// ExitFieldListSequence is called when exiting the fieldListSequence production.
	ExitFieldListSequence(c *FieldListSequenceContext)

	// ExitFieldList is called when exiting the fieldList production.
	ExitFieldList(c *FieldListContext)

	// ExitIdentList is called when exiting the identList production.
	ExitIdentList(c *IdentListContext)

	// ExitPointerType is called when exiting the pointerType production.
	ExitPointerType(c *PointerTypeContext)

	// ExitProcedureType is called when exiting the procedureType production.
	ExitProcedureType(c *ProcedureTypeContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitSimpleExpression is called when exiting the simpleExpression production.
	ExitSimpleExpression(c *SimpleExpressionContext)

	// ExitAddOperator is called when exiting the addOperator production.
	ExitAddOperator(c *AddOperatorContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitMulOperator is called when exiting the mulOperator production.
	ExitMulOperator(c *MulOperatorContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitDesignator is called when exiting the designator production.
	ExitDesignator(c *DesignatorContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitSet_ is called when exiting the set_ production.
	ExitSet_(c *Set_Context)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitExpList is called when exiting the expList production.
	ExitExpList(c *ExpListContext)

	// ExitActualParameters is called when exiting the actualParameters production.
	ExitActualParameters(c *ActualParametersContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitProcedureCall is called when exiting the procedureCall production.
	ExitProcedureCall(c *ProcedureCallContext)

	// ExitStatementSequence is called when exiting the statementSequence production.
	ExitStatementSequence(c *StatementSequenceContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitCaseStatement is called when exiting the caseStatement production.
	ExitCaseStatement(c *CaseStatementContext)

	// ExitCase_ is called when exiting the case_ production.
	ExitCase_(c *Case_Context)

	// ExitCaseLabelList is called when exiting the caseLabelList production.
	ExitCaseLabelList(c *CaseLabelListContext)

	// ExitLabelRange is called when exiting the labelRange production.
	ExitLabelRange(c *LabelRangeContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitRepeatStatement is called when exiting the repeatStatement production.
	ExitRepeatStatement(c *RepeatStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitProcedureDeclaration is called when exiting the procedureDeclaration production.
	ExitProcedureDeclaration(c *ProcedureDeclarationContext)

	// ExitProcedureHeading is called when exiting the procedureHeading production.
	ExitProcedureHeading(c *ProcedureHeadingContext)

	// ExitProcedureBody is called when exiting the procedureBody production.
	ExitProcedureBody(c *ProcedureBodyContext)

	// ExitDeclarationSequence is called when exiting the declarationSequence production.
	ExitDeclarationSequence(c *DeclarationSequenceContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitFPSection is called when exiting the fPSection production.
	ExitFPSection(c *FPSectionContext)

	// ExitFormalType is called when exiting the formalType production.
	ExitFormalType(c *FormalTypeContext)

	// ExitModule is called when exiting the module production.
	ExitModule(c *ModuleContext)

	// ExitImportList is called when exiting the importList production.
	ExitImportList(c *ImportListContext)

	// ExitImport_ is called when exiting the import_ production.
	ExitImport_(c *Import_Context)
}
