package year_2024

import (
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"strconv"
)

func init() {
	Solutions[8] = Day8{}
}

type Day8 struct{}

func (Day8) Part1(input []byte) string {
	res := make([]bool, len(input))

	cols := bytes.IndexByte(input, '\n') + 1
	rows := (len(input) + 1) / cols

	for i, v := range input {
		if v == '.' || v == '\n' {
			continue
		}

		rowI := i / cols
		colI := i % cols

		for j := i + 1; j < len(input); j++ {
			if input[j] != v {
				continue
			}

			rowJ := j / cols
			colJ := j % cols

			rowDiff := rowJ - rowI
			colDiff := colJ - colI

			rowRes1 := rowI - rowDiff
			colRes1 := colI - colDiff

			rowRes2 := rowJ + rowDiff
			colRes2 := colJ + colDiff

			if rowRes1 >= 0 && rowRes1 < rows && colRes1 >= 0 && colRes1 < (cols-1) {
				res[rowRes1*cols+colRes1] = true
			}
			if rowRes2 >= 0 && rowRes2 < rows && colRes2 >= 0 && colRes2 < (cols-1) {
				res[rowRes2*cols+colRes2] = true
			}
		}
	}

	return strconv.Itoa(_slice.CountCond(res, func(b bool) bool { return b }))
}

func (Day8) Part2(input []byte) string {
	res := make([]bool, len(input))

	cols := bytes.IndexByte(input, '\n') + 1
	rows := (len(input) + 1) / cols

	for i, v := range input {
		if v == '.' || v == '\n' {
			continue
		}

		rowI := i / cols
		colI := i % cols

		for j := i + 1; j < len(input); j++ {
			if input[j] != v {
				continue
			}

			res[i] = true
			res[j] = true

			rowJ := j / cols
			colJ := j % cols

			rowDiff := rowJ - rowI
			colDiff := colJ - colI

			rowRes1 := rowI - rowDiff
			colRes1 := colI - colDiff
			for rowRes1 >= 0 && rowRes1 < rows && colRes1 >= 0 && colRes1 < (cols-1) {
				res[rowRes1*cols+colRes1] = true
				rowRes1 -= rowDiff
				colRes1 -= colDiff
			}

			rowRes2 := rowJ + rowDiff
			colRes2 := colJ + colDiff
			for rowRes2 >= 0 && rowRes2 < rows && colRes2 >= 0 && colRes2 < (cols-1) {
				res[rowRes2*cols+colRes2] = true
				rowRes2 += rowDiff
				colRes2 += colDiff
			}
		}
	}

	return strconv.Itoa(_slice.CountCond(res, func(b bool) bool { return b }))
}
