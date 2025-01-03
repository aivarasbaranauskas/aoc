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

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	var x, maxVal int

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Text()
		if line == "" {
			maxVal = max(maxVal, x)
			x = 0
		} else {
			x += optimistic.Atoi(line)
		}
	}
	maxVal = max(maxVal, x)

	fmt.Println("Max:", maxVal)
}
