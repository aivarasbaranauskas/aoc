package year_2021

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_map"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/_string"
	"slices"
	"strconv"
	"strings"
)

func init() {
	Solutions[12] = Day12{}
}

type Day12 struct{}

func (d Day12) Part1(input []byte) string {
	nodes := make(map[string][]string)
	r := bufio.NewScanner(bytes.NewReader(input))
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, "-")

		if _, ok := nodes[spl[0]]; ok {
			nodes[spl[0]] = append(nodes[spl[0]], spl[1])
		} else {
			nodes[spl[0]] = []string{spl[1]}
		}

		if _, ok := nodes[spl[1]]; ok {
			nodes[spl[1]] = append(nodes[spl[1]], spl[0])
		} else {
			nodes[spl[1]] = []string{spl[0]}
		}
	}

	var routes []string
	d.findRoutes1(nodes, &routes, []string{"start"}, "start")

	return strconv.Itoa(len(routes))
}

func (d Day12) findRoutes1(nodes map[string][]string, routes *[]string, path []string, current string) {
	if current == "end" {
		*routes = append(*routes, strings.Join(path, ","))
		return
	}

	for _, next := range nodes[current] {
		if next == "start" {
			continue
		}
		if _string.IsLower(next) && slices.Contains(path, next) {
			continue
		}
		d.findRoutes1(nodes, routes, append(_slice.Duplicate(path), next), next)
	}
}

func (d Day12) Part2(input []byte) string {
	nodes := make(map[string][]string)
	r := bufio.NewScanner(bytes.NewReader(input))
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, "-")

		if _, ok := nodes[spl[0]]; ok {
			nodes[spl[0]] = append(nodes[spl[0]], spl[1])
		} else {
			nodes[spl[0]] = []string{spl[1]}
		}

		if _, ok := nodes[spl[1]]; ok {
			nodes[spl[1]] = append(nodes[spl[1]], spl[0])
		} else {
			nodes[spl[1]] = []string{spl[0]}
		}
	}

	routesCt := 0
	p := Day12Path{pathItems: map[string]int{}}
	p.add("start")
	d.findRoutes2(nodes, &routesCt, &p, "start")

	return strconv.Itoa(routesCt)
}

func (d Day12) findRoutes2(nodes map[string][]string, routesCt *int, path *Day12Path, current string) {
	if current == "end" {
		*routesCt = *routesCt + 1
		return
	}

	for _, next := range nodes[current] {
		if next == "start" {
			continue
		}

		if path.canGo(next) {
			nextPath := path.clone()
			nextPath.add(next)
			d.findRoutes2(nodes, routesCt, nextPath, next)
		}
	}
}

type Day12Path struct {
	pathItems            map[string]int
	hasEnteredSmallTwice bool
}

func (p *Day12Path) canGo(next string) bool {
	if !_string.IsLower(next) {
		return true
	}

	_, wasThere := p.pathItems[next]
	if wasThere && p.hasEnteredSmallTwice {
		return false
	}
	return true
}

func (p *Day12Path) add(next string) {
	if _, ok := p.pathItems[next]; ok {
		p.pathItems[next]++
		if _string.IsLower(next) {
			p.hasEnteredSmallTwice = true
		}
	} else {
		p.pathItems[next] = 1
	}
}

func (p *Day12Path) clone() *Day12Path {
	return &Day12Path{
		pathItems:            _map.Duplicate(p.pathItems),
		hasEnteredSmallTwice: p.hasEnteredSmallTwice,
	}
}
