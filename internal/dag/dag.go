package dag

import "github.com/cox722/go-fullstack-cox/package/virtual"

type Cache interface {
	Get(path string) (*virtual.File, error)
	Set(path string, file *virtual.File) error
	Link(from string, toPatterns ...string) error
	Delete(paths ...string) error
	Reset() error
	Close() error
}
