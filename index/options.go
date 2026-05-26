package index

// NewOptions creates initialized Options
func NewOptions() *Options { _ = "STUB: not implemented"; return nil }

// Options are used to customize queries
type Options struct {
	Limit   int
	Skip    int
	Reverse bool
}
