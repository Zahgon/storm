//go:build !go1.8
// +build !go1.8

package storm

func (s *sorter) Swap(i, j int) {
	_ = "STUB: not implemented"
	// skip if we encountered an earlier error
	return
}
