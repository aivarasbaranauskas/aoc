package year_2023

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
	s := 0
LineLoop:
	for i, line := range strings.Split(string(input), "\n") {
		for _, draw := range strings.Split(strings.Split(line, ": ")[1], "; ") {
			for _, ballSet := range strings.Split(draw, ", ") {
				spl := strings.Split(ballSet, " ")
				ct := optimistic.Atoi(spl[0])
				switch spl[1] {
				case "red":
					if ct > 12 {
						continue LineLoop
					}
				case "green":
					if ct > 13 {
						continue LineLoop
					}
				case "blue":
					if ct > 14 {
						continue LineLoop
					}
				}
			}
		}

		s += i + 1
	}

	return strconv.Itoa(s)
}

func (Day2) Part2(input []byte) string {
	s := 0
	for _, line := range strings.Split(string(input), "\n") {
		r, g, b := 0, 0, 0
		for _, draw := range strings.Split(strings.Split(line, ": ")[1], "; ") {
			for _, ballSet := range strings.Split(draw, ", ") {
				spl := strings.Split(ballSet, " ")
				ct := optimistic.Atoi(spl[0])
				switch spl[1] {
				case "red":
					r = max(r, ct)
				case "green":
					g = max(g, ct)
				case "blue":
					b = max(b, ct)
				}
			}
		}

		s += r * g * b
	}

	return strconv.Itoa(s)
}
