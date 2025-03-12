package year_2024

import (
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"strconv"
)

func init() {
	Solutions[6] = Day6{}
}

type Day6 struct{}

func (Day6) Part1(input []byte) string {
	directions := [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	cols := bytes.IndexByte(input, '\n') + 1
	rows := len(input)/cols + 1

	guardI := bytes.IndexByte(input, '^')
	guardRow := guardI / cols
	guardCol := guardI % cols

	direction := 0
	ct := 0

	for guardRow >= 0 && guardRow < rows && guardCol >= 0 && guardCol < cols-1 {
		if input[guardRow*cols+guardCol] != 'O' {
			ct++
			input[guardRow*cols+guardCol] = 'O'
		}

		nextRow := guardRow + directions[direction][0]
		nextCol := guardCol + directions[direction][1]

		if nextRow >= 0 && nextRow < rows && nextCol >= 0 && nextCol < cols && input[nextRow*cols+nextCol] == '#' {
			direction = (direction + 1) % 4
			continue
		}

		guardRow = nextRow
		guardCol = nextCol
	}

	return strconv.Itoa(ct)
}

func (Day6) Part2(input []byte) string {
	directions := [4][2]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	cols := bytes.IndexByte(input, '\n') + 1
	rows := len(input)/cols + 1

	obstaclesPerCol := make([][]int, cols)
	obstaclesPerRow := make([][]int, rows)
	for i, b := range input {
		if b == '#' {
			row := i / cols
			col := i % cols
			obstaclesPerRow[row] = append(obstaclesPerRow[row], col)
			obstaclesPerCol[col] = append(obstaclesPerCol[col], row)
		}
	}

	guardI := bytes.IndexByte(input, '^')
	guardRow := guardI / cols
	guardCol := guardI % cols

	direction := 0

	visited := make([][3]int, 256)

	loops := func(obstacleRow, obstacleCol int) bool {
		guardRow := guardRow
		guardCol := guardCol
		direction := (direction + 1) % 4
		visited = visited[:0]

		for guardRow >= 0 && guardRow < rows && guardCol >= 0 && guardCol < cols-1 {
			nextRow := guardRow
			nextCol := guardCol

			if directions[direction][0] == 0 {
				if guardRow == obstacleRow && _num.Sign(obstacleCol-guardCol) == directions[direction][1] {
					nextCol = obstacleCol
				} else {
					nextCol = rows/2 + ((rows/2 + 2) * directions[direction][1])
				}
				for _, col := range obstaclesPerRow[guardRow] {
					if _num.Sign(col-guardCol) == directions[direction][1] && _num.Abs(guardCol-col) < _num.Abs(guardCol-nextCol) {
						nextCol = col
					}
				}
				nextCol -= directions[direction][1]
			} else {
				if guardCol == obstacleCol && _num.Sign(obstacleRow-guardRow) == directions[direction][0] {
					nextRow = obstacleRow
				} else {
					nextRow = cols/2 + ((cols/2 + 2) * directions[direction][0])
				}
				for _, row := range obstaclesPerCol[guardCol] {
					if _num.Sign(row-guardRow) == directions[direction][0] && _num.Abs(guardRow-row) < _num.Abs(guardRow-nextRow) {
						nextRow = row
					}
				}
				nextRow -= directions[direction][0]
			}

			if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
				break
			}

			guardRow = nextRow
			guardCol = nextCol
			direction = (direction + 1) % 4

			if input[guardRow*cols+guardCol] == byte('0'+direction) {
				return true
			}
			state := [3]int{guardRow, guardCol, direction}
			for _, v := range visited {
				if v == state {
					return true
				}
			}
			visited = append(visited, state)
		}

		return false
	}

	ct := 0
	for guardRow >= 0 && guardRow < rows && guardCol >= 0 && guardCol < cols-1 {
		input[guardRow*cols+guardCol] = byte('0' + direction)

		nextRow := guardRow + directions[direction][0]
		nextCol := guardCol + directions[direction][1]

		if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols-1 {
			break
		}

		if input[nextRow*cols+nextCol] == '#' {
			direction = (direction + 1) % 4
			continue
		}

		if input[nextRow*cols+nextCol] == '.' && loops(nextRow, nextCol) {
			ct++
		}

		guardRow = nextRow
		guardCol = nextCol
	}

	return strconv.Itoa(ct)
}
