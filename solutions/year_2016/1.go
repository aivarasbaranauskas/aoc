package year_2016

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[1] = Day1{}
}

type Day1 struct{}

func (Day1) Part1(input []byte) string {
	d := 0
	ds := [...][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	p := [2]int{0, 0}
	for v := range bytes.SplitSeq(input, []byte(", ")) {
		if v[0] == 'R' {
			d++
		} else {
			d--
		}
		d = (d + 4) % 4
		l := optimistic.AtoiBFast(v[1:])
		p[0] += l * ds[d][0]
		p[1] += l * ds[d][1]
	}

	return strconv.Itoa(_num.Abs(p[0]) + _num.Abs(p[1]))
}

func (Day1) Part2(input []byte) string {
	d := 0
	ds := [...][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	p := [2]int{0, 0}
	lines := make([][3]int, 0, bytes.Count(input, []byte(",")))
	for v := range bytes.SplitSeq(input, []byte(", ")) {
		if v[0] == 'R' {
			d++
		} else {
			d--
		}
		d = (d + 4) % 4
		l := optimistic.AtoiBFast(v[1:])
		p2 := [2]int{
			p[0] + l*ds[d][0],
			p[1] + l*ds[d][1],
		}
		a, b, c := p[0], p2[0], p[1]
		if a == b {
			a, b, c = p[1], p2[1], p[0]
		}
		a, b = min(a, b), max(a, b)
		for i := len(lines) - 3; i >= 0; i -= 2 {
			if a <= lines[i][2] && lines[i][2] <= b && lines[i][0] <= c && c <= lines[i][1] {
				return strconv.Itoa(_num.Abs(lines[i][2]) + _num.Abs(c))
			}
		}

		lines = append(lines, [3]int{a, b, c})
		p = p2
	}

	panic("not found")
}
