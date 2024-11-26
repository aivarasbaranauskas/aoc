package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

const forward = "forward"
const up = "up"
const down = "down"

type command struct {
	command string
	value   int
}

//go:embed test_input.txt
var testInput string

//go:embed input.txt
var input string

func main() {
	fmt.Println("Test:")
	hp, d := process(parseData(testInput), true)
	fmt.Printf("%v * %v = %v\n", hp, d, hp*d)

	fmt.Println("Actual:")
	hp, d = process(parseData(input), false)
	fmt.Printf("%v * %v = %v\n", hp, d, hp*d)
}

func process(commands []command, debug bool) (horizontalPosition int, depth int) {
	aim := 0

	for _, c := range commands {
		switch c.command {
		case forward:
			horizontalPosition += c.value
			depth += aim * c.value
		case up:
			aim -= c.value
		case down:
			aim += c.value
		}

		if debug {
			fmt.Printf("%v %v - %v %v %v\n", c.command, c.value, horizontalPosition, depth, aim)
		}
	}
	return
}

func parseData(data string) []command {
	lines := strings.Split(data, "\n")
	m := make([]command, len(lines))
	for i, line := range lines {
		spl := strings.Split(line, " ")
		v, err := strconv.Atoi(spl[1])
		if err != nil {
			panic(err)
		}
		m[i] = command{
			command: spl[0],
			value:   v,
		}
	}

	return m
}
