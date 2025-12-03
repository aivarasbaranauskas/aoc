package solutions

import (
	"fmt"
	"slices"
	"sync"
	"testing"

	"github.com/aivarasbaranauskas/aoc/solutions/framework"
)

func Run(year, day int) error {
	solution, input, err := getSolutionAndInput(year, day)
	if err != nil {
		return err
	}

	fmt.Printf("Part 1: %s\n", solution.Part1(slices.Clone(input)))
	fmt.Printf("Part 2: %s\n", solution.Part2(slices.Clone(input)))

	return nil
}

func Bench(year, day int) error {
	solution, input, err := getSolutionAndInput(year, day)
	if err != nil {
		return err
	}

	bench(solution.Part1, input, 1)
	bench(solution.Part2, input, 2)

	return nil
}

func bench(f func([]byte) string, input []byte, part int) {
	var (
		result string
		once   sync.Once
	)

	benchResult := testing.Benchmark(func(b *testing.B) {
		inputs := make([][]byte, b.N)
		for i := range b.N {
			inputs[i] = slices.Clone(input)
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			r := f(inputs[i])
			once.Do(func() {
				result = r
			})
		}
	})

	fmt.Printf("Part %v:\n  Result: %s\n  Bench: %v\t%v\n", part, result, benchResult.String(), benchResult.MemString())
}

func getSolutionAndInput(year, day int) (framework.Solution, []byte, error) {
	solution, ok := solutions[year][day]
	if !ok {
		return nil, nil, fmt.Errorf("solution not found")
	}

	input, err := getInput(year, day)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get the input for solution: %w", err)
	}

	return solution, input, nil
}
