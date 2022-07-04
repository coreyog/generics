package collections

import "sync"

type SafeStack[T any] struct {
	arr []T
	mux *sync.RWMutex
}

func NewSafeStack[T any](arr []T) *SafeStack[T] {
	ss := &SafeStack[T]{
		mux: &sync.RWMutex{},
	}

	if len(arr) != 0 {
		ss.arr = arr[:]
	}

	return ss
}

func (s *SafeStack[T]) IsEmpty() (is bool) {
	if s == nil {
		return true
	}

	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.arr == nil || len(s.arr) == 0
}

func (s *SafeStack[T]) Push(value T) *SafeStack[T] {
	if s == nil {
		s = NewSafeStack[T](nil)
	}

	s.mux.Lock()
	defer s.mux.Unlock()

	s.arr = append(s.arr, value)

	return s
}

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

func (s *SafeStack[T]) Array() []T {
	if s == nil {
		return []T{}
	}

	s.mux.RLock()
	defer s.mux.RUnlock()

	return s.arr[:]
}

func (q *SafeStack[T]) Len() int {
	if q == nil {
		return 0
	}

	q.mux.RLock()
	defer q.mux.RUnlock()

	return len(q.arr)
}
