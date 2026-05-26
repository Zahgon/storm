package q

import (
	"errors"
	"go/token"
	"reflect"
)

// ErrUnknownField is returned when an unknown field is passed.
var ErrUnknownField = errors.New("unknown field")

type fieldMatcherDelegate struct {
	FieldMatcher
	Field string
}

// NewFieldMatcher creates a Matcher for a given field.
func NewFieldMatcher(field string, fm FieldMatcher) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

// FieldMatcher can be used in NewFieldMatcher as a simple way to create the
// most common Matcher: A Matcher that evaluates one field's value.
// For more complex scenarios, implement the Matcher interface directly.
type FieldMatcher interface {
	MatchField(v interface{}) (bool, error)
}

func (r fieldMatcherDelegate) Match(i interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r fieldMatcherDelegate) MatchValue(v *reflect.Value) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// NewField2FieldMatcher creates a Matcher for a given field1 and field2.
func NewField2FieldMatcher(field1, field2 string, tok token.Token) Matcher {
	_ = "STUB: not implemented"
	return *new(Matcher)
}

type field2fieldMatcherDelegate struct {
	Field1, Field2 string
	Tok            token.Token
}

func (r field2fieldMatcherDelegate) Match(i interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r field2fieldMatcherDelegate) MatchValue(v *reflect.Value) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
