package year_2015

import (
	"bytes"
	"strconv"
)

func init() {
	Solutions[5] = Day5{}
}

type Day5 struct{}

func (Day5) Part1(input []byte) string {
	ct := 0

Loop:
	for line := range bytes.Lines(input) {
		var vowels int
		var seenDouble bool
		for i := range line {
			if line[i] == 'a' || line[i] == 'e' || line[i] == 'i' || line[i] == 'o' || line[i] == 'u' {
				vowels++
			}

			if i < len(line)-1 {
				if line[i] == line[i+1] {
					seenDouble = true
				}

				if line[i] == 'a' && line[i+1] == 'b' ||
					line[i] == 'c' && line[i+1] == 'd' ||
					line[i] == 'p' && line[i+1] == 'q' ||
					line[i] == 'x' && line[i+1] == 'y' {
					continue Loop
				}
			}
		}

		if !seenDouble {
			continue
		}

		if vowels < 3 {
			continue
		}

		ct++
	}

	return strconv.Itoa(ct)
}

func (Day5) Part2(input []byte) string {
	ct := 0

	for line := range bytes.Lines(input) {
		var seenSpaced, seenDoubleDouble bool

		for i := range len(line) - 3 {
			if i < len(line)-4 {
				if bytes.Index(line[i+2:], line[i:i+2]) > -1 {
					seenDoubleDouble = true
				}

			}

			if line[i] == line[i+2] {
				seenSpaced = true
			}
		}

		if seenDoubleDouble && seenSpaced {
			ct++
		}
	}

	return strconv.Itoa(ct)
}
