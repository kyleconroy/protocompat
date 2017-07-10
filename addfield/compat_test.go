package addfield

import (
	"testing"

	"github.com/kyleconroy/protocompat"
	"github.com/kyleconroy/protocompat/addfield/a"
	"github.com/kyleconroy/protocompat/addfield/b"
)

func TestForwardCompat(t *testing.T) {
	a := addfielda.Foo{Bar: "1"}
	b := addfieldb.Foo{}
	protocompat.AssertCompatible(t, &a, &b)
}

func TestBackwardCompat(t *testing.T) {
	a := addfielda.Foo{}
	b := addfieldb.Foo{Bar: "1", Bat: "2"}
	protocompat.AssertCompatible(t, &b, &a)
}

func TestSourceCode(t *testing.T) {
	// Given a template, attempt to compile it with both versions
}
