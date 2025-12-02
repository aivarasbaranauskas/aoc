package year_2025

import (
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[2] = Day2{}
}

type Day2 struct{}

func (day Day2) Part1(input []byte) string {
	return day.solve(input, func(code int) bool {
		higherH := code
		lowerH := 0
		for i := 1; higherH > lowerH; i *= 10 {
			d := higherH % 10
			lowerH += d * i
			higherH /= 10
			if higherH == lowerH && d != 0 {
				return true
			}
		}
		return false
	})
}

func (day Day2) Part2(input []byte) string {
	repetitiveLengths := [][]int{
		{},
		{},
		{1},
		{1},
		{1, 2},
		{1},
		{1, 2, 3},
		{1},
		{1, 2, 4},
		{1, 3},
		{1, 2, 5},
	}

	return day.solve(input, func(code int) bool {
		codeS := strconv.Itoa(code)

	OuterLoop:
		for _, l := range repetitiveLengths[len(codeS)] {
			p1 := codeS[:l]
			for i := 1; (i+1)*l <= len(codeS); i++ {
				if p1 != codeS[i*l:(i+1)*l] {
					continue OuterLoop
				}
			}
			return true
		}
		return false
	})
}

func (Day2) solve(input []byte, isBadBarcode func(i int) bool) string {
	sum := 0

	for i := 0; i < len(input); {
		tmp := i
		for input[i] != '-' {
			i++
		}
		from := optimistic.AtoiBFast(input[tmp:i])
		i++

		tmp = i
		for i < len(input) && input[i] != ',' {
			i++
		}
		to := optimistic.AtoiBFast(input[tmp:i])
		i++

		for code := from; code <= to; code++ {
			if isBadBarcode(code) {
				sum += code
			}
		}
	}

	return strconv.Itoa(sum)
}
