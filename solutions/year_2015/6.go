package year_2015

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[6] = Day6{}
}

type Day6 struct{}

func (day Day6) Part1(input []byte) string {
	var m [1000][1000]bool

	for line := range bytes.Lines(input) {
		turnOn, toggle, x1, y1, x2, y2 := day.parseLine(line)
		if toggle {
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					m[x][y] = !m[x][y]
				}
			}
		} else {
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					m[x][y] = turnOn
				}
			}
		}
	}

	ct := 0
	for x := range 1000 {
		for y := range 1000 {
			if m[x][y] {
				ct++
			}
		}
	}

	return strconv.Itoa(ct)
}

func (day Day6) parseLine(line []byte) (turnOn, toggle bool, x1, y1, x2, y2 int) {
	if line[1] == 'u' {
		// turn
		if line[6] == 'n' {
			// on
			turnOn = true
			line = line[8:]
		} else {
			//off
			line = line[9:]
		}
	} else {
		toggle = true
		line = line[7:]
	}

	i := bytes.IndexByte(line, ',')
	x1 = optimistic.AtoiBFast(line[:i])
	line = line[i+1:]

	i = bytes.IndexByte(line, ' ')
	y1 = optimistic.AtoiBFast(line[:i])
	line = line[i+9:]

	i = bytes.IndexByte(line, ',')
	x2 = optimistic.AtoiBFast(line[:i])
	line = line[i+1:]

	y2 = optimistic.AtoiBFast(bytes.TrimSpace(line))

	return
}

func (day Day6) Part2(input []byte) string {
	var m [1000][1000]int

	for line := range bytes.Lines(input) {
		turnOn, toggle, x1, y1, x2, y2 := day.parseLine(line)
		diff := -1
		if toggle {
			diff = 2
		} else if turnOn {
			diff = 1
		}

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				m[x][y] += diff
				if m[x][y] < 0 {
					m[x][y] = 0
				}
			}
		}
	}

	brightness := 0
	for x := range 1000 {
		for y := range 1000 {
			brightness += m[x][y]
		}
	}

	return strconv.Itoa(brightness)
}
