package _slice

import "cmp"

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
		intersection = Intersect[T](a[0], Intersect[T](a[1:]...))
	}

	return
}

func Duplicate[T any](a []T) []T {
	out := make([]T, len(a))
	copy(out, a)
	return out
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

func CountCond[T any](a []T, f func(T) bool) int {
	ct := 0
	for _, x := range a {
		if f(x) {
			ct++
		}
	}
	return ct
}

func Count[T comparable](a []T, v T) int {
	ct := 0
	for _, x := range a {
		if x == v {
			ct++
		}
	}
	return ct
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

func Remove[T any](a []T, i int) ([]T, T) {
	v := a[i]
	return append(a[:i], a[i+1:]...), v
}

func MaxI[S ~[]E, E cmp.Ordered](x S) (int, E) {
	if len(x) < 1 {
		panic("slices.Max: empty list")
	}
	m := x[0]
	mi := 0
	for i := 1; i < len(x); i++ {
		if x[i] > m {
			mi = i
			m = x[i]
		}
	}
	return mi, m
}
