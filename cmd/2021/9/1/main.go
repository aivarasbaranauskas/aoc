package main

import (
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"io"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)
	fullFile, err := io.ReadAll(f)
	_a.CheckErr(err)

	heightmap := _slice.Map(
		strings.Split(string(fullFile), "\n"),
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

	fmt.Println(sum)
}
