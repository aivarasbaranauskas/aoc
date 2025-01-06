package year_2023

import (
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[4] = Day4{}
}

type Day4 struct{}

func (d Day4) Part1(input []byte) string {
	lines := strings.Split(string(input), "\n")

	s := 0
	for _, line := range lines {
		ct := d.countMatching(line)
		if ct > 0 {
			s += 1 << (ct - 1)
		}
	}

	return strconv.Itoa(s)
}

func (d Day4) Part2(input []byte) string {
	lines := strings.Split(string(input), "\n")

	q := _a.Queue[int]{}
	for i := range lines {
		q.Enqueue(i)
	}

	ctTotal := 0
	m := map[int]int{}
	for !q.Empty() {
		i := q.Dequeue()

		if _, ok := m[i]; !ok {
			m[i] = d.countMatching(lines[i])
		}

		ct := m[i]

		for j := 1; j <= ct; j++ {
			q.Enqueue(i + j)
		}

		ctTotal++
	}

	return strconv.Itoa(ctTotal)
}

func (Day4) countMatching(line string) int {
	spl := strings.Split(line, ":")
	spl = strings.Split(spl[1], "|")

	luckyNumbers := _set.New[int]()
	for _, x := range strings.Split(spl[0], " ") {
		if len(x) == 0 {
			continue
		}
		luckyNumbers.Add(optimistic.Atoi(x))
	}

	ct := 0
	for _, x := range strings.Split(spl[1], " ") {
		if len(x) == 0 {
			continue
		}

		if !luckyNumbers.Has(optimistic.Atoi(x)) {
			continue
		}

		ct++
	}
	return ct
}
