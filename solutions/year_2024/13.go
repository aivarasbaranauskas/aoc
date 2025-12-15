package year_2024

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[13] = Day13{}
}

type Day13 struct{}

func (day Day13) Part1(input []byte) string {
	return strconv.Itoa(day.solve(day.parseInput(input)))
}

func (day Day13) solve(machines [][6]int) int {
	sum := 0
	for _, machine := range machines {
		ax := machine[0]
		ay := machine[1]
		bx := machine[2]
		by := machine[3]
		x := machine[4]
		y := machine[5]

		top := ax*y - ay*x
		bot := ax*by - bx*ay
		if top%bot != 0 {
			continue
		}

		b := top / bot
		top2 := x - (bx * b)
		if top2%ax != 0 {
			continue
		}

		a := top2 / ax
		sum += 3*a + b
	}
	return sum
}

func (day Day13) parseInput(input []byte) [][6]int {
	ct := bytes.Count(input, []byte("\n\n"))
	machines := make([][6]int, ct)

	for i := 0; i < ct; i++ {
		for ; input[0] != '+'; input = input[1:] {
			// forward to start of the number
		}
		input = input[1:]
		ax := optimistic.AtoiBFast(input[:bytes.IndexByte(input, ',')])
		for ; input[0] != '+'; input = input[1:] {
			// forward to start of the number
		}
		input = input[1:]
		ay := optimistic.AtoiBFast(input[:bytes.IndexByte(input, '\n')])

		for ; input[0] != '+'; input = input[1:] {
			// forward to start of the number
		}
		input = input[1:]
		bx := optimistic.AtoiBFast(input[:bytes.IndexByte(input, ',')])
		for ; input[0] != '+'; input = input[1:] {
			// forward to start of the number
		}
		input = input[1:]
		by := optimistic.AtoiBFast(input[:bytes.IndexByte(input, '\n')])

		for ; input[0] != '='; input = input[1:] {
			// forward to start of the number
		}
		input = input[1:]
		x := optimistic.AtoiBFast(input[:bytes.IndexByte(input, ',')])
		for ; input[0] != '='; input = input[1:] {
			// forward to start of the number
		}
		input = input[1:]
		y := optimistic.AtoiBFast(input[:bytes.IndexByte(input, '\n')])

		machines[i] = [6]int{ax, ay, bx, by, x, y}
	}

	return machines
}

func (day Day13) Part2(input []byte) string {
	machines := day.parseInput(input)

	for i := range machines {
		machines[i][4] += 10000000000000
		machines[i][5] += 10000000000000
	}

	return strconv.Itoa(day.solve(machines))
}
