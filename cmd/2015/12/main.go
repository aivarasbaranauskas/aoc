package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

//go:embed input.txt
var input []byte

func main() {
	part1()
	part2()
}

func part1() {
	l := len(input)
	s := 0
	for i := 0; i < l; i++ {
		if input[i] != '-' && (input[i] < '0' || input[i] > '9') {
			continue
		}

		j := i
		if input[j] == '-' {
			if input[j+1] >= '0' && input[j+1] <= '9' {
				j++
			} else {
				continue
			}
		}
		for ; input[j] >= '0' && input[j] <= '9'; j++ {
		}
		s += optimistic.Atoi(string(input[i:j]))
		i = j
	}
	fmt.Println("Part 1:", s)
}

func part2() {
	// filter the source first
	checkB := 0
	for {
		i := bytes.Index(input[checkB:], []byte("\":\"red\""))
		if i == -1 {
			break
		}

		b := i
		d := 0
		for ; b >= 0; b-- {
			if input[b] == '}' {
				d++
			} else if input[b] == '{' {
				d--
			}
			if d < 0 {
				break
			}
		}
		if b < 0 {
			checkB = i + 1
			continue
		}
		e := i
		d = 0
		for ; e < len(input); e++ {
			if input[e] == '{' {
				d++
			} else if input[e] == '}' {
				d--
			}
			if d < 0 {
				break
			}
		}
		input = append(input[:b], input[e+1:]...)
	}

	l := len(input)
	s := 0
	for i := 0; i < l; i++ {
		if input[i] != '-' && (input[i] < '0' || input[i] > '9') {
			continue
		}

		j := i
		if input[j] == '-' {
			if input[j+1] >= '0' && input[j+1] <= '9' {
				j++
			} else {
				continue
			}
		}
		for ; input[j] >= '0' && input[j] <= '9'; j++ {
		}
		s += optimistic.Atoi(string(input[i:j]))
		i = j
	}
	fmt.Println("Part 2:", s)
}
