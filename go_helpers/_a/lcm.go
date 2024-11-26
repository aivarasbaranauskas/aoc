package _a

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(x ...int) int {
	result := x[0]

	for i := 1; i < len(x); i++ {
		result = result * x[i] / GCD(result, x[i])
	}

	return result
}
