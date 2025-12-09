package year_2025

import (
	"bytes"
	"math"
	"sort"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[8] = Day8{}
}

type Day8 struct{}

func (Day8) Part1(input []byte) string {
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

	for _, d := range distances[:1000] {
		if circuits[d.i1] == circuits[d.i2] {
			continue
		}

		curr := circuits[d.i1]
		prev := circuits[d.i2]
		for i := range circuits {
			if circuits[i] == prev {
				circuits[i] = curr
			}
		}
	}

	circuitSizes := make([]int, len(boxes))
	for i := range circuits {
		circuitSizes[circuits[i]]++
	}

	i1, i2, i3 := -1, -1, -1
	for i, v := range circuitSizes {
		if i1 == -1 || v > circuitSizes[i1] {
			i1, i2, i3 = i, i1, i2
			continue
		}
		if i2 == -1 || v > circuitSizes[i2] {
			i2, i3 = i, i2
			continue
		}
		if i3 == -1 || v > circuitSizes[i3] {
			i3 = i
		}
	}

	return strconv.Itoa(circuitSizes[i1] * circuitSizes[i2] * circuitSizes[i3])
}

func (Day8) Part2(input []byte) string {
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
