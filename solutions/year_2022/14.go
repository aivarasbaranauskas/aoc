package year_2022

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"math"
	"strconv"
	"strings"
)

func init() {
	Solutions[14] = Day14{}
}

type Day14 struct{}

func (Day14) Part1(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var (
		rockLines        [][][2]int
		maxY, maxX, minX int
	)
	minX = math.MaxInt
	for r.Scan() {
		var line [][2]int
		spl := strings.Split(r.Text(), " -> ")
		for _, dot := range spl {
			spl2 := strings.Split(dot, ",")
			x := optimistic.Atoi(spl2[0])
			y := optimistic.Atoi(spl2[1])
			maxY = max(maxY, y)
			maxX = max(maxX, x)
			minX = min(minX, x)
			line = append(line, [2]int{x, y})
		}
		rockLines = append(rockLines, line)
	}

	xSize, ySize := maxX-minX+1, maxY+1
	m := make([][]byte, ySize)
	for i := 0; i < ySize; i++ {
		m[i] = make([]byte, xSize)
		for j := 0; j < xSize; j++ {
			m[i][j] = '.'
		}
	}

	for _, rockLine := range rockLines {
		for i := 0; i < len(rockLine)-1; i++ {
			fromX := min(rockLine[i][0], rockLine[i+1][0])
			toX := max(rockLine[i][0], rockLine[i+1][0])
			for x := fromX; x <= toX; x++ {
				fromY := min(rockLine[i][1], rockLine[i+1][1])
				toY := max(rockLine[i][1], rockLine[i+1][1])
				for y := fromY; y <= toY; y++ {
					adjX := x - minX
					m[y][adjX] = '#'
				}
			}
		}
	}

	xStart, yStart := 500-minX, 0
	x, y := xStart, yStart
	ct := 0
	for {
		if m[y+1][x] == '.' {
			y++
			continue
		}
		if x-1 < 0 {
			break
		}
		if m[y+1][x-1] == '.' {
			y++
			x--
			continue
		}
		if x+1 >= xSize {
			break
		}
		if m[y+1][x+1] == '.' {
			y++
			x++
			continue
		}
		m[y][x] = 'o'
		ct++
		x, y = xStart, yStart
	}

	return strconv.Itoa(ct)
}

func (Day14) Part2(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var (
		rockLines        [][][2]int
		maxY, maxX, minX int
	)
	minX = math.MaxInt
	for r.Scan() {
		var line [][2]int
		spl := strings.Split(r.Text(), " -> ")
		for _, dot := range spl {
			spl2 := strings.Split(dot, ",")
			x := optimistic.Atoi(spl2[0])
			y := optimistic.Atoi(spl2[1])
			maxY = max(maxY, y)
			maxX = max(maxX, x)
			minX = min(minX, x)
			line = append(line, [2]int{x, y})
		}
		rockLines = append(rockLines, line)
	}

	ySize := maxY + 2
	xSize := 1 + 2*ySize
	xCenter := ySize + 1
	m := make([][]byte, ySize)
	for i := 0; i < ySize; i++ {
		m[i] = make([]byte, xSize)
		for j := 0; j < xSize; j++ {
			m[i][j] = '.'
		}
	}

	for _, rockLine := range rockLines {
		for i := 0; i < len(rockLine)-1; i++ {
			fromX := min(rockLine[i][0], rockLine[i+1][0])
			toX := max(rockLine[i][0], rockLine[i+1][0])
			for x := fromX; x <= toX; x++ {
				fromY := min(rockLine[i][1], rockLine[i+1][1])
				toY := max(rockLine[i][1], rockLine[i+1][1])
				for y := fromY; y <= toY; y++ {
					adjX := x - 500 + xCenter
					m[y][adjX] = '#'
				}
			}
		}
	}

	xStart, yStart := xCenter, 0
	x, y := xStart, yStart
	ct := 0
	for {
		if m[y][x] == 'o' {
			break
		}

		if y+1 == ySize {
			m[y][x] = 'o'
			ct++
			x, y = xStart, yStart
		}

		if m[y+1][x] == '.' {
			y++
			continue
		}
		if x-1 < 0 {
			break
		}
		if m[y+1][x-1] == '.' {
			y++
			x--
			continue
		}
		if x+1 >= xSize {
			break
		}
		if m[y+1][x+1] == '.' {
			y++
			x++
			continue
		}
		m[y][x] = 'o'
		ct++
		x, y = xStart, yStart
	}

	return strconv.Itoa(ct)
}
