package year_2015

import (
	"bytes"
	"slices"
	"sort"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[24] = Day24{}
}

type Day24 struct{}

func (day Day24) Part1(input []byte) string {
	return day.solve(input, 3)
}

func (day Day24) Part2(input []byte) string {
	return day.solve(input, 4)
}

func (day Day24) solve(input []byte, groupsN int) string {
	items := day.parse(input)
	sort.Ints(items)
	slices.Reverse(items)

	groupSize := _num.Sum(items...) / groupsN

	var (
		minGroupN, minGroupQE int
		checkRec              func(n, s, qe, iStart int)
	)

	checkRec = func(n, s, qe, i int) {
		if minGroupN > 0 {
			if n > minGroupN {
				return
			}
			if n == minGroupN {
				if s != groupSize {
					return
				}
				minGroupQE = min(minGroupQE, qe)
				return
			}
		}

		if s == groupSize {
			minGroupN = n
			minGroupQE = qe
			return
		}

		for ; i < len(items) && s+items[i] > groupSize; i++ {
			// skip those that will overflow
		}
		for ; i < len(items); i++ {
			checkRec(n+1, s+items[i], qe*items[i], i+1)
		}
	}

	checkRec(0, 0, 1, 0)

	return strconv.Itoa(minGroupQE)
}

func (day Day24) parse(input []byte) []int {
	return _slice.Map(
		bytes.Split(input, []byte("\n")),
		optimistic.AtoiBFast,
	)
}
