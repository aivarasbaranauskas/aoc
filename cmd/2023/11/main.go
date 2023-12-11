package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
)

//go:embed input.txt
var input []byte

func main() {
	m := bytes.Split(input, []byte("\n"))

	// mark
	cols := make([]bool, len(m[0]))
	rows := make([]bool, len(m))
	for i := 0; i < len(m); i++ {
		for _, c := range m[i] {
			if c == '#' {
				rows[i] = true
				break
			}
		}
	}
	for i := 0; i < len(m[0]); i++ {
		for j := 0; j < len(m); j++ {
			if m[j][i] == '#' {
				cols[i] = true
				break
			}
		}
	}

	// collect coordinates
	var galaxies [][2]int
	for i := range m {
		for j := range m[i] {
			if m[i][j] == '#' {
				galaxies = append(galaxies, [2]int{i, j})
			}
		}
	}

	//calc distances
	p1 := 0
	p2 := 0
	for i, g1 := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			g2 := galaxies[j]
			d := _num.Abs(g1[0]-g2[0]) + _num.Abs(g1[1]-g2[1])

			lB, lE := min(g1[0], g2[0]), max(g1[0], g2[0])
			cB, cE := min(g1[1], g2[1]), max(g1[1], g2[1])

			exp := 0
			for k := lB; k < lE; k++ {
				if !rows[k] {
					exp++
				}
			}
			for k := cB; k < cE; k++ {
				if !cols[k] {
					exp++
				}
			}

			p1 += (d - exp) + exp*2
			p2 += (d - exp) + exp*1000000
		}
	}

	fmt.Println("part 1:", p1)
	fmt.Println("part 2:", p2)
}
