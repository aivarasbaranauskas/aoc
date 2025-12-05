package year_2025

import (
	"bytes"
	"sort"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[5] = Day5{}
}

type Day5 struct{}

func (Day5) Part1(input []byte) string {
	spl := bytes.Split(input, []byte("\n"))
	var ranges [][2]int
	i := 0
	for ; len(spl[i]) > 0; i++ {
		spl2 := bytes.Split(spl[i], []byte{'-'})
		ranges = append(ranges, [2]int{
			optimistic.AtoiBFast(spl2[0]),
			optimistic.AtoiBFast(spl2[1]),
		})
	}

	i++

	ct := 0
	for ; i < len(spl); i++ {
		x := optimistic.AtoiBFast(spl[i])
		for _, r := range ranges {
			if r[0] <= x && x <= r[1] {
				ct++
				break
			}
		}
	}

	return strconv.Itoa(ct)
}

func (Day5) Part2(input []byte) string {
	rangesEnd := bytes.Index(input, []byte("\n\n"))
	rangesCount := bytes.Count(input[:rangesEnd], []byte("\n")) + 1
	ranges := make([][2]int, rangesCount)

	ii := 0
	for i := range rangesCount {
		for ; input[ii] != '-'; ii++ {
			ranges[i][0] = ranges[i][0]*10 + int(input[ii]-'0')
		}
		ii++
		for ; input[ii] != '\n'; ii++ {
			ranges[i][1] = ranges[i][1]*10 + int(input[ii]-'0')
		}
		ii++
	}

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		}
		return ranges[i][0] < ranges[j][0]
	})

	last := 0
	ct := 0
	for _, r := range ranges {
		if r[1] <= last {
			continue
		}

		var n int
		if r[0] > last {
			n = r[1] - r[0] + 1
		} else {
			n = r[1] - last
		}

		last = r[1]
		ct += n
	}

	return strconv.Itoa(ct)
}
