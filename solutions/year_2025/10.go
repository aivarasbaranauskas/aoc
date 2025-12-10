package year_2025

import (
	"bytes"
	"cmp"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[10] = Day10{}
}

type Day10 struct{}

func (day Day10) Part1(input []byte) string {
	lines := bytes.Split(input, []byte{'\n'})

	sum := 0
	for i := range lines {
		sum += day.configureLights(lines[i])
	}

	return strconv.Itoa(sum)
}

func (day Day10) configureLights(input []byte) int {
	spl := bytes.Split(input, []byte{' '})

	var goal int
	for i := range len(spl[0]) - 2 {
		if spl[0][i+1] == '#' {
			goal |= 1 << i
		}
	}

	buttons := make([]int, len(spl)-2)
	for i := 1; i < len(spl)-1; i++ {
		var button int
		bits := bytes.Split(spl[i][1:len(spl[i])-1], []byte{','})
		for _, b := range bits {
			button |= 1 << optimistic.AtoiBFast(b)
		}
		buttons[i-1] = button
	}

	var q _a.Queue[[3]int] // [current lights state, last added i, button presses count]
	q.Enqueue([3]int{0, -1, 0})

	for !q.Empty() {
		state := q.Dequeue()

		for i := state[1] + 1; i < len(buttons); i++ {
			s := state[0] ^ buttons[i]
			if s == goal {
				return state[2] + 1
			}
			q.Enqueue([3]int{s, i, state[2] + 1})
		}
	}

	panic("not found")
}

func (day Day10) Part2(input []byte) string {
	lines := bytes.Split(input, []byte{'\n'})

	sum := 0
	for i := range lines {
		sum += day.configureJoltage(lines[i])
	}

	return strconv.Itoa(sum)
}

func (day Day10) configureJoltage(input []byte) int {
	parse := func(numsB []byte) []int {
		bits := bytes.Split(numsB[1:len(numsB)-1], []byte{','})
		nums := make([]int, len(bits))
		for i := range bits {
			nums[i] = optimistic.AtoiBFast(bits[i])
		}
		return nums
	}

	spl := bytes.Split(input, []byte{' '})
	goal := parse(spl[len(spl)-1])
	buttons := make([][]int, len(spl)-2)
	for i := 1; i < len(spl)-1; i++ {
		buttons[i-1] = parse(spl[i])
	}

	type S struct {
		joltage              []int
		lastI, buttonPresses int
	}

	var q _a.Queue[S]
	q.Enqueue(S{joltage: make([]int, len(goal))})

	pushButton := func(joltage []int, button []int) []int {
		newJoltage := make([]int, len(joltage))
		copy(newJoltage, joltage)
		for _, i := range button {
			newJoltage[i]++
		}
		return newJoltage
	}

	for !q.Empty() {
		state := q.Dequeue()

	OuterLoop:
		for i := state.lastI; i < len(buttons); i++ {
			s := pushButton(state.joltage, buttons[i])

			eq := true
			for j := range s {
				if c := cmp.Compare(s[j], goal[j]); c != 0 {
					if c > 0 {
						// got past the goal, terminate
						continue OuterLoop
					}
					eq = false
				}
			}

			if eq {
				return state.buttonPresses + 1
			}

			// BFS won't do ehe - it's a ILP problem... TODO

			q.Enqueue(S{
				joltage:       s,
				lastI:         i,
				buttonPresses: state.buttonPresses + 1,
			})
		}
	}

	panic("not found")
}
