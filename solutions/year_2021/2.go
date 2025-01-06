package year_2021

import (
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[2] = Day2{}
}

type Day2 struct{}

func (Day2) Part1(input []byte) string {
	horizontalPosition := 0
	depth := 0

	for _, line := range strings.Split(string(input), "\n") {
		spl := strings.Split(line, " ")
		value := optimistic.Atoi(spl[1])
		switch spl[0] {
		case "forward":
			horizontalPosition += value
		case "up":
			depth -= value
		case "down":
			depth += value
		}
	}

	return strconv.Itoa(horizontalPosition * depth)
}

func (Day2) Part2(input []byte) string {
	aim := 0
	horizontalPosition := 0
	depth := 0

	for _, line := range strings.Split(string(input), "\n") {
		spl := strings.Split(line, " ")
		value := optimistic.Atoi(spl[1])
		switch spl[0] {
		case "forward":
			horizontalPosition += value
			depth += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}

	return strconv.Itoa(horizontalPosition * depth)
}
