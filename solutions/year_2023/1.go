package year_2023

import (
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"strconv"
)

func init() {
	Solutions[1] = Day1{}
}

type Day1 struct{}

func (Day1) Part1(input []byte) string {
	return strconv.Itoa(_slice.Reduce(
		_slice.Map(
			bytes.Split(input, []byte("\n")),
			func(in []byte) int {
				nums := _slice.Filter(
					in,
					func(b byte) bool {
						return b >= '0' && b <= '9'
					},
				)

				return (int(nums[0]-'0') * 10) + (int(nums[len(nums)-1] - '0'))

			}),
		func(a, b int) int { return a + b },
	))
}

func (d Day1) Part2(input []byte) string {
	p2 := 0

	for _, line := range bytes.Split(input, []byte("\n")) {
		first, last := -1, -1
		for i, b := range line {
			if b >= '0' && b <= '9' {
				first = int(b - '0')
				break
			}

			if num, ok := d.findNum(line[:i+1]); ok {
				first = num
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			b := line[i]
			if b >= '0' && b <= '9' {
				last = int(b - '0')
				break
			}

			if num, ok := d.findNumRev(line[i:]); ok {
				last = num
				break
			}
		}

		p2 += first*10 + last
	}
	return strconv.Itoa(p2)
}

func (Day1) findNum(buf []byte) (int, bool) {
	nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i, num := range nums {
		if len(buf) >= len(num) && string(buf[len(buf)-len(num):]) == num {
			return i + 1, true
		}
	}

	return 0, false
}

func (Day1) findNumRev(buf []byte) (int, bool) {
	nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i, num := range nums {
		if len(buf) >= len(num) && string(buf[:len(num)]) == num {
			return i + 1, true
		}
	}

	return 0, false
}
