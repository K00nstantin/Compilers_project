package codegen

import (
	"github.com/K00nstantin/Compilers_project/internal/ast"
	"github.com/llir/llvm/ir"
)

type Generator struct {
	module *ir.Module
}

func NewGenerator(moduleName string) *Generator {
	m := ir.NewModule()
	m.SourceFilename = moduleName + ".mod"
	return &Generator{module: m}
}

func (g *Generator) Generate(astMod *ast.Module) *ir.Module {
	return g.module
}
