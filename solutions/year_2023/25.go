package year_2023

import (
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"strconv"
	"strings"
)

func init() {
	Solutions[25] = Day25{}
}

type Day25 struct{}

func (day Day25) Part1(input []byte) string {
	g := map[string]map[string]struct{}{}

	for _, line := range strings.Split(string(input), "\n") {
		spl := strings.Split(line, ": ")
		a := spl[0]
		if _, ok := g[a]; !ok {
			g[a] = make(map[string]struct{})
		}
		spl = strings.Split(spl[1], " ")
		for _, b := range spl {
			if _, ok := g[b]; !ok {
				g[b] = make(map[string]struct{})
			}
			g[a][b] = struct{}{}
			g[b][a] = struct{}{}
		}
	}

	var start string
	for k := range g {
		start = k
		break
	}

	paths, end := day.bfs(g, start, nil)

	for i := 0; i < len(paths[end])-1; i++ {
		s := paths[end][i]
		d := paths[end][i+1]
		paths2, _ := day.bfs(g, start, map[[2]string]struct{}{{s, d}: {}, {d, s}: {}})

		for i2 := 0; i2 < len(paths2[end])-1; i2++ {
			s2 := paths2[end][i2]
			d2 := paths2[end][i2+1]
			paths3, _ := day.bfs(g, start, map[[2]string]struct{}{{s, d}: {}, {d, s}: {}, {s2, d2}: {}, {d2, s2}: {}})

			for i3 := 0; i3 < len(paths3[end])-1; i3++ {
				s3 := paths3[end][i3]
				d3 := paths3[end][i3+1]
				exclusions := map[[2]string]struct{}{
					{s, d}:   {},
					{d, s}:   {},
					{s2, d2}: {},
					{d2, s2}: {},
					{s3, d3}: {},
					{d3, s3}: {},
				}

				paths4, _ := day.bfs(g, start, exclusions)
				if len(paths4) != len(g) {
					return strconv.Itoa(len(paths4) * (len(g) - len(paths4)))
				}
			}
		}
	}
	return "???"
}

func (Day25) Part2(_ []byte) string {
	return ""
}

func (Day25) bfs(g map[string]map[string]struct{}, start string, exclusions map[[2]string]struct{}) (paths map[string][]string, end string) {
	paths = map[string][]string{
		start: {start},
	}
	q := _a.Queue[string]{}
	q.Enqueue(start)

	for !q.Empty() {
		cur := q.Dequeue()
		end = cur
		for next := range g[cur] {
			if _, ok := paths[next]; ok {
				continue
			}
			if exclusions != nil {
				if _, ok := exclusions[[2]string{cur, next}]; ok {
					continue
				}
			}
			paths[next] = append([]string{next}, paths[cur]...)
			q.Enqueue(next)
		}
	}

	return
}
