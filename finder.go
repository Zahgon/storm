package storm

import (
	"reflect"

	"github.com/asdine/storm/v3/index"
	"github.com/asdine/storm/v3/q"
	bolt "go.etcd.io/bbolt"
)

// A Finder can fetch types from BoltDB.
type Finder interface {
	// One returns one record by the specified index
	One(fieldName string, value interface{}, to interface{}) error

	// Find returns one or more records by the specified index
	Find(fieldName string, value interface{}, to interface{}, options ...func(q *index.Options)) error

	// AllByIndex gets all the records of a bucket that are indexed in the specified index
	AllByIndex(fieldName string, to interface{}, options ...func(*index.Options)) error

	// All gets all the records of a bucket.
	// If there are no records it returns no error and the 'to' parameter is set to an empty slice.
	All(to interface{}, options ...func(*index.Options)) error

	// Select a list of records that match a list of matchers. Doesn't use indexes.
	Select(matchers ...q.Matcher) Query

	// Range returns one or more records by the specified index within the specified range
	Range(fieldName string, min, max, to interface{}, options ...func(*index.Options)) error

	// Prefix returns one or more records whose given field starts with the specified prefix.
	Prefix(fieldName string, prefix string, to interface{}, options ...func(*index.Options)) error

	// Count counts all the records of a bucket
	Count(data interface{}) (int, error)
}

// One returns one record by the specified index
func (n *node) One(fieldName string, value interface{}, to interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) one(tx *bolt.Tx, bucketName, fieldName string, cfg *structConfig, to interface{}, val []byte, skipIndex bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Find returns one or more records by the specified index
func (n *node) Find(fieldName string, value interface{}, to interface{}, options ...func(q *index.Options)) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) find(tx *bolt.Tx, bucketName, fieldName string, cfg *structConfig, sink *listSink, val []byte, opts *index.Options) error {
	_ = "STUB: not implemented"
	return nil
}

// AllByIndex gets all the records of a bucket that are indexed in the specified index
func (n *node) AllByIndex(fieldName string, to interface{}, options ...func(*index.Options)) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) allByIndex(tx *bolt.Tx, fieldName string, cfg *structConfig, ref *reflect.Value, opts *index.Options) error {
	_ = "STUB: not implemented"
	return nil
}

// All gets all the records of a bucket.
// If there are no records it returns no error and the 'to' parameter is set to an empty slice.
func (n *node) All(to interface{}, options ...func(*index.Options)) error {
	_ = "STUB: not implemented"
	return nil
}

// Range returns one or more records by the specified index within the specified range
func (n *node) Range(fieldName string, min, max, to interface{}, options ...func(*index.Options)) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) rnge(tx *bolt.Tx, bucketName, fieldName string, cfg *structConfig, sink *listSink, min, max []byte, opts *index.Options) error {
	_ = "STUB: not implemented"
	return nil
}

// Prefix returns one or more records whose given field starts with the specified prefix.
func (n *node) Prefix(fieldName string, prefix string, to interface{}, options ...func(*index.Options)) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) prefix(tx *bolt.Tx, bucketName, fieldName string, cfg *structConfig, sink *listSink, prefix []byte, opts *index.Options) error {
	_ = "STUB: not implemented"
	return nil
}

// Count counts all the records of a bucket
func (n *node) Count(data interface{}) (int, error) { _ = "STUB: not implemented"; return 0, nil }
