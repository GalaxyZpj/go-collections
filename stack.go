package go_collections

type Stack[T comparable] struct {
	array []T
	Size  int
}

func (s *Stack[T]) Push(elements ...T) {
	s.array = append(s.array, elements...)
	s.Size += len(elements)
}

func (s *Stack[T]) Pop() *T {
	if s.IsEmpty() {
		return nil
	}
	ptr := &s.array[s.Size-1]
	s.array = s.array[:s.Size-1]
	s.Size--
	return ptr
}

func (s *Stack[T]) Peek() *T {
	if s.IsEmpty() {
		return nil
	}
	return &s.array[s.Size-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size == 0
}
