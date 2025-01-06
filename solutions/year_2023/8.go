package year_2023

import (
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_map"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"strconv"
	"strings"
)

func init() {
	Solutions[8] = Day8{}
}

type Day8 struct{}

func (d Day8) Part1(input []byte) string {
	steps, network := d.parseData(input)
	ct := 0
	cur := "AAA"
	ls := len(steps)

	for i := 0; cur != "ZZZ"; i = (i + 1) % ls {
		cur = network[cur][steps[i]]
		ct++
	}

	return strconv.Itoa(ct)
}

func (d Day8) Part2(input []byte) string {
	steps, network := d.parseData(input)
	starts := _slice.Filter(
		_map.Keys(network),
		func(s string) bool {
			return s[2] == 'A'
		},
	)

	ls := len(steps)
	loopSizes := make([]int, len(starts))
	for i, cur := range starts {
		ct := 0

		for stepI := 0; cur[2] != 'Z'; stepI = (stepI + 1) % ls {
			cur = network[cur][steps[stepI]]
			ct++
		}

		loopSizes[i] = ct
	}

	return strconv.Itoa(_a.LCM(loopSizes...))
}

func (Day8) parseData(input []byte) ([]int, map[string][2]string) {
	lines := strings.Split(string(input), "\n")

	steps := _slice.Map([]byte(lines[0]), func(c byte) int {
		if c == 'R' {
			return 1
		}
		return 0
	})

	network := make(map[string][2]string)
	for _, line := range lines[2:] {
		spl := strings.Split(line, " = ")
		key := spl[0]
		spl = strings.Split(spl[1], ", ")
		network[key] = [2]string{spl[0][1:], spl[1][:3]}
	}

	return steps, network
}
