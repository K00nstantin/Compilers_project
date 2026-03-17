package visitor

import (
	"fmt"

	"github.com/K00nstantin/Compilers_project/internal/parser"
	"github.com/antlr4-go/antlr/v4"
)

type MyVisitor struct {
	*parser.BaseMiniJavaVisitor
	Variables []string
}

func NewVisitor() *MyVisitor {
	return &MyVisitor{
		BaseMiniJavaVisitor: &parser.BaseMiniJavaVisitor{},
		Variables:           make([]string, 0),
	}
}

type MyListener struct {
	*parser.BaseMiniJavaListener
	ruleNames []string
}

func NewMyListener(ruleNames []string) *MyListener {
	return &MyListener{
		BaseMiniJavaListener: &parser.BaseMiniJavaListener{},
		ruleNames:            ruleNames,
	}
}

func (l *MyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("Enter rule, context type: %T\n", ctx)
	fmt.Printf("Text: %q\n", ctx.GetText())
}
