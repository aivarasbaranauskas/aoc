package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	field := parseInput()
	lastStates := map[uint32]struct{}{}

	for {
		enc := field.encode()
		if _, ok := lastStates[enc]; ok {
			break
		}
		lastStates[enc] = struct{}{}
		field.next()
	}

	fmt.Println("FOUND!")
	fmt.Printf("Score: %v\n", field.encode())
	field.print()
}

func parseInput() fieldT {
	r := bufio.NewScanner(strings.NewReader(strings.TrimSpace(input)))
	var f [][]bool
	for r.Scan() {
		line := r.Text()
		fieldLine := make([]bool, len(line))
		for i := range line {
			if line[i] == '#' {
				fieldLine[i] = true
			}
		}
		f = append(f, fieldLine)

		if len(fieldLine) != 5 {
			log.Fatal("not 5 cols in row", len(f))
		}
	}

	if len(f) != 5 {
		log.Fatal("not 5 rows")
	}

	return fieldT{f: f}
}

type fieldT struct {
	f [][]bool
}

func (f *fieldT) next() {
	tmp := emptyField()

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			var s int
			//left
			if j > 0 && f.f[i][j-1] {
				s++
			}
			//right
			if j < 4 && f.f[i][j+1] {
				s++
			}
			//up
			if i > 0 && f.f[i-1][j] {
				s++
			}
			//down
			if i < 4 && f.f[i+1][j] {
				s++
			}

			tmp[i][j] = (!f.f[i][j] && (s == 1 || s == 2)) || (f.f[i][j] && s == 1)
		}
	}

	f.f = tmp
}

func (f *fieldT) encode() (out uint32) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if f.f[i][j] {
				out |= 1 << (i*5 + j)
			}
		}
	}
	return
}

func (f *fieldT) print() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if f.f[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func emptyField() [][]bool {
	return [][]bool{
		make([]bool, 5),
		make([]bool, 5),
		make([]bool, 5),
		make([]bool, 5),
		make([]bool, 5),
	}
}
