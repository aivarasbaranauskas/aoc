package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/fatih/color"
	"log"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

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
	queue.Enqueue(s)

	cameFrom := map[[2]int][2]int{}
	fScore := map[[2]int]int{s: 0}

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
				return v[0] >= 0 && v[0] < len(m) && v[1] >= 0 && v[1] < len(m[0]) && int(m[v[0]][v[1]])-int(cC) <= 1
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

	//fmt.Println(cameFrom)
	var totalPath [][2]int
	var items []byte
	current, ok := e, true
	for ok {
		totalPath = append(totalPath, current)
		items = append(items, m[current[0]][current[1]])
		current, ok = cameFrom[current]
	}

	fmt.Println(len(items) - 1)
	fmt.Println(string(items))
	fmt.Println()
	draw(totalPath, m)
}

func draw(path [][2]int, m [][]byte) {
	whilte := color.New(color.FgWhite)
	boldWhite := whilte.Add(color.BgRed)
	p := _set.FromSlice(path)
	for i := range m {
		for j := range m[i] {
			if p.Has([2]int{i, j}) {
				_, _ = boldWhite.Print(string(m[i][j]))
			} else {
				fmt.Print(string(m[i][j]))
			}
		}
		fmt.Println()
	}
}
