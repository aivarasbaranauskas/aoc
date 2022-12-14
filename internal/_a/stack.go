package _a

import "sync"

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

type CStack[T any] struct {
	s  Stack[T]
	mu sync.Mutex
}

func (s *CStack[T]) Push(item T) {
	s.mu.Lock()
	s.s.Push(item)
	s.mu.Unlock()
}

func (s *CStack[T]) Pop() (item T) {
	s.mu.Lock()
	item = s.s.Pop()
	s.mu.Unlock()
	return
}

func (s *CStack[T]) Empty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.s.Empty()
}
