package year_2023

import (
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[12] = Day12{}
}

type Day12 struct{}

func (d Day12) Part1(input []byte) string {
	recs := d.parseInput(input)
	s := 0
	for _, rec := range recs {
		s += rec.countVariations(0, 0)
	}
	return strconv.Itoa(s)
}

func (d Day12) Part2(input []byte) string {
	recs := d.parseInput(input)
	s := 0
	for _, rec := range recs {
		rec.conds = bytes.Join([][]byte{rec.conds, rec.conds, rec.conds, rec.conds, rec.conds}, []byte("?"))
		rec.groups = append(rec.groups, append(rec.groups, append(rec.groups, append(rec.groups, rec.groups...)...)...)...)

		x := rec.countVariations(0, 0)
		s += x
	}
	return strconv.Itoa(s)
}

func (Day12) parseInput(input []byte) []Rec {
	lines := strings.Split(string(input), "\n")
	recs := make([]Rec, len(lines))
	for i := range lines {
		spl := strings.Split(lines[i], " ")
		recs[i] = Rec{
			conds:  []byte(spl[0]),
			groups: _slice.Map(strings.Split(spl[1], ","), optimistic.Atoi),
			mem:    make(map[[2]int]int),
		}
	}
	return recs
}

type Rec struct {
	conds  []byte
	groups []int
	mem    map[[2]int]int
}

func (r *Rec) countVariations(cur, curGroup int) (ct int) {
	if ctt, ok := r.mem[[2]int{cur, curGroup}]; ok {
		return ctt
	}

	if cur >= len(r.conds) {
		if curGroup == len(r.groups) {
			return 1
		}
		return 0
	}

	c := r.conds[cur]

	if c == '.' || c == '?' {
		ct += r.countVariations(cur+1, curGroup)
	}

	if (c == '?' || c == '#') &&
		curGroup < len(r.groups) &&
		(cur+r.groups[curGroup] <= len(r.conds) && bytes.Index(r.conds[cur:cur+r.groups[curGroup]], []byte(".")) == -1) &&
		(cur+r.groups[curGroup] == len(r.conds) || r.conds[cur+r.groups[curGroup]] != '#') {
		ct += r.countVariations(cur+1+r.groups[curGroup], curGroup+1)
	}

	r.mem[[2]int{cur, curGroup}] = ct
	return ct
}
