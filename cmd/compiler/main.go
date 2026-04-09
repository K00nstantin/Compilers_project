package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/K00nstantin/Compilers_project/internal/ast"
	"github.com/K00nstantin/Compilers_project/internal/parser"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, err := antlr.NewFileStream("cmd/compiler/test.oberon")
	if err != nil {
		log.Fatalf("error reading input: %v", err)
	}

	lexer := parser.NewoberonLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewoberonParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := p.Module()
	if p.HasError() {
		log.Fatalf("parser encountered syntax errors")
	}

	fmt.Println("Parse tree:")
	fmt.Println(tree.ToStringTree(p.GetRuleNames(), p))

	builder := parser.NewASTBuilder()
	result := builder.Visit(tree)
	mod, ok := result.(*ast.Module)
	if !ok || mod == nil {
		log.Fatalf("ast builder returned invalid result: %T", result)
	}
	fmt.Println("\n AST builder\n")
	dumpModule(mod, "")

}

func dumpModule(m *ast.Module, indent string) {
	fmt.Printf("%sModule name=%q endName=%q\n", indent, m.Name, m.EndName)
	next := indent + "  "
	if len(m.Imports) > 0 {
		fmt.Printf("%sImports:\n", indent)
		for _, imp := range m.Imports {
			fmt.Printf("%s  %q <- %q\n", indent, imp.Alias, imp.Module)
		}
	}
	if m.Declarations != nil {
		fmt.Printf("%sDeclarations:\n", indent)
		for _, c := range m.Declarations.Consts {
			fmt.Printf("%s  Const %s = %s\n", indent, formatIdentDef(c.Name), formatExpr(c.Value))
		}
		for _, t := range m.Declarations.Types {
			fmt.Printf("%s  Type %s = %s\n", indent, formatIdentDef(t.Name), formatTypeExpr(t.Type))
		}
		for _, v := range m.Declarations.Vars {
			names := make([]string, len(v.Names))
			for i, id := range v.Names {
				names[i] = formatIdentDef(id)
			}
			fmt.Printf("%s  Var %s: %s\n", indent, strings.Join(names, ", "), formatTypeExpr(v.Type))
		}
		for _, proc := range m.Declarations.Procedures {
			dumpProcedure(proc, next)
		}
	}
	if len(m.Body) > 0 {
		fmt.Printf("%sBody:\n", indent)
		for _, s := range m.Body {
			dumpStmt(s, next)
		}
	} else {
		fmt.Printf("%sBody: (empty)\n", indent)
	}
}

func dumpProcedure(p *ast.ProcedureDecl, indent string) {
	fmt.Printf("%sProcedure %s", indent, formatIdentDef(p.Name))
	if p.Signature != nil {
		fmt.Printf(" params=%d returnType=%q", len(p.Signature.Params), p.Signature.ReturnType)
	}
	fmt.Println()
	next := indent + "  "
	if p.Declarations != nil && (len(p.Declarations.Consts)+len(p.Declarations.Types)+len(p.Declarations.Vars)+len(p.Declarations.Procedures) > 0) {
		fmt.Printf("%s  (nested decls omitted)\n", indent)
	}
	for _, s := range p.Body {
		dumpStmt(s, next)
	}
	if p.ReturnExpr != nil {
		fmt.Printf("%s  RETURN %s\n", indent, formatExpr(p.ReturnExpr))
	}
	if p.EndName != "" {
		fmt.Printf("%s  endName=%q\n", indent, p.EndName)
	}
}

func dumpStmt(s ast.Stmt, indent string) {
	switch x := s.(type) {
	case *ast.AssignmentStmt:
		fmt.Printf("%sAssign %s := %s\n", indent, formatExpr(x.Target), formatExpr(x.Value))
	case *ast.ProcedureCallStmt:
		fmt.Printf("%sCall %s(%d args)\n", indent, formatExpr(x.Call.Callee), len(x.Call.Args))
	case *ast.IfStmt:
		fmt.Printf("%sIf (%d branches, else=%v)\n", indent, len(x.Branches), len(x.ElseBody) > 0)
	case *ast.WhileStmt:
		fmt.Printf("%sWhile (%d branches)\n", indent, len(x.Branches))
	case *ast.RepeatStmt:
		fmt.Printf("%sRepeat until %s\n", indent, formatExpr(x.Until))
	case *ast.ForStmt:
		by := ""
		if x.HasBy {
			by = fmt.Sprintf(" BY %s", formatExpr(x.By))
		}
		fmt.Printf("%sFor %s := %s TO %s%s\n", indent, x.Var, formatExpr(x.From), formatExpr(x.To), by)
	case *ast.CaseStmt:
		fmt.Printf("%sCase on %s (%d branches)\n", indent, formatExpr(x.Expr), len(x.Branches))
	default:
		fmt.Printf("%sStmt %T\n", indent, s)
	}
}

func formatIdentDef(id ast.IdentDef) string {
	if id.Exported {
		return id.Name + "*"
	}
	return id.Name
}

func formatQual(q ast.QualIdent) string {
	if q.Module != "" {
		return q.Module + "." + q.Name
	}
	return q.Name
}

func formatTypeExpr(t ast.TypeExpr) string {
	switch x := t.(type) {
	case *ast.NamedType:
		return x.Name
	case *ast.ArrayType:
		parts := make([]string, len(x.Lengths))
		for i, e := range x.Lengths {
			parts[i] = formatExpr(e)
		}
		return fmt.Sprintf("ARRAY %s OF %s", strings.Join(parts, ", "), formatTypeExpr(x.Elem))
	case *ast.RecordType:
		return fmt.Sprintf("RECORD(...%d fields)", len(x.Fields))
	case *ast.PointerType:
		return "POINTER TO " + formatTypeExpr(x.Target)
	case *ast.ProcedureType:
		return "PROCEDURE(...)"
	default:
		return fmt.Sprintf("%T", t)
	}
}

func formatExpr(e ast.Expr) string {
	if e == nil {
		return "<nil>"
	}
	switch x := e.(type) {
	case *ast.NumberExpr:
		return "num(" + x.Text + ")"
	case *ast.StringExpr:
		return fmt.Sprintf("str(%q)", x.Value)
	case *ast.BoolExpr:
		return fmt.Sprintf("%v", x.Value)
	case *ast.NilExpr:
		return "NIL"
	case *ast.BinaryExpr:
		return fmt.Sprintf("(%s %s %s)", formatExpr(x.Left), x.Op, formatExpr(x.Right))
	case *ast.UnaryExpr:
		return fmt.Sprintf("(%s%s)", x.Op, formatExpr(x.Expr))
	case *ast.DesignatorExpr:
		s := formatQual(x.Base)
		for _, sel := range x.Selectors {
			switch {
			case sel.Field != "":
				s += "." + sel.Field
			case len(sel.Index) > 0:
				parts := make([]string, len(sel.Index))
				for i, ix := range sel.Index {
					parts[i] = formatExpr(ix)
				}
				s += "[" + strings.Join(parts, ", ") + "]"
			case sel.Deref:
				s += "^"
			case sel.Type != "":
				s += "(" + sel.Type + ")"
			}
		}
		return s
	case *ast.CallExpr:
		args := make([]string, len(x.Args))
		for i, a := range x.Args {
			args[i] = formatExpr(a)
		}
		return fmt.Sprintf("%s(%s)", formatExpr(x.Callee), strings.Join(args, ", "))
	case *ast.SetExpr:
		return fmt.Sprintf("SET(%d elems)", len(x.Elements))
	default:
		return fmt.Sprintf("%T", e)
	}
}
