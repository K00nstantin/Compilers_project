package main

import (
	"fmt"
	"log"

	"github.com/K00nstantin/Compilers_project/internal/ast"
	"github.com/K00nstantin/Compilers_project/internal/parser"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, err := antlr.NewFileStream("cmd/compiler/test.minijava")
	if err != nil {
		log.Fatalf("Error while reading")
	}

	lexer := parser.NewMiniJavaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewMiniJavaParser(stream)

	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := p.Goal()
	if p.HasError() {
		log.Fatalf("parser encountered errors")
	}

	fmt.Println("Дерево разбора:")
	fmt.Println(tree.ToStringTree(p.GetRuleNames(), p))

	builder := parser.NewASTBuilder()
	result := builder.Visit(tree)
	prog, ok := result.(*ast.Program)
	if !ok || prog == nil {
		log.Fatalf("ast builder returned invalid result: %T", result)
	}

	fmt.Println("\nAST:")
	printAST(prog, 0)
}

func printAST(node interface{}, indent int) {
	spaces := func(n int) string {
		s := ""
		for i := 0; i < n; i++ {
			s += "  "
		}
		return s
	}

	switch n := node.(type) {
	case *ast.Program:
		fmt.Printf("%sProgram\n", spaces(indent))
		printAST(n.MainClass, indent+1)
		for _, c := range n.Classes {
			printAST(c, indent+1)
		}
	case *ast.MainClass:
		fmt.Printf("%sMainClass %s\n", spaces(indent), n.Name)
		printAST(n.MainMethod, indent+1)
	case *ast.MainMethod:
		fmt.Printf("%sMainMethod\n", spaces(indent))
		for _, stmt := range n.Body {
			printAST(stmt, indent+1)
		}
	case *ast.ClassDecl:
		fmt.Printf("%sClass %s", spaces(indent), n.Name)
		if n.Parent != "" {
			fmt.Printf(" extends %s", n.Parent)
		}
		fmt.Println()
		for _, f := range n.Fields {
			printAST(f, indent+1)
		}
		for _, m := range n.Methods {
			printAST(m, indent+1)
		}
	case *ast.MethodDecl:
		fmt.Printf("%sMethod %s : %s", spaces(indent), n.Name, n.ReturnType)
		if n.Public {
			fmt.Print(" (public)")
		}
		fmt.Println()
		for _, p := range n.Params {
			printAST(p, indent+1)
		}
		for _, v := range n.Vars {
			printAST(v, indent+1)
		}
		for _, stmt := range n.Body {
			printAST(stmt, indent+1)
		}
	case *ast.VarDecl:
		fmt.Printf("%sVar %s : %s\n", spaces(indent), n.Name, n.Type)
	case *ast.BlockStmt:
		fmt.Printf("%sBlock\n", spaces(indent))
		for _, stmt := range n.Stmts {
			printAST(stmt, indent+1)
		}
	case *ast.IfStmt:
		fmt.Printf("%sIf\n", spaces(indent))
		fmt.Printf("%sCond:\n", spaces(indent+1))
		printAST(n.Cond, indent+2)
		fmt.Printf("%sThen:\n", spaces(indent+1))
		printAST(n.Then, indent+2)
		if n.Else != nil {
			fmt.Printf("%sElse:\n", spaces(indent+1))
			printAST(n.Else, indent+2)
		}
	case *ast.WhileStmt:
		fmt.Printf("%sWhile\n", spaces(indent))
		fmt.Printf("%sCond:\n", spaces(indent+1))
		printAST(n.Cond, indent+2)
		fmt.Printf("%sBody:\n", spaces(indent+1))
		printAST(n.Body, indent+2)
	case *ast.PrintStmt:
		fmt.Printf("%sPrint\n", spaces(indent))
		printAST(n.Expr, indent+1)
	case *ast.AssignStmt:
		fmt.Printf("%sAssign %s =\n", spaces(indent), n.Name)
		printAST(n.Expr, indent+1)
	case *ast.ArrayAssignStmt:
		fmt.Printf("%sArrayAssign %s\n", spaces(indent), n.Name)
		fmt.Printf("%sIndex:\n", spaces(indent+1))
		printAST(n.Index, indent+2)
		fmt.Printf("%sExpr:\n", spaces(indent+1))
		printAST(n.Expr, indent+2)
	case *ast.ReturnStmt:
		fmt.Printf("%sReturn\n", spaces(indent))
		printAST(n.Expr, indent+1)
	case *ast.RecurStmt:
		fmt.Printf("%sRecur\n", spaces(indent))
		fmt.Printf("%sCond:\n", spaces(indent+1))
		printAST(n.Cond, indent+2)
		fmt.Printf("%sArgs:\n", spaces(indent+1))
		for _, arg := range n.Args {
			printAST(arg, indent+2)
		}
		fmt.Printf("%sElse:\n", spaces(indent+1))
		printAST(n.Else, indent+2)
	case *ast.ArrayAccessExpr:
		fmt.Printf("%sArrayAccess\n", spaces(indent))
		fmt.Printf("%sArray:\n", spaces(indent+1))
		printAST(n.Array, indent+2)
		fmt.Printf("%sIndex:\n", spaces(indent+1))
		printAST(n.Index, indent+2)
	case *ast.ArrayLengthExpr:
		fmt.Printf("%sArrayLength\n", spaces(indent))
		printAST(n.Array, indent+1)
	case *ast.MethodCallExpr:
		fmt.Printf("%sMethodCall %s\n", spaces(indent), n.Method)
		fmt.Printf("%sObject:\n", spaces(indent+1))
		printAST(n.Object, indent+2)
		if len(n.Args) > 0 {
			fmt.Printf("%sArgs:\n", spaces(indent+1))
			for _, arg := range n.Args {
				printAST(arg, indent+2)
			}
		}
	case *ast.NegExpr:
		fmt.Printf("%sNeg\n", spaces(indent))
		printAST(n.Expr, indent+1)
	case *ast.NotExpr:
		fmt.Printf("%sNot\n", spaces(indent))
		printAST(n.Expr, indent+1)
	case *ast.ArrayInstantiationExpr:
		fmt.Printf("%sNewIntArray\n", spaces(indent))
		printAST(n.Size, indent+1)
	case *ast.ObjectInstantiationExpr:
		fmt.Printf("%sNewObject %s\n", spaces(indent), n.ClassName)
	case *ast.BinaryOpExpr:
		fmt.Printf("%sBinaryOp %s\n", spaces(indent), n.Op)
		fmt.Printf("%sLeft:\n", spaces(indent+1))
		printAST(n.Left, indent+2)
		fmt.Printf("%sRight:\n", spaces(indent+1))
		printAST(n.Right, indent+2)
	case *ast.IntLitExpr:
		fmt.Printf("%sInt %d\n", spaces(indent), n.Value)
	case *ast.BoolLitExpr:
		fmt.Printf("%sBool %t\n", spaces(indent), n.Value)
	case *ast.IdentifierExpr:
		fmt.Printf("%sId %s\n", spaces(indent), n.Name)
	case *ast.ThisExpr:
		fmt.Printf("%sThis\n", spaces(indent))
	case *ast.ParenExpr:
		fmt.Printf("%sParen\n", spaces(indent))
		printAST(n.Expr, indent+1)
	default:
		fmt.Printf("%sUnknown node type: %T\n", spaces(indent), n)
	}
}
