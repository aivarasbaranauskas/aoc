package optimistic

import (
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"strconv"
)

func Atoi(s string) int {
	return _a.E(strconv.Atoi(s))
}

func AtoiB(s []byte) int {
	return _a.E(strconv.Atoi(string(s)))
}

func AtoiBFast(s []byte) int {
	n := 0
	for _, ch := range s {
		n = n*10 + int(ch-'0')
	}
	return n
}

func ParseInt(s string, base int, bitSize int) int64 {
	return _a.E(strconv.ParseInt(s, base, bitSize))
}
