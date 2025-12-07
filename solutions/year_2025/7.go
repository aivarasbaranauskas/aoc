package year_2025

import (
	"bytes"
	"strconv"
)

func init() {
	Solutions[7] = Day7{}
}

type Day7 struct{}

func (Day7) Part1(input []byte) string {
	m := bytes.Split(input, []byte{'\n'})
	ct := 0

	for r := range len(m) - 1 {
		for c := range m[r] {
			if m[r][c] != 'S' && m[r][c] != '|' {
				continue
			}

			if m[r+1][c] == '.' {
				m[r+1][c] = '|'
				continue
			}

			if m[r+1][c] != '^' {
				continue
			}

			ct++
			if c > 1 && m[r+1][c-1] == '.' {
				m[r+1][c-1] = '|'
			}
			if c < len(m[0])-1 && m[r+1][c+1] == '.' {
				m[r+1][c+1] = '|'
			}
		}
	}

	return strconv.Itoa(ct)
}

func (Day7) Part2(input []byte) string {
	m := bytes.Split(input, []byte{'\n'})
	buf := make([]int, len(m[0]))
	tmp := 0

	for i := range m[0] {
		if m[0][i] == 'S' {
			buf[i] = 1
			break
		}
	}

	for r := range len(m) {
		tmp = 0
		for c := range m[r] {
			if m[r][c] != '^' {
				buf[c] += tmp
				tmp = 0
				continue
			}

			if c >= 1 {
				buf[c-1] += buf[c]
			}
			if c < len(m[0])-1 {
				tmp = buf[c]
			}
			buf[c] = 0
		}
	}

	sum := 0
	for i := range buf {
		sum += buf[i]
	}

	return strconv.Itoa(sum)
}
