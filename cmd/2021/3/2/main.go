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
	oxygen := getOxygenRating(getInput(), 12)
	co2 := getCO2Rating(getInput(), 12)
	fmt.Println("Oxygen:", oxygen)
	fmt.Println("CO2:", co2)
	fmt.Println("Life support rating:", oxygen*co2)
}

func getOxygenRating(input []uint64, l int) uint64 {
	inputCopy := make([]uint64, len(input))
	copy(inputCopy, input)
	return getRating(inputCopy, l, func(onesC, zeroesC int) bool { return onesC >= zeroesC })
}

func getCO2Rating(input []uint64, l int) uint64 {
	inputCopy := make([]uint64, len(input))
	copy(inputCopy, input)
	return getRating(inputCopy, l, func(onesC, zeroesC int) bool { return onesC < zeroesC })
}

func getRating(input []uint64, l int, checkF func(onesC, zeroesC int) bool) uint64 {
	for p := l - 1; p >= 0; p-- {
		if len(input) == 1 {
			break
		}

		mask := uint64(1 << p)
		onesC := 0
		zeroesC := 0
		for _, ip := range input {
			if ip&mask == mask {
				onesC++
			} else {
				zeroesC++
			}
		}

		doPut1 := checkF(onesC, zeroesC)

		var newI []uint64
		for i := range input {
			if (input[i]&mask == mask) == doPut1 {
				newI = append(newI, input[i])
			}
		}
		input = newI
	}

	if len(input) > 1 {
		panic(fmt.Sprint(input))
	}

	return input[0]
}

func getInput() []uint64 {
	lines := strings.Split(input, "\n")
	m := make([]uint64, len(lines))
	for i, line := range lines {
		m[i] = uint64(optimistic.ParseInt(line, 2, 64))
	}
	return m
}
