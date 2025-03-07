package year_2024

import (
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"sort"
	"strconv"
)

func init() {
	Solutions[1] = Day1{}
}

type Day1 struct{}

func (Day1) Part1(input []byte) string {
	var l1, l2 [1000]int
	for i := 0; i*14 < len(input); i++ {
		iv := i * 14
		l1[i] = optimistic.AtoiBFast(input[iv : iv+5])
		l2[i] = optimistic.AtoiBFast(input[iv+8 : iv+13])
	}

	sort.Ints(l1[:])
	sort.Ints(l2[:])

	s := 0
	for i, v := range l1 {
		s += _num.Abs(v - l2[i])
	}

	return strconv.Itoa(s)
}

func (Day1) Part2(input []byte) string {
	var l [1000]int
	m := make(map[int]int, 1000)
	for i := 0; i*14 < len(input); i++ {
		iv := i * 14
		l[i] = optimistic.AtoiBFast(input[iv : iv+5])
		v2 := optimistic.AtoiBFast(input[iv+8 : iv+13])
		if _, ok := m[v2]; ok {
			m[v2]++
		} else {
			m[v2] = 1
		}
	}

	s := 0
	for _, v := range l {
		if freq, ok := m[v]; ok {
			s += v * freq
		}
	}
	return strconv.Itoa(s)
}
