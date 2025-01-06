package year_2023

import (
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"math"
	"strconv"
	"strings"
)

func init() {
	Solutions[5] = Day5{}
}

type Day5 struct{}

func (d Day5) Part1(input []byte) string {
	seeds, ms := d.parseData(input)
	minLoc := math.MaxInt

	for _, seed := range seeds {
	SectionLoop:
		for _, m := range ms {
			for _, l := range m {
				if l[1] <= seed && seed <= l[1]+l[2] {
					seed = l[0] + (seed - l[1])
					continue SectionLoop
				}
			}
		}
		minLoc = min(minLoc, seed)
	}

	return strconv.Itoa(minLoc)
}

func (d Day5) Part2(input []byte) string {
	seeds, ms := d.parseData(input)
	var mms [][]struct {
		from   [2]int
		offset int
	}
	for _, a := range ms {
		var aa []struct {
			from   [2]int
			offset int
		}
		for _, b := range a {
			aa = append(aa, struct {
				from   [2]int
				offset int
			}{
				from: [2]int{
					b[1],
					b[1] + b[2] - 1,
				},
				offset: b[0] - b[1],
			})
		}
		mms = append(mms, aa)
	}

	var rs [][2]int
	for i := 0; i < len(seeds); i += 2 {
		rs = append(rs, [2]int{seeds[i], seeds[i] + seeds[i+1] - 1})
	}

	for _, mm := range mms {
		var tmpTransformed [][2]int
		for _, r := range rs {
			tmpR := [][2]int{r}
		Loop:
			for _, m := range mm {
				for i, r := range tmpR {
					if r[0] <= m.from[1] && m.from[0] <= r[1] {
						from := max(r[0], m.from[0])
						to := min(r[1], m.from[1])
						tmpTransformed = append(tmpTransformed, [2]int{from + m.offset, to + m.offset})
						if r[0] < from {
							tmpR = append(tmpR, [2]int{r[0], from - 1})
						}
						if to < r[1] {
							tmpR = append(tmpR, [2]int{to + 1, r[1]})
						}

						copy(tmpR[i:], tmpR[i+1:])
						tmpR = tmpR[:len(tmpR)-1]

						continue Loop
					}
				}
			}
			tmpTransformed = append(tmpTransformed, tmpR...)
		}
		rs = tmpTransformed
	}

	minLoc := math.MaxInt
	for _, r := range rs {
		minLoc = min(minLoc, r[0])
	}

	return strconv.Itoa(minLoc)
}

func (d Day5) parseData(input []byte) ([]int, [][][]int) {
	sections := strings.Split(string(input), "\n\n")
	seeds := _slice.Map(strings.Split(strings.Split(sections[0], ": ")[1], " "), optimistic.Atoi)

	var ms [][][]int
	for i := 1; i < len(sections); i++ {
		ms = append(ms, d.parseMap(sections[i]))
	}

	return seeds, ms
}

func (Day5) parseMap(section string) [][]int {
	var m [][]int
	for _, line := range strings.Split(section, "\n")[1:] {
		m = append(m, _slice.Map(strings.Split(line, " "), optimistic.Atoi))
	}
	return m
}
