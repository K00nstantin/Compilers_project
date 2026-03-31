// Code generated from /home/konstantin/go/Compilers_project/grammar/oberon.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // oberon
import "github.com/antlr4-go/antlr/v4"

type BaseoberonVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseoberonVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitQualident(ctx *QualidentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitIdentdef(ctx *IdentdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitReal(ctx *RealContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitScaleFactor(ctx *ScaleFactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitConstDeclaration(ctx *ConstDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitConstExpression(ctx *ConstExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitType_(ctx *Type_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitLength(ctx *LengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitRecordType(ctx *RecordTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitBaseType(ctx *BaseTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitFieldListSequence(ctx *FieldListSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitFieldList(ctx *FieldListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitIdentList(ctx *IdentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitPointerType(ctx *PointerTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitProcedureType(ctx *ProcedureTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitRelation(ctx *RelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitSimpleExpression(ctx *SimpleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitAddOperator(ctx *AddOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitMulOperator(ctx *MulOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitDesignator(ctx *DesignatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitSet_(ctx *Set_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitExpList(ctx *ExpListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitActualParameters(ctx *ActualParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitProcedureCall(ctx *ProcedureCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitStatementSequence(ctx *StatementSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitCaseStatement(ctx *CaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitCase_(ctx *Case_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitCaseLabelList(ctx *CaseLabelListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitLabelRange(ctx *LabelRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitRepeatStatement(ctx *RepeatStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitProcedureDeclaration(ctx *ProcedureDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitProcedureHeading(ctx *ProcedureHeadingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitProcedureBody(ctx *ProcedureBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitDeclarationSequence(ctx *DeclarationSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitFormalParameters(ctx *FormalParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitFPSection(ctx *FPSectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitFormalType(ctx *FormalTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitModule(ctx *ModuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitImportList(ctx *ImportListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseoberonVisitor) VisitImport_(ctx *Import_Context) interface{} {
	return v.VisitChildren(ctx)
}
