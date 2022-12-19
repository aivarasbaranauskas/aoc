package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"log"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	m := make([][][]bool, 101)
	for i := range m {
		m[i] = make([][]bool, 101)
		for j := range m[i] {
			m[i][j] = make([]bool, 101)
		}
	}

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " ")
		b := spl[0] == "on"
		spl = strings.Split(spl[1], ",")

		splx := strings.Split(spl[0][2:], "..")
		xFrom := _num.Max(-50, optimistic.Atoi(splx[0]))
		xTo := _num.Min(50, optimistic.Atoi(splx[1]))

		sply := strings.Split(spl[1][2:], "..")
		yFrom := _num.Max(-50, optimistic.Atoi(sply[0]))
		yTo := _num.Min(50, optimistic.Atoi(sply[1]))

		splz := strings.Split(spl[2][2:], "..")
		zFrom := _num.Max(-50, optimistic.Atoi(splz[0]))
		zTo := _num.Min(50, optimistic.Atoi(splz[1]))

		for x := xFrom; x <= xTo; x++ {
			for y := yFrom; y <= yTo; y++ {
				for z := zFrom; z <= zTo; z++ {
					m[x+50][y+50][z+50] = b
				}
			}
		}
	}

	var ct int
	for _, a := range m {
		for _, b := range a {
			for _, c := range b {
				if c {
					ct++
				}
			}
		}
	}
	fmt.Println(ct)
}
