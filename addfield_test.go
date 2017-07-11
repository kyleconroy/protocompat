package protocompat

import (
	"testing"

	"github.com/kyleconroy/protocompat/one"
	"github.com/kyleconroy/protocompat/two"
)

func TestAddFieldForward(t *testing.T) {
	a := one.AddField{Bar: "1"}
	b := two.AddField{}
	assertCompatible(t, &a, &b)
}

func TestAddFieldBackward(t *testing.T) {
	a := one.AddField{}
	b := two.AddField{Bar: "1", Bat: "2"}
	assertCompatible(t, &b, &a)
}
