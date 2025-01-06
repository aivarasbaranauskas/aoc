package year_2021

import (
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[1] = Day1{}
}

type Day1 struct{}

func (Day1) Part1(input []byte) string {
	spl := strings.Split(string(input), "\n")
	m := make([]int, len(spl))
	for i, v := range spl {
		m[i] = optimistic.Atoi(v)
	}

	ct := 0
	for i := 0; i < len(m)-1; i++ {
		if m[i] < m[i+1] {
			ct++
		}
	}

	return strconv.Itoa(ct)
}

func (Day1) Part2(input []byte) string {
	spl := strings.Split(string(input), "\n")
	m := make([]int, len(spl))
	for i, v := range spl {
		m[i] = optimistic.Atoi(v)
	}

	ct := 0
	for i := 0; i < len(m)-3; i++ {
		sumA := m[i] + m[i+1] + m[i+2]
		sumB := m[i+1] + m[i+2] + m[i+3]
		if sumA < sumB {
			ct++
		}
	}

	return strconv.Itoa(ct)
}
