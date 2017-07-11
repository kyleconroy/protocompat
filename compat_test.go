package protocompat

import (
	"testing"

	"github.com/golang/protobuf/proto"
)

func check(a, b proto.Message) error {
	data, err := proto.Marshal(a)
	// Something is serioualy wrong if we can't marshal the message
	if err != nil {
		panic(err)
	}
	return proto.Unmarshal(data, b)
}

func assertCompatible(t *testing.T, a, b proto.Message) {
	if err := check(a, b); err != nil {
		t.Errorf("expected messages to be compatible; %s", err)
	}
}

func assertIncompatible(t *testing.T, a, b proto.Message) {
	if err := check(a, b); err == nil {
		t.Error("expected messages to be incompatible, but no error")
	}
}
