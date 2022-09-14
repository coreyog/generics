package collections

// Set is a set of unique elements of type T.
type Set[T comparable] map[T]struct{}

// MakeStableSet returns a stable set of unique elements of type T.
func MakeStableSet[T comparable](values ...T) []T {
	set := Set[T]{}
	out := make([]T, 0, len(values))

	for _, x := range values {
		if !set.InSet(x) {
			set.Add(x)
			out = append(out, x)
		}
	}

	return out
}

// ToSet converts a slice of type T to a set of type T.
func ToSet[T comparable](values ...T) Set[T] {
	s := Set[T]{}

	for _, a := range values {
		s.Add(a)
	}

	return s
}

// IsEmpty returns true if the set is nil or has a length of 0.
func (s *Set[T]) IsEmpty() bool {
	return s == nil || len(*s) == 0
}

// Add adds elements to the set.
func (s *Set[T]) Add(x ...T) {
	if s == nil || *s == nil {
		*s = make(map[T]struct{})
	}

	for _, e := range x {
		(*s)[e] = struct{}{}
	}
}

// Delete removes elements from the set. If the element is not in the set, it is ignored.
func (s *Set[T]) Delete(x ...T) {
	if s == nil || *s == nil {
		return
	}

	for _, e := range x {
		delete(*s, e)
	}
}

// InSet returns true if the element is in the set.
func (s *Set[T]) InSet(x T) (in bool) {
	if s == nil || *s == nil {
		return false
	}

	_, in = (*s)[x]

	return in
}

// Slice returns a slice of the elements in the set.
func (s *Set[T]) Slice() []T {
	if s == nil {
		return []T{}
	}

	arr := make([]T, 0, len(*s))

	for x := range *s {
		arr = append(arr, x)
	}

	return arr
}
