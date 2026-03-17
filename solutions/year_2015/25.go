package year_2015

import (
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[25] = Day25{}
}

type Day25 struct{}

func (Day25) Part1(input []byte) string {
	row := optimistic.AtoiBFast(input[80:84])
	col := optimistic.AtoiBFast(input[93:97])
	diag := row + col - 1
	whichInOrder := diag*(diag-1)/2 + col
	val := 20151125
	for range whichInOrder - 1 {
		val = val * 252533 % 33554393
	}

	return strconv.Itoa(val)
}

func (Day25) Part2(_ []byte) string {
	return ""
}
