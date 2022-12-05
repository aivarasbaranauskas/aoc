package main

import (
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fullFile, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	heightmap := _a.Map(
		strings.Split(string(fullFile), "\n"),
		func(line string) []int {
			return _a.Map(
				strings.Split(line, ""),
				optimistic.Atoi,
			)
		},
	)

	var sizes []int

	for i := range heightmap {
		for j := range heightmap[i] {
			if heightmap[i][j] != 9 {
				sizes = append(sizes, basinSize(heightmap, i, j))
			}
		}
	}

	sort.Ints(sizes)
	fmt.Println(_a.Product(sizes[len(sizes)-3:]...))
}

func basinSize(heightmap [][]int, i, j int) int {
	if i < 0 || j < 0 || i >= len(heightmap) || j >= len(heightmap[i]) {
		return 0
	}
	if heightmap[i][j] == 9 {
		return 0
	}
	heightmap[i][j] = 9
	return 1 + basinSize(heightmap, i-1, j) + basinSize(heightmap, i+1, j) + basinSize(heightmap, i, j-1) + basinSize(heightmap, i, j+1)
}
