package protocompat

import (
	"testing"

	"github.com/kyleconroy/protocompat/one"
	"github.com/kyleconroy/protocompat/two"
)

func TestChangeTagForward(t *testing.T) {
	foo := one.ChangeTag{Bar: "1"}
	bar := two.ChangeTag{}
	assertCompatible(t, &foo, &bar)
}

func TestChangeTagBackward(t *testing.T) {
	foo := one.ChangeTag{}
	bar := two.ChangeTag{Bar: "1"}
	assertCompatible(t, &bar, &foo)
}
