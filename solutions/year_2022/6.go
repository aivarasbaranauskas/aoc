package year_2022

import (
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"strconv"
)

func init() {
	Solutions[6] = Day6{}
}

type Day6 struct{}

func (Day6) Part1(input []byte) string {
	var o []byte
	for i, c := range input {
		o = append(o, c)
		if i >= 3 && o[i] != o[i-1] && o[i] != o[i-2] && o[i] != o[i-3] && o[i-1] != o[i-2] && o[i-1] != o[i-3] && o[i-2] != o[i-3] {
			break
		}
	}
	return strconv.Itoa(len(o))
}

func (Day6) Part2(input []byte) string {
	for i := range input {
		if i >= 14 {
			if _set.FromSlice(input[i-14:i]).Len() == 14 {
				return strconv.Itoa(i)
			}
		}
	}
	return ""
}
