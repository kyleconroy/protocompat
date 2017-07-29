package protocompat

import (
	"testing"

	"github.com/kyleconroy/protocompat/one"
	"github.com/kyleconroy/protocompat/two"
)

func TestRemoveFieldForward(t *testing.T) {
	a := one.RemoveField{Bar: "1", Bat: "2"}
	b := two.RemoveField{}
	assertCompatible(t, &b, &a)
}

func TestRemoveFieldBackward(t *testing.T) {
	a := two.RemoveField{}
	b := one.RemoveField{Bar: "1"}
	assertCompatible(t, &a, &b)
}
