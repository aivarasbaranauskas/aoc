package main

import (
	"bufio"
	"embed"
	"fmt"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/a"
)

//go:embed input.txt
var inputData embed.FS

func parseNum(line []byte) int {
	numStr := string(line[1:])
	num, err := strconv.Atoi(numStr)
	a.CheckErr(err)

	return num
}

func main() {
	f, err := inputData.Open("input.txt")
	a.CheckErr(err)
	scanner := bufio.NewScanner(f)

	freq := 0
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
	}

	fmt.Println("Final frequency:", freq)
}
