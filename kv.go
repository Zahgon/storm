package storm

import (
	bolt "go.etcd.io/bbolt"
)

// KeyValueStore can store and fetch values by key
type KeyValueStore interface {
	// Get a value from a bucket
	Get(bucketName string, key interface{}, to interface{}) error
	// Set a key/value pair into a bucket
	Set(bucketName string, key interface{}, value interface{}) error
	// Delete deletes a key from a bucket
	Delete(bucketName string, key interface{}) error
	// GetBytes gets a raw value from a bucket.
	GetBytes(bucketName string, key interface{}) ([]byte, error)
	// SetBytes sets a raw value into a bucket.
	SetBytes(bucketName string, key interface{}, value []byte) error
	// KeyExists reports the presence of a key in a bucket.
	KeyExists(bucketName string, key interface{}) (bool, error)
}

// GetBytes gets a raw value from a bucket.
func (n *node) GetBytes(bucketName string, key interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBytes gets a raw value from a bucket.
func (n *node) getBytes(tx *bolt.Tx, bucketName string, id []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetBytes sets a raw value into a bucket.
func (n *node) SetBytes(bucketName string, key interface{}, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) setBytes(tx *bolt.Tx, bucketName string, id, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// save node configuration in the bucket

// Get a value from a bucket
func (n *node) Get(bucketName string, key interface{}, to interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Set a key/value pair into a bucket
func (n *node) Set(bucketName string, key interface{}, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete deletes a key from a bucket
func (n *node) Delete(bucketName string, key interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) delete(tx *bolt.Tx, bucketName string, id []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// KeyExists reports the presence of a key in a bucket.
func (n *node) KeyExists(bucketName string, key interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
