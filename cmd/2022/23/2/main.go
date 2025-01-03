package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"math"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	elves := map[Pos]struct{}{}

	r := bufio.NewScanner(f)
	y := 0
	for r.Scan() {
		for x, c := range []byte(r.Text()) {
			if c == '#' {
				elves[Pos{x, y}] = struct{}{}
			}
		}
		y++
	}

	// 1 - north
	// 2 - south
	// 3 - west
	// 4 - east
	dirs := []int{1, 2, 3, 4}
	moved := true
	ct := 0
	for moved {
		moved = false
		ct++
		// first half - generate propositions
		// [to] -> [from elf 1, from elf 2, ...]
		propositions := map[Pos][]Pos{}
	ElfLoop:
		for elf := range elves {
			shouldMove := false
		L:
			for a := -1; a <= 1; a++ {
				for b := -1; b <= 1; b++ {
					if a == b && b == 0 {
						continue
					}
					if _, ok := elves[Pos{elf.x + a, elf.y + b}]; ok {
						shouldMove = true
						break L
					}
				}
			}

			if !shouldMove {
				continue
			}
			moved = true

			for _, d := range dirs {
				switch d {
				case 1:
					// north
					newP := Pos{elf.x, elf.y - 1}
					_, a := elves[Pos{elf.x - 1, elf.y - 1}]
					_, b := elves[newP]
					_, c := elves[Pos{elf.x + 1, elf.y - 1}]
					if !a && !b && !c {
						if _, ok := propositions[newP]; ok {
							propositions[newP] = append(propositions[newP], elf)
						} else {
							propositions[newP] = []Pos{elf}
						}
						continue ElfLoop
					}
				case 2:
					// south
					newP := Pos{elf.x, elf.y + 1}
					_, a := elves[Pos{elf.x - 1, elf.y + 1}]
					_, b := elves[newP]
					_, c := elves[Pos{elf.x + 1, elf.y + 1}]
					if !a && !b && !c {
						if _, ok := propositions[newP]; ok {
							propositions[newP] = append(propositions[newP], elf)
						} else {
							propositions[newP] = []Pos{elf}
						}
						continue ElfLoop
					}
				case 3:
					// west
					newP := Pos{elf.x - 1, elf.y}
					_, a := elves[Pos{elf.x - 1, elf.y - 1}]
					_, b := elves[newP]
					_, c := elves[Pos{elf.x - 1, elf.y + 1}]
					if !a && !b && !c {
						if _, ok := propositions[newP]; ok {
							propositions[newP] = append(propositions[newP], elf)
						} else {
							propositions[newP] = []Pos{elf}
						}
						continue ElfLoop
					}
				case 4:
					// east
					newP := Pos{elf.x + 1, elf.y}
					_, a := elves[Pos{elf.x + 1, elf.y - 1}]
					_, b := elves[newP]
					_, c := elves[Pos{elf.x + 1, elf.y + 1}]
					if !a && !b && !c {
						if _, ok := propositions[newP]; ok {
							propositions[newP] = append(propositions[newP], elf)
						} else {
							propositions[newP] = []Pos{elf}
						}
						continue ElfLoop
					}
				}
			}
		}

		dirs = []int{dirs[1], dirs[2], dirs[3], dirs[0]}

		// second half - execute propositions
		for dest, proposers := range propositions {
			if len(proposers) == 1 {
				delete(elves, proposers[0])
				elves[dest] = struct{}{}
			}
		}
		//printE(elves)
	}

	fmt.Println(ct)
}

type Pos struct {
	x, y int
}

func printE(elves map[Pos]struct{}) {
	minX, maxX := 0, math.MinInt
	minY, maxY := 0, math.MinInt
	for elf := range elves {
		minX = min(minX, elf.x)
		maxX = max(maxX, elf.x)
		minY = min(minY, elf.y)
		maxY = max(maxY, elf.y)
	}

	w, h := maxX-minX+1, maxY-minY+1
	m := make([][]byte, h)
	for i := range m {
		m[i] = bytes.Repeat([]byte("."), w)
	}
	for elf := range elves {
		m[elf.y-minY][elf.x-minX] = '#'
	}
	for i := range m {
		fmt.Println(string(m[i]))
	}
	fmt.Println()
}
