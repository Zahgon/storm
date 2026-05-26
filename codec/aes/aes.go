package aes

import (
	"crypto/cipher"

	"github.com/asdine/storm/v3/codec"
)

const name = "aes-"

// AES is an Codec that encrypts the data and uses a sub marshaller to actually serialize the data
type AES struct {
	subMarshaller codec.MarshalUnmarshaler
	aesGCM        cipher.AEAD
}

// NewAES creates a new AES encryption marshaller. It takes a sub marshaller to actually serialize the data and a 16/32 bytes private key to
// encrypt all data using AES in GCM block mode.
func NewAES(subMarshaller codec.MarshalUnmarshaler, key []byte) (*AES, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Name returns a compound of the inner marshaller prefixed by 'aes-'
func (c *AES) Name() string {
	_ = "STUB: not implemented"
	// Return a dynamic name, because the marshalling will also fail if the inner marshalling changes.
	return ""
}

// Marshal marshals the given data object to an encrypted byte array
func (c *AES) Marshal(v interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal unmarshals the given encrypted byte array to the given type
func (c *AES) Unmarshal(data []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }
