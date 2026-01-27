package year_2015

import (
	"bytes"
	"math"
	"slices"
	"sort"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[17] = Day17{}
}

type Day17 struct{}

func (day Day17) Part1(input []byte) string {
	m := day.parse(input)
	return strconv.Itoa(day.count(m, 150))
}

func (day Day17) count(m []int, n int) int {
	ct := 0
	for i, v := range m {
		if v > n {
			continue
		}
		if v == n {
			ct++
			continue
		}

		ct += day.count(m[i+1:], n-v)
	}
	return ct
}

func (day Day17) Part2(input []byte) string {
	m := day.parse(input)
	_, ct := day.count2(m, 0, 150)
	return strconv.Itoa(ct)
}

func (day Day17) count2(m []int, used int, n int) (int, int) {
	ct := 0
	minUsed := math.MaxInt
	for i, v := range m {
		if v > n {
			continue
		}

		if v == n {
			if minUsed != used+1 {
				ct = 1
				minUsed = used + 1
			} else {
				ct++
			}
			continue
		}

		if minUsed != used+1 {
			minUsedD, ctD := day.count2(m[i+1:], used+1, n-v)
			if minUsedD < minUsed {
				minUsed = minUsedD
				ct = ctD
			} else if minUsedD == minUsed {
				ct += ctD
			}
		}
	}

	return minUsed, ct
}

func (day Day17) parse(input []byte) []int {
	n := bytes.Count(input, []byte{'\n'}) + 1
	m := make([]int, 0, n)

	for line := range bytes.Lines(input) {
		line = bytes.TrimSpace(line)
		m = append(m, optimistic.AtoiBFast(line))
	}

	sort.Ints(m)
	slices.Reverse(m)

	return m
}
