package _a

// CheckErr panics if err is not nil. Useful for naive and fast error checking.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func E[T any](x T, err error) T {
	CheckErr(err)
	return x
}
