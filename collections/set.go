package collections

type Set[T comparable] map[T]struct{}

func ToSet[T comparable](arr []T) Set[T] {
	s := Set[T]{}

	for _, a := range arr {
		s.Add(a)
	}

	return s
}

func (s *Set[T]) IsEmpty() bool {
	return s == nil || len(*s) == 0
}

func (s *Set[T]) Add(x T) {
	if s == nil || *s == nil {
		*s = make(map[T]struct{})
	}

	(*s)[x] = struct{}{}
}

func (s *Set[T]) Delete(x T) {
	if s == nil || *s == nil {
		return
	}

	delete(*s, x)
}

func (s *Set[T]) InSet(x T) (in bool) {
	if s == nil || *s == nil {
		return false
	}

	_, in = (*s)[x]

	return in
}

func (s *Set[T]) Array() []T {
	arr := []T{}

	for x := range *s {
		arr = append(arr, x)
	}

	return arr
}
