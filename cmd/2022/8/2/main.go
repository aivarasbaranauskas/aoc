package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
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
	var maxScenicScore int

	for x := 1; x < l-1; x++ {
		for y := 1; y < w-1; y++ {
			tree := trees[x][y]

			// down
			down := l - x - 1
			for i := x + 1; i < l; i++ {
				if tree <= trees[i][y] {
					down = i - x
					break
				}
			}

			// up
			up := x
			for i := x - 1; i >= 0; i-- {
				if tree <= trees[i][y] {
					up = x - i
					break
				}
			}

			// right
			right := w - y - 1
			for i := y + 1; i < w; i++ {
				if tree <= trees[x][i] {
					right = i - y
					break
				}
			}

			// left
			left := y
			for i := y - 1; i >= 0; i-- {
				if tree <= trees[x][i] {
					left = y - i
					break
				}
			}
			maxScenicScore = _num.Max(maxScenicScore, down*up*left*right)
			fmt.Println(x, y, maxScenicScore, down, up, left, right)
		}
	}

	fmt.Println(maxScenicScore)
}
