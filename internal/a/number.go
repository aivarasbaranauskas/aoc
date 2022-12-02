package a

type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Max[T Numeric](numbers ...T) (max T) {
	max = numbers[0]
	for i := 1; i < len(numbers); i++ {
		if max < numbers[i] {
			max = numbers[i]
		}
	}
	return
}

func Min[T Numeric](numbers ...T) (min T) {
	min = numbers[0]
	for i := 1; i < len(numbers); i++ {
		if min > numbers[i] {
			min = numbers[i]
		}
	}
	return
}

func Sum[T Numeric](numbers ...T) (sum T) {
	for i := range numbers {
		sum += numbers[i]
	}
	return
}
