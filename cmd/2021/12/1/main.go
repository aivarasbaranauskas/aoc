package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/_string"
	"slices"
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

	var routes []string
	findRoutes(nodes, &routes, []string{"start"}, "start")

	fmt.Println(len(routes))
}

func findRoutes(nodes map[string][]string, routes *[]string, path []string, current string) {
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
		findRoutes(nodes, routes, append(_slice.Duplicate(path), next), next)
	}
}
