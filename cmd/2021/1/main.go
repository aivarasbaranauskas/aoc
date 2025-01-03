package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(countIncreases(getMeasurements()))
	fmt.Println(countIncreasesSlidingWindow(getMeasurements()))
}

func countIncreases(measurements []int) int {
	ct := 0
	for i := 0; i < len(measurements)-1; i++ {
		if measurements[i] < measurements[i+1] {
			ct++
		}
	}
	return ct
}

func countIncreasesSlidingWindow(measurements []int) int {
	ct := 0
	for i := 0; i < len(measurements)-3; i++ {
		sumA := measurements[i] + measurements[i+1] + measurements[i+2]
		sumB := measurements[i+1] + measurements[i+2] + measurements[i+3]
		if sumA < sumB {
			ct++
		}
	}
	return ct
}

func getMeasurements() []int {
	spl := strings.Split(input, "\n")
	m := make([]int, len(spl))
	for i, v := range spl {
		m[i] = optimistic.Atoi(v)
	}
	return m
}
