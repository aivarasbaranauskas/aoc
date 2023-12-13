package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_matrix"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"slices"
)

//go:embed input.txt
var input []byte

func main() {
	maps := bytes.Split(input, []byte("\n\n"))

	p1 := 0
	p2 := 0

	for _, rawMap := range maps {
		_mapRaw := bytes.Split(rawMap, []byte("\n"))
		_map := _slice.Map(_mapRaw, func(tin []byte) []int {
			return _slice.Map(tin, func(tin byte) int {
				if tin == '#' {
					return 1
				}
				return 0
			})
		})

		p1 += 100*doCalcP1Rows(_map) + doCalcP1Rows(_matrix.Transpose(_map))
		p2 += 100*doCalcP2Rows(_map) + doCalcP2Rows(_matrix.Transpose(_map))
	}

	fmt.Println("part 1:", p1)
	fmt.Println("part 2:", p2)
}

func doCalcP1Rows(_map [][]int) int {
RowsLoop:
	for i := 0; i < len(_map)-1; i++ {
		for j := 0; i-j >= 0 && i+1+j < len(_map); j++ {
			if !slices.Equal(_map[i-j], _map[i+1+j]) {
				continue RowsLoop
			}
		}
		return i + 1
	}

	return 0
}

func doCalcP2Rows(_map [][]int) int {
RowsLoop:
	for i := 0; i < len(_map)-1; i++ {
		var (
			tmp  []int
			tmpI int
		)
		for j := 0; i-j >= 0 && i+1+j < len(_map); j++ {
			if !slices.Equal(_map[i-j], _map[i+1+j]) {
				if tmp == nil && diffBy1(_map[i-j], _map[i+1+j]) {
					tmpI = i - j
					tmp = _map[i-j]
					_map[i-j] = _map[i+1+j]
					continue
				}

				if tmp != nil {
					_map[tmpI] = tmp
				}
				continue RowsLoop
			}
		}
		if tmp != nil {
			_map[tmpI] = tmp
			return i + 1
		}
	}

	return 0
}

func diffBy1(r1, r2 []int) bool {
	diffCt := 0
	for i, v := range r1 {
		if v != r2[i] {
			diffCt++
			if diffCt > 1 {
				return false
			}
		}
	}
	return diffCt == 1
}
