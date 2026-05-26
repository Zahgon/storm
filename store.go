package storm

import (
	"reflect"

	"github.com/asdine/storm/v3/index"
	bolt "go.etcd.io/bbolt"
)

// TypeStore stores user defined types in BoltDB.
type TypeStore interface {
	Finder
	// Init creates the indexes and buckets for a given structure
	Init(data interface{}) error

	// ReIndex rebuilds all the indexes of a bucket
	ReIndex(data interface{}) error

	// Save a structure
	Save(data interface{}) error

	// Update a structure
	Update(data interface{}) error

	// UpdateField updates a single field
	UpdateField(data interface{}, fieldName string, value interface{}) error

	// Drop a bucket
	Drop(data interface{}) error

	// DeleteStruct deletes a structure from the associated bucket
	DeleteStruct(data interface{}) error
}

// Init creates the indexes and buckets for a given structure
func (n *node) Init(data interface{}) error { _ = "STUB: not implemented"; return nil }

func (n *node) init(tx *bolt.Tx, cfg *structConfig) error {
	bucket, err := n.CreateBucketIfNotExists(tx, cfg.Name)
	if err != nil {
		return err
	}

	// save node configuration in the bucket
	_, err = newMeta(bucket, n)
	if err != nil {
		return err
	}

	for fieldName, fieldCfg := range cfg.Fields {
		if fieldCfg.Index == "" {
			continue
		}
		switch fieldCfg.Index {
		case tagUniqueIdx:
			_, err = index.NewUniqueIndex(bucket, []byte(indexPrefix+fieldName))
		case tagIdx:
			_, err = index.NewListIndex(bucket, []byte(indexPrefix+fieldName))
		default:
			err = ErrIdxNotFound
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (n *node) ReIndex(data interface{}) error { _ = "STUB: not implemented"; return nil }

func (n *node) reIndex(tx *bolt.Tx, data interface{}, cfg *structConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Save a structure
func (n *node) Save(data interface{}) error { _ = "STUB: not implemented"; return nil }

func (n *node) save(tx *bolt.Tx, cfg *structConfig, data interface{}, update bool) error {
	_ = "STUB: not implemented"
	return nil
}

// save node configuration in the bucket

// Update a structure
func (n *node) Update(data interface{}) error { _ = "STUB: not implemented"; return nil }

// UpdateField updates a single field
func (n *node) UpdateField(data interface{}, fieldName string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) update(data interface{}, fn func(*reflect.Value, *reflect.Value, *structConfig) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Drop a bucket
func (n *node) Drop(data interface{}) error { _ = "STUB: not implemented"; return nil }

func (n *node) drop(tx *bolt.Tx, bucketName string) error { _ = "STUB: not implemented"; return nil }

// DeleteStruct deletes a structure from the associated bucket
func (n *node) DeleteStruct(data interface{}) error { _ = "STUB: not implemented"; return nil }

func (n *node) deleteStruct(tx *bolt.Tx, cfg *structConfig, id []byte) error {
	_ = "STUB: not implemented"
	return nil
}
