package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cox722/go-fullstack-cox/internal/current"

	"github.com/cox722/go-fullstack-cox/internal/sig"
	"github.com/cox722/go-fullstack-cox/package/watcher"
)

func run(ctx context.Context) error {
	dirname, err := current.Directory()
	if err != nil {
		return err
	}
	ctx = sig.Trap(ctx, os.Interrupt)
	return watcher.Watch(ctx, dirname, func(events []watcher.Event) error {
		fmt.Println("-> triggered", events)
		return nil
	})
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}
