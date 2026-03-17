package main

import (
	"fmt"
	"log"

	"github.com/K00nstantin/Compilers_project/internal/parser"
	"github.com/K00nstantin/Compilers_project/internal/visitor"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input, err := antlr.NewFileStream("/home/konstantin/go/Compilers_project/cmd/compiler/test.minijava")
	if err != nil {
		log.Fatalf("Error while reading")
	}

	lexer := parser.NewMiniJavaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewMiniJavaParser(stream)

	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := p.Goal()

	fmt.Println("Дерево разбора:")
	fmt.Println(tree.ToStringTree(p.GetRuleNames(), p))

	ruleNames := p.GetRuleNames()
	listener := visitor.NewMyListener(ruleNames)
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	myVisitor := visitor.NewVisitor()

	tree.Accept(myVisitor)
	fmt.Println("\nСобранные переменные:")
	for _, v := range myVisitor.Variables {
		fmt.Println(v)
	}
}
