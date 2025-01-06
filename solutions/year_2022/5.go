package year_2022

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"slices"
	"strings"
)

func init() {
	Solutions[5] = Day5{}
}

type Day5 struct{}

func (d Day5) Part1(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))

	stacks := d.readStacks(r)

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

	return out
}

func (d Day5) Part2(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))

	stacks := d.readStacks(r)

	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " ")
		ct := optimistic.Atoi(spl[1])
		from := optimistic.Atoi(spl[3]) - 1
		to := optimistic.Atoi(spl[5]) - 1

		lastFrom := len(stacks[from]) - ct
		stacks[to] = append(stacks[to], stacks[from][lastFrom:]...)
		stacks[from] = stacks[from][:lastFrom]
	}

	var out string
	for _, stack := range stacks {
		out += string(stack[len(stack)-1])
	}

	return out
}

func (Day5) readStacks(r *bufio.Scanner) [][]byte {
	var stacks [][]byte
	for r.Scan() {
		line := r.Bytes()
		if len(line) == 0 {
			break
		}
		if stacks == nil {
			stacks = make([][]byte, len(line)/4+1)
		}
		for i := 0; i <= len(line)/4; i++ {
			c := line[i*4+1]
			if c != ' ' {
				stacks[i] = append(stacks[i], c)
			}
		}
	}
	for i, v := range stacks {
		slices.Reverse(v)
		stacks[i] = v[1:]
	}
	return stacks
}
