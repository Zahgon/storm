// Package json contains a codec to encode and decode entities in JSON format
package json

const name = "json"

// Codec that encodes to and decodes from JSON.
var Codec = new(jsonCodec)

type jsonCodec int

func (j jsonCodec) Marshal(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j jsonCodec) Unmarshal(b []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }

func (j jsonCodec) Name() string { _ = "STUB: not implemented"; return "" }
