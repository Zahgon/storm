package storm

import (
	"github.com/asdine/storm/v3/codec"
	"github.com/asdine/storm/v3/codec/json"
	bolt "go.etcd.io/bbolt"
)

const (
	dbinfo         = "__storm_db"
	metadataBucket = "__storm_metadata"
)

// Defaults to json
var defaultCodec = json.Codec

// Open opens a database at the given path with optional Storm options.
func Open(path string, stormOptions ...func(*Options) error) (*DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip if UseDB option is used

// DB is the wrapper around BoltDB. It contains an instance of BoltDB and uses it to perform all the
// needed operations
type DB struct {
	// The root node that points to the root bucket.
	Node

	// Bolt is still easily accessible
	Bolt *bolt.DB
}

// Close the database
func (s *DB) Close() error { _ = "STUB: not implemented"; return nil }

func (s *DB) checkVersion() error { _ = "STUB: not implemented"; return nil }

// for now, we only set the current version if it doesn't exist.
// v1 and v2 database files are compatible.

// toBytes turns an interface into a slice of bytes
func toBytes(key interface{}, codec codec.MarshalUnmarshaler) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func numbertob(v interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func numberfromb(raw []byte) (int64, error) { _ = "STUB: not implemented"; return 0, nil }
