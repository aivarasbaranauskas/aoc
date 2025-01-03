package main

import (
	"bufio"
	"embed"
	"errors"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"io"
	"log"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	var (
		maxX, maxY int
		lines      [][2][2]int
	)

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		spl := strings.Split(line, " -> ")
		splStart := strings.Split(spl[0], ",")
		splEnd := strings.Split(spl[1], ",")
		x1 := optimistic.Atoi(splStart[0])
		y1 := optimistic.Atoi(splStart[1])
		x2 := optimistic.Atoi(splEnd[0])
		y2 := optimistic.Atoi(splEnd[1])

		maxX = max(maxX, x1, x2)
		maxY = max(maxY, y1, y2)
		lines = append(lines, [2][2]int{{x1, y1}, {x2, y2}})
	}
	if err = s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalln(err)
	}

	maxX++
	maxY++
	matrix := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		matrix[i] = make([]int, maxY)
	}

	for _, line := range lines {
		x1, y1, x2, y2 := line[0][0], line[0][1], line[1][0], line[1][1]
		if x1 != x2 && y1 != y2 {
			// diagonal
			xMod, yMod := 1, 1
			if x1 > x2 {
				xMod = -1
			}
			if y1 > y2 {
				yMod = -1
			}

			x := x1
			y := y1
			for x != x2+xMod && y != y2+yMod {
				matrix[x][y]++
				x += xMod
				y += yMod
			}
		} else {
			xFrom, xTo := min(x1, x2), max(x1, x2)
			yFrom, yTo := min(y1, y2), max(y1, y2)
			for x := xFrom; x <= xTo; x++ {
				for y := yFrom; y <= yTo; y++ {
					matrix[x][y]++
				}
			}
		}
	}

	var ct int
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if matrix[x][y] > 1 {
				ct++
			}
		}
	}

	fmt.Println(ct)
}
