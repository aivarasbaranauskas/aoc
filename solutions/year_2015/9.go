package year_2015

import (
	"bytes"
	"maps"
	"math"
	"slices"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[9] = Day9{}
}

type Day9 struct{}

func (day Day9) Part1(input []byte) string {
	m := day.parseGraph(input)
	nodes := slices.Collect(maps.Keys(m))
	return strconv.Itoa(
		day.find(
			m,
			nodes,
			math.MaxInt,
			func(a int, b int) int {
				return min(a, b)
			},
		),
	)
}

func (day Day9) Part2(input []byte) string {
	m := day.parseGraph(input)
	nodes := slices.Collect(maps.Keys(m))
	return strconv.Itoa(
		day.find(
			m,
			nodes,
			0,
			func(a int, b int) int {
				return max(a, b)
			},
		),
	)
}

func (day Day9) find(
	m map[string]map[string]int,
	nodes []string,
	init int,
	comF func(int, int) int,
) int {
	var find func(visited []bool, current int, n int, sum int) int

	find = func(visited []bool, current int, n int, sum int) int {
		if n == len(visited) {
			return sum
		}

		shortest := init

		for i, v := range visited {
			if v {
				continue
			}

			visited[i] = true
			shortest = comF(shortest, find(visited, i, n+1, sum+m[nodes[current]][nodes[i]]))
			visited[i] = false
		}
		return shortest
	}

	shortest := init

	for i := range nodes {
		visited := make([]bool, len(nodes))
		visited[i] = true
		shortest = comF(shortest, find(visited, i, 1, 0))
	}
	return shortest
}

func (day Day9) parseGraph(input []byte) map[string]map[string]int {
	m := map[string]map[string]int{}

	for line := range bytes.Lines(input) {
		line = bytes.TrimSpace(line)

		spl := bytes.Split(line, []byte{' '})

		from := string(spl[0])
		to := string(spl[2])
		dist := optimistic.AtoiBFast(spl[4])

		if _, ok := m[from]; !ok {
			m[from] = map[string]int{}
		}
		m[from][to] = dist

		if _, ok := m[to]; !ok {
			m[to] = map[string]int{}
		}
		m[to][from] = dist
	}

	return m
}
