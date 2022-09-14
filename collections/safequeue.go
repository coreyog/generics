package collections

import "sync"

// SafeQueue is a thread-safe FIFO queue.
type SafeQueue[T any] struct {
	arr []T
	mux *sync.RWMutex
}

// NewSafeQueue constructs a SafeQueue and queues all parameters in order.
func NewSafeQueue[T any](arr ...T) *SafeQueue[T] {
	sq := &SafeQueue[T]{
		mux: &sync.RWMutex{},
	}

	if len(arr) != 0 {
		sq.arr = arr[:]
	}

	return sq
}

// IsEmpty returns true if the queue is nil or has a length of 0.
func (q *SafeQueue[T]) IsEmpty() (is bool) {
	if q == nil {
		return true
	}

	q.mux.RLock()
	defer q.mux.RUnlock()

	return q.arr == nil || len(q.arr) == 0
}

// Enqueue adds an element to the end of the queue.
func (q *SafeQueue[T]) Enqueue(x ...T) *SafeQueue[T] {
	if q == nil {
		q = NewSafeQueue[T]()
	}

	q.mux.Lock()
	defer q.mux.Unlock()

	q.arr = append(q.arr, x...)

	return q
}

// Dequeue removes and returns the first element of the queue.
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

// Peek returns the first element of the queue without removing it.
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

// Slice returns a copy of the queue as a slice.
func (q *SafeQueue[T]) Slice() []T {
	q.mux.RLock()
	defer q.mux.RUnlock()

	return q.arr[:]
}

// Len returns the length of the queue.
func (q *SafeQueue[T]) Len() int {
	if q == nil {
		return 0
	}

	q.mux.RLock()
	defer q.mux.RUnlock()

	return len(q.arr)
}
