package year_2025

import (
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/_num"
)

func init() {
	Solutions[1] = Day1{}
}

type Day1 struct{}

func (Day1) Part1(input []byte) string {
	pos := 50
	ct := 0
	i := 0

	for i < len(input) {
		direction := input[i]
		i++

		num := 0
		for ; i < len(input) && input[i] != '\n'; i++ {
			num = num*10 + int(input[i]-'0')
		}
		i++

		if direction == 'R' {
			pos += num
		} else {
			pos -= num
		}

		pos = ((pos % 100) + 100) % 100

		if pos == 0 {
			ct++
		}
	}

	return strconv.Itoa(ct)
}

func (Day1) Part2(input []byte) string {
	pos := 50
	ct := 0
	i := 0

	for i < len(input) {
		direction := input[i]
		i++

		num := 0
		for ; i < len(input) && input[i] != '\n'; i++ {
			num = num*10 + int(input[i]-'0')
		}
		i++

		prevPos := pos
		if direction == 'R' {
			pos += num
		} else {
			pos -= num
		}

		if pos == 0 {
			ct++
		} else {
			ct += _num.Abs(pos / 100)
			if pos < 0 && prevPos != 0 {
				ct++
			}
		}

		pos = ((pos % 100) + 100) % 100
	}

	return strconv.Itoa(ct)
}
