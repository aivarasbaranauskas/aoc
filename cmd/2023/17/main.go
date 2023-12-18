package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"math"
	"strings"
)

//go:embed input.txt
var input string

var m [][]int

func init() {
	m = _slice.Map(
		strings.Split(input, "\n"),
		func(line string) []int {
			return _slice.Map(
				strings.Split(line, ""),
				optimistic.Atoi,
			)
		},
	)
}

func main() {
	//fmt.Println("part 1:", p1())
	fmt.Println("part 2:", p2())
}

func p2() int {
	mem := make(map[MemKey]Mem)
	walk2(&mem)

	s := math.MaxInt
	end := [2]int{len(m) - 1, len(m[0]) - 1}
	//var h [][2]int
	for i := 4; i < 10; i++ {
		memK := MemKey{
			p:   end,
			d:   [2]int{1, 0},
			dCt: i,
		}
		if sT, ok := mem[memK]; ok {
			if sT.s < s {
				s = sT.s
				//h = sT.h
			}
		}
		memK.d = [2]int{0, 1}
		if sT, ok := mem[memK]; ok {
			if sT.s < s {
				s = sT.s
				//h = sT.h
			}
		}
	}

	//_ = h
	//fmt.Println(h)
	//x := 0
	//for i := range h {
	//	fmt.Print(m[h[i][0]][h[i][1]], " ")
	//	x += m[h[i][0]][h[i][1]]
	//}
	//fmt.Println()
	//fmt.Println(x)
	//
	//for l := range m {
	//	for c := range m[l] {
	//		if slices.Contains(h, [2]int{l, c}) {
	//			fmt.Print(".")
	//		} else {
	//			fmt.Print(m[l][c])
	//		}
	//	}
	//	fmt.Println()
	//}

	return s + m[end[0]][end[1]] - m[0][0]
}

func p1() int {
	mem := make(map[MemKey]Mem)
	walk1(&mem)

	s := math.MaxInt
	end := [2]int{len(m) - 1, len(m[0]) - 1}
	//var h [][2]int
	for i := 0; i < 3; i++ {
		memK := MemKey{
			p:   end,
			d:   [2]int{1, 0},
			dCt: i,
		}
		if sT, ok := mem[memK]; ok {
			if sT.s < s {
				s = sT.s
				//h = sT.h
			}
		}
		memK.d = [2]int{0, 1}
		if sT, ok := mem[memK]; ok {
			if sT.s < s {
				s = sT.s
				//h = sT.h
			}
		}
	}

	//_ = h
	//fmt.Println(h)
	//x := 0
	//for i := range h {
	//	fmt.Print(m[h[i][0]][h[i][1]], " ")
	//	x += m[h[i][0]][h[i][1]]
	//}
	//fmt.Println()
	//fmt.Println(x)
	//
	//for l := range m {
	//	for c := range m[l] {
	//		if slices.Contains(h, [2]int{l, c}) {
	//			fmt.Print(".")
	//		} else {
	//			fmt.Print(m[l][c])
	//		}
	//	}
	//	fmt.Println()
	//}

	return s + m[end[0]][end[1]] - m[0][0]
}

type MemKey struct {
	p, d [2]int
	dCt  int
}

type Mem struct {
	s int
	//h [][2]int
}

type Step struct {
	p, d [2]int
	dCt  int
	s    int
	//h    [][2]int
}

func walk1(mem *map[MemKey]Mem) {
	var q _a.Queue[Step]
	q.Enqueue(Step{
		p:   [2]int{0, 0},
		d:   [2]int{0, 0},
		dCt: 0,
		s:   0,
		//h:   [][2]int{{0, 0}},
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
			//h: s.h,
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
				//h:   append(_slice.Duplicate(s.h), [2]int{s.p[0] + 1, s.p[1]}),
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
				//h:   append(_slice.Duplicate(s.h), [2]int{s.p[0] - 1, s.p[1]}),
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
				//h:   append(_slice.Duplicate(s.h), [2]int{s.p[0], s.p[1] + 1}),
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
				//h:   append(_slice.Duplicate(s.h), [2]int{s.p[0], s.p[1] - 1}),
			})
		}
	}
}

func walk2(mem *map[MemKey]Mem) {
	var q _a.Queue[Step]
	q.Enqueue(Step{
		p:   [2]int{0, 0},
		d:   [2]int{0, 0},
		dCt: 0,
		s:   0,
		//h:   [][2]int{{0, 0}},
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
			//h: s.h,
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
					//h:   append(_slice.Duplicate(s.h), next),
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
				//h:   append(_slice.Duplicate(s.h), [2]int{s.p[0] + d[0], s.p[1] + d[1]}, [2]int{s.p[0] + 2*d[0], s.p[1] + 2*d[1]}, [2]int{s.p[0] + 3*d[0], s.p[1] + 3*d[1]}, next),
			})
		}
	}
}
