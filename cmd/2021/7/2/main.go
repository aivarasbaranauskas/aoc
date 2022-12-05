package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/a"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	spl := strings.Split(input, ",")
	positions := make([]int, len(spl))
	var err error
	for i, v := range spl {
		positions[i], err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
	}

	maxPos := a.Max(positions...)
	minSum := math.MaxInt
	for i := 0; i < maxPos; i++ {
		var sum int
		for _, position := range positions {
			distance := a.Abs(i - position)
			if distance > 0 {
				sum += distance * (1 + distance) / 2
			}
		}
		minSum = a.Min(minSum, sum)
	}

	fmt.Println(minSum)
}
