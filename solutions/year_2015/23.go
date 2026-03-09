package year_2015

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[23] = Day23{}
}

type Day23 struct{}

func (day Day23) Part1(input []byte) string {
	return day.run(input, 0, 0)
}

func (day Day23) Part2(input []byte) string {
	return day.run(input, 1, 0)
}

func (day Day23) run(input []byte, a, b uint) string {
	lines := bytes.Split(input, []byte("\n"))

	var i int

Loop:
	for i < len(lines) {
		instruction := lines[i][:3]
		reg := lines[i][4]
		switch string(instruction) {
		case "hlf":
			if reg == 'a' {
				a /= 2
			} else {
				b /= 2
			}
		case "tpl":
			if reg == 'a' {
				a *= 3
			} else {
				b *= 3
			}
		case "inc":
			if reg == 'a' {
				a++
			} else {
				b++
			}
		case "jmp":
			offset := optimistic.AtoiBFast(lines[i][5:])
			if reg == '+' {
				i += offset
			} else {
				i -= offset
			}
			continue Loop
		case "jie":
			var isEven bool
			if reg == 'a' {
				isEven = a%2 == 0
			} else {
				isEven = b%2 == 0
			}
			if isEven {
				offset := optimistic.AtoiBFast(lines[i][8:])
				sign := lines[i][7]
				if sign == '+' {
					i += offset
				} else {
					i -= offset
				}
				continue Loop
			}
		case "jio":
			var isOne bool
			if reg == 'a' {
				isOne = a == 1
			} else {
				isOne = b == 1
			}
			if isOne {
				offset := optimistic.AtoiBFast(lines[i][8:])
				sign := lines[i][7]
				if sign == '+' {
					i += offset
				} else {
					i -= offset
				}
				continue Loop
			}
		}
		i++
	}

	return strconv.Itoa(int(b))
}
