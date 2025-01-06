package year_2021

import (
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"math"
	"slices"
	"strconv"
	"strings"
)

func init() {
	Solutions[7] = Day7{}
}

type Day7 struct{}

func (Day7) Part1(input []byte) string {
	spl := strings.Split(string(input), ",")
	positions := make([]int, len(spl))
	for i, v := range spl {
		positions[i] = optimistic.Atoi(v)
	}

	maxPos := slices.Max(positions)
	minSum := math.MaxInt
	for i := 0; i < maxPos; i++ {
		var sum int
		for _, position := range positions {
			sum += _num.Abs(i - position)
		}
		minSum = min(minSum, sum)
	}

	return strconv.Itoa(minSum)
}

func (Day7) Part2(input []byte) string {
	spl := strings.Split(string(input), ",")
	positions := make([]int, len(spl))
	for i, v := range spl {
		positions[i] = optimistic.Atoi(v)
	}

	maxPos := slices.Max(positions)
	minSum := math.MaxInt
	for i := 0; i < maxPos; i++ {
		var sum int
		for _, position := range positions {
			distance := _num.Abs(i - position)
			if distance > 0 {
				sum += distance * (1 + distance) / 2
			}
		}
		minSum = min(minSum, sum)
	}

	return strconv.Itoa(minSum)
}
