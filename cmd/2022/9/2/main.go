package main

import (
	"bufio"
	"embed"
	"fmt"
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
	rope := [10][2]int{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}
	visited := map[[2]int]bool{[2]int{0, 0}: true}
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " ")
		c := optimistic.Atoi(spl[1])

		x, y := 0, 0
		switch spl[0] {
		case "R":
			x = 1
		case "L":
			x = -1
		case "D":
			y = -1
		case "U":
			y = 1
		}

		for i := 0; i < c; i++ {
			rope[0] = [2]int{rope[0][0] + x, rope[0][1] + y}
			for j := 1; j < 10; j++ {
				diff := [2]int{rope[j-1][0] - rope[j][0], rope[j-1][1] - rope[j][1]}
				switch diff {
				case [2]int{2, 0}:
					rope[j][0]++
				case [2]int{-2, 0}:
					rope[j][0]--
				case [2]int{0, 2}:
					rope[j][1]++
				case [2]int{0, -2}:
					rope[j][1]--

				case [2]int{2, 2}:
					rope[j][0]++
					rope[j][1]++
				case [2]int{-2, 2}:
					rope[j][0]--
					rope[j][1]++
				case [2]int{-2, -2}:
					rope[j][0]--
					rope[j][1]--
				case [2]int{2, -2}:
					rope[j][0]++
					rope[j][1]--

				case [2]int{2, 1}:
					rope[j][0]++
					rope[j][1]++
				case [2]int{2, -1}:
					rope[j][0]++
					rope[j][1]--
				case [2]int{-2, 1}:
					rope[j][0]--
					rope[j][1]++
				case [2]int{-2, -1}:
					rope[j][0]--
					rope[j][1]--

				case [2]int{1, 2}:
					rope[j][0]++
					rope[j][1]++
				case [2]int{-1, 2}:
					rope[j][0]--
					rope[j][1]++
				case [2]int{1, -2}:
					rope[j][0]++
					rope[j][1]--
				case [2]int{-1, -2}:
					rope[j][0]--
					rope[j][1]--
				}
			}
			visited[rope[9]] = true
			fmt.Println(rope)
		}
		fmt.Println()
	}

	fmt.Println(len(visited))
}
