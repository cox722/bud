package app

import (
	"fmt"
	"io/fs"

	"github.com/cox722/go-fullstack-cox/framework"
	"github.com/cox722/go-fullstack-cox/internal/bail"
	"github.com/cox722/go-fullstack-cox/package/di"
	"github.com/cox722/go-fullstack-cox/package/gomod"
	"github.com/cox722/go-fullstack-cox/package/imports"
	"github.com/cox722/go-fullstack-cox/package/vfs"
)

func Load(fsys fs.FS, injector *di.Injector, module *gomod.Module, flag *framework.Flag) (*State, error) {
	if err := vfs.Exist(fsys, "bud/internal/web/web.go"); err != nil {
		return nil, err
	}
	return (&loader{
		fsys:     fsys,
		injector: injector,
		module:   module,
		flag:     flag,
		imports:  imports.New(),
	}).Load()
}

type loader struct {
	fsys     fs.FS
	injector *di.Injector
	module   *gomod.Module
	flag     *framework.Flag

	imports *imports.Set
	bail.Struct
}

func (l *loader) Load() (state *State, err error) {
	defer l.Recover2(&err, "app: unable to load state")
	state = new(State)
	l.imports.AddStd("os", "context", "errors")
	l.imports.AddNamed("commander", "github.com/cox722/go-fullstack-cox/package/commander")
	l.imports.AddNamed("console", "github.com/cox722/go-fullstack-cox/package/log/console")
	l.imports.AddNamed("levelfilter", "github.com/cox722/go-fullstack-cox/package/log/levelfilter")
	l.imports.AddNamed("log", "github.com/cox722/go-fullstack-cox/package/log")
	l.imports.AddNamed("budhttp", "github.com/cox722/go-fullstack-cox/package/budhttp")
	l.imports.Add(l.module.Import("bud/internal/web"))
	state.Provider = l.loadProvider()
	state.Flag = l.flag
	state.Imports = l.imports.List()
	return state, nil
}

func (l *loader) loadProvider() *di.Provider {
	jsVM := di.ToType("github.com/cox722/go-fullstack-cox/package/js", "VM")
	// TODO: the public generator should be able to configure this
	publicFS := di.ToType("github.com/cox722/go-fullstack-cox/framework/public/publicrt", "FS")
	viewFS := di.ToType("github.com/cox722/go-fullstack-cox/framework/view/viewrt", "FS")
	transpilerFS := di.ToType("github.com/cox722/go-fullstack-cox/runtime/transpiler", "FS")
	fn := &di.Function{
		Name:    "loadWeb",
		Imports: l.imports,
		Target:  l.module.Import("bud", "program"),
		Params: []*di.Param{
			{Import: "github.com/cox722/go-fullstack-cox/package/log", Type: "Log"},
			{Import: "github.com/cox722/go-fullstack-cox/package/gomod", Type: "*Module"},
			{Import: "github.com/cox722/go-fullstack-cox/package/budhttp", Type: "Client"},
			{Import: "github.com/cox722/go-fullstack-cox/package/remotefs", Type: "*Client"},
			{Import: "context", Type: "Context"},
		},
		Results: []di.Dependency{
			di.ToType(l.module.Import("bud/internal/web"), "*Server"),
			&di.Error{},
		},
		Aliases: di.Aliases{
			transpilerFS: di.ToType("github.com/cox722/go-fullstack-cox/package/remotefs", "*Client"),
			publicFS:     di.ToType("github.com/cox722/go-fullstack-cox/runtime/transpiler", "*Proxy"),
			viewFS:       di.ToType("github.com/cox722/go-fullstack-cox/package/remotefs", "*Client"),
			jsVM:         di.ToType("github.com/cox722/go-fullstack-cox/package/budhttp", "Client"),
		},
	}
	if l.flag.Embed {
		fn.Aliases[jsVM] = di.ToType("github.com/cox722/go-fullstack-cox/package/js/v8", "*VM")
		fn.Aliases[publicFS] = di.ToType(l.module.Import("bud/internal/web/public"), "FS")
		fn.Aliases[viewFS] = di.ToType(l.module.Import("bud/internal/web/view"), "FS")
	}
	provider, err := l.injector.Wire(fn)
	if err != nil {
		// Intentionally don't wrap this error, it gets swallowed up too easily
		l.Bail(fmt.Errorf("app: unable to wire. %s", err))
	}
	// Add generated imports
	for _, imp := range provider.Imports {
		l.imports.AddNamed(imp.Name, imp.Path)
	}
	return provider
}
