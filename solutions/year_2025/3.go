package year_2025

import (
	"bytes"
	"slices"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/_slice"
)

func init() {
	Solutions[3] = Day3{}
}

type Day3 struct{}

func (Day3) Part1(input []byte) string {
	m := bytes.Split(input, []byte("\n"))
	sum := 0

	for _, line := range m {
		d1i, d1 := _slice.MaxI(line[:len(line)-1])
		d2 := slices.Max(line[d1i+1:])
		sum += int(d1-'0')*10 + int(d2-'0')
	}

	return strconv.Itoa(sum)
}

func (Day3) Part2(input []byte) string {
	m := bytes.Split(input, []byte("\n"))
	sum := 0

	for _, line := range m {
		num := 0
		for i := 0; i < 12; i++ {
			id, d := _slice.MaxI(line[:len(line)-(11-i)])
			line = line[id+1:]
			num = num*10 + int(d-'0')
		}
		sum += num
	}

	return strconv.Itoa(sum)
}
