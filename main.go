package main

import (
	"flag"
	"fmt"

	"os"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/solutions"
)

import _ "github.com/joho/godotenv/autoload"

func main() {
	var bench bool

	flag.BoolVar(&bench, "bench", false, "Run benchmarks")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Usage: aoc <year> <day>")
		os.Exit(1)
	}

	year, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println("Invalid year")
		os.Exit(1)
	}

	day, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		fmt.Println("Invalid day")
		os.Exit(1)
	}

	if bench {
		err = solutions.Bench(year, day)
	} else {
		err = solutions.Run(year, day)
	}

	if err != nil {
		fmt.Printf("Error running solution: %v\n", err)
	}
}
