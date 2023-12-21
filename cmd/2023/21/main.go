package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
)

//go:embed input.txt
var input []byte

type Point [2]int

var directions = []Point{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func main() {
	part1()
	part2()
}

func part2() {
	m := bytes.Split(input, []byte("\n"))

	var s Point

	for li, line := range m {
		if ci := bytes.IndexByte(line, 'S'); ci != -1 {
			s = Point{li, ci}
			break
		}
	}

	doCount := func(n int) int {
		mem := map[Point]int{s: 0}
		q := _a.Queue[Point]{}
		q.Enqueue(s)

		for !q.Empty() {
			ps := q.Dequeue()

			if mem[ps]+1 > n {
				continue
			}

			for _, d := range directions {
				p := Point{
					ps[0] + d[0],
					ps[1] + d[1],
				}

				pAdj := Point{
					((p[0] % len(m)) + len(m)) % len(m),
					((p[1] % len(m[0])) + len(m[0])) % len(m[0]),
				}

				if m[pAdj[0]][pAdj[1]] == '#' {
					continue
				}
				if _, ok := mem[p]; ok {
					continue
				}

				mem[p] = mem[ps] + 1
				q.Enqueue(p)
			}
		}

		ct := 0
		for _, memC := range mem {
			if memC%2 == n%2 {
				ct++
			}
		}
		return ct
	}

	fmt.Println("part 2:")
	fmt.Println("x=0:", doCount(65))
	fmt.Println("x=1:", doCount(196))
	fmt.Println("x=2:", doCount(327))
	fmt.Println("x:", (26501365-65)/len(m))
	fmt.Println("And do Newton Interpolating Polynomial")
}

func part1() {
	m := bytes.Split(input, []byte("\n"))

	mem := make(map[Point]int)

	q := _a.Queue[Point]{}
	for li, line := range m {
		if ci := bytes.IndexByte(line, 'S'); ci != -1 {
			p := Point{li, ci}
			q.Enqueue(p)
			mem[p] = 0
			break
		}
	}

	n := 64

	for !q.Empty() {
		ps := q.Dequeue()

		if mem[ps]+1 > n {
			continue
		}

		for _, d := range directions {
			p := Point{
				ps[0] + d[0],
				ps[1] + d[1],
			}

			if p[0] < 0 || p[0] >= len(m) || p[1] < 0 || p[1] >= len(m[0]) {
				continue
			}
			if m[p[0]][p[1]] == '#' {
				continue
			}
			if _, ok := mem[p]; ok {
				continue
			}

			mem[p] = mem[ps] + 1
			q.Enqueue(p)
		}
	}

	p1 := 0
	for _, memC := range mem {
		if memC%2 == n%2 {
			p1++
		}
	}

	fmt.Println("part 1:", p1)
}
