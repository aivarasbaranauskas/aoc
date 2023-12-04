package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
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
	lines := strings.Split(input, "\n")

	q := _a.Queue[int]{}
	for i := range lines {
		q.Enqueue(i)
	}

	ctTotal := 0
	m := map[int]int{}
	for !q.Empty() {
		i := q.Dequeue()

		if _, ok := m[i]; !ok {
			m[i] = countMatching(lines[i])
		}

		ct := m[i]

		for j := 1; j <= ct; j++ {
			q.Enqueue(i + j)
		}

		ctTotal++
	}
	fmt.Println("part 2:", ctTotal)
}

func p1() {
	lines := strings.Split(input, "\n")

	s := 0
	for _, line := range lines {
		ct := countMatching(line)
		if ct > 0 {
			s += 1 << (ct - 1)
		}
	}

	fmt.Println("part 1:", s)
}

func countMatching(line string) int {
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
