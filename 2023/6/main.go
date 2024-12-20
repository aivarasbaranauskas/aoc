package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/go_helpers/_slice"
	"github.com/aivarasbaranauskas/aoc/go_helpers/o"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")
	times := _slice.Map(
		_slice.Filter(
			strings.Split(strings.Split(lines[0], ":")[1], " "),
			func(s string) bool {
				return s != ""
			},
		),
		o.Atoi,
	)
	distances := _slice.Map(
		_slice.Filter(
			strings.Split(strings.Split(lines[1], ":")[1], " "),
			func(s string) bool {
				return s != ""
			},
		),
		o.Atoi,
	)

	p1 := 1
	for i := range times {
		p1 *= calcNumWinWays(times[i], distances[i])
	}

	time := o.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	distance := o.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))

	fmt.Println("part 1:", p1)
	fmt.Println("part 2:", calcNumWinWays(time, distance))
}

func calcNumWinWays(time, distance int) int {
	ct := 0
	for i := 0; i <= time; i++ {
		d := (time - i) * i
		if d > distance {
			ct++
		}
	}
	return ct
}
