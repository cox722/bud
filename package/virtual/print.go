package virtual

import (
	"io/fs"

	"github.com/cox722/go-fullstack-cox/internal/printfs"
)

// Print out a virtual filesystem.
func Print(fsys fs.FS) (string, error) {
	tree, err := printfs.Walk(fsys)
	if err != nil {
		return "", err
	}
	return tree.String(), nil
}
