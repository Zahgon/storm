package internal

import (
	bolt "go.etcd.io/bbolt"
)

// Cursor that can be reversed
type Cursor struct {
	C       *bolt.Cursor
	Reverse bool
}

// First element
func (c *Cursor) First() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

// Next element
func (c *Cursor) Next() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

// RangeCursor that can be reversed
type RangeCursor struct {
	C         *bolt.Cursor
	Reverse   bool
	Min       []byte
	Max       []byte
	CompareFn func([]byte, []byte) int
}

// First element
func (c *RangeCursor) First() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

// If Seek doesn't find a key it goes to the next.
// If so, we need to get the previous one to avoid
// including bigger values. #218

// Next element
func (c *RangeCursor) Next() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

// Continue tells if the loop needs to continue
func (c *RangeCursor) Continue(val []byte) bool { _ = "STUB: not implemented"; return false }

// PrefixCursor that can be reversed
type PrefixCursor struct {
	C       *bolt.Cursor
	Reverse bool
	Prefix  []byte
}

// First element
func (c *PrefixCursor) First() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

// Next element
func (c *PrefixCursor) Next() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

// Continue tells if the loop needs to continue
func (c *PrefixCursor) Continue(val []byte) bool { _ = "STUB: not implemented"; return false }
