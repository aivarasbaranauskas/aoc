package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	var m [][]byte

	r := bufio.NewScanner(f)
	for r.Scan() {
		m = append(m, []byte(r.Text()))
	}
	moved := true
	ct := 0
	w, h := len(m[0]), len(m)
	//printM(m)
	for moved {
		moved = false
		ct++

		tmp := copyM(m)
		for l := 0; l < h; l++ {
			for c := 0; c < w; c++ {
				if m[l][c] == '>' {
					cn := (c + 1) % w
					if m[l][cn] == '.' {
						tmp[l][c], tmp[l][cn] = '.', '>'
						moved = true
						c++
					}
				}
			}
		}

		m = tmp
		tmp = copyM(m)
		for c := 0; c < w; c++ {
			for l := 0; l < h; l++ {
				if m[l][c] == 'v' {
					ln := (l + 1) % h
					if m[l][c] == 'v' && m[ln][c] == '.' {
						tmp[l][c], tmp[ln][c] = '.', 'v'
						moved = true
						l++
					}
				}
			}
		}
		m = tmp

		//printM(m)
	}

	printM(m)
	fmt.Println(ct)
}

func copyM(m [][]byte) [][]byte {
	m2 := make([][]byte, len(m))
	for i := range m {
		m2[i] = make([]byte, len(m[i]))
		copy(m2[i], m[i])
	}
	return m2
}

func printM(m [][]byte) {
	for _, r := range m {
		fmt.Println(string(r))
	}
	fmt.Println()
}
