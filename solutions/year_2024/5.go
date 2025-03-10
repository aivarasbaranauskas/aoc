package year_2024

import (
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
)

func init() {
	Solutions[5] = Day5{}
}

type Day5 struct{}

func (day Day5) Part1(input []byte) string {
	rules, i := day.parseRules(input)
	s := 0
	var updateOriginal []int
Loop:
	for ; i < len(input); i++ {
		var update [100]int
		updateOriginal = updateOriginal[:0]
		for {
			v := optimistic.AtoiBFast(input[i : i+2])
			updateOriginal = append(updateOriginal, v)
			update[v] = len(updateOriginal)

			i += 2
			if i >= len(input) || input[i] == '\n' {
				break
			}
			i++
		}

		for ir := range rules {
			a := rules[ir][0]
			b := rules[ir][1]
			if update[b] > 0 && update[a] > update[b] {
				continue Loop
			}
		}
		s += updateOriginal[len(updateOriginal)/2]
	}

	return strconv.Itoa(s)
}

func (day Day5) Part2(input []byte) string {
	rules, i := day.parseRules(input)

	s := 0
	var updateOriginal []int
	for ; i < len(input); i++ {
		var update [100]int
		updateOriginal = updateOriginal[:0]
		for {
			v := optimistic.AtoiBFast(input[i : i+2])
			updateOriginal = append(updateOriginal, v)
			update[v] = len(updateOriginal)

			i += 2
			if i >= len(input) || input[i] == '\n' {
				break
			}
			i++
		}

		good := true
		for ir := range rules {
			a := rules[ir][0]
			b := rules[ir][1]
			if update[b] > 0 && update[a] > update[b] {
				good = false
				break
			}
		}
		if good {
			continue
		}

	Loop:
		for ir := range rules {
			a := rules[ir][0]
			b := rules[ir][1]
			if update[b] > 0 && update[a] > update[b] {
				from, to := update[b], update[a]
				for i := from - 1; i < to-1; i++ {
					update[updateOriginal[i]]++
				}
				update[a] = from

				copy(updateOriginal[from:], updateOriginal[from-1:to-1])
				updateOriginal[from-1] = a
				goto Loop
			}
		}

		s += updateOriginal[len(updateOriginal)/2]
	}

	return strconv.Itoa(s)
}

func (Day5) parseRules(input []byte) (rules [][2]int, i int) {
	ib := bytes.Index(input, []byte("\n\n"))
	rules = make([][2]int, 0, ib/6+1)

	for ; input[i] != '\n'; i += 6 {
		rules = append(
			rules,
			[2]int{
				optimistic.AtoiBFast(input[i : i+2]),
				optimistic.AtoiBFast(input[i+3 : i+5]),
			},
		)
	}
	i++
	return
}
