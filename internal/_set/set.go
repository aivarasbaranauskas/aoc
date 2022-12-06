package _set

type Set[T comparable] map[T]struct{}

func FromSlice[T comparable](s []T) Set[T] {
	set := make(Set[T], len(s))
	for _, item := range s {
		set[item] = struct{}{}
	}
	return set
}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) Has(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}
