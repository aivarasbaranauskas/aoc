package year_2021

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"io"
	"log"
	"strconv"
	"strings"
)

func init() {
	Solutions[5] = Day5{}
}

type Day5 struct{}

func (Day5) Part1(input []byte) string {
	var (
		maxX, maxY int
		lines      [][2][2]int
	)

	s := bufio.NewScanner(bytes.NewReader(input))
	for s.Scan() {
		line := s.Text()
		spl := strings.Split(line, " -> ")
		splStart := strings.Split(spl[0], ",")
		splEnd := strings.Split(spl[1], ",")
		x1 := optimistic.Atoi(splStart[0])
		y1 := optimistic.Atoi(splStart[1])
		x2 := optimistic.Atoi(splEnd[0])
		y2 := optimistic.Atoi(splEnd[1])

		if x1 != x2 && y1 != y2 {
			// Skip diagonal lines
			continue
		}

		maxX = max(maxX, x1, x2)
		maxY = max(maxY, y1, y2)
		lines = append(lines, [2][2]int{{x1, y1}, {x2, y2}})
	}
	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
		log.Fatalln(err)
	}

	maxX++
	maxY++
	matrix := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		matrix[i] = make([]int, maxY)
	}

	for _, line := range lines {
		xFrom, xTo := min(line[0][0], line[1][0]), max(line[0][0], line[1][0])
		yFrom, yTo := min(line[0][1], line[1][1]), max(line[0][1], line[1][1])
		for x := xFrom; x <= xTo; x++ {
			for y := yFrom; y <= yTo; y++ {
				matrix[x][y]++
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

	return strconv.Itoa(ct)
}

func (Day5) Part2(input []byte) string {
	var (
		maxX, maxY int
		lines      [][2][2]int
	)

	s := bufio.NewScanner(bytes.NewReader(input))
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
	if err := s.Err(); err != nil && !errors.Is(err, io.EOF) {
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

	return strconv.Itoa(ct)
}
