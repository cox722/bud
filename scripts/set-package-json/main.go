package main

import (
	"os"
	"path/filepath"

	"github.com/cox722/go-fullstack-cox/internal/npm"
	"github.com/cox722/go-fullstack-cox/internal/versions"
	"github.com/cox722/go-fullstack-cox/package/gomod"
	"github.com/cox722/go-fullstack-cox/package/log/console"
)

func main() {
	if err := run(); err != nil {
		console.Error(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	dir, err := gomod.Absolute(".")
	if err != nil {
		return err
	}
	// Update the dependencies in ./cox722/package.json
	if err := npm.Set(filepath.Join(dir, "cox722"), map[string]string{
		"dependencies.svelte":              versions.Svelte,
		"dependencies.react":               versions.React,
		"dependencies.react-dom":           versions.React,
		"devDependencies.@types/react":     versions.React,
		"devDependencies.@types/react-dom": versions.React,
	}); err != nil {
		return err
	}
	// Update the dependencies in .
	if err := npm.Set(dir, map[string]string{
		"devDependencies.svelte":           versions.Svelte,
		"devDependencies.react":            versions.React,
		"devDependencies.react-dom":        versions.React,
		"devDependencies.@types/react":     versions.React,
		"devDependencies.@types/react-dom": versions.React,
	}); err != nil {
		return err
	}
	return nil
}
