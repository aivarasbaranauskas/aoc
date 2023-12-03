package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	p1()
	p2()
}

func p2() {
	s := 0
	for _, line := range strings.Split(input, "\n") {
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

	fmt.Println("part 2:", s)
}

func p1() {
	s := 0
LineLoop:
	for i, line := range strings.Split(input, "\n") {
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

	fmt.Println("part 1:", s)
}
