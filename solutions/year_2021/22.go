package year_2021

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[22] = Day22{}
}

type Day22 struct{}

func (Day22) Part1(input []byte) string {
	m := make([][][]bool, 101)
	for i := range m {
		m[i] = make([][]bool, 101)
		for j := range m[i] {
			m[i][j] = make([]bool, 101)
		}
	}

	r := bufio.NewScanner(bytes.NewReader(input))
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " ")
		b := spl[0] == "on"
		spl = strings.Split(spl[1], ",")

		splx := strings.Split(spl[0][2:], "..")
		xFrom := max(-50, optimistic.Atoi(splx[0]))
		xTo := min(50, optimistic.Atoi(splx[1]))

		sply := strings.Split(spl[1][2:], "..")
		yFrom := max(-50, optimistic.Atoi(sply[0]))
		yTo := min(50, optimistic.Atoi(sply[1]))

		splz := strings.Split(spl[2][2:], "..")
		zFrom := max(-50, optimistic.Atoi(splz[0]))
		zTo := min(50, optimistic.Atoi(splz[1]))

		for x := xFrom; x <= xTo; x++ {
			for y := yFrom; y <= yTo; y++ {
				for z := zFrom; z <= zTo; z++ {
					m[x+50][y+50][z+50] = b
				}
			}
		}
	}

	var ct int
	for _, a := range m {
		for _, b := range a {
			for _, c := range b {
				if c {
					ct++
				}
			}
		}
	}
	return strconv.Itoa(ct)
}

func (Day22) Part2(input []byte) string {
	var cubes []*Day22Cube

	r := bufio.NewScanner(bytes.NewReader(input))
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " ")
		cube := &Day22Cube{
			mode: spl[0] == "on",
		}
		spl = strings.Split(spl[1], ",")

		splx := strings.Split(spl[0][2:], "..")
		cube.x1 = optimistic.Atoi(splx[0])
		cube.x2 = optimistic.Atoi(splx[1])

		sply := strings.Split(spl[1][2:], "..")
		cube.y1 = optimistic.Atoi(sply[0])
		cube.y2 = optimistic.Atoi(sply[1])

		splz := strings.Split(spl[2][2:], "..")
		cube.z1 = optimistic.Atoi(splz[0])
		cube.z2 = optimistic.Atoi(splz[1])

		var toAdd []*Day22Cube

		if cube.mode {
			toAdd = append(toAdd, cube)
		}

		for _, c := range cubes {
			if inter := intersection(cube, c); inter != nil {
				toAdd = append(toAdd, inter)
			}
		}

		cubes = append(cubes, toAdd...)
	}

	var ct int
	for _, cube := range cubes {
		if cube.mode {
			ct += cube.Size()
		} else {
			ct -= cube.Size()
		}
	}
	return strconv.Itoa(ct)
}

type Day22Cube struct {
	x1, x2, y1, y2, z1, z2 int
	mode                   bool
}

func (c *Day22Cube) Size() int {
	return (c.x2 - c.x1 + 1) * (c.y2 - c.y1 + 1) * (c.z2 - c.z1 + 1)
}

func intersection(s, t *Day22Cube) *Day22Cube {
	c := &Day22Cube{
		x1:   max(s.x1, t.x1),
		x2:   min(s.x2, t.x2),
		y1:   max(s.y1, t.y1),
		y2:   min(s.y2, t.y2),
		z1:   max(s.z1, t.z1),
		z2:   min(s.z2, t.z2),
		mode: !t.mode,
	}

	if c.x1 > c.x2 || c.y1 > c.y2 || c.z1 > c.z2 {
		return nil
	}
	return c
}
