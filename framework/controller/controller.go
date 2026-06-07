package controller

import (

	// Embed templates

	_ "embed"
	"fmt"

	"github.com/cox722/go-fullstack-cox/package/di"
	"github.com/cox722/go-fullstack-cox/package/genfs"
	"github.com/cox722/go-fullstack-cox/package/gomod"
	"github.com/cox722/go-fullstack-cox/package/gotemplate"
	"github.com/cox722/go-fullstack-cox/package/parser"
)

//go:embed controller.gotext
var template string

var generator = gotemplate.MustParse("framework/controller/controller.gotext", template)

// Generate the controller template from state
func Generate(state *State) ([]byte, error) {
	return generator.Generate(state)
}

// New controller generator
func New(injector *di.Injector, module *gomod.Module, parser *parser.Parser) *Generator {
	return &Generator{injector, module, parser}
}

// Generator for controllers
type Generator struct {
	injector *di.Injector
	module   *gomod.Module
	parser   *parser.Parser
}

func (g *Generator) GenerateFile(fsys genfs.FS, file *genfs.File) error {
	state, err := Load(fsys, g.injector, g.module, g.parser)
	if err != nil {
		return fmt.Errorf("controller: unable to load. %w", err)
	}
	code, err := Generate(state)
	if err != nil {
		return err
	}
	file.Data = code
	return nil
}
