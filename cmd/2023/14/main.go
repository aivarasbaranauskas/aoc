package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	fmt.Println("part 1:", p1())
	fmt.Println("part 2:", p2())
}

func p2() int {
	m := bytes.Split(input, []byte("\n"))
	n := 1000000000
	mem := map[string]int{}
	for i := 0; i < n; i++ {
		north(&m)
		west(&m)
		south(&m)
		east(&m)
		s := string(bytes.Join(m, nil))
		if iM, ok := mem[s]; ok {
			// found a cycle
			i += ((n - i) / (i - iM)) * (i - iM)
			mem = map[string]int{}
		}
		mem[s] = i
	}
	return score(m)
}

func p1() int {
	m := bytes.Split(input, []byte("\n"))
	north(&m)
	return score(m)
}

func north(m *[][]byte) {
	for c := range (*m)[0] {
		p := -1
		for l := 0; l < len(*m); l++ {
			switch (*m)[l][c] {
			case '.':
				if p == -1 {
					p = l
				}
			case 'O':
				if p > -1 {
					(*m)[p][c] = 'O'
					(*m)[l][c] = '.'
					p++
				}
			case '#':
				p = -1
			}
		}
	}
}

func west(m *[][]byte) {
	for l := range *m {
		p := -1
		for c := 0; c < len((*m)[0]); c++ {
			switch (*m)[l][c] {
			case '.':
				if p == -1 {
					p = c
				}
			case 'O':
				if p > -1 {
					(*m)[l][p] = 'O'
					(*m)[l][c] = '.'
					p++
				}
			case '#':
				p = -1
			}
		}
	}
}

func south(m *[][]byte) {
	for c := range (*m)[0] {
		p := -1
		for l := len(*m) - 1; l >= 0; l-- {
			switch (*m)[l][c] {
			case '.':
				if p == -1 {
					p = l
				}
			case 'O':
				if p > -1 {
					(*m)[p][c] = 'O'
					(*m)[l][c] = '.'
					p--
				}
			case '#':
				p = -1
			}
		}
	}
}

func east(m *[][]byte) {
	for l := range *m {
		p := -1
		for c := len((*m)[0]) - 1; c >= 0; c-- {
			switch (*m)[l][c] {
			case '.':
				if p == -1 {
					p = c
				}
			case 'O':
				if p > -1 {
					(*m)[l][p] = 'O'
					(*m)[l][c] = '.'
					p--
				}
			case '#':
				p = -1
			}
		}
	}
}

func score(m [][]byte) (s int) {
	ln := len(m)
	for i, l := range m {
		ct := bytes.Count(l, []byte("O"))
		s += (ln - i) * ct
	}
	return
}
