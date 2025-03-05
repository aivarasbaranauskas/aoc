package year_2024

import (
	"bytes"
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
	lines := bytes.Split(input, []byte("\n"))
	l1, l2 := make([]int, len(lines)), make([]int, len(lines))
	for i, line := range lines {
		spl := bytes.Split(line, []byte("   "))
		l1[i] = optimistic.Atoi(string(spl[0]))
		l2[i] = optimistic.Atoi(string(spl[1]))
	}
	sort.Ints(l1)
	sort.Ints(l2)
	s := 0
	for i, v := range l1 {
		s += _num.Abs(v - l2[i])
	}
	return strconv.Itoa(s)
}

func (Day1) Part2(input []byte) string {
	lines := bytes.Split(input, []byte("\n"))
	l, m := make([]int, len(lines)), make(map[int]int)
	for i, line := range lines {
		spl := bytes.Split(line, []byte("   "))
		l[i] = optimistic.Atoi(string(spl[0]))
		v2 := optimistic.Atoi(string(spl[1]))
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
