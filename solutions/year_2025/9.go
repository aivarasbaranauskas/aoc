package year_2025

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[9] = Day9{}
}

type Day9 struct{}

func (Day9) Part1(input []byte) string {
	lines := bytes.Split(input, []byte{'\n'})
	tiles := make([][2]int, len(lines))
	for i, line := range lines {
		spl := bytes.Split(line, []byte{','})
		tiles[i][0] = optimistic.AtoiBFast(spl[0])
		tiles[i][1] = optimistic.AtoiBFast(spl[1])
	}

	biggest := 0
	for i, t1 := range tiles {
		for _, t2 := range tiles[i+1:] {
			biggest = max(biggest, (_num.Abs(t1[0]-t2[0])+1)*(_num.Abs(t1[1]-t2[1])+1))
		}
	}

	return strconv.Itoa(biggest)
}

func (Day9) Part2(input []byte) string {
	lines := bytes.Split(input, []byte{'\n'})
	tiles := make([][2]int, len(lines))
	width, height := 0, 0
	for i, line := range lines {
		spl := bytes.Split(line, []byte{','})
		tiles[i][0] = optimistic.AtoiBFast(spl[0])
		tiles[i][1] = optimistic.AtoiBFast(spl[1])
		width = max(width, tiles[i][0])
		height = max(height, tiles[i][1])
	}

	isAllInside := func(t1, t2 [2]int) bool {
		fromX, toX := min(t1[0], t2[0]), max(t1[0], t2[0])
		fromY, toY := min(t1[1], t2[1]), max(t1[1], t2[1])

		checkLine := func(t1, t2 [2]int) bool {
			if t1[1] == t2[1] {
				// horizontal
				a := t2[1]
				fromT, toT := min(t1[0], t2[0]), max(t1[0], t2[0])
				if fromY < a && a < toY && fromX < toT && fromT < toX {
					return false
				}
			} else {
				a := t2[0]
				fromT, toT := min(t1[1], t2[1]), max(t1[1], t2[1])
				if fromX < a && a < toX && fromY < toT && fromT < toY {
					return false
				}
			}
			return true
		}

		for i := range len(tiles) - 1 {
			if !checkLine(tiles[i], tiles[i+1]) {
				return false
			}
		}

		return checkLine(tiles[0], tiles[len(tiles)-1])
	}

	biggest := 0
	for i, t1 := range tiles {
		for _, t2 := range tiles[i+1:] {
			if !isAllInside(t1, t2) {
				continue
			}

			biggest = max(biggest, (_num.Abs(t1[0]-t2[0])+1)*(_num.Abs(t1[1]-t2[1])+1))
		}
	}

	return strconv.Itoa(biggest)
}
