package storm

// Tx is a transaction.
type Tx interface {
	// Commit writes all changes to disk.
	Commit() error

	// Rollback closes the transaction and ignores all previous updates.
	Rollback() error
}

// Begin starts a new transaction.
func (n node) Begin(writable bool) (Node, error) { _ = "STUB: not implemented"; return *new(Node), nil }

// Rollback closes the transaction and ignores all previous updates.
func (n *node) Rollback() error { _ = "STUB: not implemented"; return nil }

// Commit writes all changes to disk.
func (n *node) Commit() error { _ = "STUB: not implemented"; return nil }
