package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/go_helpers/o"
	"log"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var x, maxV int

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Text()
		if line == "" {
			if x > maxV {
				maxV = x
			}
			x = 0
		} else {
			x += o.Atoi(line)
		}
	}
	if x > maxV {
		maxV = x
	}

	fmt.Println("Max:", maxV)
}
