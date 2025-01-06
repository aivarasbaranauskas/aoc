package year_2022

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_matrix"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"slices"
	"strconv"
	"strings"
)

func init() {
	Solutions[8] = Day8{}
}

type Day8 struct{}

func (Day8) Part1(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var trees [][]int
	for r.Scan() {
		trees = append(trees, _slice.Map(strings.Split(r.Text(), ""), optimistic.Atoi))
	}

	treesT := _matrix.Transpose(trees)
	l := len(trees)
	w := len(trees[0])
	vCt := 2*w + 2*l - 4

	for x := 1; x < l-1; x++ {
		for y := 1; y < w-1; y++ {
			if trees[x][y] > min(slices.Max(treesT[y][x+1:]), slices.Max(treesT[y][:x]), slices.Max(trees[x][y+1:]), slices.Max(trees[x][:y])) {
				vCt++
			}
		}
	}

	return strconv.Itoa(vCt)
}

func (Day8) Part2(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var (
		trees [][]int
	)
	for r.Scan() {
		trees = append(trees, _slice.Map(strings.Split(r.Text(), ""), optimistic.Atoi))
	}

	l := len(trees)
	w := len(trees[0])
	var maxScenicScore int

	for x := 1; x < l-1; x++ {
		for y := 1; y < w-1; y++ {
			tree := trees[x][y]

			// down
			down := l - x - 1
			for i := x + 1; i < l; i++ {
				if tree <= trees[i][y] {
					down = i - x
					break
				}
			}

			// up
			up := x
			for i := x - 1; i >= 0; i-- {
				if tree <= trees[i][y] {
					up = x - i
					break
				}
			}

			// right
			right := w - y - 1
			for i := y + 1; i < w; i++ {
				if tree <= trees[x][i] {
					right = i - y
					break
				}
			}

			// left
			left := y
			for i := y - 1; i >= 0; i-- {
				if tree <= trees[x][i] {
					left = y - i
					break
				}
			}
			maxScenicScore = max(maxScenicScore, down*up*left*right)
		}
	}

	return strconv.Itoa(maxScenicScore)
}
