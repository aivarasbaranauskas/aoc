package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
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

	r := bufio.NewScanner(f)
	var (
		trees [][]int
	)
	for r.Scan() {
		trees = append(trees, _slice.Map(strings.Split(r.Text(), ""), optimistic.Atoi))
	}

	l := len(trees)
	w := len(trees[0])
	var vCt int

	for x := 0; x < l; x++ {
		for y := 0; y < w; y++ {
			tree := trees[x][y]

			// down
			visible := true
			for i := x + 1; i < l; i++ {
				if tree <= trees[i][y] {
					visible = false
					break
				}
			}
			if visible {
				vCt++
				continue
			}

			// up
			visible = true
			for i := x - 1; i >= 0; i-- {
				if tree <= trees[i][y] {
					visible = false
					break
				}
			}
			if visible {
				vCt++
				continue
			}

			// right
			visible = true
			for i := y + 1; i < w; i++ {
				if tree <= trees[x][i] {
					visible = false
					break
				}
			}
			if visible {
				vCt++
				continue
			}

			// left
			visible = true
			for i := y - 1; i >= 0; i-- {
				if tree <= trees[x][i] {
					visible = false
					break
				}
			}
			if visible {
				vCt++
				continue
			}
		}
	}

	fmt.Println(vCt)
}
