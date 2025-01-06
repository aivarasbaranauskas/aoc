package year_2022

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"sort"
	"strconv"
	"strings"
)

func init() {
	Solutions[15] = Day15{}
}

type Day15 struct{}

func (d Day15) Part1(input []byte) string {
	var s, b [][2]int
	n := 0

	r := bufio.NewScanner(bytes.NewReader(input))
	for r.Scan() {
		spl := strings.Split(r.Text(), " ")
		s = append(s, [2]int{
			d.gI(spl[2][:len(spl[2])-1]),
			d.gI(spl[3][:len(spl[3])-1]),
		})
		b = append(b, [2]int{
			d.gI(spl[8][:len(spl[8])-1]),
			d.gI(spl[9][:len(spl[9])]),
		})
		n++
	}

	// inspected row
	iR := 2000000
	ps := _set.New[[2]int]()
	for i := 0; i < n; i++ {
		d := _num.Abs(s[i][0]-b[i][0]) + _num.Abs(s[i][1]-b[i][1])
		mx := d - _num.Abs(iR-s[i][1])
		for k := -mx; k <= mx; k++ {
			ps.Add([2]int{
				s[i][0] + k,
				iR,
			})
		}
	}

	for i := 0; i < n; i++ {
		ps.Remove(b[i])
	}

	return strconv.Itoa(ps.Len())
}

func (Day15) gI(s string) int {
	spl := strings.Split(s, "=")
	return optimistic.Atoi(spl[1])
}

func (d Day15) Part2(input []byte) string {
	var sb [][2][2]int
	n := 0

	r := bufio.NewScanner(bytes.NewReader(input))
	for r.Scan() {
		spl := strings.Split(r.Text(), " ")
		sb = append(sb, [2][2]int{
			{
				d.gI(spl[2][:len(spl[2])-1]),
				d.gI(spl[3][:len(spl[3])-1]),
			},
			{
				d.gI(spl[8][:len(spl[8])-1]),
				d.gI(spl[9][:len(spl[9])]),
			},
		})
		n++
	}

	sort.Slice(sb, func(i, j int) bool {
		return sb[i][0][1] < sb[j][0][1]
	})

	var dBs []int
	for i := 0; i < n; i++ {
		dBs = append(dBs, _num.Abs(sb[i][0][0]-sb[i][1][0])+_num.Abs(sb[i][0][1]-sb[i][1][1]))
	}

	// search field
	iR := 4000000

	for y := 0; y <= iR; y++ {
		var ranges [][2]int
		for i := 0; i < n; i++ {
			if sb[i][0][1]-dBs[i] > y || y > sb[i][0][1]+dBs[i] {
				continue
			}

			mx := dBs[i] - _num.Abs(y-sb[i][0][1])
			ranges = append(ranges, [2]int{
				max(sb[i][0][0]-mx, 0),
				min(sb[i][0][0]+mx, iR),
			})
		}

		sort.Slice(ranges, func(i, j int) bool {
			if ranges[i][0] == ranges[j][0] {
				return ranges[i][1] < ranges[j][1]
			}
			return ranges[i][0] < ranges[j][0]
		})

		rr := ranges[0]
		if rr[0] > 0 {
			x := 0
			return strconv.Itoa(x*4000000 + y)
		}
		for i := 1; i < len(ranges)-1; i++ {
			if rr[1] < ranges[i][0] {
				x := rr[1] + 1
				return strconv.Itoa(x*4000000 + y)
			} else {
				rr[1] = max(rr[1], ranges[i][1])
			}
		}
	}

	return "not found"
}
