package _a

import (
	"bufio"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"io"
	"strings"
)

func ReadIntMatrix(r io.Reader) [][]int {
	s := bufio.NewScanner(r)
	var m [][]int
	for s.Scan() {
		m = append(m, _slice.Map(strings.Split(s.Text(), ""), optimistic.Atoi))
	}
	return m
}
