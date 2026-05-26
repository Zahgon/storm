package storm

import (
	"reflect"

	"github.com/asdine/storm/v3/index"
	bolt "go.etcd.io/bbolt"
)

// Storm tags
const (
	tagID        = "id"
	tagIdx       = "index"
	tagUniqueIdx = "unique"
	tagInline    = "inline"
	tagIncrement = "increment"
	indexPrefix  = "__storm_index_"
)

type fieldConfig struct {
	Name           string
	Index          string
	IsZero         bool
	IsID           bool
	Increment      bool
	IncrementStart int64
	IsInteger      bool
	Value          *reflect.Value
	ForceUpdate    bool
}

// structConfig is a structure gathering all the relevant informations about a model
type structConfig struct {
	Name   string
	Fields map[string]*fieldConfig
	ID     *fieldConfig
}

func extract(s *reflect.Value, mi ...*structConfig) (*structConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractField(value *reflect.Value, field *reflect.StructField, m *structConfig, isChild bool) error {
	_ = "STUB: not implemented"
	return nil
}

// we don't need to save this field

// the field is named ID and no ID field has been detected before

func extractSingleField(ref *reflect.Value, fieldName string) (*structConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getIndex(bucket *bolt.Bucket, idxKind string, fieldName string) (index.Index, error) {
	_ = "STUB: not implemented"
	return *new(index.Index), nil
}

func isZero(v *reflect.Value) bool { _ = "STUB: not implemented"; return false }

func isInteger(v *reflect.Value) bool { _ = "STUB: not implemented"; return false }
