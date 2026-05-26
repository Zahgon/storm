// Package sereal contains a codec to encode and decode entities using Sereal
package sereal

const name = "sereal"

// Codec that encodes to and decodes using Sereal.
// The Sereal codec has some interesting features, one of them being
// serialization of object references, including circular references.
// See https://github.com/Sereal/Sereal
var Codec = new(serealCodec)

type serealCodec int

func (c serealCodec) Marshal(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c serealCodec) Unmarshal(b []byte, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c serealCodec) Name() string { _ = "STUB: not implemented"; return "" }
