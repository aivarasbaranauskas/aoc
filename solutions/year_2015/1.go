package year_2015

import "strconv"

func init() {
	Solutions[1] = Day1{}
}

type Day1 struct{}

func (Day1) Part1(input []byte) string {
	floor := 0
	for i := range input {
		if input[i] == '(' {
			floor++
		} else {
			floor--
		}
	}
	return strconv.Itoa(floor)
}

func (Day1) Part2(input []byte) string {
	floor := 0
	for i := range input {
		if input[i] == '(' {
			floor++
		} else {
			floor--
		}
		if floor == -1 {
			return strconv.Itoa(i + 1)
		}
	}
	return "not found"
}
