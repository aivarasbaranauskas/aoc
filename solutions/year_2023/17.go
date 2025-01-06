package year_2023

import (
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"math"
	"strconv"
	"strings"
)

func init() {
	Solutions[17] = Day17{}
}

type Day17 struct{}

func (d Day17) Part1(input []byte) string {
	m := _slice.Map(
		strings.Split(string(input), "\n"),
		func(line string) []int {
			return _slice.Map(
				strings.Split(line, ""),
				optimistic.Atoi,
			)
		},
	)
	mem := make(map[MemKey]Mem)
	d.walk1(m, &mem)

	s := math.MaxInt
	end := [2]int{len(m) - 1, len(m[0]) - 1}
	for i := 0; i < 3; i++ {
		memK := MemKey{
			p:   end,
			d:   [2]int{1, 0},
			dCt: i,
		}
		if sT, ok := mem[memK]; ok {
			if sT.s < s {
				s = sT.s
			}
		}
		memK.d = [2]int{0, 1}
		if sT, ok := mem[memK]; ok {
			if sT.s < s {
				s = sT.s
			}
		}
	}

	return strconv.Itoa(s + m[end[0]][end[1]] - m[0][0])
}

func (d Day17) Part2(input []byte) string {
	m := _slice.Map(
		strings.Split(string(input), "\n"),
		func(line string) []int {
			return _slice.Map(
				strings.Split(line, ""),
				optimistic.Atoi,
			)
		},
	)
	mem := make(map[MemKey]Mem)
	d.walk2(m, &mem)

	s := math.MaxInt
	end := [2]int{len(m) - 1, len(m[0]) - 1}
	for i := 4; i < 10; i++ {
		memK := MemKey{
			p:   end,
			d:   [2]int{1, 0},
			dCt: i,
		}
		if sT, ok := mem[memK]; ok {
			if sT.s < s {
				s = sT.s
			}
		}
		memK.d = [2]int{0, 1}
		if sT, ok := mem[memK]; ok {
			if sT.s < s {
				s = sT.s
			}
		}
	}

	return strconv.Itoa(s + m[end[0]][end[1]] - m[0][0])
}

type MemKey struct {
	p, d [2]int
	dCt  int
}

type Mem struct {
	s int
}

type Step struct {
	p, d [2]int
	dCt  int
	s    int
}

func (Day17) walk1(m [][]int, mem *map[MemKey]Mem) {
	var q _a.Queue[Step]
	q.Enqueue(Step{
		p:   [2]int{0, 0},
		d:   [2]int{0, 0},
		dCt: 0,
		s:   0,
	})

	for !q.Empty() {
		s := q.Dequeue()
		if s.dCt > 3 {
			// to long straight
			continue
		}
		if s.p[0] < 0 || s.p[0] >= len(m) || s.p[1] < 0 || s.p[1] >= len(m[0]) {
			// out of bounds
			continue
		}

		memK := MemKey{
			p:   s.p,
			d:   s.d,
			dCt: s.dCt,
		}

		if t, ok := (*mem)[memK]; ok && s.s >= t.s {
			// lower value already found in this pos
			continue
		}
		(*mem)[memK] = Mem{
			s: s.s,
		}
		newS := s.s + m[s.p[0]][s.p[1]]

		if s.d != [2]int{-1, 0} {
			newdCt := 1
			if s.d == [2]int{1, 0} {
				newdCt = s.dCt + 1
			}
			q.Enqueue(Step{
				p:   [2]int{s.p[0] + 1, s.p[1]},
				d:   [2]int{1, 0},
				dCt: newdCt,
				s:   newS,
			})
		}
		if s.d != [2]int{1, 0} {
			newdCt := 1
			if s.d == [2]int{-1, 0} {
				newdCt = s.dCt + 1
			}
			q.Enqueue(Step{
				p:   [2]int{s.p[0] - 1, s.p[1]},
				d:   [2]int{-1, 0},
				dCt: newdCt,
				s:   newS,
			})
		}
		if s.d != [2]int{0, -1} {
			newdCt := 1
			if s.d == [2]int{0, 1} {
				newdCt = s.dCt + 1
			}
			q.Enqueue(Step{
				p:   [2]int{s.p[0], s.p[1] + 1},
				d:   [2]int{0, 1},
				dCt: newdCt,
				s:   newS,
			})
		}
		if s.d != [2]int{0, 1} {
			newdCt := 1
			if s.d == [2]int{0, -1} {
				newdCt = s.dCt + 1
			}
			q.Enqueue(Step{
				p:   [2]int{s.p[0], s.p[1] - 1},
				d:   [2]int{0, -1},
				dCt: newdCt,
				s:   newS,
			})
		}
	}
}

func (Day17) walk2(m [][]int, mem *map[MemKey]Mem) {
	var q _a.Queue[Step]
	q.Enqueue(Step{
		p:   [2]int{0, 0},
		d:   [2]int{0, 0},
		dCt: 0,
		s:   0,
	})

	for !q.Empty() {
		s := q.Dequeue()
		if s.dCt > 10 {
			// to long straight
			continue
		}
		if s.p[0] < 0 || s.p[0] >= len(m) || s.p[1] < 0 || s.p[1] >= len(m[0]) {
			// out of bounds
			continue
		}

		memK := MemKey{
			p:   s.p,
			d:   s.d,
			dCt: s.dCt,
		}

		if t, ok := (*mem)[memK]; ok && s.s >= t.s {
			// lower value already found in this pos
			continue
		}
		(*mem)[memK] = Mem{
			s: s.s,
		}
		ds := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

		for _, d := range ds {
			if s.d[0] == -1*d[0] && s.d[1] == -1*d[1] {
				// no going back
				continue
			}

			if s.d == d {
				// straight
				next := [2]int{s.p[0] + d[0], s.p[1] + d[1]}
				q.Enqueue(Step{
					p:   next,
					d:   d,
					dCt: s.dCt + 1,
					s:   s.s + m[s.p[0]][s.p[1]],
				})
				continue
			}

			// turn
			newS := s.s
			for i := 0; i < 4; i++ {
				if s.p[0]+i*d[0] < 0 || s.p[0]+i*d[0] >= len(m) || s.p[1]+i*d[1] < 0 || s.p[1]+i*d[1] >= len(m[0]) {
					// out of bounds
					break
				}
				newS += m[s.p[0]+i*d[0]][s.p[1]+i*d[1]]
			}
			next := [2]int{s.p[0] + 4*d[0], s.p[1] + 4*d[1]}
			q.Enqueue(Step{
				p:   next,
				d:   d,
				dCt: 4,
				s:   newS,
			})
		}
	}
}
