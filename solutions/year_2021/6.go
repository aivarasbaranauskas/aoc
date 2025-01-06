package year_2021

import (
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"strconv"
)

func init() {
	Solutions[6] = Day6{}
}

type Day6 struct{}

func (d Day6) Part1(input []byte) string {
	return d.solve(input, 80)
}

func (d Day6) Part2(input []byte) string {
	return d.solve(input, 256)
}

func (Day6) solve(input []byte, iterations int) string {
	var fishes [9]int
	for i := 0; i < len(input); i += 2 {
		fishes[input[i]-'0']++
	}

	var tmp [9]int
	for range iterations {
		copy(tmp[:], fishes[1:])
		tmp[6] += fishes[0]
		tmp[8] = fishes[0]
		copy(fishes[:], tmp[:])
	}

	return strconv.Itoa(_num.Sum(fishes[:]...))
}
