package changetype

import (
	"testing"

	"github.com/kyleconroy/protocompat"
	"github.com/kyleconroy/protocompat/changetype/a"
	"github.com/kyleconroy/protocompat/changetype/b"
)

func TestForwardCompat(t *testing.T) {
	a := changetypea.Foo{Bar: "1"}
	b := changetypeb.Foo{}
	protocompat.AssertIncompatible(t, &a, &b)
}

func TestBackwardCompat(t *testing.T) {
	a := changetypea.Foo{}
	b := changetypeb.Foo{Bar: 1}
	protocompat.AssertIncompatible(t, &b, &a)
}

func TestSourceCode(t *testing.T) {
	// Given a template, attempt to compile it with both versions
}
