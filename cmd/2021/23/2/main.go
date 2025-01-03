package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"math"
)

//go:embed input.txt
var input []byte

var POW10 = [4]int{1, 10, 100, 1000}

func main() {
	M := bytes.Split(input, []byte("\n"))
	m := [27]byte{
		'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.',
		11: M[2][3], 15: M[2][5], 19: M[2][7], 23: M[2][9],
		12: 'D', 16: 'C', 20: 'B', 24: 'A',
		13: 'D', 17: 'B', 21: 'A', 25: 'C',
		14: M[3][3], 18: M[3][5], 22: M[3][7], 26: M[3][9],
	}

	fmt.Println(aStar(m))
}

var GOAL = [27]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', 'A', 'A', 'A', 'A', 'B', 'B', 'B', 'B', 'C', 'C', 'C', 'C', 'D', 'D', 'D', 'D'}

func aStar(m [27]byte) int {
	// The set of discovered nodes that may need to be (re-)expanded.
	// Initially, only the start node is known.
	// This is usually implemented as a min-heap or priority queue rather than a hash-set.
	openSet := _set.FromSlice([][27]byte{m})

	// For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from start
	// to n currently known.
	cameFrom := make(map[[27]byte][27]byte)

	// For node n, gScore[n] is the cost of the cheapest path from start to n currently known.
	gScore := map[[27]byte]int{m: 0}

	// For node n, fScore[n] := gScore[n] + h(n). fScore[n] represents our current best guess as to
	// how cheap a path could be from start to finish if it goes through n.
	fScore := map[[27]byte]int{m: heuristic(m)}

	for !openSet.IsEmpty() {
		// This operation can occur in O(Log(N)) time if openSet is a min-heap or a priority queue
		current := openSet.MinBy(func(m [27]byte) int {
			if x, ok := fScore[m]; ok {
				return x
			}
			return math.MaxInt
		})
		if current == GOAL {
			return gScore[current]
		}

		openSet.Remove(current)

		for _, move := range genMoves(current) {
			from, to, dist := move[0], move[1], move[2]
			nextM := current
			nextM[from], nextM[to] = '.', current[from]
			// d(current,neighbor) is the weight of the edge from current to neighbor
			// tentative_gScore is the distance from start to the neighbor through current
			aaa := current[from] - 'A'
			tentativeGscore := gScore[current] + dist*POW10[aaa]
			if nextGScore, ok := gScore[nextM]; !ok || tentativeGscore < nextGScore {
				// This path to neighbor is better than any previous one. Record it!
				cameFrom[nextM] = current
				gScore[nextM] = tentativeGscore
				fScore[nextM] = tentativeGscore + heuristic(nextM)
				if !openSet.Has(nextM) {
					openSet.Add(nextM)
				}
			}
		}
	}

	panic("NOT FOUND")
}

func genMoves(m [27]byte) (moves [][3]int) { // [from, to, len]
	tryGenMove := func(hallway, room int, toRoom bool) {
		from := min(hallway, 2+2*room)
		to := max(hallway, 2+2*room)
		for i := from; i < to; i++ {
			if i != hallway && m[i] != '.' {
				// path obstructed
				return
			}
		}
		pathLen := to - from

		if toRoom {
			// From hallway to room.
			vacantIdx := 0
			for i := 0; i < 4; i++ {
				p := 11 + room*4 + i
				if m[p] != '.' {
					if int(m[p]-'A') != room {
						// not ready to fill the room yet
						return
					}
				} else {
					vacantIdx = i
				}
			}
			moves = append(moves, [3]int{
				hallway,
				11 + room*4 + vacantIdx,
				pathLen + vacantIdx + 1,
			})
		} else {
			occupiedIdx := -1
			var targetRoom int
			for i := 0; i < 4; i++ {
				p := 11 + room*4 + i
				if m[p] != '.' {
					occupiedIdx = i
					targetRoom = int(m[p] - 'A')
					break
				}
			}
			if occupiedIdx == -1 {
				return
			}
			directRoute := room == targetRoom
			if !(directRoute && pathLen > 2) {
				moves = append(moves, [3]int{
					11 + room*4 + occupiedIdx,
					hallway,
					pathLen + occupiedIdx + 1,
				})
			}
		}
	}

	for _, hallway := range []int{0, 1, 3, 5, 7, 9, 10} {
		if m[hallway] != '.' {
			tryGenMove(hallway, int(m[hallway]-'A'), true)
		} else {
			for room := 0; room < 4; room++ {
				tryGenMove(hallway, room, false)
			}
		}
	}

	return moves
}

func heuristic(m [27]byte) (s int) {
	var notInRoom [4]int
	for i := 0; i < 11; i++ {
		if m[i] != '.' {
			targetRoom := int(m[i] - 'A')
			notInRoom[targetRoom]++
			dist := _num.Abs(i - (2 + 2*targetRoom))
			s += dist * POW10[targetRoom]
		}
	}

	for room := 0; room < 4; room++ {
		for offset := 0; offset < 4; offset++ {
			i := 11 + 4*room + offset
			if m[i] != '.' {
				targetRoom := int(m[i] - 'A')
				if targetRoom != room {
					notInRoom[targetRoom]++
					hallwayPathLen := 2 * _num.Abs(room-targetRoom)
					exitDist := 1 + offset
					s += (exitDist + hallwayPathLen) * POW10[targetRoom]
				}
			}
		}
	}

	for i, k := range notInRoom {
		s += (k * (k + 1) / 2) * POW10[i]
	}

	return
}
