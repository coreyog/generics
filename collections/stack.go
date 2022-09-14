package collections

// Stack is a LIFO stack.
type Stack[T any] []T

// IsEmpty returns true if the stack is nil or has a length of 0.
func (s *Stack[T]) IsEmpty() (is bool) {
	return s == nil || len(*s) == 0
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(x ...T) {
	*s = append(*s, x...)
}

// Pop removes and returns the top element of the stack.
func (s *Stack[T]) Pop() (elem T, ok bool) {
	if s == nil || len(*s) == 0 {
		return [1]T{}[0], false
	}

	last := len(*s) - 1
	elem = (*s)[last]
	*s = (*s)[:last]

	return elem, true
}

// Peek returns the top element of the stack without removing it.
func (s *Stack[T]) Peek() (elem T, ok bool) {
	if s == nil || len(*s) == 0 {
		return [1]T{}[0], false
	}

	return (*s)[len(*s)-1], true
}
