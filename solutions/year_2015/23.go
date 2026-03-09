package year_2015

import (
	"bytes"
	"fmt"
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

func (day Day23) parse(input []byte) [][3]int {
	// [instruction, register, offset]
	// instructions:
	//   0: hlf
	//   1: tpl
	//   2: inc
	//   3: jmp
	//   4: jie
	//   5: jio
	// registers:
	//   0: a
	//   1: b

	lines := bytes.Split(input, []byte("\n"))
	instructions := make([][3]int, len(lines))

	for i, line := range lines {
		// instruction
		switch string(line[:3]) {
		case "hlf":
			instructions[i][0] = 0
		case "tpl":
			instructions[i][0] = 1
		case "inc":
			instructions[i][0] = 2
		case "jmp":
			instructions[i][0] = 3
		case "jie":
			instructions[i][0] = 4
		case "jio":
			instructions[i][0] = 5
		default:
			panic(fmt.Sprintf("invalid instruction at line %d: %s", i))
		}

		// register
		if instructions[i][0] != 3 && line[4] == 'b' {
			instructions[i][1] = 1
		}

		// offset
		if instructions[i][0] >= 3 {
			var offsetBytes []byte
			if instructions[i][0] == 3 {
				offsetBytes = line[4:]
			} else {
				offsetBytes = line[7:]
			}

			instructions[i][2] = optimistic.AtoiBFast(offsetBytes[1:])
			if offsetBytes[0] == '-' {
				instructions[i][2] *= -1
			}
		}
	}

	return instructions
}

func (day Day23) run(input []byte, a, b uint) string {
	instructions := day.parse(input)

	var i int
	var regs [2]uint
	regs[0] = a
	regs[1] = b

Loop:
	for i < len(instructions) {
		instruction := instructions[i]
		switch instruction[0] {
		case 0:
			regs[instruction[1]] /= 2
		case 1:
			regs[instruction[1]] *= 3
		case 2:
			regs[instruction[1]]++
		case 3:
			i += instruction[2]
			continue Loop
		case 4:
			if regs[instruction[1]]%2 == 0 {
				i += instruction[2]
				continue Loop
			}
		case 5:
			if regs[instruction[1]] == 1 {
				i += instruction[2]
				continue Loop
			}
		}
		i++
	}

	return strconv.Itoa(int(regs[1]))
}
