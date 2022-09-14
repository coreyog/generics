package collections

import "sync"

// SafeStack is a thread-safe LIFO stack.
type SafeStack[T any] struct {
	arr []T
	mux *sync.RWMutex
}

// NewSafeStack constructs a SafeStack and pushes all parameters in order.
func NewSafeStack[T any](arr ...T) *SafeStack[T] {
	ss := &SafeStack[T]{
		mux: &sync.RWMutex{},
	}

	if len(arr) != 0 {
		ss.arr = arr[:]
	}

	return ss
}

// IsEmpty returns true if the stack is nil or has a length of 0.
func (s *SafeStack[T]) IsEmpty() (is bool) {
	if s == nil {
		return true
	}

	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.arr == nil || len(s.arr) == 0
}

// Push adds elements to the top of the stack.
func (s *SafeStack[T]) Push(value ...T) *SafeStack[T] {
	if s == nil {
		s = NewSafeStack[T]()
	}

	s.mux.Lock()
	defer s.mux.Unlock()

	s.arr = append(s.arr, value...)

	return s
}

// Pop removes and returns the top element of the stack.
func (s *SafeStack[T]) Pop() (elem T, ok bool) {
	if s == nil {
		return [1]T{}[0], false
	}

	s.mux.Lock()
	defer s.mux.Unlock()

	if s.arr == nil || len(s.arr) == 0 {
		return [1]T{}[0], false
	}

	last := len(s.arr) - 1
	elem = s.arr[last]
	s.arr = s.arr[:last]

	return elem, true
}

// Peek returns the top element of the stack without removing it.
func (s *SafeStack[T]) Peek() (elem T, ok bool) {
	if s == nil {
		return [1]T{}[0], false
	}

	s.mux.RLock()
	defer s.mux.RUnlock()

	if s.arr == nil || len(s.arr) == 0 {
		return [1]T{}[0], false
	}

	return s.arr[0], true
}

// Slice returns a copy of the stack as a slice.
func (s *SafeStack[T]) Slice() []T {
	if s == nil {
		return []T{}
	}

	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.arr[:]
}

// Len returns the length of the stack.
func (q *SafeStack[T]) Len() int {
	if q == nil {
		return 0
	}

	q.mux.RLock()
	defer q.mux.RUnlock()

	return len(q.arr)
}
