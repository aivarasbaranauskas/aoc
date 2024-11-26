package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/go_helpers/o"
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

	var cubes []*Cube

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Text()
		fmt.Println(line)
		spl := strings.Split(line, " ")
		cube := &Cube{
			mode: spl[0] == "on",
		}
		spl = strings.Split(spl[1], ",")

		splx := strings.Split(spl[0][2:], "..")
		cube.x1 = o.Atoi(splx[0])
		cube.x2 = o.Atoi(splx[1])

		sply := strings.Split(spl[1][2:], "..")
		cube.y1 = o.Atoi(sply[0])
		cube.y2 = o.Atoi(sply[1])

		splz := strings.Split(spl[2][2:], "..")
		cube.z1 = o.Atoi(splz[0])
		cube.z2 = o.Atoi(splz[1])

		var toAdd []*Cube

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
	fmt.Println(ct, len(cubes))
}

type Cube struct {
	x1, x2, y1, y2, z1, z2 int
	mode                   bool
}

func (c *Cube) Size() int {
	return (c.x2 - c.x1 + 1) * (c.y2 - c.y1 + 1) * (c.z2 - c.z1 + 1)
}

func intersection(s, t *Cube) *Cube {
	c := &Cube{
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
