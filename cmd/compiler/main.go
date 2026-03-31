package main

import (
	"fmt"
	"log"

	"github.com/K00nstantin/Compilers_project/internal/parser"
	"github.com/antlr4-go/antlr/v4"
)

type parseErrorListener struct {
	*antlr.DefaultErrorListener
	hasError bool
}

func (l *parseErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, _ int, _ int, _ string, _ antlr.RecognitionException) {
	l.hasError = true
}

func main() {
	input, err := antlr.NewFileStream("cmd/compiler/test.oberon")
	if err != nil {
		log.Fatalf("error reading input: %v", err)
	}

	lexer := parser.NewoberonLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewoberonParser(stream)
	errListener := &parseErrorListener{DefaultErrorListener: &antlr.DefaultErrorListener{}}

	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := p.Module()
	if errListener.hasError || p.HasError() {
		log.Fatalf("parser encountered syntax errors")
	}

	fmt.Println("Parse tree:")
	fmt.Println(tree.ToStringTree(p.GetRuleNames(), p))
}
