package year_2024

import (
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
)

func init() {
	Solutions[3] = Day3{}
}

type Day3 struct{}

func (day Day3) Part1(input []byte) string {
	return strconv.Itoa(day.collectMultiplications(input))
}

func (day Day3) Part2(input []byte) string {
	i := 0
	s := 0
	dont := []byte("don't()")
	do := []byte("do()")
	for i < len(input) {
		next := bytes.Index(input[i:], dont)
		if next == -1 {
			s += day.collectMultiplications(input[i:])
			break
		}

		s += day.collectMultiplications(input[i : i+next])
		i += next
		next = bytes.Index(input[i+7:], do)
		if next == -1 {
			break
		}
		i += next
	}

	return strconv.Itoa(s)
}

func (day Day3) collectMultiplications(input []byte) int {
	s := 0
	for i := 0; i < len(input); i++ {
		if input[i] == 'm' &&
			input[i+1] == 'u' &&
			input[i+2] == 'l' &&
			input[i+3] == '(' &&
			input[i+4] >= '0' &&
			input[i+4] <= '9' {
			//	might be multiplication
			i += 4
			aB := i
			for i < len(input) && input[i] >= '0' && input[i] <= '9' {
				i++
			}
			aE := i
			if input[i] != ',' {
				continue
			}
			i++
			if input[i] < '0' || input[i] > '9' {
				continue
			}
			bB := i
			for i < len(input) && input[i] >= '0' && input[i] <= '9' {
				i++
			}
			bE := i
			if input[i] != ')' {
				continue
			}

			a := optimistic.AtoiB(input[aB:aE])
			b := optimistic.AtoiB(input[bB:bE])
			s += a * b
		}
	}

	return s
}
