package _num

import "unsafe"

type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type SignedNumeric interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Sum[T Numeric](numbers ...T) (sum T) {
	for i := range numbers {
		sum += numbers[i]
	}
	return
}

func Product[T Numeric](numbers ...T) (p T) {
	p = 1
	for i := range numbers {
		p *= numbers[i]
	}
	return
}

func Sign[T SignedNumeric](x T) T {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func Abs[T SignedNumeric](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func FastBoolToInt(b bool) int {
	return int(*(*byte)(unsafe.Pointer(&b)))
}
