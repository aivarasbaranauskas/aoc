package year_2021

import (
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"sort"
	"strconv"
	"strings"
)

func init() {
	Solutions[9] = Day9{}
}

type Day9 struct{}

func (Day9) Part1(input []byte) string {
	heightmap := _slice.Map(
		strings.Split(string(input), "\n"),
		func(line string) []int {
			return _slice.Map(
				strings.Split(line, ""),
				optimistic.Atoi,
			)
		},
	)

	var sum int

	for i := range heightmap {
		for j := range heightmap[i] {
			top, left, right, bottom := true, true, true, true

			if i > 0 {
				top = heightmap[i][j] < heightmap[i-1][j]
			}
			if i < len(heightmap)-1 {
				bottom = heightmap[i][j] < heightmap[i+1][j]
			}
			if j > 0 {
				left = heightmap[i][j] < heightmap[i][j-1]
			}
			if j < len(heightmap[i])-1 {
				right = heightmap[i][j] < heightmap[i][j+1]
			}

			if top && left && right && bottom {
				sum += heightmap[i][j] + 1
			}
		}
	}

	return strconv.Itoa(sum)
}

func (d Day9) Part2(input []byte) string {
	heightmap := _slice.Map(
		strings.Split(string(input), "\n"),
		func(line string) []int {
			return _slice.Map(
				strings.Split(line, ""),
				optimistic.Atoi,
			)
		},
	)

	var sizes []int

	for i := range heightmap {
		for j := range heightmap[i] {
			if heightmap[i][j] != 9 {
				sizes = append(sizes, d.basinSize(heightmap, i, j))
			}
		}
	}

	sort.Ints(sizes)
	return strconv.Itoa(_num.Product(sizes[len(sizes)-3:]...))
}

func (d Day9) basinSize(heightmap [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(heightmap) || j >= len(heightmap[i]) {
		return 0
	}
	if heightmap[i][j] == 9 {
		return 0
	}
	heightmap[i][j] = 9
	return 1 + d.basinSize(heightmap, i-1, j) + d.basinSize(heightmap, i+1, j) + d.basinSize(heightmap, i, j-1) + d.basinSize(heightmap, i, j+1)
}
