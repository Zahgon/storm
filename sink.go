package storm

import (
	"reflect"

	"github.com/asdine/storm/v3/q"
	bolt "go.etcd.io/bbolt"
)

type item struct {
	value  *reflect.Value
	bucket *bolt.Bucket
	k      []byte
	v      []byte
}

func newSorter(n Node, snk sink) *sorter { _ = "STUB: not implemented"; return nil }

type sorter struct {
	node    Node
	sink    sink
	list    []*item
	skip    int
	limit   int
	orderBy []string
	reverse bool
	err     chan error
	done    chan struct{}
}

func (s *sorter) filter(tree q.Matcher, bucket *bolt.Bucket, k, v []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// add directly to sink, we'll apply skip/limits after sorting

func (s *sorter) add(itm *item) (stop bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (s *sorter) compareValue(left reflect.Value, right reflect.Value) int {
	_ = "STUB: not implemented"
	return 0
}

func (s *sorter) less(leftElem reflect.Value, rightElem reflect.Value) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *sorter) flush() error { _ = "STUB: not implemented"; return nil }

func (s *sorter) Len() int {
	_ = "STUB: not implemented"
	// skip if we encountered an earlier error
	return 0
}

func (s *sorter) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// skip if we encountered an earlier error
	return false
}

type sink interface {
	bucketName() string
	flush() error
	add(*item) error
	readOnly() bool
}

type reflectSink interface {
	elem() reflect.Value
}

type sliceSink interface {
	slice() reflect.Value
	setSlice(reflect.Value)
	reset()
}

func newListSink(node Node, to interface{}) (*listSink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type listSink struct {
	node     Node
	ref      reflect.Value
	results  reflect.Value
	elemType reflect.Type
	name     string
	isPtr    bool
	idx      int
}

func (l *listSink) slice() reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func (l *listSink) setSlice(s reflect.Value) { _ = "STUB: not implemented"; return }

func (l *listSink) reset() { _ = "STUB: not implemented"; return }

func (l *listSink) elem() reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func (l *listSink) bucketName() string { _ = "STUB: not implemented"; return "" }

func (l *listSink) add(i *item) error { _ = "STUB: not implemented"; return nil }

func (l *listSink) flush() error { _ = "STUB: not implemented"; return nil }

func (l *listSink) readOnly() bool { _ = "STUB: not implemented"; return false }

func newFirstSink(node Node, to interface{}) (*firstSink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type firstSink struct {
	node  Node
	ref   reflect.Value
	found bool
}

func (f *firstSink) elem() reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func (f *firstSink) bucketName() string { _ = "STUB: not implemented"; return "" }

func (f *firstSink) add(i *item) error { _ = "STUB: not implemented"; return nil }

func (f *firstSink) flush() error { _ = "STUB: not implemented"; return nil }

func (f *firstSink) readOnly() bool { _ = "STUB: not implemented"; return false }

func newDeleteSink(node Node, kind interface{}) (*deleteSink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type deleteSink struct {
	node    Node
	ref     reflect.Value
	removed int
}

func (d *deleteSink) elem() reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func (d *deleteSink) bucketName() string { _ = "STUB: not implemented"; return "" }

func (d *deleteSink) add(i *item) error { _ = "STUB: not implemented"; return nil }

func (d *deleteSink) flush() error { _ = "STUB: not implemented"; return nil }

func (d *deleteSink) readOnly() bool { _ = "STUB: not implemented"; return false }

func newCountSink(node Node, kind interface{}) (*countSink, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type countSink struct {
	node    Node
	ref     reflect.Value
	counter int
}

func (c *countSink) elem() reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func (c *countSink) bucketName() string { _ = "STUB: not implemented"; return "" }

func (c *countSink) add(i *item) error { _ = "STUB: not implemented"; return nil }

func (c *countSink) flush() error { _ = "STUB: not implemented"; return nil }

func (c *countSink) readOnly() bool { _ = "STUB: not implemented"; return false }

func newRawSink() *rawSink { _ = "STUB: not implemented"; return nil }

type rawSink struct {
	results [][]byte
	execFn  func([]byte, []byte) error
}

func (r *rawSink) add(i *item) error { _ = "STUB: not implemented"; return nil }

func (r *rawSink) bucketName() string { _ = "STUB: not implemented"; return "" }

func (r *rawSink) flush() error { _ = "STUB: not implemented"; return nil }

func (r *rawSink) readOnly() bool { _ = "STUB: not implemented"; return false }

func newEachSink(to interface{}) (*eachSink, error) { _ = "STUB: not implemented"; return nil, nil }

type eachSink struct {
	ref    reflect.Value
	execFn func(interface{}) error
}

func (e *eachSink) elem() reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

func (e *eachSink) bucketName() string { _ = "STUB: not implemented"; return "" }

func (e *eachSink) add(i *item) error { _ = "STUB: not implemented"; return nil }

func (e *eachSink) flush() error { _ = "STUB: not implemented"; return nil }

func (e *eachSink) readOnly() bool { _ = "STUB: not implemented"; return false }
