// Package q contains a list of Matchers used to compare struct fields with values
package q

import (
	"go/token"
	"reflect"
)

// A Matcher is used to test against a record to see if it matches.
type Matcher interface {
	// Match is used to test the criteria against a structure.
	Match(interface{}) (bool, error)
}

// A ValueMatcher is used to test against a reflect.Value.
type ValueMatcher interface {
	// MatchValue tests if the given reflect.Value matches.
	// It is useful when the reflect.Value of an object already exists.
	MatchValue(*reflect.Value) (bool, error)
}

type cmp struct {
	value interface{}
	token token.Token
}

func (c *cmp) MatchField(v interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

type trueMatcher struct{}

func (*trueMatcher) Match(i interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (*trueMatcher) MatchValue(v *reflect.Value) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type or struct {
	children []Matcher
}

func (c *or) Match(i interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (c *or) MatchValue(v *reflect.Value) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type and struct {
	children []Matcher
}

func (c *and) Match(i interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (c *and) MatchValue(v *reflect.Value) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type strictEq struct {
	field string
	value interface{}
}

func (s *strictEq) MatchField(v interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type in struct {
	list interface{}
}

func (i *in) MatchField(v interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

type not struct {
	children []Matcher
}

func (n *not) Match(i interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (n *not) MatchValue(v *reflect.Value) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Eq matcher, checks if the given field is equal to the given value
func Eq(field string, v interface{}) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// EqF matcher, checks if the given field is equal to the given field
func EqF(field1, field2 string) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// StrictEq matcher, checks if the given field is deeply equal to the given value
func StrictEq(field string, v interface{}) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// Gt matcher, checks if the given field is greater than the given value
func Gt(field string, v interface{}) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// GtF matcher, checks if the given field is greater than the given field
func GtF(field1, field2 string) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// Gte matcher, checks if the given field is greater than or equal to the given value
func Gte(field string, v interface{}) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// GteF matcher, checks if the given field is greater than or equal to the given field
func GteF(field1, field2 string) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// Lt matcher, checks if the given field is lesser than the given value
func Lt(field string, v interface{}) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// LtF matcher, checks if the given field is lesser than the given field
func LtF(field1, field2 string) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// Lte matcher, checks if the given field is lesser than or equal to the given value
func Lte(field string, v interface{}) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// LteF matcher, checks if the given field is lesser than or equal to the given field
func LteF(field1, field2 string) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// In matcher, checks if the given field matches one of the value of the given slice.
// v must be a slice.
func In(field string, v interface{}) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// True matcher, always returns true
func True() Matcher {
	_ = "STUB: not implemented"
	return *

	// Or matcher, checks if at least one of the given matchers matches the record
	new(Matcher)
}

func Or(matchers ...Matcher) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// And matcher, checks if all of the given matchers matches the record
func And(matchers ...Matcher) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }

// Not matcher, checks if all of the given matchers return false
func Not(matchers ...Matcher) Matcher { _ = "STUB: not implemented"; return *new(Matcher) }
