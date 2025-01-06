package year_2023

import (
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"strconv"
)

func init() {
	Solutions[3] = Day3{}
}

type Day3 struct{}

func (d Day3) Part1(input []byte) string {
	m := bytes.Split(input, []byte("\n"))
	s := 0

	hasAdjF := func(line, col int) (has bool) {
		for l := max(0, line-1); l < min(line+2, len(m)); l++ {
			for c := max(0, col-1); c < min(col+2, len(m[line])); c++ {
				has = has || d.isSym(m[l][c])
			}
		}
		return
	}

	for i := range m {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == '.' {
				continue
			}

			if d.isNum(m[i][j]) {
				hasAdj := false
				n := 0
				for ; j < len(m[i]) && d.isNum(m[i][j]); j++ {
					hasAdj = hasAdj || hasAdjF(i, j)
					n = n*10 + int(m[i][j]-'0')
				}
				if hasAdj {
					s += n
				}
			}
		}
	}

	return strconv.Itoa(s)
}

func (d Day3) Part2(input []byte) string {
	m := bytes.Split(input, []byte("\n"))
	s := 0

	getNum := func(l, col int) (int, int) {
		line := m[l]
		beg := col
		for ; beg > 0 && d.isNum(line[beg-1]); beg-- {
		}
		num := 0
		for i := beg; i < len(line) && d.isNum(line[i]); i++ {
			num = num*10 + int(line[i]-'0')
		}
		return num, beg
	}

	for i := range m {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] != '*' {
				continue
			}

			ct := 0
			mul := 1
			for l := max(0, i-1); l < min(i+2, len(m)); l++ {
				numsBegs := _set.New[int]()
				for c := max(0, j-1); c < min(j+2, len(m[i])); c++ {
					if d.isNum(m[l][c]) {
						num, beg := getNum(l, c)
						if !numsBegs.Has(beg) {
							ct++
							mul *= num
							numsBegs.Add(beg)
						}
					}
				}
			}
			if ct == 2 {
				s += mul
			}
		}
	}

	return strconv.Itoa(s)
}

func (d Day3) isSym(b byte) bool {
	return b != '.' && !d.isNum(b)
}

func (Day3) isNum(b byte) bool {
	return '0' <= b && b <= '9'
}
