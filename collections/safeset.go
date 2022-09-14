package collections

import "sync"

// SafeSet is a thread-safe set.
type SafeSet[T comparable] struct {
	sync.Map
}

// NewSafeSet constructs a SafeSet and adds all parameters.
func NewSafeSet[T comparable](arr ...T) *SafeSet[T] {
	s := &SafeSet[T]{}

	for _, a := range arr {
		s.Add(a)
	}

	return s
}

// IsEmpty returns true if the set is nil or has a length of 0.
func (s *SafeSet[T]) IsEmpty() bool {
	if s == nil {
		return true
	}

	containsSomething := false
	s.Range(func(value T) bool {
		containsSomething = true
		return false
	})

	return !containsSomething
}

// Add adds elements to the set.
func (s *SafeSet[T]) Add(x ...T) *SafeSet[T] {
	if s == nil {
		s = &SafeSet[T]{}
	}

	for _, e := range x {
		s.Store(e, struct{}{})
	}

	return s
}

// Delete removes elements from the set. If the element is not in the set, it is ignored.
func (s *SafeSet[T]) Delete(x ...T) {
	if s == nil {
		return
	}

	for _, e := range x {
		s.Map.Delete(e)
	}
}

// InSet returns true if the element is in the set.
func (s *SafeSet[T]) InSet(x T) (in bool) {
	if s == nil {
		return false
	}

	_, in = s.Map.Load(x)

	return in
}

// Len returns the number of elements in the set.
func (s *SafeSet[T]) Len() int {
	if s == nil {
		return 0
	}

	count := 0
	s.Range(func(value T) bool {
		count++
		return true
	})

	return count
}

// Range calls f sequentially for each key and value present in the set. If f returns false, range stops the iteration.
func (s *SafeSet[T]) Range(f func(value T) bool) {
	if s == nil {
		return
	}

	s.Map.Range(func(key, value interface{}) bool {
		return f(key.(T))
	})
}

// Slice returns a slice of the elements in the set.
func (s *SafeSet[T]) Slice() []T {
	arr := []T{}

	s.Range(func(value T) bool {
		arr = append(arr, value)
		return true
	})

	return arr
}
