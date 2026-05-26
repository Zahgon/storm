// Package msgpack contains a codec to encode and decode entities in msgpack format
package msgpack

const name = "msgpack"

// Codec that encodes to and decodes from msgpack.
var Codec = new(msgpackCodec)

type msgpackCodec int

func (m msgpackCodec) Marshal(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m msgpackCodec) Unmarshal(b []byte, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (m msgpackCodec) Name() string { _ = "STUB: not implemented"; return "" }
