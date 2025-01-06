package solutions

import (
	"fmt"
	"github.com/aivarasbaranauskas/aoc/solutions/framework"
	"github.com/aivarasbaranauskas/aoc/solutions/year_2015"
	"github.com/aivarasbaranauskas/aoc/solutions/year_2018"
	"github.com/aivarasbaranauskas/aoc/solutions/year_2019"
	"github.com/aivarasbaranauskas/aoc/solutions/year_2021"
	"github.com/aivarasbaranauskas/aoc/solutions/year_2022"
	"github.com/aivarasbaranauskas/aoc/solutions/year_2023"
)

var solutions = map[int]map[int]framework.Solution{
	2015: year_2015.Solutions,
	2018: year_2018.Solutions,
	2019: year_2019.Solutions,
	2021: year_2021.Solutions,
	2022: year_2022.Solutions,
	2023: year_2023.Solutions,
}

func Run(year, day int) {
	solution, ok := solutions[year][day]
	if !ok {
		fmt.Println("Solution not found")
		return
	}

	input, err := getInput(year, day)
	if err != nil {
		fmt.Println("Failed to get the input for solution:", err)
		return
	}

	fmt.Printf("Part 1: %s\n", solution.Part1(input))
	fmt.Printf("Part 2: %s\n", solution.Part2(input))
}
