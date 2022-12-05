package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"log"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var x, max int

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Text()
		if line == "" {
			if x > max {
				max = x
			}
			x = 0
		} else {
			x += optimistic.Atoi(line)
		}
	}
	if x > max {
		max = x
	}

	fmt.Println("Max:", max)
}
