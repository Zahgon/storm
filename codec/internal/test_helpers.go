package internal

import (
	"encoding/gob"
	"testing"

	"github.com/asdine/storm/v3/codec"
)

type testStruct struct {
	Name string
}

// RoundtripTester is a test helper to test a MarshalUnmarshaler
func RoundtripTester(t *testing.T, c codec.MarshalUnmarshaler, vals ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func init() {
	gob.Register(&testStruct{})
}
