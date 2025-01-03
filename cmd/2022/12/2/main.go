package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	var m [][]byte
	var s, e [2]int

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Text()
		m = append(m, []byte(line))
		if si := strings.Index(line, "S"); si >= 0 {
			s[0], s[1] = len(m)-1, si
			m[s[0]][s[1]] = 'a'
		}
		if si := strings.Index(line, "E"); si >= 0 {
			e[0], e[1] = len(m)-1, si
			m[e[0]][e[1]] = 'z'
		}
	}

	queue := _a.Queue[[2]int]{}
	queue.Enqueue(e)

	cameFrom := map[[2]int][2]int{}
	fScore := map[[2]int]int{e: 0}

	for !queue.Empty() {
		//fmt.Println(queue.Len())
		current := queue.Dequeue()
		cC := m[current[0]][current[1]]
		neighbors := _slice.Filter(
			[][2]int{
				{current[0] - 1, current[1]},
				{current[0] + 1, current[1]},
				{current[0], current[1] - 1},
				{current[0], current[1] + 1},
			},
			func(v [2]int) bool {
				return v[0] >= 0 && v[0] < len(m) && v[1] >= 0 && v[1] < len(m[0]) && int(cC)-int(m[v[0]][v[1]]) <= 1
			},
		)

		for _, neighbor := range neighbors {
			tentativeFScore := fScore[current] + 1
			if neighborFScore, ok := fScore[neighbor]; !ok || (ok && tentativeFScore < neighborFScore) {
				// This path to neighbor is better than any previous one. Record it!
				cameFrom[neighbor] = current
				fScore[neighbor] = tentativeFScore
				queue.Enqueue(neighbor)
			}
		}
	}

	minVal := 100000
	for i := range m {
		for j := range m[i] {
			if m[i][j] == 'a' {
				if sc, ok := fScore[[2]int{i, j}]; ok {
					minVal = min(minVal, sc)
				}
			}
		}
	}

	fmt.Println(fScore[e], minVal)
}
