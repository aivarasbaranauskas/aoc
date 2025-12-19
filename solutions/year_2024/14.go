package year_2024

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[14] = Day14{
		limX: 101,
		limY: 103,
	}
}

type Day14 struct {
	limX, limY int
	m          [][]bool
}

func (day Day14) Part1(input []byte) string {
	robots := day.parseInput(input)
	day.simulate(robots, 100)
	qx, qy := (day.limX-1)/2, (day.limY-1)/2
	var qs [2][2]int

	for i := range robots {
		if robots[i][0] == qx || robots[i][1] == qy {
			continue
		}
		qs[_num.FastBoolToInt(robots[i][0] < qx)][_num.FastBoolToInt(robots[i][1] < qy)]++
	}

	return strconv.Itoa(qs[0][0] * qs[0][1] * qs[1][0] * qs[1][1])
}

func (day Day14) simulate(robots [][4]int, t int) {
	for i := range robots {
		robots[i][0] = (((robots[i][0] + t*robots[i][2]) % day.limX) + day.limX) % day.limX
		robots[i][1] = (((robots[i][1] + t*robots[i][3]) % day.limY) + day.limY) % day.limY
	}
}

func (day Day14) parseInput(input []byte) [][4]int {
	ct := bytes.Count(input, []byte{'\n'}) + 1
	robots := make([][4]int, ct)

	for i := range robots {
		input = input[2:]
		idx := bytes.IndexByte(input, ',')
		px := optimistic.AtoiB(input[:idx])
		input = input[idx+1:]

		idx = bytes.IndexByte(input, ' ')
		py := optimistic.AtoiB(input[:idx])
		input = input[idx+3:]

		idx = bytes.IndexByte(input, ',')
		vx := optimistic.AtoiB(input[:idx])
		input = input[idx+1:]

		idx = bytes.IndexByte(input, '\n')
		var vy int
		if idx == -1 {
			vy = optimistic.AtoiB(input)
		} else {
			vy = optimistic.AtoiB(input[:idx])
			input = input[idx+1:]
		}

		robots[i] = [4]int{px, py, vx, vy}
	}

	return robots
}

func (day Day14) Part2(input []byte) string {
	robots := day.parseInput(input)
	n := 1
	tmp := make([][4]int, len(robots))
	reader := bufio.NewReader(os.Stdin)

Loop:
	for {
		copy(tmp, robots)
		day.simulate(tmp, n)
		day.draw(tmp)

		fmt.Print("good? ->")
		text, _ := reader.ReadString('\n')
		switch text {
		case "y":
			break Loop
		case "b":
			n--
		default:
			n++
		}
	}

	return strconv.Itoa(n)
}

func (day Day14) draw(robots [][4]int) {
	if day.m == nil {
		day.m = make([][]bool, day.limY)
		for i := range day.m {
			day.m[i] = make([]bool, day.limX)
		}
	}

	for i := range day.m {
		for j := range day.m[i] {
			day.m[i][j] = false
		}
	}

	for i := range robots {
		day.m[robots[i][1]][robots[i][0]] = true
	}

	for i := range day.m {
		for j := range day.m[i] {
			if day.m[i][j] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
