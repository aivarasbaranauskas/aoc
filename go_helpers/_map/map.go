package _map

import "github.com/aivarasbaranauskas/aoc/go_helpers/_num"

func Keys[TKey comparable, TValue any](m map[TKey]TValue) []TKey {
	a := make([]TKey, 0, len(m))
	for key := range m {
		a = append(a, key)
	}
	return a
}

func Values[TKey comparable, TValue any](m map[TKey]TValue) []TValue {
	a := make([]TValue, 0, len(m))
	for _, value := range m {
		a = append(a, value)
	}
	return a
}

func IsSet[TKey comparable, TValue any](m map[TKey]TValue, key TKey) (ok bool) {
	_, ok = m[key]
	return
}

func Duplicate[TKey comparable, TValue any](m map[TKey]TValue) map[TKey]TValue {
	newM := make(map[TKey]TValue, len(m))
	for key, value := range m {
		newM[key] = value
	}
	return newM
}

func Max[TKey comparable, TValue _num.Numeric](m map[TKey]TValue) (kMax TKey, vMax TValue) {
	var f bool

	for key, value := range m {
		if !f {
			f = true
			kMax, vMax = key, value
			continue
		}
		if vMax < value {
			kMax, vMax = key, value
		}
	}
	return
}
