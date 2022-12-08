package _map

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

func IsSet[TKey comparable, TValue any](m map[TKey]TValue, key TKey) bool {
	_, ok := m[key]
	return ok
}

func Duplicate[TKey comparable, TValue any](m map[TKey]TValue) map[TKey]TValue {
	newM := make(map[TKey]TValue, len(m))
	for key, value := range m {
		newM[key] = value
	}
	return newM
}
