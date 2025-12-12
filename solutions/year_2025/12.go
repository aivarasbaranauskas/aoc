package year_2025

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[12] = Day12{}
}

type Day12 struct{}

func (Day12) Part1(input []byte) string {
	parts := bytes.Split(input, []byte("\n\n"))
	lines := bytes.Split(parts[len(parts)-1], []byte("\n"))

	partSizes := make([]int, len(parts)-1)
	for i := range partSizes {
		partSizes[i] = bytes.Count(parts[i], []byte{'#'})
	}

	ct := 0

	for _, line := range lines {
		spl1 := bytes.Split(line, []byte(": "))
		spl2 := bytes.Split(spl1[0], []byte("x"))
		spl3 := bytes.Split(spl1[1], []byte(" "))

		area := optimistic.AtoiBFast(spl2[0]) * optimistic.AtoiBFast(spl2[1])

		minimumArea := 0
		for i, v := range spl3 {
			minimumArea += partSizes[i] * optimistic.AtoiBFast(v)
		}

		if minimumArea < area {
			ct++
		}
	}

	return strconv.Itoa(ct)
}

func (Day12) Part2(_ []byte) string {
	return ""
}
