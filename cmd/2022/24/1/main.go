package main

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

var width, height int
var windsSteps []map[P][]byte

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	ffb, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(ffb), "\n")
	winds := map[P][]byte{}

	height, width = len(lines), len(lines[0])

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			c := lines[y][x]
			if c != '.' {
				if _, ok := winds[P{x, y}]; ok {
					winds[P{x, y}] = append(winds[P{x, y}], c)
				} else {
					winds[P{x, y}] = []byte{c}
				}
			}
		}
	}

	windsSteps = append(windsSteps, winds)

	//printWinds(winds)
	fmt.Println(aStar())
}

func printWinds(winds map[P][]byte) {
	m := make([][]byte, height)
	m[0] = bytes.Repeat([]byte("#"), width)
	m[0][1] = '.'
	for i := 1; i < height-1; i++ {
		m[i] = bytes.Repeat([]byte("."), width)
		m[i][0] = '#'
		m[i][width-1] = '#'
	}
	m[height-1] = bytes.Repeat([]byte("#"), width)
	m[height-1][width-2] = '.'

	for p, ws := range winds {
		if len(ws) > 1 {
			m[p.y][p.x] = byte('0' + len(ws))
		} else {
			m[p.y][p.x] = ws[0]
		}
	}

	for _, l := range m {
		fmt.Println(string(l))
	}
	fmt.Println()
}

func windsAtStep(step int) map[P][]byte {
	step = step % ((width - 2) * (height - 2))
	for len(windsSteps) <= step {
		windsSteps = append(windsSteps, nextWindsStep(windsSteps[len(windsSteps)-1]))
	}
	return windsSteps[step]
}

func nextWindsStep(winds map[P][]byte) map[P][]byte {
	w2 := map[P][]byte{}
	for p, ws := range winds {
		for _, w := range ws {
			newP := p
			switch w {
			case '>':
				newP.x++
				if newP.x == width-1 {
					newP.x = 1
				}
			case '<':
				newP.x--
				if newP.x == 0 {
					newP.x = width - 2
				}
			case 'v':
				newP.y++
				if newP.y == height-1 {
					newP.y = 1
				}
			case '^':
				newP.y--
				if newP.y == 0 {
					newP.y = height - 2
				}
			}

			if _, ok := w2[newP]; ok {
				w2[newP] = append(w2[newP], w)
			} else {
				w2[newP] = []byte{w}
			}
		}
	}

	//printWinds(w2)

	return w2
}

type P struct {
	x, y int
}

type State struct {
	x, y, step int
}

func aStar() int {
	start := State{1, 0, 0}
	openSet := _set.FromSlice([]State{start})
	cameFrom := make(map[State]State)
	gScore := map[State]int{start: 0}
	fScore := map[State]int{start: heuristic(start)}
	for !openSet.IsEmpty() {
		current := openSet.MinBy(func(m State) int {
			if x, ok := fScore[m]; ok {
				return x
			}
			return math.MaxInt
		})
		if current.x == width-2 && current.y == height-1 {
			return gScore[current]
		}

		openSet.Remove(current)

		for _, nextState := range genNextStates(current) {
			// d(current,neighbor) is the weight of the edge from current to neighbor
			// tentative_gScore is the distance from start to the neighbor through current
			tentativeGScore := gScore[current] + 1
			if nextGScore, ok := gScore[nextState]; !ok || tentativeGScore < nextGScore {
				// This path to neighbor is better than any previous one. Record it!
				cameFrom[nextState] = current
				gScore[nextState] = tentativeGScore
				fScore[nextState] = tentativeGScore + heuristic(nextState)
				if !openSet.Has(nextState) {
					openSet.Add(nextState)
				}
			}
		}
	}

	panic("NOT FOUND")
}

func validStep(s State) bool {
	if s.x == 1 && s.y == 0 {
		// start
		return true
	}
	if s.x == width-2 && s.y == height-1 {
		// end
		return true
	}
	if s.x <= 0 || s.x >= width-1 || s.y <= 0 || s.y >= height-1 {
		return false
	}

	winds := windsAtStep(s.step)
	_, isWind := winds[P{s.x, s.y}]
	return !isWind
}

func genNextStates(s State) []State {
	var states []State

	s.step++
	if validStep(s) {
		states = append(states, s)
	}

	st := s
	st.x++
	if validStep(st) {
		states = append(states, st)
	}

	st = s
	st.x--
	if validStep(st) {
		states = append(states, st)
	}

	st = s
	st.y++
	if validStep(st) {
		states = append(states, st)
	}

	st = s
	st.y--
	if validStep(st) {
		states = append(states, st)
	}

	return states
}

func heuristic(s State) int {
	return width + height - s.x - s.y - 3
}
