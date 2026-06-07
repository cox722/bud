package public

import (
	"errors"
	"io/fs"
	"path"
	"strings"

	"github.com/cox722/go-fullstack-cox/package/valid"
	"github.com/cox722/go-fullstack-cox/runtime/transpiler"

	"github.com/cox722/go-fullstack-cox/framework"
	"github.com/cox722/go-fullstack-cox/package/finder"

	"github.com/cox722/go-fullstack-cox/internal/bail"
	"github.com/cox722/go-fullstack-cox/package/imports"
)

func Load(fsys fs.FS, flag *framework.Flag) (*State, error) {
	loader := &loader{
		fsys:    fsys,
		flag:    flag,
		imports: imports.New(),
	}
	return loader.Load()
}

type loader struct {
	bail.Struct
	flag    *framework.Flag
	fsys    fs.FS
	imports *imports.Set
}

// Load the command state
func (l *loader) Load() (state *State, err error) {
	defer l.Recover(&err)
	paths, err := finder.Find(l.fsys, "public/**", func(fullpath string, isDir bool) (entries []string) {
		if isDir {
			return nil
		}
		if valid.PublicFile(path.Base(fullpath)) {
			entries = append(entries, fullpath)
		}
		return entries
	})
	if err != nil {
		return nil, err
	} else if len(paths) == 0 {
		return nil, fs.ErrNotExist
	}
	state = new(State)
	// Load the files from paths
	state.Files = l.loadFiles(paths)
	// Default imports
	l.imports.AddNamed("virtual", "github.com/cox722/go-fullstack-cox/package/virtual")
	l.imports.AddNamed("publicrt", "github.com/cox722/go-fullstack-cox/framework/public/publicrt")
	l.imports.AddNamed("router", "github.com/cox722/go-fullstack-cox/package/router")
	l.imports.AddNamed("http", "net/http")
	l.imports.AddNamed("fs", "io/fs")
	// Add the imports
	state.Imports = l.imports.List()
	return state, nil
}

func (l *loader) loadFiles(paths []string) (files []*File) {
	for _, path := range paths {
		files = append(files, l.loadFile(path))
	}
	return files
}

func (l *loader) loadFile(fpath string) *File {
	file := new(File)
	file.Path = fpath
	file.Route = strings.TrimPrefix(fpath, "public")
	if l.flag.Embed {
		data, err := transpiler.TranspileFile(l.fsys, fpath, path.Ext(fpath))
		if err != nil {
			if !errors.Is(err, fs.ErrNotExist) {
				l.Bail(err)
			}
			data, err = fs.ReadFile(l.fsys, fpath)
			if err != nil {
				l.Bail(err)
			}
		}
		file.Data = data
	}
	return file
}
