package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	r := bufio.NewScanner(f)
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

	matrix := make([][]bool, maxY+1)
	for i := 0; i <= maxY; i++ {
		matrix[i] = make([]bool, maxX+1)
	}

	for _, point := range points {
		matrix[point[1]][point[0]] = true
	}

	var folds [][2]int
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

	matrix = fold(folds[0], matrix)
	fmt.Println(countDots(matrix))
}

func fold(f [2]int, matrix [][]bool) [][]bool {
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

func countDots(matrix [][]bool) (ct int) {
	for _, row := range matrix {
		for _, cell := range row {
			if cell {
				ct++
			}
		}
	}
	return
}
