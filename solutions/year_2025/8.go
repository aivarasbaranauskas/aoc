package year_2025

import (
	"bytes"
	"math"
	"slices"
	"sort"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[8] = Day8{}
}

type Day8 struct{}

func (day Day8) Part1(input []byte) string {
	lines := bytes.Split(input, []byte{'\n'})
	boxes := make([][3]int, len(lines))
	for i, line := range lines {
		spl := bytes.Split(line, []byte{','})
		boxes[i][0] = optimistic.AtoiBFast(spl[0])
		boxes[i][1] = optimistic.AtoiBFast(spl[1])
		boxes[i][2] = optimistic.AtoiBFast(spl[2])
	}

	type Light struct {
		d      float64
		i1, i2 int
	}

	distances := make([]Light, len(boxes)*(len(boxes)-1)/2)
	di := 0
	for i := range len(boxes) - 1 {
		for j := i + 1; j < len(boxes); j++ {
			distances[di].d = math.Sqrt(math.Pow(float64(boxes[i][0]-boxes[j][0]), 2) +
				math.Pow(float64(boxes[i][1]-boxes[j][1]), 2) +
				math.Pow(float64(boxes[i][2]-boxes[j][2]), 2),
			)
			distances[di].i1 = i
			distances[di].i2 = j
			di++
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].d < distances[j].d
	})

	distances = distances[:1000]
	connections := make(map[int][]int, 1000)
	for _, d := range distances {
		l, _ := connections[d.i1]
		connections[d.i1] = append(l, d.i2)
		l, _ = connections[d.i2]
		connections[d.i2] = append(l, d.i1)
	}

	visited := make([]bool, len(boxes))
	var sizes []int

	for i := range boxes {
		if visited[i] {
			continue
		}

		if _, ok := connections[i]; !ok {
			continue
		}

		sizes = append(sizes, day.visit(connections, visited, i))
	}

	slices.Sort(sizes)
	sizes = sizes[len(sizes)-3:]
	result := sizes[0] * sizes[1] * sizes[2]

	return strconv.Itoa(result)
}

func (day Day8) visit(connections map[int][]int, visited []bool, i int) int {
	if visited[i] {
		return 0
	}

	visited[i] = true
	ct := 1
	for _, next := range connections[i] {
		ct += day.visit(connections, visited, next)
	}
	return ct
}

func (day Day8) Part2(input []byte) string {
	lines := bytes.Split(input, []byte{'\n'})
	boxes := make([][3]int, len(lines))
	for i, line := range lines {
		spl := bytes.Split(line, []byte{','})
		boxes[i][0] = optimistic.AtoiBFast(spl[0])
		boxes[i][1] = optimistic.AtoiBFast(spl[1])
		boxes[i][2] = optimistic.AtoiBFast(spl[2])
	}

	type Light struct {
		d      float64
		i1, i2 int
	}

	distances := make([]Light, len(boxes)*(len(boxes)-1)/2)
	di := 0
	for i := range len(boxes) - 1 {
		for j := i + 1; j < len(boxes); j++ {
			distances[di].d = math.Sqrt(math.Pow(float64(boxes[i][0]-boxes[j][0]), 2) +
				math.Pow(float64(boxes[i][1]-boxes[j][1]), 2) +
				math.Pow(float64(boxes[i][2]-boxes[j][2]), 2),
			)
			distances[di].i1 = i
			distances[di].i2 = j
			di++
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].d < distances[j].d
	})

	circuits := make([]int, len(boxes))
	for i := range circuits {
		circuits[i] = i
	}

	ct := len(boxes)

	for _, d := range distances {
		if circuits[d.i1] == circuits[d.i2] {
			continue
		}

		if ct == 2 {
			return strconv.Itoa(boxes[d.i1][0] * boxes[d.i2][0])
		}

		curr := circuits[d.i1]
		prev := circuits[d.i2]
		for i := range circuits {
			if circuits[i] == prev {
				circuits[i] = curr
			}
		}

		ct--
	}

	return "not found"
}
