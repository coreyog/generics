package collections

import "sync"

type SafeSet[T comparable] struct {
	sync.Map
}

func NewSafeSet[T comparable](arr []T) *SafeSet[T] {
	s := &SafeSet[T]{}

	for _, a := range arr {
		s.Add(a)
	}

	return s
}

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

func (s *SafeSet[T]) Add(x T) *SafeSet[T] {
	if s == nil {
		s = &SafeSet[T]{}
	}

	s.Store(x, struct{}{})

	return s
}

func (s *SafeSet[T]) Delete(x T) {
	if s == nil {
		return
	}

	s.Map.Delete(x)
}

func (s *SafeSet[T]) InSet(x T) (in bool) {
	if s == nil {
		return false
	}

	_, in = s.Map.Load(x)

	return in
}

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

func (s *SafeSet[T]) Range(f func(value T) bool) {
	if s == nil {
		return
	}

	s.Map.Range(func(key, value interface{}) bool {
		return f(key.(T))
	})
}

func (s *SafeSet[T]) Array() []T {
	arr := []T{}

	s.Range(func(value T) bool {
		arr = append(arr, value)
		return true
	})

	return arr
}
