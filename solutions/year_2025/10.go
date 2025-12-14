package year_2025

import (
	"bytes"
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
	for _, v := range lines {
		sum += day.configureJoltage(v)
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
	buttonsI := make([]int, len(spl)-2)
	for i := 1; i < len(spl)-1; i++ {
		buttons[i-1] = parse(spl[i])
		for _, v := range buttons[i-1] {
			buttonsI[i-1] |= 1 << v
		}
	}

	combos := map[int][]int{0: {0}} // buttons masks
	var q _a.Queue[[4]int]          // [current state, last clicked button i, buttons presses count, buttons mask]
	q.Enqueue([4]int{0, -1, 0, 0})

	for !q.Empty() {
		state := q.Dequeue()

		for i := state[1] + 1; i < len(buttons); i++ {
			s := state[0] ^ buttonsI[i]
			c := state[2] + 1
			mask := state[3] | (1 << i)
			q.Enqueue([4]int{s, i, c, mask})

			if cc, ok := combos[s]; ok {
				combos[s] = append(cc, mask)
			} else {
				combos[s] = []int{mask}
			}
		}
	}

	var findMin func([]int, int, int) (int, bool)

	findMin = func(goal []int, ct, mul int) (int, bool) {
		allZero := true
		for _, v := range goal {
			allZero = allZero && (v == 0)
		}
		if allZero {
			return ct, true
		}

		goalMask := 0
		for i, v := range goal {
			goalMask |= (v % 2) << i
		}

		matchingCombos, ok := combos[goalMask]
		if !ok {
			return 0, false
		}

		goalTmp := make([]int, len(goal))
		found := false
		minV := 1 << 30
		for _, buttonsMask := range matchingCombos {
			copy(goalTmp, goal)
			ctTmp := ct
			for i, button := range buttons {
				if buttonsMask&(1<<i) > 0 {
					ctTmp += mul
					for _, j := range button {
						goalTmp[j]--
					}
				}
			}
			for i := range goalTmp {
				goalTmp[i] /= 2
			}
			rez, ok := findMin(goalTmp, ctTmp, mul*2)
			if ok {
				found = true
				minV = min(minV, rez)
			}
		}
		return minV, found
	}

	ct, ok := findMin(goal, 0, 1)
	if !ok {
		panic("not found")
	}

	return ct
}
