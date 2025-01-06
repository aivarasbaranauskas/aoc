package year_2022

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"sort"
	"strconv"
)

func init() {
	Solutions[1] = Day1{}
}

type Day1 struct{}

func (Day1) Part1(input []byte) string {
	var x, maxVal int

	r := bufio.NewScanner(bytes.NewReader(input))
	for r.Scan() {
		line := r.Text()
		if line == "" {
			maxVal = max(maxVal, x)
			x = 0
		} else {
			x += optimistic.Atoi(line)
		}
	}
	maxVal = max(maxVal, x)

	return strconv.Itoa(maxVal)
}

func (Day1) Part2(input []byte) string {
	var (
		x int
		s []int
	)

	r := bufio.NewScanner(bytes.NewReader(input))
	for r.Scan() {
		line := r.Text()
		if line == "" {
			s = append(s, x)
			x = 0
		} else {
			x += optimistic.Atoi(line)
		}
	}
	s = append(s, x)

	sort.Ints(s)
	l := len(s)

	return strconv.Itoa(s[l-3] + s[l-2] + s[l-1])
}
