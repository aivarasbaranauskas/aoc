package year_2015

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/_map"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[13] = Day13{}
}

type Day13 struct{}

func (day Day13) Part1(input []byte) string {
	m := day.parse(input)
	names := _map.Keys(m)
	n := len(names)
	perms := day.getPermutations(n)
	maxScore := 0

	for _, perm := range perms {
		score := m[names[perm[n-1]]][names[perm[0]]]
		for i := range n - 1 {
			score += m[names[perm[i]]][names[perm[i+1]]]
		}
		maxScore = max(maxScore, score)
	}

	return strconv.Itoa(maxScore)
}

func (day Day13) Part2(input []byte) string {
	m := day.parse(input)
	names := _map.Keys(m)
	n := len(names)
	perms := day.getPermutations(n)
	maxScore := 0

	for _, perm := range perms {
		score := m[names[perm[n-1]]][names[perm[0]]]
		minScore := m[names[perm[n-1]]][names[perm[0]]]
		for i := range n - 1 {
			score += m[names[perm[i]]][names[perm[i+1]]]
			minScore = min(minScore, m[names[perm[i]]][names[perm[i+1]]])
		}
		score -= minScore
		maxScore = max(maxScore, score)
	}

	return strconv.Itoa(maxScore)
}

func (day Day13) getPermutations(n int) [][]int {
	ct := 1
	for i := 2; i < n; i++ {
		ct *= i
	}

	ps := make([][]int, 0, ct)

	var recF func(p []int)
	recF = func(p []int) {
		if len(p) == n {
			if p[n-1] > p[1] {
				pCopy := make([]int, len(p))
				copy(pCopy, p)
				ps = append(ps, pCopy)
			}
		} else {
		Loop:
			for i := 0; i < n; i++ {
				for _, v := range p {
					if v == i {
						continue Loop
					}
				}
				recF(append(p, i))
			}
		}
	}

	init := make([]int, 1, n)
	init[0] = 0
	recF(init)

	return ps
}

func (day Day13) parse(input []byte) map[string]map[string]int {
	m := make(map[string]map[string]int)

	for line := range bytes.Lines(input) {
		line = bytes.TrimSpace(line)
		firstSpacePos := bytes.IndexByte(line, ' ')

		name1 := string(line[:firstSpacePos])
		name2 := string(line[bytes.LastIndexByte(line, ' ')+1 : len(line)-1])

		isGain := line[firstSpacePos+7] == 'g' // offset 'would ' for first letter of gain or lose
		amount := optimistic.AtoiBFast(line[firstSpacePos+12 : firstSpacePos+12+bytes.IndexByte(line[firstSpacePos+12:], ' ')])
		if !isGain {
			amount *= -1
		}

		if _, ok := m[name1]; !ok {
			m[name1] = make(map[string]int)
			m[name1][name2] = amount
		} else {
			m[name1][name2] += amount
		}

		if _, ok := m[name2]; !ok {
			m[name2] = make(map[string]int)
			m[name2][name1] = amount
		} else {
			m[name2][name1] += amount
		}
	}

	return m
}
