package year_2015

import (
	"bytes"
	"strconv"
)

func init() {
	Solutions[8] = Day8{}
}

type Day8 struct{}

func (Day8) Part1(input []byte) string {
	sumLen := 0
	sumLenDecoded := 0

	for line := range bytes.Lines(input) {
		line = bytes.TrimSpace(line)
		sumLen += len(line)
		sumLenDecoded += len(line) - 2

		for i := 0; i < len(line)-1; i++ {
			if line[i] == '\\' {
				if line[i+1] == '"' || line[i+1] == '\\' {
					sumLenDecoded--
					i++
				} else {
					sumLenDecoded -= 3
					i += 3
				}
			}
		}
	}

	return strconv.Itoa(sumLen - sumLenDecoded)
}

func (Day8) Part2(input []byte) string {
	sumLen := 0
	sumLenEncoded := 0

	for line := range bytes.Lines(input) {
		line = bytes.TrimSpace(line)
		sumLen += len(line)
		sumLenEncoded += len(line) + bytes.Count(line, []byte("\\")) + bytes.Count(line, []byte("\"")) + 2
	}

	return strconv.Itoa(sumLenEncoded - sumLen)
}
