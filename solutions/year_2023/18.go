package year_2023

import (
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[18] = Day18{}
}

type Day18 struct{}

func (day Day18) Part1(input []byte) string {
	p := [2]int{0, 0}
	vertices := [][2]int{p}
	ds := map[string][2]int{
		"R": {0, 1},
		"L": {0, -1},
		"D": {1, 0},
		"U": {-1, 0},
	}
	for _, line := range strings.Split(string(input), "\n") {
		spl := strings.Split(line, " ")
		d := ds[spl[0]]
		n := optimistic.Atoi(spl[1])

		p[0] += d[0] * n
		p[1] += d[1] * n
		vertices = append(vertices, p)
	}

	return strconv.Itoa(day.calcArea(vertices))
}

func (day Day18) Part2(input []byte) string {
	p := [2]int{0, 0}
	vertices := [][2]int{p}
	ds := map[byte][2]int{
		'0': {0, 1},
		'2': {0, -1},
		'1': {1, 0},
		'3': {-1, 0},
	}
	for _, line := range strings.Split(string(input), "\n") {
		spl := strings.Split(line, " ")
		hex := spl[2][2:8]
		n, _ := strconv.ParseInt(hex[:5], 16, 64)
		d := ds[hex[5]]

		p[0] += d[0] * int(n)
		p[1] += d[1] * int(n)
		vertices = append(vertices, p)
	}

	return strconv.Itoa(day.calcArea(vertices))
}

func (Day18) calcArea(v [][2]int) int {
	// A = 0.5 * |(x1*y2 - x2*y1) + (x2*y3 - x3*y2) + ... + (xn*y1 - x1*yn)|
	// to compensate for non0 width line add (perimeter/2+1)
	area := 0
	per := 0
	for i := 0; i < len(v)-1; i++ {
		area += v[i][0]*v[i+1][1] - v[i+1][0]*v[i][1]
		per += _num.Abs(v[i][0]-v[i+1][0]) + _num.Abs(v[i][1]-v[i+1][1])
	}
	area += v[len(v)-1][0]*v[0][1] - v[0][0]*v[len(v)-1][1]
	area = _num.Abs(area)/2 + per/2 + 1
	return area
}
