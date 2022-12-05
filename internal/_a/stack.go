package _a

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (item T) {
	l := len(s.items) - 1
	item = s.items[l]
	s.items = s.items[:l]
	return
}

func (s *Stack[T]) Empty() bool {
	return len(s.items) == 0
}
