package year_2022

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[18] = Day18{}
}

type Day18 struct{}

func (d Day18) Part1(input []byte) string {
	var points [][3]int

	r := bufio.NewScanner(bytes.NewReader(input))
	for r.Scan() {
		spl := strings.Split(r.Text(), ",")
		points = append(
			points,
			[3]int{
				optimistic.Atoi(spl[0]),
				optimistic.Atoi(spl[1]),
				optimistic.Atoi(spl[2]),
			},
		)
	}

	sides := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			if d.manhattan(points[i], points[j]) == 1 {
				sides++
			}
		}
	}

	return strconv.Itoa(len(points)*6 - sides*2)
}

func (Day18) manhattan(a, b [3]int) int {
	return _num.Abs(a[0]-b[0]) + _num.Abs(a[1]-b[1]) + _num.Abs(a[2]-b[2])
}

func (Day18) Part2(input []byte) string {
	points := _set.New[[3]int]()
	var xh, yh, zh [2]int

	r := bufio.NewScanner(bytes.NewReader(input))
	for r.Scan() {
		spl := strings.Split(r.Text(), ",")
		p := [3]int{
			optimistic.Atoi(spl[0]),
			optimistic.Atoi(spl[1]),
			optimistic.Atoi(spl[2]),
		}
		points.Add(p)
		xh[0] = min(xh[0], p[0])
		xh[1] = max(xh[1], p[0])
		yh[0] = min(yh[0], p[1])
		yh[1] = max(yh[1], p[1])
		zh[0] = min(zh[0], p[2])
		zh[1] = max(zh[1], p[2])
	}

	xh[0]--
	xh[1]++
	yh[0]--
	yh[1]++
	zh[0]--
	zh[1]++

	visited := _set.New[[3]int]()
	toVisit := _a.Queue[[3]int]{}
	toVisit.Enqueue([3]int{xh[0], yh[0], zh[0]})
	ct := 0

	for !toVisit.Empty() {
		current := toVisit.Dequeue()
		if visited.Has(current) {
			continue
		}
		if points.Has(current) {
			ct++
			continue
		}
		visited.Add(current)
		if current[0]-1 >= xh[0] {
			toVisit.Enqueue([3]int{
				current[0] - 1,
				current[1],
				current[2],
			})
		}
		if current[0]+1 <= xh[1] {
			toVisit.Enqueue([3]int{
				current[0] + 1,
				current[1],
				current[2],
			})
		}

		if current[1]-1 >= yh[0] {
			toVisit.Enqueue([3]int{
				current[0],
				current[1] - 1,
				current[2],
			})
		}
		if current[1]+1 <= yh[1] {
			toVisit.Enqueue([3]int{
				current[0],
				current[1] + 1,
				current[2],
			})
		}

		if current[2]-1 >= zh[0] {
			toVisit.Enqueue([3]int{
				current[0],
				current[1],
				current[2] - 1,
			})
		}
		if current[2]+1 <= zh[1] {
			toVisit.Enqueue([3]int{
				current[0],
				current[1],
				current[2] + 1,
			})
		}
	}

	return strconv.Itoa(ct)
}
