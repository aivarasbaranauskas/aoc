package main

import (
	"fmt"
	"github.com/aivarasbaranauskas/aoc/solutions"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: aoc <year> <day>")
		os.Exit(1)
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid year")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid day")
		os.Exit(1)
	}

	solutions.Run(year, day)
}
