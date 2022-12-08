package _slice

func Map[Tin, Tout any](in []Tin, f func(Tin) Tout) []Tout {
	out := make([]Tout, len(in))
	for i, x := range in {
		out[i] = f(x)
	}
	return out
}

func Reduce[T any](in []T, f func(T, T) T) (result T) {
	l := len(in)
	if l == 0 {
		return
	}
	result = in[0]
	for i := 1; i < l; i++ {
		result = f(result, in[i])
	}
	return
}

func ReduceInit[Tin, Tout any](in []Tin, initial Tout, f func(Tout, Tin) Tout) (result Tout) {
	l := len(in)
	if l == 0 {
		return
	}
	result = initial
	for i := 0; i < l; i++ {
		result = f(result, in[i])
	}
	return
}

func Intersect[T comparable](a ...[]T) (intersection []T) {
	switch len(a) {
	case 1:
		intersection = a[0]
	case 2:
		for _, a1 := range a[0] {
			for _, a2 := range a[1] {
				if a1 == a2 {
					intersection = append(intersection, a1)
				}
			}
		}
	default:
		intersection = Intersect(a[0], Intersect(a[1:]...))
	}

	return
}

func Duplicate[T any](a []T) []T {
	out := make([]T, len(a))
	copy(out, a)
	return out
}

func Has[T comparable](a []T, x T) bool {
	for i := range a {
		if a[i] == x {
			return true
		}
	}
	return false
}

func First[T any](a []T, check func(T) bool) int {
	for i := 0; i < len(a); i++ {
		if check(a[i]) {
			return i
		}
	}
	return -1
}

func CountUnique[T comparable](a []T) map[T]int {
	m := map[T]int{}
	for _, x := range a {
		if _, ok := m[x]; ok {
			m[x]++
		} else {
			m[x] = 1
		}
	}
	return m
}

func Filter[T any](a []T, f func(T) bool) []T {
	var a1 []T
	for _, value := range a {
		if f(value) {
			a1 = append(a1, value)
		}
	}
	return a1
}
