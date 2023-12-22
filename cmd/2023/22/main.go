package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_map"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

type Cube [3]int

type Brick [2]Cube

func (b Brick) Overlaps(b2 Brick) bool {
	xOverlaps := b[0][0] <= b2[1][0] && b2[0][0] <= b[1][0]
	yOverlaps := b[0][1] <= b2[1][1] && b2[0][1] <= b[1][1]
	return xOverlaps && yOverlaps
}

func main() {
	part1()
	part2()
}

func part2() {
	bricks := parseInput()

	// sort by bottom z
	slices.SortFunc(
		bricks,
		func(a, b Brick) int {
			return a[0][2] - b[0][2]
		},
	)

	var landedBricks []Brick
	for _, brick := range bricks {
		z := 0
		for i := len(landedBricks) - 1; i >= 0; i-- {
			if brick.Overlaps(landedBricks[i]) {
				z = max(z, landedBricks[i][1][2])
			}
		}
		z++
		landedBricks = append(landedBricks, Brick{
			Cube{
				brick[0][0],
				brick[0][1],
				z,
			},
			Cube{
				brick[1][0],
				brick[1][1],
				z + brick[1][2] - brick[0][2],
			},
		})
	}

	sims := map[int]int{}

	for _, brick := range landedBricks {
		if brick[0][2] == 1 {
			// it's on the ground
			continue
		}

		var isSupportedBy []int
		for j, b := range landedBricks {
			if b[1][2]+1 == brick[0][2] && brick.Overlaps(b) {
				isSupportedBy = append(isSupportedBy, j)
			}
		}
		if len(isSupportedBy) == 1 {
			// found a chain reaction!
			ignoredI := isSupportedBy[0]
			if _, ok := sims[ignoredI]; !ok {
				sims[ignoredI] = simulateFall(landedBricks, ignoredI)
			}
		}
	}

	fmt.Println("part 2:", _num.Sum(_map.Values(sims)...))
}

func simulateFall(bricks []Brick, ignoredI int) (ct int) {
	var landedBricks []Brick
	for _, brick := range bricks {
		z := 0
		for i := len(landedBricks) - 1; i >= 0; i-- {
			if i != ignoredI && brick.Overlaps(landedBricks[i]) {
				z = max(z, landedBricks[i][1][2])
			}
		}
		z++
		if z != brick[0][2] {
			ct++
		}

		landedBricks = append(landedBricks, Brick{
			Cube{
				brick[0][0],
				brick[0][1],
				z,
			},
			Cube{
				brick[1][0],
				brick[1][1],
				z + brick[1][2] - brick[0][2],
			},
		})
	}
	return
}

func part1() {
	bricks := parseInput()

	// sort by bottom z
	slices.SortFunc(
		bricks,
		func(a, b Brick) int {
			return a[0][2] - b[0][2]
		},
	)

	var landedBricks []Brick
	for _, brick := range bricks {
		z := 0
		for i := len(landedBricks) - 1; i >= 0; i-- {
			if brick.Overlaps(landedBricks[i]) {
				z = max(z, landedBricks[i][1][2])
			}
		}
		z++
		landedBricks = append(landedBricks, Brick{
			Cube{
				brick[0][0],
				brick[0][1],
				z,
			},
			Cube{
				brick[1][0],
				brick[1][1],
				z + brick[1][2] - brick[0][2],
			},
		})
	}

	var canNotBeDisintegrated []int

	for _, brick := range landedBricks {
		if brick[0][2] == 1 {
			// it's on the ground
			continue
		}

		var isSupportedBy []int
		for j, b := range landedBricks {
			if b[1][2]+1 == brick[0][2] && brick.Overlaps(b) {
				isSupportedBy = append(isSupportedBy, j)
			}
		}
		if len(isSupportedBy) == 1 {
			canNotBeDisintegrated = append(canNotBeDisintegrated, isSupportedBy[0])
		}
	}

	canNotBeDisintegratedCt := len(_slice.CountUnique(canNotBeDisintegrated))

	fmt.Println("part 1:", len(bricks)-canNotBeDisintegratedCt)
}

func parseInput() []Brick {
	return _slice.Map(strings.Split(input, "\n"), parseBrick)
}

func parseBrick(s string) Brick {
	spl := strings.Split(s, "~")
	a := parseCube(spl[0])
	b := parseCube(spl[1])
	// rearrange so a<b in all axes
	return Brick{
		Cube{
			min(a[0], b[0]),
			min(a[1], b[1]),
			min(a[2], b[2]),
		},
		Cube{
			max(a[0], b[0]),
			max(a[1], b[1]),
			max(a[2], b[2]),
		},
	}
}

func parseCube(s string) Cube {
	spl := strings.Split(s, ",")
	return Cube{
		optimistic.Atoi(spl[0]),
		optimistic.Atoi(spl[1]),
		optimistic.Atoi(spl[2]),
	}
}
