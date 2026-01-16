package _a

func Ptr[T any](v T) *T {
	return &v
}
