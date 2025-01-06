package year_2022

import (
	"bufio"
	"bytes"
	"math"
	"strconv"
)

func init() {
	Solutions[23] = Day23{}
}

type Day23 struct{}

type Day23Pos struct {
	x, y int
}

func (Day23) Part1(input []byte) string {
	elves := map[Day23Pos]struct{}{}

	r := bufio.NewScanner(bytes.NewReader(input))
	y := 0
	for r.Scan() {
		for x, c := range []byte(r.Text()) {
			if c == '#' {
				elves[Day23Pos{x, y}] = struct{}{}
			}
		}
		y++
	}

	// 1 - north
	// 2 - south
	// 3 - west
	// 4 - east
	dirs := []int{1, 2, 3, 4}
	for i := 0; i < 10; i++ {
		// first half - generate propositions
		// [to] -> [from elf 1, from elf 2, ...]
		propositions := map[Day23Pos][]Day23Pos{}
	ElfLoop:
		for elf := range elves {
			shouldMove := false
		L:
			for a := -1; a <= 1; a++ {
				for b := -1; b <= 1; b++ {
					if a == b && b == 0 {
						continue
					}
					if _, ok := elves[Day23Pos{elf.x + a, elf.y + b}]; ok {
						shouldMove = true
						break L
					}
				}
			}

			if !shouldMove {
				continue
			}

			for _, d := range dirs {
				switch d {
				case 1:
					// north
					newP := Day23Pos{elf.x, elf.y - 1}
					_, a := elves[Day23Pos{elf.x - 1, elf.y - 1}]
					_, b := elves[newP]
					_, c := elves[Day23Pos{elf.x + 1, elf.y - 1}]
					if !a && !b && !c {
						if _, ok := propositions[newP]; ok {
							propositions[newP] = append(propositions[newP], elf)
						} else {
							propositions[newP] = []Day23Pos{elf}
						}
						continue ElfLoop
					}
				case 2:
					// south
					newP := Day23Pos{elf.x, elf.y + 1}
					_, a := elves[Day23Pos{elf.x - 1, elf.y + 1}]
					_, b := elves[newP]
					_, c := elves[Day23Pos{elf.x + 1, elf.y + 1}]
					if !a && !b && !c {
						if _, ok := propositions[newP]; ok {
							propositions[newP] = append(propositions[newP], elf)
						} else {
							propositions[newP] = []Day23Pos{elf}
						}
						continue ElfLoop
					}
				case 3:
					// west
					newP := Day23Pos{elf.x - 1, elf.y}
					_, a := elves[Day23Pos{elf.x - 1, elf.y - 1}]
					_, b := elves[newP]
					_, c := elves[Day23Pos{elf.x - 1, elf.y + 1}]
					if !a && !b && !c {
						if _, ok := propositions[newP]; ok {
							propositions[newP] = append(propositions[newP], elf)
						} else {
							propositions[newP] = []Day23Pos{elf}
						}
						continue ElfLoop
					}
				case 4:
					// east
					newP := Day23Pos{elf.x + 1, elf.y}
					_, a := elves[Day23Pos{elf.x + 1, elf.y - 1}]
					_, b := elves[newP]
					_, c := elves[Day23Pos{elf.x + 1, elf.y + 1}]
					if !a && !b && !c {
						if _, ok := propositions[newP]; ok {
							propositions[newP] = append(propositions[newP], elf)
						} else {
							propositions[newP] = []Day23Pos{elf}
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
	}

	minX, maxX := math.MaxInt, math.MinInt
	minY, maxY := math.MaxInt, math.MinInt
	for elf := range elves {
		minX = min(minX, elf.x)
		maxX = max(maxX, elf.x)
		minY = min(minY, elf.y)
		maxY = max(maxY, elf.y)
	}

	return strconv.Itoa((maxX-minX+1)*(maxY-minY+1) - len(elves))
}

func (Day23) Part2(input []byte) string {
	elves := map[Day23Pos]struct{}{}

	r := bufio.NewScanner(bytes.NewReader(input))
	y := 0
	for r.Scan() {
		for x, c := range []byte(r.Text()) {
			if c == '#' {
				elves[Day23Pos{x, y}] = struct{}{}
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
		propositions := map[Day23Pos][]Day23Pos{}
	ElfLoop:
		for elf := range elves {
			shouldMove := false
		L:
			for a := -1; a <= 1; a++ {
				for b := -1; b <= 1; b++ {
					if a == b && b == 0 {
						continue
					}
					if _, ok := elves[Day23Pos{elf.x + a, elf.y + b}]; ok {
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
					newP := Day23Pos{elf.x, elf.y - 1}
					_, a := elves[Day23Pos{elf.x - 1, elf.y - 1}]
					_, b := elves[newP]
					_, c := elves[Day23Pos{elf.x + 1, elf.y - 1}]
					if !a && !b && !c {
						if _, ok := propositions[newP]; ok {
							propositions[newP] = append(propositions[newP], elf)
						} else {
							propositions[newP] = []Day23Pos{elf}
						}
						continue ElfLoop
					}
				case 2:
					// south
					newP := Day23Pos{elf.x, elf.y + 1}
					_, a := elves[Day23Pos{elf.x - 1, elf.y + 1}]
					_, b := elves[newP]
					_, c := elves[Day23Pos{elf.x + 1, elf.y + 1}]
					if !a && !b && !c {
						if _, ok := propositions[newP]; ok {
							propositions[newP] = append(propositions[newP], elf)
						} else {
							propositions[newP] = []Day23Pos{elf}
						}
						continue ElfLoop
					}
				case 3:
					// west
					newP := Day23Pos{elf.x - 1, elf.y}
					_, a := elves[Day23Pos{elf.x - 1, elf.y - 1}]
					_, b := elves[newP]
					_, c := elves[Day23Pos{elf.x - 1, elf.y + 1}]
					if !a && !b && !c {
						if _, ok := propositions[newP]; ok {
							propositions[newP] = append(propositions[newP], elf)
						} else {
							propositions[newP] = []Day23Pos{elf}
						}
						continue ElfLoop
					}
				case 4:
					// east
					newP := Day23Pos{elf.x + 1, elf.y}
					_, a := elves[Day23Pos{elf.x + 1, elf.y - 1}]
					_, b := elves[newP]
					_, c := elves[Day23Pos{elf.x + 1, elf.y + 1}]
					if !a && !b && !c {
						if _, ok := propositions[newP]; ok {
							propositions[newP] = append(propositions[newP], elf)
						} else {
							propositions[newP] = []Day23Pos{elf}
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
	}

	return strconv.Itoa(ct)
}
