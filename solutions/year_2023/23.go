package year_2023

import (
	"bytes"
	"strconv"
)

func init() {
	Solutions[23] = Day23{}
}

type Day23 struct{}

func (day Day23) Part1(input []byte) string {
	mRaw := bytes.Split(input, []byte("\n"))
	m := make(map[Point]byte, len(mRaw)*len(mRaw[0]))
	for li, line := range mRaw {
		for ci, c := range line {
			m[Point{li, ci}] = c
		}
	}
	start := Point{0, 1}
	end := Point{len(mRaw) - 1, len(mRaw[0]) - 2}

	return strconv.Itoa(day.longestWalk(m, map[Point]struct{}{}, start, end))
}

func (day Day23) Part2(input []byte) string {
	mRaw := bytes.Split(input, []byte("\n"))
	m := make(map[Point]byte, len(mRaw)*len(mRaw[0]))
	for li, line := range mRaw {
		for ci, c := range line {
			m[Point{li, ci}] = c
		}
	}
	start := Point{0, 1}
	end := Point{len(mRaw) - 1, len(mRaw[0]) - 2}

	return strconv.Itoa(day.longestWalk2(m, start, end))
}

func (day Day23) longestWalk(m map[Point]byte, mem map[Point]struct{}, p, end Point) int {
	if p == end {
		return len(mem)
	}

	c, ok := m[p]
	if !ok {
		return 0
	}

	if c == '.' {
		var n int

		for _, d := range []Point{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
			pNext := p.walk(d)
			if _, ok := mem[pNext]; ok {
				continue
			}

			mem[pNext] = struct{}{}
			n = max(n, day.longestWalk(m, mem, pNext, end))
			delete(mem, pNext)
		}

		return n
	}

	var pNext Point

	switch m[p] {
	case '>':
		pNext = p.walk(Point{0, 1})
	case 'v':
		pNext = p.walk(Point{1, 0})
	case '<':
		pNext = p.walk(Point{0, -1})
	case '^':
		pNext = p.walk(Point{-1, 0})
	}

	if _, ok := mem[pNext]; ok {
		return 0
	}

	mem[pNext] = struct{}{}
	n := day.longestWalk(m, mem, pNext, end)
	delete(mem, pNext)
	return n
}

func (day Day23) longestWalk2(m map[Point]byte, start, end Point) int {
	// transform into weighted graph

	// [from][to]length
	g := map[Point]map[Point]int{}

	day.populateG(&g, m, start, start)

	return day.longestWalk2G(g, map[Point]struct{}{}, start, end, 0)
}

func (day Day23) longestWalk2G(g map[Point]map[Point]int, visited map[Point]struct{}, p, end Point, n int) int {
	if p == end {
		return n
	}

	if _, ok := visited[p]; ok {
		return 0
	}
	visited[p] = struct{}{}

	var nMax int
	for pNext, nNext := range g[p] {
		nMax = max(nMax, day.longestWalk2G(g, visited, pNext, end, n+nNext))
	}

	delete(visited, p)

	return nMax
}

func (day Day23) populateG(g *map[Point]map[Point]int, m map[Point]byte, start, p Point) {
	pPrev := start
	n := 0
	if start == p {
		// initial start
		n--
	}

	pNexts := make([]Point, 0, 4)
	for {
		for _, d := range []Point{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
			pNext := p.walk(d)
			if _, ok := m[pNext]; !ok || m[pNext] == '#' || pNext == pPrev {
				continue
			}
			pNexts = append(pNexts, pNext)
		}
		n++
		if len(pNexts) != 1 {
			break
		}
		pPrev = p
		p = pNexts[0]
		pNexts = pNexts[:0]
	}

	if _, ok := (*g)[start]; !ok {
		(*g)[start] = make(map[Point]int)
	}
	if _, ok := (*g)[p]; !ok {
		(*g)[p] = make(map[Point]int)
	}
	if _, ok := (*g)[start][p]; ok {
		//must have been visited from other end
		return
	}
	(*g)[start][p] = n
	(*g)[p][start] = n
	for _, pNext := range pNexts {
		day.populateG(g, m, p, pNext)
	}
}

func (p Point) walk(d Point) Point {
	return Point{
		p[0] + d[0],
		p[1] + d[1],
	}
}
