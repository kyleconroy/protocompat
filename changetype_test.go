package protocompat

import (
	"testing"

	"github.com/kyleconroy/protocompat/one"
	"github.com/kyleconroy/protocompat/two"
)

func TestChangeTypeForward(t *testing.T) {
	foo := one.ChangeType{Bar: "1"}
	bar := two.ChangeType{}
	assertIncompatible(t, &foo, &bar)
}

func TestChangeTypeBackward(t *testing.T) {
	foo := one.ChangeType{}
	bar := two.ChangeType{Bar: 1}
	assertIncompatible(t, &bar, &foo)
}
