package year_2021

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[13] = Day13{}
}

type Day13 struct{}

func (d Day13) Part1(input []byte) string {
	matrix, folds := d.readData(input)
	matrix = d.fold(folds[0], matrix)
	ct := 0
	for _, row := range matrix {
		for _, cell := range row {
			if cell {
				ct++
			}
		}
	}
	return strconv.Itoa(ct)
}

func (d Day13) Part2(input []byte) string {
	matrix, folds := d.readData(input)
	for _, f := range folds {
		matrix = d.fold(f, matrix)
	}
	s := strings.Builder{}
	for _, row := range matrix {
		s.WriteByte('\n')
		for _, cell := range row {
			if cell {
				s.WriteByte('#')
			} else {
				s.WriteByte(' ')
			}
		}
	}
	return s.String()
}

func (Day13) readData(input []byte) (matrix [][]bool, folds [][2]int) {
	r := bufio.NewScanner(bytes.NewReader(input))
	var points [][2]int
	var maxX, maxY int
	for r.Scan() {
		line := strings.TrimSpace(r.Text())
		if line == "" {
			break
		}
		spl := strings.Split(line, ",")
		x := optimistic.Atoi(spl[0])
		y := optimistic.Atoi(spl[1])
		points = append(points, [2]int{x, y})
		maxX = max(maxX, x)
		maxY = max(maxY, y)
	}

	matrix = make([][]bool, maxY+1)
	for i := 0; i <= maxY; i++ {
		matrix[i] = make([]bool, maxX+1)
	}

	for _, point := range points {
		matrix[point[1]][point[0]] = true
	}

	for r.Scan() {
		line := strings.TrimSpace(r.Text())
		if line == "" {
			break
		}
		spl := strings.Split(line, " ")
		spl = strings.Split(spl[2], "=")
		if spl[0] == "x" {
			folds = append(folds, [2]int{0, optimistic.Atoi(spl[1])})
		} else {
			folds = append(folds, [2]int{1, optimistic.Atoi(spl[1])})
		}
	}

	return
}

func (Day13) fold(f [2]int, matrix [][]bool) [][]bool {
	if f[0] == 0 {
		for y := range matrix {
			for x := f[1] + 1; x < len(matrix[y]); x++ {
				newX := f[1] - (x - f[1])
				matrix[y][newX] = matrix[y][newX] || matrix[y][x]
			}
			matrix[y] = matrix[y][:f[1]]
		}
	} else {
		for y := f[1] + 1; y < len(matrix); y++ {
			for x := range matrix[y] {
				newY := f[1] - (y - f[1])
				matrix[newY][x] = matrix[newY][x] || matrix[y][x]
			}
		}
		matrix = matrix[:f[1]]
	}
	return matrix
}
