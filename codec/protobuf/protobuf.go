// Package protobuf contains a codec to encode and decode entities in Protocol Buffer
package protobuf

import (
	"errors"
)

const name = "protobuf"

// More details on Protocol Buffers https://github.com/golang/protobuf
var (
	Codec                       = new(protobufCodec)
	errNotProtocolBufferMessage = errors.New("value isn't a Protocol Buffers Message")
)

type protobufCodec int

// Encode value with protocol buffer.
// If type isn't a Protocol buffer Message, json encoder will be used instead.
func (c protobufCodec) Marshal(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// toBytes() may need to encode non-protobuf type, if that occurs use json

func (c protobufCodec) Unmarshal(b []byte, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// toBytes() may have encoded non-protobuf type, if that occurs use json

func (c protobufCodec) Name() string { _ = "STUB: not implemented"; return "" }
