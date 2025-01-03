package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_map"
	"github.com/aivarasbaranauskas/aoc/internal/_string"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	nodes := make(map[string][]string)
	r := bufio.NewScanner(f)
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
	p := Path{pathItems: map[string]int{}}
	p.add("start")
	findRoutes(nodes, &routesCt, &p, "start")

	fmt.Println(routesCt)
}

func findRoutes(nodes map[string][]string, routesCt *int, path *Path, current string) {
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
			findRoutes(nodes, routesCt, nextPath, next)
		}
	}
}

type Path struct {
	pathItems            map[string]int
	hasEnteredSmallTwice bool
}

func (p *Path) canGo(next string) bool {
	if !_string.IsLower(next) {
		return true
	}

	_, wasThere := p.pathItems[next]
	if wasThere && p.hasEnteredSmallTwice {
		return false
	}
	return true
}

func (p *Path) add(next string) {
	if _, ok := p.pathItems[next]; ok {
		p.pathItems[next]++
		if _string.IsLower(next) {
			p.hasEnteredSmallTwice = true
		}
	} else {
		p.pathItems[next] = 1
	}
}

func (p *Path) clone() *Path {
	return &Path{
		pathItems:            _map.Duplicate(p.pathItems),
		hasEnteredSmallTwice: p.hasEnteredSmallTwice,
	}
}
