package year_2021

import (
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[17] = Day17{}
}

type Day17 struct{}

func (d Day17) Part1(input []byte) string {
	spl := strings.Split(strings.Split(string(input), ": ")[1], ", ")
	splX := strings.Split(spl[0][2:], "..")
	splY := strings.Split(spl[1][2:], "..")

	targetXFrom, targetXTo := optimistic.Atoi(splX[0]), optimistic.Atoi(splX[1])
	targetYFrom, targetYTo := optimistic.Atoi(splY[0]), optimistic.Atoi(splY[1])

	/**
	given t
	x = max(0, (t * (2 * velocityX - t + 1)) / 2)
	y = (t * (2 * velocityY - t + 1)) / 2
	*/

	//find all vY
	maxMaxY := _num.Abs(min(targetYFrom, targetYTo))
	var vYs []int
	var vYTs [][]int
	for vY := 0; vY <= maxMaxY; vY++ {
		y := 0
		v := vY
		var ts []int
		for t := 0; ; t++ {
			y += v
			v--
			if targetYFrom <= y && y <= targetYTo {
				ts = append(ts, t)
			}
			if y < targetYFrom {
				break
			}
		}
		if len(ts) > 0 {
			vYs = append(vYs, vY)
			vYTs = append(vYTs, ts)
		}
	}

	//fmt.Println(vYs)

	// Find all vY for each vX and then find a (vX, vY) pair with the highest point
	theHighestVY := -1
YLOOP:
	for i, vY := range vYs {
		ts := vYTs[i]
		h := (vY * (2*vY - vY + 1)) / 2
		if h < theHighestVY {
			continue
		}

		for _, t := range ts {
			for vX := 0; ; vX++ {
				if d.simulateX(vX, t, targetXFrom, targetXTo) {
					theHighestVY = h
					continue YLOOP
				}
			}
		}
	}

	return strconv.Itoa(theHighestVY)
}

func (d Day17) Part2(input []byte) string {
	spl := strings.Split(strings.Split(string(input), ": ")[1], ", ")
	splX := strings.Split(spl[0][2:], "..")
	splY := strings.Split(spl[1][2:], "..")

	targetXFrom, targetXTo := optimistic.Atoi(splX[0]), optimistic.Atoi(splX[1])
	targetYFrom, targetYTo := optimistic.Atoi(splY[0]), optimistic.Atoi(splY[1])

	/**
	given t
	x = max(0, (t * (2 * velocityX - t + 1)) / 2)
	y = (t * (2 * velocityY - t + 1)) / 2
	*/

	//find all vY
	maxMaxY := _num.Abs(min(targetYFrom, targetYTo))
	var vYs []int
	var vYTs [][]int
	for vY := -maxMaxY; vY <= maxMaxY; vY++ {
		y := 0
		v := vY
		var ts []int
		for t := 0; ; t++ {
			y += v
			v--
			if targetYFrom <= y && y <= targetYTo {
				ts = append(ts, t)
			}
			if y < targetYFrom {
				break
			}
		}
		if len(ts) > 0 {
			vYs = append(vYs, vY)
			vYTs = append(vYTs, ts)
		}
	}

	// Find all vY for each vX and then find a (vX, vY) pair with the highest point
	vels := _set.New[[2]int]()
	for i, vY := range vYs {
		ts := vYTs[i]
		for _, t := range ts {
			for vX := 0; vX <= targetXTo; vX++ {
				if d.simulateX(vX, t, targetXFrom, targetXTo) {
					vels.Add([2]int{vX, vY})
				}
			}
		}
	}

	return strconv.Itoa(vels.Len())
}

func (Day17) simulateX(vX, t, targetXFrom, targetXTo int) bool {
	x := 0
	for i := 0; i <= t; i++ {
		x = x + vX
		vX = max(vX-1, 0)
	}
	return targetXFrom <= x && x <= targetXTo
}
