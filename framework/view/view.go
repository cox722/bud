package view

import (
	_ "embed"

	"github.com/cox722/go-fullstack-cox/framework"
	"github.com/cox722/go-fullstack-cox/framework/transform/transformrt"
	"github.com/cox722/go-fullstack-cox/package/genfs"
	"github.com/cox722/go-fullstack-cox/package/gomod"
	"github.com/cox722/go-fullstack-cox/package/gotemplate"
)

//go:embed view.gotext
var template string

var generator = gotemplate.MustParse("framework/view/view.gotext", template)

// Generate the view from state
func Generate(state *State) ([]byte, error) {
	return generator.Generate(state)
}

func New(module *gomod.Module, transform *transformrt.Map, flag *framework.Flag) *Generator {
	return &Generator{
		flag:      flag,
		module:    module,
		transform: transform,
	}
}

type Generator struct {
	flag      *framework.Flag
	module    *gomod.Module
	transform *transformrt.Map
}

func (c *Generator) GenerateFile(fsys genfs.FS, file *genfs.File) error {
	state, err := Load(fsys, c.module, c.transform, c.flag)
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
