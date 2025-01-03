package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strings"
)

//go:embed input.txt
var input string

const l = 12

func main() {
	gamma := getGamma(getInput())
	epsilon := ^gamma & 0b111111111111
	fmt.Printf("Gamma:   %8v, %012b\n", gamma, gamma)
	fmt.Printf("Epsilon: %8v, %012b\n", epsilon, epsilon)
	fmt.Println("Power consumption:", gamma*epsilon)
}

func getGamma(input []uint64) (gamma uint64) {
	onesC := make([]int, l)
	zeroesC := make([]int, l)
	for _, ip := range input {
		for p := 0; p < l; p++ {
			f := uint64(1 << p)
			if ip&f == f {
				onesC[p]++
			} else {
				zeroesC[p]++
			}
		}
	}

	for i := 0; i < l; i++ {
		if onesC[i] > zeroesC[i] {
			gamma += 1 << i
		}
	}

	return
}

func getInput() []uint64 {
	lines := strings.Split(input, "\n")
	m := make([]uint64, len(lines))
	for i, line := range lines {
		m[i] = uint64(optimistic.ParseInt(line, 2, 64))
	}
	return m
}
