package year_2018

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
)

func init() {
	Solutions[1] = Day1{}
}

type Day1 struct{}

func (Day1) Part1(input []byte) string {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	freq := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			panic(fmt.Errorf("empty line"))
		}

		if line[0] == '+' {
			freq += parseNum(line)
		} else if line[0] == '-' {
			freq -= parseNum(line)
		} else {
			panic(fmt.Errorf("unknown symbol %v", line[0]))
		}
	}

	return strconv.Itoa(freq)
}

func (Day1) Part2(input []byte) string {
	alreadyEncountered := make(map[int]bool)

	r := bytes.NewReader(input)
	scanner := bufio.NewScanner(r)

	freq := 0
	for {
		for scanner.Scan() {
			line := scanner.Bytes()
			if len(line) == 0 {
				panic(fmt.Errorf("empty line"))
			}

			if line[0] == '+' {
				freq += parseNum(line)
			} else if line[0] == '-' {
				freq -= parseNum(line)
			} else {
				panic(fmt.Errorf("unknown symbol %v", line[0]))
			}

			if _, ok := alreadyEncountered[freq]; ok {
				return strconv.Itoa(freq)
			}

			alreadyEncountered[freq] = true
		}

		// Reset input
		scanner = bufio.NewScanner(bytes.NewReader(input))
	}
}

func parseNum(line []byte) int {
	numStr := string(line[1:])
	return optimistic.Atoi(numStr)
}
