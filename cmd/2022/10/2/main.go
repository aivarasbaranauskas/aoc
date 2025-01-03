package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	r := bufio.NewScanner(f)
	x := 1
	c := 1
	var crt [240]bool
	for r.Scan() {
		line := r.Text()
		if line == "noop" {
			//here
			if x-1 <= c%40-1 && c%40-1 <= x+1 {
				crt[c-1] = true
			}
			c++
			continue
		}

		//here
		if x-1 <= c%40-1 && c%40-1 <= x+1 {
			crt[c-1] = true
		}

		c++
		//here
		if x-1 <= c%40-1 && c%40-1 <= x+1 {
			crt[c-1] = true
		}
		c++

		spl := strings.Split(line, " ")
		x += optimistic.Atoi(spl[1])
	}
	for i := 0; i <= 5; i++ {
		for j := 0; j < 40; j++ {
			if crt[i*40+j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
