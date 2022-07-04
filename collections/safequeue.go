package collections

import "sync"

type SafeQueue[T any] struct {
	arr []T
	mux *sync.RWMutex
}

func NewSafeQueue[T any](arr []T) *SafeQueue[T] {
	sq := &SafeQueue[T]{
		mux: &sync.RWMutex{},
	}

	if len(arr) != 0 {
		sq.arr = arr[:]
	}

	return sq
}

func (q *SafeQueue[T]) IsEmpty() (is bool) {
	if q == nil {
		return true
	}

	q.mux.RLock()
	defer q.mux.RUnlock()

	return q.arr == nil || len(q.arr) == 0
}

func (q *SafeQueue[T]) Enqueue(x T) *SafeQueue[T] {
	if q == nil {
		q = NewSafeQueue[T](nil)
	}

	q.mux.Lock()
	defer q.mux.Unlock()

	q.arr = append(q.arr, x)

	return q
}

func (q *SafeQueue[T]) Dequeue() (elem T, ok bool) {
	if q == nil {
		return [1]T{}[0], false
	}

	q.mux.Lock()
	defer q.mux.Unlock()

	if q.arr == nil || len(q.arr) == 0 {
		return [1]T{}[0], false
	}

	elem = q.arr[0]
	q.arr = q.arr[1:]

	return elem, true
}

func (q *SafeQueue[T]) Peek() (elem T, ok bool) {
	if q == nil {
		return [1]T{}[0], false
	}

	q.mux.RLock()
	defer q.mux.RUnlock()

	if q.arr == nil || len(q.arr) == 0 {
		return [1]T{}[0], false
	}

	return q.arr[0], true
}

func (q *SafeQueue[T]) Array() []T {
	q.mux.RLock()
	defer q.mux.RUnlock()

	return q.arr[:]
}

func (q *SafeQueue[T]) Len() int {
	if q == nil {
		return 0
	}

	q.mux.RLock()
	defer q.mux.RUnlock()

	return len(q.arr)
}
