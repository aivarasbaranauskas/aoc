package year_2025

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/_slice"
)

func init() {
	Solutions[11] = Day11{}
}

type Day11 struct{}

// aaa: you hhh
// you: bbb ccc
// bbb: ddd eee
// ccc: ddd eee fff
// ddd: ggg
// eee: out
// fff: out
// ggg: out
// hhh: ccc fff iii
// iii: out

func (day Day11) Part1(input []byte) string {
	m := day.parseInput(input)

	var walk func(string) int

	walk = func(s string) int {
		if s == "out" {
			return 1
		}
		sum := 0
		for _, next := range m[s] {
			sum += walk(next)
		}
		return sum
	}

	return strconv.Itoa(walk("you"))
}

func (day Day11) parseInput(input []byte) map[string][]string {
	lines := bytes.Split(input, []byte{'\n'})
	m := make(map[string][]string, len(lines))
	for _, line := range lines {
		spl1 := bytes.Split(line, []byte(": "))
		spl2 := bytes.Split(spl1[1], []byte(" "))
		m[string(spl1[0])] = _slice.Map(spl2, func(tin []byte) string {
			return string(tin)
		})
	}
	return m
}

func (day Day11) Part2(input []byte) string {
	m := day.parseInput(input)
	type State struct {
		k        string
		fft, dac bool
	}
	mem := map[State]int{}

	var walk func(string, bool, bool) int

	walk = func(s string, fft, dac bool) int {
		if s == "out" {
			if fft && dac {
				return 1
			}
			return 0
		}

		if s == "fft" {
			fft = true
		}
		if s == "dac" {
			dac = true
		}

		state := State{s, fft, dac}
		if n, ok := mem[state]; ok {
			return n
		}

		sum := 0
		for _, next := range m[s] {
			sum += walk(next, fft, dac)
		}

		mem[state] = sum
		return sum
	}

	return strconv.Itoa(walk("svr", false, false))
}
