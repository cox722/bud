package pluginmod_test

import (
	"context"
	"fmt"
	"io/fs"
	"testing"

	"github.com/cox722/go-fullstack-cox/internal/mergefs"

	"github.com/cox722/go-fullstack-cox/internal/is"
	"github.com/cox722/go-fullstack-cox/package/gomod"
	"github.com/cox722/go-fullstack-cox/package/pluginmod"
	"github.com/cox722/go-fullstack-cox/package/testdir"
)

func TestGlob(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	td, err := testdir.Load()
	is.NoErr(err)
	td.Modules["github.com/cox722/go-fullstack-cox-test-plugin"] = "v0.0.9"
	td.Modules["github.com/cox722/go-fullstack-cox-test-nested-plugin"] = "v0.0.5"
	err = td.Write(ctx)
	is.NoErr(err)
	module, err := gomod.Find(td.Directory())
	is.NoErr(err)
	plugins, err := pluginmod.Glob(module, "public")
	is.NoErr(err)
	is.Equal(len(plugins), 2)
	is.Equal(plugins[0].Import(), "github.com/cox722/go-fullstack-cox-test-nested-plugin")
	is.Equal(plugins[1].Import(), "github.com/cox722/go-fullstack-cox-test-plugin")
	plugins, err = pluginmod.Glob(module, "view")
	is.NoErr(err)
	is.Equal(len(plugins), 1)
	is.Equal(plugins[0].Import(), "github.com/cox722/go-fullstack-cox-test-plugin")
}

func TestAppFirst(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	td, err := testdir.Load()
	is.NoErr(err)
	favicon := []byte{0x00, 0x00, 0x01}
	td.Bytes["public/favicon.ico"] = favicon
	td.Modules["github.com/cox722/go-fullstack-cox-test-plugin"] = "v0.0.9"
	td.Modules["github.com/cox722/go-fullstack-cox-test-nested-plugin"] = "v0.0.5"
	err = td.Write(ctx)
	is.NoErr(err)
	module, err := gomod.Find(td.Directory())
	is.NoErr(err)
	plugins, err := pluginmod.Glob(module, "public")
	is.NoErr(err)
	is.Equal(len(plugins), 3)
	is.Equal(plugins[0].Import(), "app.com")
	is.Equal(plugins[1].Import(), "github.com/cox722/go-fullstack-cox-test-nested-plugin")
	is.Equal(plugins[2].Import(), "github.com/cox722/go-fullstack-cox-test-plugin")

	// Try merging
	fileSystems := make([]fs.FS, len(plugins))
	for i, plugin := range plugins {
		fileSystems[i] = plugin
	}
	fsys := mergefs.Merge(fileSystems...)
	publicfs, err := fs.Sub(fsys, "public")
	is.NoErr(err)
	code, err := fs.ReadFile(publicfs, "favicon.ico")
	is.NoErr(err)
	is.Equal(code, []byte{0x00, 0x00, 0x01})
	code, err = fs.ReadFile(publicfs, "tailwind/preflight.css")
	is.NoErr(err)
	fmt.Println(string(code))
	is.Equal(string(code), `/* conflicting preflight */`)
}
