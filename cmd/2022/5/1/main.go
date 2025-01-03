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

	stacks := [][]byte{
		{'H', 'R', 'B', 'D', 'Z', 'F', 'L', 'S'},
		{'T', 'B', 'M', 'Z', 'R'},
		{'Z', 'L', 'C', 'H', 'N', 'S'},
		{'S', 'C', 'F', 'J'},
		{'P', 'G', 'H', 'W', 'R', 'Z', 'B'},
		{'V', 'J', 'Z', 'G', 'D', 'N', 'M', 'T'},
		{'G', 'L', 'N', 'W', 'F', 'S', 'P', 'Q'},
		{'M', 'Z', 'R'},
		{'M', 'C', 'L', 'G', 'V', 'R', 'T'},
	}

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " ")
		ct := optimistic.Atoi(spl[1])
		from := optimistic.Atoi(spl[3]) - 1
		to := optimistic.Atoi(spl[5]) - 1

		for i := 0; i < ct; i++ {
			lastFrom := len(stacks[from]) - 1
			stacks[to] = append(stacks[to], stacks[from][lastFrom])
			stacks[from] = stacks[from][:lastFrom]
		}
	}

	var out string
	for _, stack := range stacks {
		out += string(stack[len(stack)-1])
	}

	fmt.Println(out)
}
