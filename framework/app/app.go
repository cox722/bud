package app

import (
	_ "embed"

	"github.com/cox722/go-fullstack-cox/framework"
	"github.com/cox722/go-fullstack-cox/package/di"
	"github.com/cox722/go-fullstack-cox/package/genfs"

	"github.com/cox722/go-fullstack-cox/package/gomod"
	"github.com/cox722/go-fullstack-cox/package/gotemplate"
)

//go:embed app.gotext
var template string

var generator = gotemplate.MustParse("framework/app/app.gotext", template)

func Generate(state *State) ([]byte, error) {
	return generator.Generate(state)
}

func New(injector *di.Injector, module *gomod.Module, flag *framework.Flag) *Generator {
	return &Generator{flag, injector, module}
}

type Generator struct {
	flag     *framework.Flag
	injector *di.Injector
	module   *gomod.Module
}

func (g *Generator) GenerateFile(fsys genfs.FS, file *genfs.File) error {
	state, err := Load(fsys, g.injector, g.module, g.flag)
	if err != nil {
		return err
	}
	code, err := Generate(state)
	if err != nil {
		return err
	}
	file.Data = code
	return nil
}
