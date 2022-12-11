package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"sort"
	"strings"
)

type Monkey struct {
	Items     []int
	Operation func(int) int
	Test      func(int) int
	Inspected int
}

var monkeys []*Monkey

//go:embed input.txt
var input string

func init() {
	for _, monkeyData := range strings.Split(input, "\n\n") {
		lines := strings.Split(monkeyData, "\n")

		div := optimistic.Atoi(strings.Split(lines[3], "by ")[1])
		divTrue := optimistic.Atoi(strings.Split(lines[4], "monkey ")[1])
		divFalse := optimistic.Atoi(strings.Split(lines[5], "monkey ")[1])

		var op func(int) int

		spl := strings.Split(strings.TrimSpace(lines[2]), " ")
		if spl[5] == "old" {
			if spl[4] == "+" {
				op = func(i int) int { return i + i }
			} else if spl[4] == "*" {
				op = func(i int) int { return i * i }
			}
		} else {
			x := optimistic.Atoi(spl[5])
			if spl[4] == "+" {
				op = func(i int) int { return i + x }
			} else if spl[4] == "*" {
				op = func(i int) int { return i * x }
			}
		}

		monkeys = append(monkeys, &Monkey{
			Items:     _slice.Map(strings.Split(strings.Split(lines[1], ": ")[1], ", "), optimistic.Atoi),
			Operation: op,
			Test: func(i int) int {
				if i%div == 0 {
					return divTrue
				}
				return divFalse
			},
		})
	}
}

func main() {
	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			for _, w := range monkey.Items {
				w = monkey.Operation(w) / 3
				next := monkey.Test(w)
				monkeys[next].Items = append(monkeys[next].Items, w)
				monkey.Inspected++
			}
			monkey.Items = monkey.Items[:0]
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspected < monkeys[j].Inspected
	})
	fmt.Println(monkeys[6].Inspected, monkeys[7].Inspected)
	fmt.Println(monkeys[6].Inspected * monkeys[7].Inspected)
}
