package year_2023

import (
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[6] = Day6{}
}

type Day6 struct{}

func (d Day6) Part1(input []byte) string {
	times, distances := d.parseData(input)
	p1 := 1
	for i := range times {
		p1 *= d.calcNumWinWays(times[i], distances[i])
	}
	return strconv.Itoa(p1)
}

func (d Day6) Part2(input []byte) string {
	lines := strings.Split(string(input), "\n")
	time := optimistic.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	distance := optimistic.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))
	return strconv.Itoa(d.calcNumWinWays(time, distance))
}

func (Day6) calcNumWinWays(time, distance int) int {
	ct := 0
	for i := 0; i <= time; i++ {
		d := (time - i) * i
		if d > distance {
			ct++
		}
	}
	return ct
}

func (Day6) parseData(input []byte) ([]int, []int) {
	lines := strings.Split(string(input), "\n")
	times := _slice.Map(
		_slice.Filter(
			strings.Split(strings.Split(lines[0], ":")[1], " "),
			func(s string) bool {
				return s != ""
			},
		),
		optimistic.Atoi,
	)
	distances := _slice.Map(
		_slice.Filter(
			strings.Split(strings.Split(lines[1], ":")[1], " "),
			func(s string) bool {
				return s != ""
			},
		),
		optimistic.Atoi,
	)

	return times, distances
}
