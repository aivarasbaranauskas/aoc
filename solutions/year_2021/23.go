package year_2021

import (
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"math"
	"strconv"
)

func init() {
	Solutions[23] = Day23{}
}

type Day23 struct{}

var Day23POW10 = [4]int{1, 10, 100, 1000}

func (d Day23) Part1(input []byte) string {
	M := bytes.Split(input, []byte("\n"))
	m := []byte{
		'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.',
		11: M[2][3], 13: M[2][5], 15: M[2][7], 17: M[2][9],
		12: M[3][3], 14: M[3][5], 16: M[3][7], 18: M[3][9],
	}

	result := d.aStar(
		string(m),
		"...........AABBCCDD",
		2,
	)

	return strconv.Itoa(result)
}

func (d Day23) Part2(input []byte) string {
	M := bytes.Split(input, []byte("\n"))
	m := []byte{
		'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.',
		11: M[2][3], 15: M[2][5], 19: M[2][7], 23: M[2][9],
		12: 'D', 16: 'C', 20: 'B', 24: 'A',
		13: 'D', 17: 'B', 21: 'A', 25: 'C',
		14: M[3][3], 18: M[3][5], 22: M[3][7], 26: M[3][9],
	}

	result := d.aStar(
		string(m),
		"...........AAAABBBBCCCCDDDD",
		4,
	)

	return strconv.Itoa(result)
}

func (d Day23) aStar(m string, GOAL string, depth int) int {
	openSet := _set.FromSlice([]string{m})
	cameFrom := make(map[string]string)
	gScore := map[string]int{m: 0}
	fScore := map[string]int{m: d.heuristic(m, depth)}

	for !openSet.IsEmpty() {
		// This operation can occur in O(Log(N)) time if openSet is a min-heap or a priority queue
		current := openSet.MinBy(func(m string) int {
			if x, ok := fScore[m]; ok {
				return x
			}
			return math.MaxInt
		})
		if current == GOAL {
			return gScore[current]
		}

		openSet.Remove(current)

		for _, move := range d.genMoves([]byte(current), depth) {
			from, to, dist := move[0], move[1], move[2]
			next := []byte(current)
			next[from], next[to] = '.', current[from]
			nextM := string(next)
			aaa := current[from] - 'A'
			tentativeGscore := gScore[current] + dist*Day23POW10[aaa]
			if nextGScore, ok := gScore[nextM]; !ok || tentativeGscore < nextGScore {
				// This path to neighbor is better than any previous one. Record it!
				cameFrom[nextM] = current
				gScore[nextM] = tentativeGscore
				fScore[nextM] = tentativeGscore + d.heuristic(nextM, depth)
				if !openSet.Has(nextM) {
					openSet.Add(nextM)
				}
			}
		}
	}

	panic("NOT FOUND")
}

func (Day23) genMoves(m []byte, d int) (moves [][3]int) { // [from, to, len]
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
			for i := 0; i < d; i++ {
				p := 11 + room*d + i
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
				11 + room*d + vacantIdx,
				pathLen + vacantIdx + 1,
			})
		} else {
			occupiedIdx := -1
			var targetRoom int
			for i := 0; i < d; i++ {
				p := 11 + room*d + i
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
					11 + room*d + occupiedIdx,
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

func (Day23) heuristic(m string, d int) (s int) {
	var notInRoom [4]int
	for i := 0; i < 11; i++ {
		if m[i] != '.' {
			targetRoom := int(m[i] - 'A')
			notInRoom[targetRoom]++
			dist := _num.Abs(i - (2 + 2*targetRoom))
			s += dist * Day23POW10[targetRoom]
		}
	}

	for room := 0; room < 4; room++ {
		for offset := 0; offset < d; offset++ {
			i := 11 + d*room + offset
			if m[i] != '.' {
				targetRoom := int(m[i] - 'A')
				if targetRoom != room {
					notInRoom[targetRoom]++
					hallwayPathLen := 2 * _num.Abs(room-targetRoom)
					exitDist := 1 + offset
					s += (exitDist + hallwayPathLen) * Day23POW10[targetRoom]
				}
			}
		}
	}

	for i, k := range notInRoom {
		s += (k * (k + 1) / 2) * Day23POW10[i]
	}

	return
}
