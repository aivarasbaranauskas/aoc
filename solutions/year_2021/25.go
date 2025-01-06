package year_2021

import (
	"bufio"
	"bytes"
	"strconv"
)

func init() {
	Solutions[25] = Day25{}
}

type Day25 struct{}

func (d Day25) Part1(input []byte) string {
	var m [][]byte

	r := bufio.NewScanner(bytes.NewReader(input))
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

		tmp := d.copyM(m)
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
		tmp = d.copyM(m)
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
	}

	return strconv.Itoa(ct)
}

func (Day25) copyM(m [][]byte) [][]byte {
	m2 := make([][]byte, len(m))
	for i := range m {
		m2[i] = make([]byte, len(m[i]))
		copy(m2[i], m[i])
	}
	return m2
}

func (Day25) Part2(_ []byte) string {
	return ""
}
