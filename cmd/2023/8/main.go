package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_map"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	steps := _slice.Map([]byte(lines[0]), func(c byte) int {
		if c == 'R' {
			return 1
		}
		return 0
	})

	network := parseNetwork(lines[2:])

	p1 := countStepsP1(steps, network)
	fmt.Println("part 1:", p1)
	p2 := countStepsP2(steps, network)
	fmt.Println("part 2:", p2)
}

func countStepsP2(steps []int, network Network) int {
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

	return _a.LCM(loopSizes...)
}

func countStepsP1(steps []int, network Network) int {
	ct := 0
	cur := "AAA"
	ls := len(steps)

	for i := 0; cur != "ZZZ"; i = (i + 1) % ls {
		cur = network[cur][steps[i]]
		ct++
	}

	return ct
}

type Network map[string][2]string

func parseNetwork(lines []string) Network {
	m := make(Network)

	for _, line := range lines {
		spl := strings.Split(line, " = ")
		key := spl[0]
		spl = strings.Split(spl[1], ", ")
		m[key] = [2]string{spl[0][1:], spl[1][:3]}
	}

	return m
}
