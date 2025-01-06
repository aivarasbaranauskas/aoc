package year_2023

import (
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_matrix"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"slices"
	"strconv"
)

func init() {
	Solutions[13] = Day13{}
}

type Day13 struct{}

func (d Day13) Part1(input []byte) string {
	maps := bytes.Split(input, []byte("\n\n"))

	p1 := 0

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

		p1 += 100*d.doCalcP1Rows(_map) + d.doCalcP1Rows(_matrix.Transpose(_map))
	}

	return strconv.Itoa(p1)
}

func (d Day13) Part2(input []byte) string {
	maps := bytes.Split(input, []byte("\n\n"))

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

		p2 += 100*d.doCalcP2Rows(_map) + d.doCalcP2Rows(_matrix.Transpose(_map))
	}

	return strconv.Itoa(p2)
}

func (Day13) doCalcP1Rows(_map [][]int) int {
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

func (d Day13) doCalcP2Rows(_map [][]int) int {
RowsLoop:
	for i := 0; i < len(_map)-1; i++ {
		var (
			tmp  []int
			tmpI int
		)
		for j := 0; i-j >= 0 && i+1+j < len(_map); j++ {
			if !slices.Equal(_map[i-j], _map[i+1+j]) {
				if tmp == nil && d.diffBy1(_map[i-j], _map[i+1+j]) {
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

func (Day13) diffBy1(r1, r2 []int) bool {
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
