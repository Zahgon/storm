package storm

import (
	bolt "go.etcd.io/bbolt"
)

// A BucketScanner scans a Node for a list of buckets
type BucketScanner interface {
	// PrefixScan scans the root buckets for keys matching the given prefix.
	PrefixScan(prefix string) []Node
	// PrefixScan scans the buckets in this node for keys matching the given prefix.
	RangeScan(min, max string) []Node
}

// PrefixScan scans the buckets in this node for keys matching the given prefix.
func (n *node) PrefixScan(prefix string) []Node { _ = "STUB: not implemented"; return nil }

func (n *node) prefixScan(tx *bolt.Tx, prefix string) []Node { _ = "STUB: not implemented"; return nil }

// RangeScan scans the buckets in this node  over a range such as a sortable time range.
func (n *node) RangeScan(min, max string) []Node { _ = "STUB: not implemented"; return nil }

func (n *node) rangeScan(tx *bolt.Tx, min, max string) []Node {
	_ = "STUB: not implemented"
	return nil
}

func (n *node) cursor(tx *bolt.Tx) *bolt.Cursor { _ = "STUB: not implemented"; return nil }
