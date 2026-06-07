package app

import (
	"github.com/cox722/go-fullstack-cox/framework"
	"github.com/cox722/go-fullstack-cox/package/di"
	"github.com/cox722/go-fullstack-cox/package/imports"
)

type State struct {
	Imports  []*imports.Import
	Provider *di.Provider
	Flag     *framework.Flag
}
