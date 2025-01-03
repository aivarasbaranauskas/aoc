package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"math"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	spl := strings.Split(input, ",")
	positions := make([]int, len(spl))
	for i, v := range spl {
		positions[i] = optimistic.Atoi(v)
	}

	maxPos := slices.Max(positions)
	minSum := math.MaxInt
	for i := 0; i < maxPos; i++ {
		var sum int
		for _, position := range positions {
			distance := _num.Abs(i - position)
			if distance > 0 {
				sum += distance * (1 + distance) / 2
			}
		}
		minSum = min(minSum, sum)
	}

	fmt.Println(minSum)
}
