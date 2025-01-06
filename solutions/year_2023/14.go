package year_2023

import (
	"bytes"
	"strconv"
)

func init() {
	Solutions[14] = Day14{}
}

type Day14 struct{}

func (d Day14) Part1(input []byte) string {
	m := bytes.Split(input, []byte("\n"))
	d.north(&m)
	return strconv.Itoa(d.score(m))
}

func (d Day14) Part2(input []byte) string {
	m := bytes.Split(input, []byte("\n"))
	n := 1000000000
	mem := map[string]int{}
	for i := 0; i < n; i++ {
		d.north(&m)
		d.west(&m)
		d.south(&m)
		d.east(&m)
		s := string(bytes.Join(m, nil))
		if iM, ok := mem[s]; ok {
			// found a cycle
			i += ((n - i) / (i - iM)) * (i - iM)
			mem = map[string]int{}
		}
		mem[s] = i
	}
	return strconv.Itoa(d.score(m))
}

func (Day14) north(m *[][]byte) {
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

func (Day14) west(m *[][]byte) {
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

func (Day14) south(m *[][]byte) {
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

func (Day14) east(m *[][]byte) {
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

func (Day14) score(m [][]byte) (s int) {
	ln := len(m)
	for i, l := range m {
		ct := bytes.Count(l, []byte("O"))
		s += (ln - i) * ct
	}
	return
}
