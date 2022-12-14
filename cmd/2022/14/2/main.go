package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"log"
	"math"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	r := bufio.NewScanner(f)
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
			maxY = _num.Max(maxY, y)
			maxX = _num.Max(maxX, x)
			minX = _num.Min(minX, x)
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
			fromX := _num.Min(rockLine[i][0], rockLine[i+1][0])
			toX := _num.Max(rockLine[i][0], rockLine[i+1][0])
			for x := fromX; x <= toX; x++ {
				fromY := _num.Min(rockLine[i][1], rockLine[i+1][1])
				toY := _num.Max(rockLine[i][1], rockLine[i+1][1])
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

	for _, l := range m {
		for _, c := range l {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
	fmt.Println(ct)
}
