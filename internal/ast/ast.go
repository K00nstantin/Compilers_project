package ast

type Program struct {
	MainClass *MainClass
	Classes   []*ClassDecl
}

type MainClass struct {
	Name       string
	MainMethod *MainMethod
}

type MainMethod struct {
	Body []Stmt
}

type ClassDecl struct {
	Name    string
	Parent  string
	Fields  []*VarDecl
	Methods []*MethodDecl
}

type MethodDecl struct {
	Public     bool
	ReturnType string
	Name       string
	Params     []*VarDecl
	Vars       []*VarDecl
	Body       []Stmt
}

type VarDecl struct {
	Type string
	Name string
}

type Stmt interface {
	stmtNode()
}

type BlockStmt struct {
	Stmts []Stmt
}

func (b *BlockStmt) stmtNode() {}

type IfStmt struct {
	Cond Expression
	Then Stmt
	Else Stmt
}

func (i *IfStmt) stmtNode() {}

type WhileStmt struct {
	Cond Expression
	Body Stmt
}

func (w *WhileStmt) stmtNode() {}

type PrintStmt struct {
	Expr Expression
}

func (p *PrintStmt) stmtNode() {}

type AssignStmt struct {
	Name string
	Expr Expression
}

func (a *AssignStmt) stmtNode() {}

type ArrayAssignStmt struct {
	Name  string
	Index Expression
	Expr  Expression
}

func (a *ArrayAssignStmt) stmtNode() {}

type ReturnStmt struct {
	Expr Expression
}

func (r *ReturnStmt) stmtNode() {}

type RecurStmt struct {
	Cond Expression
	Args []Expression
	Else Expression
}

func (r *RecurStmt) stmtNode() {}

type Expression interface {
	exprNode()
}

type ArrayAccessExpr struct {
	Array Expression
	Index Expression
}

func (a *ArrayAccessExpr) exprNode() {}

type ArrayLengthExpr struct {
	Array Expression
}

func (a *ArrayLengthExpr) exprNode() {}

type MethodCallExpr struct {
	Object Expression
	Method string
	Args   []Expression
}

func (m *MethodCallExpr) exprNode() {}

type NegExpr struct {
	Expr Expression
}

func (n *NegExpr) exprNode() {}

type NotExpr struct {
	Expr Expression
}

func (n *NotExpr) exprNode() {}

type ArrayInstantiationExpr struct {
	Size Expression
}

func (a *ArrayInstantiationExpr) exprNode() {}

type ObjectInstantiationExpr struct {
	ClassName string
}

func (o *ObjectInstantiationExpr) exprNode() {}

type BinaryOpExpr struct {
	Left  Expression
	Op    string
	Right Expression
}

func (b *BinaryOpExpr) exprNode() {}

type IntLitExpr struct {
	Value int
}

func (i *IntLitExpr) exprNode() {}

type BoolLitExpr struct {
	Value bool
}

func (b *BoolLitExpr) exprNode() {}

type IdentifierExpr struct {
	Name string
}

func (i *IdentifierExpr) exprNode() {}

type ThisExpr struct{}

func (t *ThisExpr) exprNode() {}

type ParenExpr struct {
	Expr Expression
}

func (p *ParenExpr) exprNode() {}
