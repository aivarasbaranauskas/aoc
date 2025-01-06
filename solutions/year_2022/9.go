package year_2022

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[9] = Day9{}
}

type Day9 struct{}

func (Day9) Part1(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
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

	return strconv.Itoa(len(visited))
}

func (Day9) Part2(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
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
				diff := [2]int{
					rope[j-1][0] - rope[j][0],
					rope[j-1][1] - rope[j][1],
				}

				if _num.Abs(diff[0]) > 1 || _num.Abs(diff[1]) > 1 {
					rope[j][0] += _num.Sign(diff[0])
					rope[j][1] += _num.Sign(diff[1])
				}
			}
			visited[rope[9]] = true
		}
	}

	return strconv.Itoa(len(visited))
}
