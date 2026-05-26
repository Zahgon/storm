package storm

import (
	"github.com/asdine/storm/v3/q"
	bolt "go.etcd.io/bbolt"
)

// Select a list of records that match a list of matchers. Doesn't use indexes.
func (n *node) Select(matchers ...q.Matcher) Query { _ = "STUB: not implemented"; return *new(Query) }

// Query is the low level query engine used by Storm. It allows to operate searches through an entire bucket.
type Query interface {
	// Skip matching records by the given number
	Skip(int) Query

	// Limit the results by the given number
	Limit(int) Query

	// Order by the given fields, in descending precedence, left-to-right.
	OrderBy(...string) Query

	// Reverse the order of the results
	Reverse() Query

	// Bucket specifies the bucket name
	Bucket(string) Query

	// Find a list of matching records
	Find(interface{}) error

	// First gets the first matching record
	First(interface{}) error

	// Delete all matching records
	Delete(interface{}) error

	// Count all the matching records
	Count(interface{}) (int, error)

	// Returns all the records without decoding them
	Raw() ([][]byte, error)

	// Execute the given function for each raw element
	RawEach(func([]byte, []byte) error) error

	// Execute the given function for each element
	Each(interface{}, func(interface{}) error) error
}

func newQuery(n *node, tree q.Matcher) *query { _ = "STUB: not implemented"; return nil }

type query struct {
	limit   int
	skip    int
	reverse bool
	tree    q.Matcher
	node    *node
	bucket  string
	orderBy []string
}

func (q *query) Skip(nb int) Query { _ = "STUB: not implemented"; return *new(Query) }

func (q *query) Limit(nb int) Query { _ = "STUB: not implemented"; return *new(Query) }

func (q *query) OrderBy(field ...string) Query { _ = "STUB: not implemented"; return *new(Query) }

func (q *query) Reverse() Query { _ = "STUB: not implemented"; return *new(Query) }

func (q *query) Bucket(bucketName string) Query { _ = "STUB: not implemented"; return *new(Query) }

func (q *query) Find(to interface{}) error { _ = "STUB: not implemented"; return nil }

func (q *query) First(to interface{}) error { _ = "STUB: not implemented"; return nil }

func (q *query) Delete(kind interface{}) error { _ = "STUB: not implemented"; return nil }

func (q *query) Count(kind interface{}) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (q *query) Raw() ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (q *query) RawEach(fn func([]byte, []byte) error) error { _ = "STUB: not implemented"; return nil }

func (q *query) Each(kind interface{}, fn func(interface{}) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (q *query) runQuery(sink sink) error { _ = "STUB: not implemented"; return nil }

func (q *query) query(tx *bolt.Tx, sink sink) error { _ = "STUB: not implemented"; return nil }
