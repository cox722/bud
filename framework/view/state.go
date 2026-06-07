package view

import (
	"github.com/cox722/go-fullstack-cox/internal/embed"
	"github.com/cox722/go-fullstack-cox/package/imports"
)

type State struct {
	Imports []*imports.Import
	Routes  []string
	Embeds  []*embed.File
}
