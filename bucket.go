package storm

import bolt "go.etcd.io/bbolt"

// CreateBucketIfNotExists creates the bucket below the current node if it doesn't
// already exist.
func (n *node) CreateBucketIfNotExists(tx *bolt.Tx, bucket string) (*bolt.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBucket returns the given bucket below the current node.
func (n *node) GetBucket(tx *bolt.Tx, children ...string) *bolt.Bucket {
	_ = "STUB: not implemented"
	return nil
}
