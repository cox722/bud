package afsrt

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/cox722/go-fullstack-cox/internal/dag"
	"github.com/cox722/go-fullstack-cox/package/genfs"
	"github.com/cox722/go-fullstack-cox/package/virtual"

	"golang.org/x/sync/errgroup"

	"github.com/cox722/go-fullstack-cox/package/gomod"
	"github.com/cox722/go-fullstack-cox/package/log"
	"github.com/cox722/go-fullstack-cox/package/log/console"
	"github.com/cox722/go-fullstack-cox/package/log/levelfilter"
	"github.com/cox722/go-fullstack-cox/package/remotefs"

	"github.com/cox722/go-fullstack-cox/internal/extrafile"
	"github.com/cox722/go-fullstack-cox/package/socket"
)

func Logger(level string) (log.Log, error) {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	logger := log.New(levelfilter.New(console.New(os.Stderr), lvl))
	return logger, nil
}

// GenFS creates a new filesystem
func GenFS(module *gomod.Module, log log.Log) (genfs.FileSystem, error) {
	fsys := virtual.Exclude(module, func(path string) bool {
		return path == "bud" || strings.HasPrefix(path, "bud/")
	})
	cache, err := dag.Load(log, module.Directory("bud/bud.db"))
	if err != nil {
		return nil, fmt.Errorf("afs: unable to load cache. %w", err)
	}
	return genfs.New(cache, fsys, log), nil
}

// Serve the remote filesystem
func Serve(ctx context.Context, log log.Log, fsys fs.FS, path string) error {
	// First try to load the listener from the parent process.
	ln, err := listen(log, path)
	if err != nil {
		return err
	}
	eg, ctx := errgroup.WithContext(ctx)
	// Handle any immediate errors from remotefs
	eg.Go(func() error {
		return remotefs.Serve(fsys, ln)
	})
	// Any errors in the group will trigger ctx to be canceled, closing the
	// listener. The listener will also be closed if the outside context is
	// canceled.
	eg.Go(func() error {
		<-ctx.Done()
		return ln.Close()
	})
	// Wait for both goroutines to finish
	return eg.Wait()
}

func listen(log log.Log, path string) (socket.Listener, error) {
	files := extrafile.Load("AFS")
	if len(files) > 0 {
		log.Debug("afs: serving from AFS file listener passed in from the parent process")
		return socket.From(files[0])
	}
	ln, err := socket.ListenUp(path, 5)
	if err != nil {
		return nil, err
	}
	log.Debug("afs: serving from", ln.Addr())
	return ln, nil
}
