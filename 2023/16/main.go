package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

var (
	m [][]byte
)

func init() {
	m = bytes.Split(input, []byte("\n"))
}

func main() {
	fmt.Println("part 1:", explore([2]int{0, 1}, [2]int{0, 0}))

	p2 := 0

	for i := range m {
		x1 := explore([2]int{0, 1}, [2]int{i, 0})
		x2 := explore([2]int{0, -1}, [2]int{i, len(m[i]) - 1})
		p2 = max(p2, x1, x2)
	}

	for i := range m[0] {
		x1 := explore([2]int{1, 0}, [2]int{0, i})
		x2 := explore([2]int{-1, 0}, [2]int{len(m) - 1, i})
		p2 = max(p2, x1, x2)
	}

	fmt.Println("part 2:", p2)
}

func explore(d, p [2]int) int {
	mem := make(map[[2][2]int]struct{})
	beams := make([][]bool, len(m))
	for i := range beams {
		beams[i] = make([]bool, len(m[i]))
	}

	exploreR(&mem, &beams, d, p)

	s := 0
	for i := range beams {
		for j := range beams[i] {
			if beams[i][j] {
				s++
			}
		}
	}
	return s
}

func exploreR(mem *map[[2][2]int]struct{}, beams *[][]bool, d, p [2]int) {
	for p[0] >= 0 && p[0] < len(m) && p[1] >= 0 && p[1] < len(m[0]) {
		if _, ok := (*mem)[[2][2]int{d, p}]; ok {
			return
		}
		(*beams)[p[0]][p[1]] = true
		(*mem)[[2][2]int{d, p}] = struct{}{}
		switch m[p[0]][p[1]] {
		case '.':
			// skip
		case '\\':
			d[0], d[1] = d[1], d[0]
		case '/':
			d[0], d[1] = -1*d[1], -1*d[0]
		case '|':
			if d[0] == 0 {
				exploreR(mem, beams, [2]int{1, 0}, p)
				exploreR(mem, beams, [2]int{-1, 0}, p)
				return
			}
		case '-':
			if d[1] == 0 {
				exploreR(mem, beams, [2]int{0, 1}, p)
				exploreR(mem, beams, [2]int{0, -1}, p)
				return
			}
		}

		p[0], p[1] = p[0]+d[0], p[1]+d[1]
	}
}
