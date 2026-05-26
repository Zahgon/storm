package index

import (
	bolt "go.etcd.io/bbolt"
)

// NewListIndex loads a ListIndex
func NewListIndex(parent *bolt.Bucket, indexName []byte) (*ListIndex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListIndex is an index that references values and the corresponding IDs.
type ListIndex struct {
	Parent      *bolt.Bucket
	IndexBucket *bolt.Bucket
	IDs         *UniqueIndex
}

// Add a value to the list index
func (idx *ListIndex) Add(newValue []byte, targetID []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove a value from the unique index
func (idx *ListIndex) Remove(value []byte) error { _ = "STUB: not implemented"; return nil }

// RemoveID removes an ID from the list index
func (idx *ListIndex) RemoveID(targetID []byte) error { _ = "STUB: not implemented"; return nil }

// Get the first ID corresponding to the given value
func (idx *ListIndex) Get(value []byte) []byte { _ = "STUB: not implemented"; return nil }

// All the IDs corresponding to the given value
func (idx *ListIndex) All(value []byte, opts *Options) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AllRecords returns all the IDs of this index
func (idx *ListIndex) AllRecords(opts *Options) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Range returns the ids corresponding to the given range of values
func (idx *ListIndex) Range(min []byte, max []byte, opts *Options) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Prefix returns the ids whose values have the given prefix.
func (idx *ListIndex) Prefix(prefix []byte, opts *Options) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generatePrefix(value []byte) []byte { _ = "STUB: not implemented"; return nil }
