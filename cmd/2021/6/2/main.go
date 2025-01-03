package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	spl := strings.Split(input, ",")
	ages := make([]int, len(spl))
	for i, v := range spl {
		ages[i] = optimistic.Atoi(v)
	}

	var fishes [9]int
	for _, age := range ages {
		fishes[age]++
	}

	for d := 0; d < 256; d++ {
		tmp := [9]int{
			6: fishes[0],
			8: fishes[0],
		}
		for i := 0; i < 8; i++ {
			tmp[i] += fishes[i+1]
		}
		copy(fishes[:], tmp[:])
	}

	fmt.Println(_num.Sum(fishes[:]...))
}
