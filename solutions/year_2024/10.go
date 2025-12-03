package year_2024

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/_set"
)

func init() {
	Solutions[10] = Day10{}
}

type Day10 struct{}

func (day Day10) Part1(input []byte) string {
	m := bytes.Split(input, []byte("\n"))

	score := 0

	for r := range m {
		for c := range m[r] {
			if m[r][c] != '0' {
				continue
			}

			ends := _set.New[int]()
			day.findUniqueTrailEnds(ends, m, r, c)
			score += ends.Len()
		}
	}

	return strconv.Itoa(score)
}

func (day Day10) findUniqueTrailEnds(ends *_set.Set[int], m [][]byte, r, c int) {
	if m[r][c] == '9' {
		ends.Add(r*len(m) + c)
		return
	}

	nextC := m[r][c] + 1

	if r+1 < len(m) && m[r+1][c] == nextC {
		day.findUniqueTrailEnds(ends, m, r+1, c)
	}
	if r > 0 && m[r-1][c] == nextC {
		day.findUniqueTrailEnds(ends, m, r-1, c)
	}
	if c+1 < len(m[0]) && m[r][c+1] == nextC {
		day.findUniqueTrailEnds(ends, m, r, c+1)
	}
	if c > 0 && m[r][c-1] == nextC {
		day.findUniqueTrailEnds(ends, m, r, c-1)
	}
}

func (day Day10) Part2(input []byte) string {
	m := bytes.Split(input, []byte("\n"))

	score := 0

	for r := range m {
		for c := range m[r] {
			if m[r][c] != '0' {
				continue
			}

			score += day.countUniqueTrails(m, r, c)
		}
	}

	return strconv.Itoa(score)
}

func (day Day10) countUniqueTrails(m [][]byte, r, c int) int {
	if m[r][c] == '9' {
		return 1
	}

	nextC := m[r][c] + 1
	ct := 0

	if r+1 < len(m) && m[r+1][c] == nextC {
		ct += day.countUniqueTrails(m, r+1, c)
	}
	if r > 0 && m[r-1][c] == nextC {
		ct += day.countUniqueTrails(m, r-1, c)
	}
	if c+1 < len(m[0]) && m[r][c+1] == nextC {
		ct += day.countUniqueTrails(m, r, c+1)
	}
	if c > 0 && m[r][c-1] == nextC {
		ct += day.countUniqueTrails(m, r, c-1)
	}

	return ct
}
