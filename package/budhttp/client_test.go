package budhttp_test

import (
	"context"
	"testing"

	"github.com/cox722/go-fullstack-cox/framework"
	"github.com/cox722/go-fullstack-cox/internal/dag"

	"github.com/cox722/go-fullstack-cox/package/genfs"
	"github.com/cox722/go-fullstack-cox/package/socket"

	"github.com/cox722/go-fullstack-cox/package/log/testlog"

	"github.com/cox722/go-fullstack-cox/framework/transform/transformrt"
	"github.com/cox722/go-fullstack-cox/framework/view/dom"
	"github.com/cox722/go-fullstack-cox/framework/view/nodemodules"
	"github.com/cox722/go-fullstack-cox/framework/view/ssr"
	"github.com/cox722/go-fullstack-cox/internal/is"
	"github.com/cox722/go-fullstack-cox/internal/pubsub"
	"github.com/cox722/go-fullstack-cox/package/budhttp"
	"github.com/cox722/go-fullstack-cox/package/budhttp/budsvr"
	"github.com/cox722/go-fullstack-cox/package/gomod"
	v8 "github.com/cox722/go-fullstack-cox/package/js/v8"
	"github.com/cox722/go-fullstack-cox/package/svelte"
	"github.com/cox722/go-fullstack-cox/package/testdir"
)

func loadServer(bus pubsub.Client, dir string) (*budsvr.Server, error) {
	flag := new(framework.Flag)
	log := testlog.New()
	vm, err := v8.Load()
	if err != nil {
		return nil, err
	}
	svelteCompiler, err := svelte.Load(vm)
	if err != nil {
		return nil, err
	}
	transforms, err := transformrt.Default(log, svelteCompiler)
	if err != nil {
		return nil, err
	}
	module, err := gomod.Find(dir)
	if err != nil {
		return nil, err
	}
	budln, err := socket.Listen(":0")
	if err != nil {
		return nil, err
	}
	cache, err := dag.Load(log, ":memory:")
	if err != nil {
		return nil, err
	}
	gfs := genfs.New(cache, module, log)
	gfs.FileServer("bud/view", dom.New(module, transforms))
	gfs.FileServer("bud/node_modules", nodemodules.New(module))
	gfs.FileGenerator("bud/view/_ssr.js", ssr.New(module, transforms))
	return budsvr.New(budln, bus, flag, gfs, log, vm), nil
}

func TestEvents(t *testing.T) {
	ctx := context.Background()
	is := is.New(t)
	log := testlog.New()
	td, err := testdir.Load()
	is.NoErr(err)
	is.NoErr(td.Write(ctx))
	ps := pubsub.New()
	server, err := loadServer(ps, td.Directory())
	is.NoErr(err)
	server.Start(ctx)
	defer server.Close()
	client, err := budhttp.Load(log, server.Address())
	is.NoErr(err)
	sub := ps.Subscribe("ready")
	defer sub.Close()
	err = client.Publish("ready", []byte("test"))
	is.NoErr(err)
	select {
	case payload := <-sub.Wait():
		is.Equal(string(payload), "test")
	default:
		t.Fatalf("missing event")
	}
}

func TestScript(t *testing.T) {
	ctx := context.Background()
	is := is.New(t)
	log := testlog.New()
	td, err := testdir.Load()
	is.NoErr(err)
	is.NoErr(td.Write(ctx))
	ps := pubsub.New()
	server, err := loadServer(ps, td.Directory())
	is.NoErr(err)
	server.Start(ctx)
	defer server.Close()
	client, err := budhttp.Load(log, server.Address())
	is.NoErr(err)
	err = client.Script("script.js", "function a() { return 1 }")
	is.NoErr(err)
	err = client.Script("script.js", "function b() { return 1")
	is.True(err != nil)
	is.In(err.Error(), "SyntaxError: Unexpected end of input")
}

func TestScriptEval(t *testing.T) {
	ctx := context.Background()
	is := is.New(t)
	log := testlog.New()
	td, err := testdir.Load()
	is.NoErr(err)
	is.NoErr(td.Write(ctx))
	ps := pubsub.New()
	server, err := loadServer(ps, td.Directory())
	is.NoErr(err)
	server.Start(ctx)
	defer server.Close()
	client, err := budhttp.Load(log, server.Address())
	is.NoErr(err)
	err = client.Script("script.js", "function a() { return 1 }")
	is.NoErr(err)
	val, err := client.Eval("script.js", "a()")
	is.NoErr(err)
	is.Equal(val, "1")
}
