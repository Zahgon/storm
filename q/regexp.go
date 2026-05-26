package q

import (
	"regexp"
	"sync"
)

// Re creates a regexp matcher. It checks if the given field matches the given regexp.
// Note that this only supports fields of type string or []byte.
func Re(field string, re string) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

var regexpCache = struct {
	sync.RWMutex
	m map[string]*regexp.Regexp
}{m: make(map[string]*regexp.Regexp)}

type regexpMatcher struct {
	r   *regexp.Regexp
	err error
}

func (r *regexpMatcher) MatchField(v interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
