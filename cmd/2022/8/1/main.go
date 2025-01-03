package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_matrix"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"slices"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	r := bufio.NewScanner(f)
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

	fmt.Println(vCt)
}
