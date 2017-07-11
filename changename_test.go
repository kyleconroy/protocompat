package protocompat

import (
	"testing"

	"github.com/kyleconroy/protocompat/one"
	"github.com/kyleconroy/protocompat/two"
)

func TestChangeNameForward(t *testing.T) {
	foo := one.ChangeName{Bat: "1"}
	bar := two.ChangeNombre{}
	assertCompatible(t, &foo, &bar)
}

func TestChangeNameBackward(t *testing.T) {
	foo := one.ChangeName{}
	bar := two.ChangeNombre{Bat: "1"}
	assertCompatible(t, &bar, &foo)
}
