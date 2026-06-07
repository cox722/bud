package cli

import (
	"context"

	"github.com/cox722/go-fullstack-cox/internal/once"
	"github.com/cox722/go-fullstack-cox/package/commander"
)

type Custom struct {
	Closer *once.Closer
	Help   bool
	Args   []string
}

func (c *CLI) Custom(ctx context.Context, in *Custom) error {
	if in.Help {
		return commander.Usage()
	}
	return commander.Usage()
}
