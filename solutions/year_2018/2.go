package year_2018

import (
	"strconv"
	"strings"
)

func init() {
	Solutions[2] = Day2{}
}

type Day2 struct{}

func (Day2) Part1(input []byte) string {
	doubles := 0
	triples := 0
	for _, line := range strings.Split(string(input), "\n") {
		abcCounter := make(map[rune]int)
		val2 := false
		val3 := false
		for _, char := range line {
			if _, exists := abcCounter[char]; exists {
				abcCounter[char]++
			} else {
				abcCounter[char] = 1
			}
		}
		for _, value := range abcCounter {
			if value == 2 && val2 == false {
				doubles++
				val2 = true
			} else if value == 3 && val3 == false {
				triples++
				val3 = true
			}
		}
	}

	checksum := doubles * triples
	return strconv.Itoa(checksum)
}

func (Day2) Part2(input []byte) string {
	var abc []string
	for _, line := range strings.Split(string(input), "\n") {
		abc = append(abc, line)
	}
	for i := range abc {
		for j := range abc {
			diff := 0
			solution := ""
			for k := range abc[i] {
				if abc[i][k] == abc[j][k] {
					solution += string(abc[i][k])
				} else {
					diff++
				}
			}

			if diff == 1 {
				return solution
			}
		}
	}
	panic("No solution found")
}
