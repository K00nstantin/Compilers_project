package ast

type Module struct {
	Name         string
	EndName      string
	Imports      []*Import
	Declarations *DeclarationBlock
	Body         []Stmt
}

type Import struct {
	Alias  string
	Module string
}

type DeclarationBlock struct {
	Consts     []*ConstDecl
	Types      []*TypeDecl
	Vars       []*VarDecl
	Procedures []*ProcedureDecl
}

type IdentDef struct {
	Name     string
	Exported bool
}

type ConstDecl struct {
	Name  IdentDef
	Value Expr
}

type TypeDecl struct {
	Name IdentDef
	Type TypeExpr
}

type VarDecl struct {
	Names []IdentDef
	Type  TypeExpr
}

type ProcedureDecl struct {
	Name         IdentDef
	Signature    *ProcedureSignature
	Declarations *DeclarationBlock
	Body         []Stmt
	ReturnExpr   Expr
	EndName      string
	NestingLevel int            // 0 – модуль, 1 – вложенная в модуль, и т.д.
	Parent       *ProcedureDecl // внешняя процедура (nil для уровня модуля)
}

type ProcedureSignature struct {
	Params     []*ParamSection
	ReturnType string
}

type ParamSection struct {
	ByRef bool
	Names []string
	Type  TypeExpr
}

type TypeExpr interface {
	typeExprNode()
}

type NamedType struct {
	Name string
}

func (*NamedType) typeExprNode() {}

type ArrayType struct {
	Lengths []Expr
	Elem    TypeExpr
}

func (*ArrayType) typeExprNode() {}

type FieldDecl struct {
	Names []IdentDef
	Type  TypeExpr
}

type RecordType struct {
	Base       string
	Fields     []*FieldDecl
	FieldOrder []string
}

func (*RecordType) typeExprNode() {}

type PointerType struct {
	Target TypeExpr
}

func (*PointerType) typeExprNode() {}

type ProcedureType struct {
	Signature *ProcedureSignature
}

func (*ProcedureType) typeExprNode() {}

type Stmt interface {
	stmtNode()
}

type AssignmentStmt struct {
	Target *DesignatorExpr
	Value  Expr
}

func (*AssignmentStmt) stmtNode() {}

type ProcedureCallStmt struct {
	Call *CallExpr
}

func (*ProcedureCallStmt) stmtNode() {}

type IfBranch struct {
	Cond Expr
	Body []Stmt
}

type IfStmt struct {
	Branches []*IfBranch
	ElseBody []Stmt
}

func (*IfStmt) stmtNode() {}

type CaseLabel struct {
	From Expr
	To   Expr
}

type CaseBranch struct {
	Labels []*CaseLabel
	Body   []Stmt
}

type CaseStmt struct {
	Expr     Expr
	Branches []*CaseBranch
}

func (*CaseStmt) stmtNode() {}

type WhileStmt struct {
	Branches []*IfBranch
}

func (*WhileStmt) stmtNode() {}

type RepeatStmt struct {
	Body  []Stmt
	Until Expr
}

func (*RepeatStmt) stmtNode() {}

type ForStmt struct {
	Var   string
	From  Expr
	To    Expr
	By    Expr
	Body  []Stmt
	HasBy bool
}

func (*ForStmt) stmtNode() {}

type Expr interface {
	exprNode()
}

type BinaryExpr struct {
	Left  Expr
	Op    string
	Right Expr
}

func (*BinaryExpr) exprNode() {}

type UnaryExpr struct {
	Op   string
	Expr Expr
}

func (*UnaryExpr) exprNode() {}

type NumberExpr struct {
	Text   string
	IsReal bool
	IsHex  bool
}

func (*NumberExpr) exprNode() {}

type StringExpr struct {
	Value string
}

func (*StringExpr) exprNode() {}

type BoolExpr struct {
	Value bool
}

func (*BoolExpr) exprNode() {}

type NilExpr struct{}

func (*NilExpr) exprNode() {}

type SetExpr struct {
	Elements []*SetElement
}

func (*SetExpr) exprNode() {}

type SetElement struct {
	From Expr
	To   Expr
}

type CallExpr struct {
	Callee Expr
	Args   []Expr
}

func (*CallExpr) exprNode() {}

type QualIdent struct {
	Module string
	Name   string
}

type Selector struct {
	Field string
	Index []Expr
	Deref bool
	Type  string
}

type DesignatorExpr struct {
	Base      QualIdent
	Selectors []Selector
}

func (*DesignatorExpr) exprNode() {}

type IsExpr struct {
	Expr     Expr   // левая часть (выражение)
	TypeName string // полное имя типа (например, "Point" или "Module.Point")
}

func (*IsExpr) exprNode() {}
