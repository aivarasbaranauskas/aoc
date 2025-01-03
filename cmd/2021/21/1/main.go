package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	spl := _slice.Map(strings.Split(input, "\n"), func(tin string) int {
		return optimistic.Atoi(strings.Split(tin, ": ")[1])
	})
	p1, p2 := spl[0], spl[1]
	p1--
	p2--
	p1s, p2s := 0, 0
	i := 0
	ct := 0
	turn := 1
	for p1s < 1000 && p2s < 1000 {
		if turn == 1 {
			p1 += i + 1
			i = (i + 1) % 100
			p1 += i + 1
			i = (i + 1) % 100
			p1 += i + 1
			i = (i + 1) % 100
			p1 = p1 % 10
			p1s += p1 + 1
			turn = 2
		} else {
			p2 += i + 1
			i = (i + 1) % 100
			p2 += i + 1
			i = (i + 1) % 100
			p2 += i + 1
			i = (i + 1) % 100
			p2 = p2 % 10
			p2s += p2 + 1
			turn = 1
		}
		ct += 3
	}

	fmt.Println(p1s, p2s, ct)
	fmt.Println(min(p1s, p2s) * ct)
}
