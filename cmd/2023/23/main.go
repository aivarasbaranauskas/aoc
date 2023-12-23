package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

type Point [2]int

func (p Point) walk(d Point) Point {
	return Point{
		p[0] + d[0],
		p[1] + d[1],
	}
}

var Up = Point{-1, 0}
var Down = Point{1, 0}
var Right = Point{0, 1}
var Left = Point{0, -1}

var directions = []Point{Up, Down, Right, Left}

func main() {
	mRaw := bytes.Split(input, []byte("\n"))
	m := make(map[Point]byte, len(mRaw)*len(mRaw[0]))
	for li, line := range mRaw {
		for ci, c := range line {
			m[Point{li, ci}] = c
		}
	}
	start := Point{0, 1}
	end := Point{len(mRaw) - 1, len(mRaw[0]) - 2}

	p1 := longestWalk(m, map[Point]struct{}{}, start, end)
	fmt.Println("part 1:", p1)

	p2 := longestWalk2(m, start, end)
	fmt.Println("part 2:", p2)
}

func longestWalk(m map[Point]byte, mem map[Point]struct{}, p, end Point) int {
	if p == end {
		return len(mem)
	}

	c, ok := m[p]
	if !ok {
		return 0
	}

	if c == '.' {
		var n int

		for _, d := range directions {
			pNext := p.walk(d)
			if _, ok := mem[pNext]; ok {
				continue
			}

			mem[pNext] = struct{}{}
			n = max(n, longestWalk(m, mem, pNext, end))
			delete(mem, pNext)
		}

		return n
	}

	var pNext Point

	switch m[p] {
	case '>':
		pNext = p.walk(Right)
	case 'v':
		pNext = p.walk(Down)
	case '<':
		pNext = p.walk(Left)
	case '^':
		pNext = p.walk(Up)
	}

	if _, ok := mem[pNext]; ok {
		return 0
	}

	mem[pNext] = struct{}{}
	n := longestWalk(m, mem, pNext, end)
	delete(mem, pNext)
	return n
}

func longestWalk2(m map[Point]byte, start, end Point) int {
	// transform into weighted graph

	// [from][to]length
	g := map[Point]map[Point]int{}

	populateG(&g, m, start, start)

	return longestWalk2G(g, map[Point]struct{}{}, start, end, 0)
}

func longestWalk2G(g map[Point]map[Point]int, visited map[Point]struct{}, p, end Point, n int) int {
	if p == end {
		return n
	}

	if _, ok := visited[p]; ok {
		return 0
	}
	visited[p] = struct{}{}

	var nMax int
	for pNext, nNext := range g[p] {
		nMax = max(nMax, longestWalk2G(g, visited, pNext, end, n+nNext))
	}

	delete(visited, p)

	return nMax
}

func populateG(g *map[Point]map[Point]int, m map[Point]byte, start, p Point) {
	pPrev := start
	n := 0
	if start == p {
		// initial start
		n--
	}

	pNexts := make([]Point, 0, 4)
	for {
		for _, d := range directions {
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
		populateG(g, m, p, pNext)
	}
}
