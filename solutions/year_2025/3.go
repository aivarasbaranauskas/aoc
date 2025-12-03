package year_2025

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[3] = Day3{}
}

type Day3 struct{}

func (Day3) Part1(input []byte) string {
	sum := 0

	for i := 0; i < len(input); i++ {
		d1 := input[i]
		d2 := input[i+1]
		i++

		for ; i < len(input) && input[i] != '\n'; i++ {
			if input[i] > d1 && i+1 < len(input) && input[i+1] != '\n' {
				d1 = input[i]
				d2 = input[i+1]
			} else if input[i] > d2 {
				d2 = input[i]
			}
		}

		sum += int(d1-'0')*10 + int(d2-'0')
	}

	return strconv.Itoa(sum)
}

func (Day3) Part2(input []byte) string {
	sum := 0

	for len(input) > 0 {
		var buf [12]byte

		lineEnd := bytes.IndexByte(input, '\n')
		if lineEnd == -1 {
			lineEnd = len(input)
		}

	OuterLoop:
		for i := 0; i < lineEnd; i++ {
			for j := 0; j < 12; j++ {
				if 12-j > lineEnd-i {
					continue
				}
				if buf[j] == 0 {
					buf[j] = input[i]
					continue OuterLoop
				}
				if input[i] > buf[j] {
					buf[j] = input[i]
					for j2 := j + 1; j2 < 12; j2++ {
						buf[j2] = 0
					}
					continue OuterLoop
				}
			}
		}

		sum += optimistic.AtoiBFast(buf[:])

		if lineEnd < len(input) {
			input = input[lineEnd+1:]
		} else {
			input = input[lineEnd:]
		}
	}

	return strconv.Itoa(sum)
}
