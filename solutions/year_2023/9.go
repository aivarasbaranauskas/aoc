package year_2023

import (
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[9] = Day9{}
}

type Day9 struct{}

func (d Day9) Part1(input []byte) string {
	lines := _slice.Map(strings.Split(string(input), "\n"),
		func(line string) []int {
			return _slice.Map(
				strings.Split(line, " "),
				optimistic.Atoi,
			)
		},
	)

	p1 := 0
	for _, line := range lines {
		var m [][]int
		for !d.allZero(line) {
			m = append(m, line)
			line = d.stepDown(line)
		}

		for i := len(m) - 1; i >= 0; i-- {
			p1 += m[i][len(m[i])-1]
		}
	}

	return strconv.Itoa(p1)
}

func (d Day9) Part2(input []byte) string {
	lines := _slice.Map(strings.Split(string(input), "\n"),
		func(line string) []int {
			return _slice.Map(
				strings.Split(line, " "),
				optimistic.Atoi,
			)
		},
	)

	p2 := 0
	for _, line := range lines {
		var m [][]int
		for !d.allZero(line) {
			m = append(m, line)
			line = d.stepDown(line)
		}

		x2 := 0
		for i := len(m) - 1; i >= 0; i-- {
			x2 = m[i][0] - x2
		}
		p2 += x2
	}

	return strconv.Itoa(p2)
}

func (Day9) stepDown(seq []int) []int {
	m := make([]int, len(seq)-1)
	for i := 0; i < len(seq)-1; i++ {
		m[i] = seq[i+1] - seq[i]
	}
	return m
}

func (Day9) allZero(m []int) bool {
	return len(m) == _slice.CountCond(m, func(i int) bool {
		return i == 0
	})
}
