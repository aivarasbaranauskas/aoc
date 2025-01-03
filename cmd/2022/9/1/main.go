package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	r := bufio.NewScanner(f)
	head, tail := [2]int{0, 0}, [2]int{0, 0}
	visited := map[[2]int]bool{tail: true}
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " ")
		c := optimistic.Atoi(spl[1])

		x, y := 0, 0
		switch spl[0] {
		case "R":
			x = -1
		case "L":
			x = 1
		case "D":
			y = -1
		case "U":
			y = 1
		}

		for i := 0; i < c; i++ {
			head = [2]int{head[0] + x, head[1] + y}
			if _num.Abs(head[0]-tail[0]) > 1 || _num.Abs(head[1]-tail[1]) > 1 {
				tail = [2]int{head[0] - x, head[1] - y}
				visited[tail] = true
			}
		}
	}

	fmt.Println(len(visited))
}
