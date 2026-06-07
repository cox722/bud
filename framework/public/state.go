package public

import (
	"github.com/cox722/go-fullstack-cox/internal/embed"
	"github.com/cox722/go-fullstack-cox/package/imports"
)

type State struct {
	Imports []*imports.Import
	Files   []*File
}

type File struct {
	Path  string
	Route string
	Data  embed.Data
}
