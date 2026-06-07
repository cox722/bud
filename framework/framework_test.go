package framework_test

import (
	"testing"

	"github.com/cox722/go-fullstack-cox/framework"
	"github.com/cox722/go-fullstack-cox/internal/is"
)

func TestString(t *testing.T) {
	is := is.New(t)
	f := framework.Flag{
		Embed:  true,
		Minify: true,
		Hot:    false,
	}
	flags := f.Flags()
	is.Equal(flags[0], "--embed=true")
	is.Equal(flags[1], "--minify=true")
	is.Equal(flags[2], "--hot=false")
}
