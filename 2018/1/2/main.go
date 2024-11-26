package main

import (
	"bufio"
	"embed"
	"fmt"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/go_helpers/_a"
)

//go:embed input.txt
var inputData embed.FS

var alreadyEncountered map[int]bool

func parseNum(line []byte) int {
	numStr := string(line[1:])
	num, err := strconv.Atoi(numStr)
	_a.CheckErr(err)

	return num
}

func main() {
	alreadyEncountered = make(map[int]bool)

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

			if _, ok := alreadyEncountered[freq]; ok {
				fmt.Println("Duplicate frequency:", freq)
				return
			}

			alreadyEncountered[freq] = true
		}

		// Reset file
		f, err = inputData.Open("input.txt")
		_a.CheckErr(err)
		scanner = bufio.NewScanner(f)
	}
}
