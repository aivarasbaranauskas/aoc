package year_2024

import (
	"bytes"
	"strconv"
)

func init() {
	Solutions[4] = Day4{}
}

type Day4 struct{}

func (Day4) Part1(input []byte) string {
	directions := [8][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{1, -1},
		{-1, 1},
		{1, 1},
		{-1, -1},
	}

	ct := 0
	cols := bytes.IndexByte(input, '\n') + 1
	rows := len(input)/cols + 1

	for i := range input {
		if input[i] != 'X' {
			continue
		}
		row, col := i/cols, i%cols

		for _, direction := range directions {
			xRow, xCol := row+3*direction[0], col+3*direction[1]

			if xRow < 0 || xRow >= rows || xCol < 0 || xCol > cols-1 {
				continue
			}

			mRow, mCol := row+direction[0], col+direction[1]
			aRow, aCol := row+2*direction[0], col+2*direction[1]

			mi := mRow*cols + mCol
			ai := aRow*cols + aCol
			xi := xRow*cols + xCol
			if input[mi] == 'M' && input[ai] == 'A' && input[xi] == 'S' {
				ct++
			}
		}
	}

	return strconv.Itoa(ct)
}

func (Day4) Part2(input []byte) string {
	ct := 0
	cols := bytes.IndexByte(input, '\n') + 1
	rows := len(input)/cols + 1

	for row := 1; row < cols-2; row++ {
		for col := 1; col < rows-1; col++ {
			i := row*141 + col
			if input[i] != 'A' {
				continue
			}

			if ((input[i-142] == 'M' && input[i+142] == 'S') || (input[i-142] == 'S' && input[i+142] == 'M')) &&
				((input[i-140] == 'M' && input[i+140] == 'S') || (input[i-140] == 'S' && input[i+140] == 'M')) {
				ct++
			}
		}
	}
	return strconv.Itoa(ct)
}
