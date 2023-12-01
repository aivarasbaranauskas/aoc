package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
)

//go:embed input.txt
var input []byte

func main() {
	p1 := _slice.Reduce(
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
	)

	p2 := 0

	for _, line := range bytes.Split(input, []byte("\n")) {
		first, last := -1, -1
		for i, b := range line {
			if b >= '0' && b <= '9' {
				first = int(b - '0')
				break
			}

			if num, ok := findNum(line[:i+1]); ok {
				first = num
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			b := line[i]
			if b >= '0' && b <= '9' {
				last = int(b - '0')
				break
			}

			if num, ok := findNumRev(line[i:]); ok {
				last = num
			}
		}

		p2 += first*10 + last
	}

	fmt.Println("part 1:", p1)
	fmt.Println("part 2:", p2)
}

func findNum(buf []byte) (int, bool) {
	nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i, num := range nums {
		if len(buf) >= len(num) && string(buf[len(buf)-len(num):]) == num {
			return i + 1, true
		}
	}

	return 0, false
}

func findNumRev(buf []byte) (int, bool) {
	nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i, num := range nums {
		if len(buf) >= len(num) && string(buf[:len(num)]) == num {
			return i + 1, true
		}
	}

	return 0, false
}
