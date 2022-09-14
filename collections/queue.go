package collections

// Queue is a FIFO queue.
type Queue[T any] []T

// IsEmpty returns true if the queue is nil or has a length of 0.
func (q *Queue[T]) IsEmpty() (is bool) {
	return q == nil || len(*q) == 0
}

// Enqueue adds an element to the end of the queue.
func (q *Queue[T]) Enqueue(x ...T) {
	*q = append(*q, x...)
}

// Dequeue removes and returns the first element of the queue.
func (q *Queue[T]) Dequeue() (elem T, ok bool) {
	if q == nil || len(*q) == 0 {
		return [1]T{}[0], false
	}

	elem = (*q)[0]
	*q = (*q)[1:]

	return elem, true
}

// Peek returns the first element of the queue without removing it.
func (q *Queue[T]) Peek() (elem T, ok bool) {
	if q == nil || len(*q) == 0 {
		return [1]T{}[0], false
	}

	return (*q)[0], true
}
