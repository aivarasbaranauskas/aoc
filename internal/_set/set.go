package _set

type Set[T comparable] map[T]struct{}

func New[T comparable]() *Set[T] {
	return &Set[T]{}
}

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

func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s Set[T]) Remove(item T) {
	delete(s, item)
}

func (s Set[T]) ToSlice() []T {
	sl := make([]T, 0, len(s))
	for item := range s {
		sl = append(sl, item)
	}
	return sl
}

func (s Set[T]) MinBy(f func(T) int) T {
	if len(s) == 0 {
		panic("empty set")
	}
	var (
		x        T
		v        int
		notFirst bool
	)

	for item := range s {
		if !notFirst {
			notFirst = true
			x = item
			v = f(item)
			continue
		}
		if vI := f(item); vI < v {
			v = vI
			x = item
		}
	}
	return x
}

func (s Set[T]) Len() int {
	return len(s)
}
