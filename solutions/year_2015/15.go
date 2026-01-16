package year_2015

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[15] = Day15{}
}

type Day15 struct{}

func (day Day15) Part1(input []byte) string {
	m := day.parse(input)

	maxSCore := day.bruteForce(m, Day15S{}, 100, nil)

	return strconv.Itoa(maxSCore)
}

func (day Day15) Part2(input []byte) string {

	m := day.parse(input)

	wantedCalories := 500
	maxSCore := day.bruteForce(m, Day15S{}, 100, &wantedCalories)

	return strconv.Itoa(maxSCore)
}

func (day Day15) bruteForce(m []Day15S, current Day15S, teaspoonsLeft int, wantedCalories *int) int {
	if len(m) == 0 {
		if wantedCalories != nil && *wantedCalories != current.calories {
			return 0
		}

		score := max(0, current.capacity) *
			max(0, current.durability) *
			max(0, current.flavor) *
			max(0, current.texture)
		return score
	}

	maxScore := 0
	for i := 0; i < teaspoonsLeft; i++ {
		maxScore = max(maxScore, day.bruteForce(m[1:], Day15S{
			capacity:   current.capacity + i*m[0].capacity,
			durability: current.durability + i*m[0].durability,
			flavor:     current.flavor + i*m[0].flavor,
			texture:    current.texture + i*m[0].texture,
			calories:   current.calories + i*m[0].calories,
		}, teaspoonsLeft-i, wantedCalories))
	}
	return maxScore
}

type Day15S struct {
	capacity, durability, flavor, texture, calories int
}

func (day Day15) parse(input []byte) []Day15S {
	n := bytes.Count(input, []byte{'\n'}) + 1
	m := make([]Day15S, n)
	for i := 0; i < n; i++ {
		lineEnd := bytes.Index(input, []byte{'\n'})
		var line []byte
		if lineEnd == -1 {
			line = input
			input = nil
		} else {
			line = input[:lineEnd]
			input = input[lineEnd+1:]
		}

		p := bytes.IndexByte(line, ':')
		spl := bytes.Split(line[p+2:], []byte(", "))

		m[i].capacity = optimistic.AtoiB(spl[0][bytes.IndexByte(spl[0], ' ')+1:])
		m[i].durability = optimistic.AtoiB(spl[1][bytes.IndexByte(spl[1], ' ')+1:])
		m[i].flavor = optimistic.AtoiB(spl[2][bytes.IndexByte(spl[2], ' ')+1:])
		m[i].texture = optimistic.AtoiB(spl[3][bytes.IndexByte(spl[3], ' ')+1:])
		m[i].calories = optimistic.AtoiB(spl[4][bytes.IndexByte(spl[4], ' ')+1:])
	}

	return m
}
