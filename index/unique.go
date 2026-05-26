package index

import (
	bolt "go.etcd.io/bbolt"
)

// NewUniqueIndex loads a UniqueIndex
func NewUniqueIndex(parent *bolt.Bucket, indexName []byte) (*UniqueIndex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UniqueIndex is an index that references unique values and the corresponding ID.
type UniqueIndex struct {
	Parent      *bolt.Bucket
	IndexBucket *bolt.Bucket
}

// Add a value to the unique index
func (idx *UniqueIndex) Add(value []byte, targetID []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove a value from the unique index
func (idx *UniqueIndex) Remove(value []byte) error { _ = "STUB: not implemented"; return nil }

// RemoveID removes an ID from the unique index
func (idx *UniqueIndex) RemoveID(id []byte) error { _ = "STUB: not implemented"; return nil }

// Get the id corresponding to the given value
func (idx *UniqueIndex) Get(value []byte) []byte { _ = "STUB: not implemented"; return nil }

// All returns all the ids corresponding to the given value
func (idx *UniqueIndex) All(value []byte, opts *Options) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AllRecords returns all the IDs of this index
func (idx *UniqueIndex) AllRecords(opts *Options) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Range returns the ids corresponding to the given range of values
func (idx *UniqueIndex) Range(min []byte, max []byte, opts *Options) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prefix returns the ids whose values have the given prefix.
func (idx *UniqueIndex) Prefix(prefix []byte, opts *Options) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// first returns the first ID of this index
func (idx *UniqueIndex) first() []byte { _ = "STUB: not implemented"; return nil }
