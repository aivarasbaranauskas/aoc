package year_2022

import (
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"math"
	"strconv"
	"strings"
)

func init() {
	Solutions[24] = &Day24{}
}

type Day24 struct {
	width, height int
	windsSteps    []map[P][]byte
}

func (d *Day24) Part1(input []byte) string {
	d.parseData(input)
	return strconv.Itoa(d.aStar(State{1, 0, 0}, d.width-2, d.height-1))
}

func (d *Day24) Part2(input []byte) string {
	d.parseData(input)
	forward := d.aStar(State{1, 0, 0}, d.width-2, d.height-1)
	back := d.aStar(State{d.width - 2, d.height - 1, forward}, 1, 0)
	forwardAgain := d.aStar(State{1, 0, back + forward}, d.width-2, d.height-1)
	return strconv.Itoa(forward + back + forwardAgain)
}

func (d *Day24) parseData(input []byte) {
	lines := strings.Split(string(input), "\n")
	winds := map[P][]byte{}

	d.height, d.width = len(lines), len(lines[0])

	for y := 1; y < d.height-1; y++ {
		for x := 1; x < d.width-1; x++ {
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

	d.windsSteps = []map[P][]byte{winds}
}

func (d *Day24) windsAtStep(step int) map[P][]byte {
	step = step % ((d.width - 2) * (d.height - 2))
	for len(d.windsSteps) <= step {
		d.windsSteps = append(d.windsSteps, d.nextWindsStep(d.windsSteps[len(d.windsSteps)-1]))
	}
	return d.windsSteps[step]
}

func (d *Day24) nextWindsStep(winds map[P][]byte) map[P][]byte {
	w2 := map[P][]byte{}
	for p, ws := range winds {
		for _, w := range ws {
			newP := p
			switch w {
			case '>':
				newP.x++
				if newP.x == d.width-1 {
					newP.x = 1
				}
			case '<':
				newP.x--
				if newP.x == 0 {
					newP.x = d.width - 2
				}
			case 'v':
				newP.y++
				if newP.y == d.height-1 {
					newP.y = 1
				}
			case '^':
				newP.y--
				if newP.y == 0 {
					newP.y = d.height - 2
				}
			}

			if _, ok := w2[newP]; ok {
				w2[newP] = append(w2[newP], w)
			} else {
				w2[newP] = []byte{w}
			}
		}
	}

	return w2
}

type P struct {
	x, y int
}

type State struct {
	x, y, step int
}

func (d *Day24) aStar(start State, goalX, goalY int) int {
	openSet := _set.FromSlice([]State{start})
	cameFrom := make(map[State]State)
	gScore := map[State]int{start: 0}
	fScore := map[State]int{start: d.heuristic(start)}
	for !openSet.IsEmpty() {
		current := openSet.MinBy(func(m State) int {
			if x, ok := fScore[m]; ok {
				return x
			}
			return math.MaxInt
		})
		if current.x == goalX && current.y == goalY {
			return gScore[current]
		}

		openSet.Remove(current)

		for _, nextState := range d.genNextStates(current) {
			// d(current,neighbor) is the weight of the edge from current to neighbor
			// tentative_gScore is the distance from start to the neighbor through current
			tentativeGScore := gScore[current] + 1
			if nextGScore, ok := gScore[nextState]; !ok || tentativeGScore < nextGScore {
				// This path to neighbor is better than any previous one. Record it!
				cameFrom[nextState] = current
				gScore[nextState] = tentativeGScore
				fScore[nextState] = tentativeGScore + d.heuristic(nextState)
				if !openSet.Has(nextState) {
					openSet.Add(nextState)
				}
			}
		}
	}

	panic("NOT FOUND")
}

func (d *Day24) validStep(s State) bool {
	if s.x == 1 && s.y == 0 {
		// start
		return true
	}
	if s.x == d.width-2 && s.y == d.height-1 {
		// end
		return true
	}
	if s.x <= 0 || s.x >= d.width-1 || s.y <= 0 || s.y >= d.height-1 {
		return false
	}

	winds := d.windsAtStep(s.step)
	_, isWind := winds[P{s.x, s.y}]
	return !isWind
}

func (d *Day24) genNextStates(s State) []State {
	var states []State

	s.step++
	if d.validStep(s) {
		states = append(states, s)
	}

	st := s
	st.x++
	if d.validStep(st) {
		states = append(states, st)
	}

	st = s
	st.x--
	if d.validStep(st) {
		states = append(states, st)
	}

	st = s
	st.y++
	if d.validStep(st) {
		states = append(states, st)
	}

	st = s
	st.y--
	if d.validStep(st) {
		states = append(states, st)
	}

	return states
}

func (d *Day24) heuristic(s State) int {
	return d.width + d.height - s.x - s.y - 3
}
