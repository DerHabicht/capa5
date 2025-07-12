package collections

type Set[T comparable] map[T]bool

// Add() places an element in the set.
func (s *Set[T]) Add(t T) {
	if *s == nil {
		m := make(Set[T])
		*s = m
	}
	(*s)[t] = true
}

// Contains() returns true if the given element is in the set, false otherwise.
func (s *Set[T]) Contains(t T) bool {
	_, ok := (*s)[t]

	return ok
}

// Intersect(t Set[T]) returns the intersection of this set with set t.
func (s *Set[T]) Intersect(t Set[T]) Set[T] {
	var inter Set[T]

	var l Set[T]
	if len(*s) < len(t) {
		l = *s
	} else {
		l = t
	}

	for k := range l {
		if s.Contains(k) && t.Contains(k) {
			inter.Add(k)
		}
	}

	return inter
}

// Remove() pulls the given element out of the set, if present. If not present
// this has no effect.
func (s *Set[T]) Remove(t T) {
	delete(*s, t)
}

// Sub(t Set[T]) returns the s - t difference set.
func (s *Set[T]) Sub(t Set[T]) Set[T] {
	var diff Set[T]

	for k := range *s {
		if !t.Contains(k) {
			diff.Add(k)
		}
	}

	return diff
}

// Union(t Set[T]) returns the union of this set and set T.
func (s *Set[T]) Union(t Set[T]) Set[T] {
	var union Set[T]

	for k := range *s {
		union.Add(k)
	}

	for k := range t {
		union.Add(k)
	}

	return union
}