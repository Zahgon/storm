// Package gob contains a codec to encode and decode entities in Gob format
package gob

const name = "gob"

// Codec serializing objects using the gob package.
// See https://golang.org/pkg/encoding/gob/
var Codec = new(gobCodec)

type gobCodec int

func (c gobCodec) Marshal(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c gobCodec) Unmarshal(b []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }

func (c gobCodec) Name() string { _ = "STUB: not implemented"; return "" }
