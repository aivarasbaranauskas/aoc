package _a

type Queue[T any] struct {
	items []T
}

func (s *Queue[T]) Enqueue(item T) {
	s.items = append(s.items, item)
}

func (s *Queue[T]) Dequeue() (item T) {
	item = s.items[0]
	s.items = s.items[1:]
	return
}

func (s *Queue[T]) Len() int {
	return len(s.items)
}

func (s *Queue[T]) Empty() bool {
	return len(s.items) == 0
}
