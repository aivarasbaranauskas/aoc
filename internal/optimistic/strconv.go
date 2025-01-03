package optimistic

import (
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"strconv"
)

func Atoi(s string) int {
	return _a.E(strconv.Atoi(s))
}

func ParseInt(s string, base int, bitSize int) int64 {
	return _a.E(strconv.ParseInt(s, base, bitSize))
}
