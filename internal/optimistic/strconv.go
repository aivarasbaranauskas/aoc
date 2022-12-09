package optimistic

import (
	"strconv"
)

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func ParseInt(s string, base int, bitSize int) int64 {
	i, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		panic(err)
	}
	return i
}

func ParseUint(s string, base int, bitSize int) uint64 {
	i, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		panic(err)
	}
	return i
}
