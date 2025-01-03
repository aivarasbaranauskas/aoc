package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

//go:embed input.txt
var inputData embed.FS

var alreadyEncautered map[int]bool

func parseNum(line []byte) int {
	numStr := string(line[1:])
	return optimistic.Atoi(numStr)
}

func main() {
	alreadyEncautered = make(map[int]bool)

	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)
	scanner := bufio.NewScanner(f)

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

			if _, ok := alreadyEncautered[freq]; ok {
				fmt.Println("Duplicate frequency:", freq)
				return
			}

			alreadyEncautered[freq] = true
		}

		// Reset file
		f, err = inputData.Open("input.txt")
		_a.CheckErr(err)
		scanner = bufio.NewScanner(f)
	}
}
