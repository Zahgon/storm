package storm

import (
	bolt "go.etcd.io/bbolt"
)

const (
	metaCodec = "codec"
)

func newMeta(b *bolt.Bucket, n Node) (*meta, error) { _ = "STUB: not implemented"; return nil, nil }

type meta struct {
	node   Node
	bucket *bolt.Bucket
}

func (m *meta) increment(field *fieldConfig) error { _ = "STUB: not implemented"; return nil }
