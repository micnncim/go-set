// Package set defines various methods for a set.
package set

import (
	"fmt"
)

// Set is a set of comparables.
type Set[V comparable] struct {
	m map[V]struct{}
}

// New returns a Set from the given values.
func New[V comparable](v ...V) *Set[V] {
	s := &Set[V]{make(map[V]struct{})}

	s.Insert(v...)

	return s
}

// Clone returns a new Set that a copy of `s`.
func (s *Set[V]) Clone() *Set[V] {
	t := New[V]()

	t.Insert(s.Values()...)

	return t
}

// Delete removes the given values from `s`.
func (s *Set[V]) Delete(v ...V) {
	for _, x := range v {
		delete(s.m, x)
	}
}

// Difference returns a Set whose values are in `s` and not in `t`.
//
// For example:
//
//	s = {a1, a2, a3}
//	t = {a1, a2, a4, a5}
//	s.Difference(t) = {a3}
//	t.Difference(s) = {a4, a5}
func (s *Set[V]) Difference(t *Set[V]) *Set[V] {
	u := New[V]()

	for k := range s.m {
		if !t.Contains(k) {
			u.Insert(k)
		}
	}

	return u
}

// Intersection returns a new Set whose values are included in both `s` and `t`.
//
// For example:
//
//	s = {a1, a2}
//	t = {a2, a3}
//	s.Intersection(t) = {a2}
func (s *Set[V]) Intersection(t *Set[V]) *Set[V] {
	u := New[V]()

	var walk, other *Set[V]

	if s.Len() < t.Len() {
		walk = s
		other = t
	} else {
		walk = t
		other = s
	}

	for k := range walk.m {
		if other.Contains(k) {
			u.Insert(k)
		}
	}

	return u
}

// Equal returns true iff `s` is equal to `t`.
//
// Two sets are equal if their underlying values are identical not considering
// order.
func (s *Set[V]) Equal(t *Set[V]) bool {
	return len(s.m) == len(t.m) && s.IsSuperset(t)
}

// Contains returns true iff `s` contains a given value.
func (s *Set[V]) Contains(v V) bool {
	_, ok := s.m[v]
	return ok
}

// ContainsAll returns true iff `s` contains all the given values.
func (s *Set[V]) ContainsAll(v ...V) bool {
	for _, x := range v {
		if !s.Contains(x) {
			return false
		}
	}

	return true
}

// ContainsAny returns true iff `s` contains any of the given values.
func (s *Set[V]) ContainsAny(v ...V) bool {
	for _, x := range v {
		if s.Contains(x) {
			return true
		}
	}

	return false
}

// Insert adds the given values to `s`.
func (s *Set[V]) Insert(v ...V) {
	for _, x := range v {
		s.m[x] = struct{}{}
	}
}

// IsSuperset returns true iff `t` is a superset of `s`.
func (s *Set[V]) IsSuperset(t *Set[V]) bool {
	for k := range t.m {
		if !s.Contains(k) {
			return false
		}
	}

	return true
}

// Len returns the size of `s`.
func (s *Set[V]) Len() int {
	return len(s.m)
}

// PopAny returns a single value randomly chosen and removes it from `s`.
func (s *Set[V]) PopAny() (v V, _ bool) {
	for k := range s.m {
		delete(s.m, k)
		return k, true
	}

	return v, false
}

// String implements fmt.Stringer.
func (s *Set[V]) String() string {
	return fmt.Sprint(s.Values())
}

// Values returns the underlying values of `s` as a slice.
func (s *Set[V]) Values() []V {
	v := make([]V, 0, len(s.m))

	for k := range s.m {
		v = append(v, k)
	}

	return v
}

// Union returns a new Set whose values are included in either `s` or `t`.
//
// For example:
//
//	s = {a1, a2}
//	t = {a3, a4}
//	s.Union(t) = {a1, a2, a3, a4}
//	t.Union(s) = {a1, a2, a3, a4}
func (s *Set[V]) Union(t *Set[V]) *Set[V] {
	u := s.Clone()

	for k := range t.m {
		u.Insert(k)
	}

	return u
}
