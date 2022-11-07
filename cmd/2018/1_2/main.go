package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var alreadyEncautered map[int]bool

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseNum(line []byte) int {
	numStr := string(line[1:])
	num, err := strconv.Atoi(numStr)
	check(err)

	return num
}

func main() {
	alreadyEncautered = make(map[int]bool)

	f, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(f)

	freq := 0
	for {
		for scanner.Scan() {
			line := scanner.Bytes()
			if len(line) == 0 {
				panic(fmt.Errorf("Empty line"))
			}

			if line[0] == '+' {
				freq += parseNum(line)
			} else if line[0] == '-' {
				freq -= parseNum(line)
			} else {
				panic(fmt.Errorf("Unknown symbol %v", line[0]))
			}

			if _, ok := alreadyEncautered[freq]; ok {
				fmt.Println("Duplicate frequency:", freq)
				return
			}

			alreadyEncautered[freq] = true
		}

		// Reset file
		f, err = os.Open("input.txt")
		check(err)
		scanner = bufio.NewScanner(f)
	}
}
