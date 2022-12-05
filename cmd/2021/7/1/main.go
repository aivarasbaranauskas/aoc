package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
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

	maxPos := _a.Max(positions...)
	minSum := math.MaxInt
	for i := 0; i < maxPos; i++ {
		var sum int
		for _, position := range positions {
			sum += _a.Abs(i - position)
		}
		minSum = _a.Min(minSum, sum)
	}

	fmt.Println(minSum)
}
