package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := _slice.Map(strings.Split(input, "\n"),
		func(line string) []int {
			return _slice.Map(
				strings.Split(line, " "),
				optimistic.Atoi,
			)
		},
	)

	p1 := 0
	p2 := 0
	for _, line := range lines {
		var m [][]int
		for !allZero(line) {
			m = append(m, line)
			line = stepDown(line)
		}

		x1 := 0
		x2 := 0
		for i := len(m) - 1; i >= 0; i-- {
			x1 += m[i][len(m[i])-1]
			x2 = m[i][0] - x2
		}
		p1 += x1
		p2 += x2
	}

	fmt.Println("part 1:", p1)
	fmt.Println("part 2:", p2)
}

func stepDown(seq []int) []int {
	m := make([]int, len(seq)-1)
	for i := 0; i < len(seq)-1; i++ {
		m[i] = seq[i+1] - seq[i]
	}
	return m
}

func allZero(m []int) bool {
	return len(m) == _slice.CountCond(m, func(i int) bool {
		return i == 0
	})
}
