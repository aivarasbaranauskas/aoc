package year_2015

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[2] = Day2{}
}

type Day2 struct{}

func (Day2) Part1(input []byte) string {
	lines := bytes.Split(input, []byte{'\n'})
	sum := 0

	for _, line := range lines {
		spl := bytes.Split(line, []byte{'x'})
		a, b, c := optimistic.AtoiB(spl[0]), optimistic.AtoiB(spl[1]), optimistic.AtoiB(spl[2])
		a, b, c = a*b, b*c, c*a
		smallest := min(a, b, c)
		sum += smallest + 2*(a+b+c)
	}

	return strconv.Itoa(sum)
}

func (Day2) Part2(input []byte) string {
	lines := bytes.Split(input, []byte{'\n'})
	sum := 0

	for _, line := range lines {
		spl := bytes.Split(line, []byte{'x'})
		a, b, c := optimistic.AtoiB(spl[0]), optimistic.AtoiB(spl[1]), optimistic.AtoiB(spl[2])
		sum += a*b*c + 2*(a+b+c-max(a, b, c))
	}

	return strconv.Itoa(sum)
}
