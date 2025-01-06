package year_2021

import (
	"bytes"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_matrix"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"math"
	"sort"
	"strconv"
)

func init() {
	Solutions[15] = Day15{}
}

type Day15 struct{}

func (Day15) Part1(input []byte) string {
	m := _matrix.ReadIntMatrix(bytes.NewReader(input))
	l := len(m)
	end := [2]int{l - 1, l - 1}

	openSet := _set.New[[2]int]()
	openSet.Add([2]int{0, 0})

	cameFrom := map[[2]int][2]int{}
	fScore := map[[2]int]int{{0, 0}: m[0][0]}

	for openSet.Len() > 0 {
		openSetSl := openSet.ToSlice()
		sort.Slice(openSetSl, func(i, j int) bool {
			fScoreI, ok := fScore[openSetSl[i]]
			if !ok {
				fScoreI = math.MaxInt
			}
			fScoreJ, ok := fScore[openSetSl[i]]
			if !ok {
				fScoreJ = math.MaxInt
			}
			return fScoreI < fScoreJ
		})
		current := openSetSl[0]

		openSet.Remove(current)
		neighbors := _slice.Filter(
			[][2]int{
				{current[0] - 1, current[1]},
				{current[0] + 1, current[1]},
				{current[0], current[1] - 1},
				{current[0], current[1] + 1},
			},
			func(v [2]int) bool {
				return v[0] >= 0 && v[0] <= end[0] && v[1] >= 0 && v[1] <= end[1]
			},
		)

		for _, neighbor := range neighbors {
			tentativeFScore := fScore[current] + m[neighbor[0]][neighbor[1]]
			if neighborFScore, ok := fScore[neighbor]; !ok || tentativeFScore < neighborFScore {
				// This path to neighbor is better than any previous one. Record it!
				cameFrom[neighbor] = current
				fScore[neighbor] = tentativeFScore
				if !openSet.Has(neighbor) {
					openSet.Add(neighbor)
				}
			}
		}

	}

	var totalPath [][2]int
	var items []int
	current, ok := end, true
	for ok {
		totalPath = append(totalPath, current)
		items = append(items, m[current[0]][current[1]])
		current, ok = cameFrom[current]
	}
	return strconv.Itoa(fScore[end] - fScore[[2]int{0, 0}])
}

func (Day15) Part2(input []byte) string {
	m := _matrix.ReadIntMatrix(bytes.NewReader(input))
	lOrig := len(m)
	l := lOrig * 5
	tmp := make([][]int, l)
	for x := 0; x < l; x++ {
		tmp[x] = make([]int, l)
		for y := 0; y < l; y++ {
			tmp[x][y] = m[x%lOrig][y%lOrig] + (x / lOrig) + (y / lOrig)
			if tmp[x][y] > 9 {
				tmp[x][y] -= 9
			}
		}
	}
	m = tmp
	end := [2]int{l - 1, l - 1}

	queue := _a.Queue[[2]int]{}
	queue.Enqueue([2]int{0, 0})

	cameFrom := map[[2]int][2]int{}
	fScore := map[[2]int]int{{0, 0}: m[0][0]}

	for !queue.Empty() {
		fmt.Println(queue.Len())
		current := queue.Dequeue()
		neighbors := _slice.Filter(
			[][2]int{
				{current[0] - 1, current[1]},
				{current[0] + 1, current[1]},
				{current[0], current[1] - 1},
				{current[0], current[1] + 1},
			},
			func(v [2]int) bool {
				return v[0] >= 0 && v[0] <= end[0] && v[1] >= 0 && v[1] <= end[1]
			},
		)

		for _, neighbor := range neighbors {
			tentativeFScore := fScore[current] + m[neighbor[0]][neighbor[1]]
			if neighborFScore, ok := fScore[neighbor]; !ok || tentativeFScore < neighborFScore {
				// This path to neighbor is better than any previous one. Record it!
				cameFrom[neighbor] = current
				fScore[neighbor] = tentativeFScore
				queue.Enqueue(neighbor)
			}
		}

	}

	var totalPath [][2]int
	var items []int
	current, ok := end, true
	for ok {
		totalPath = append(totalPath, current)
		items = append(items, m[current[0]][current[1]])
		current, ok = cameFrom[current]
	}
	return strconv.Itoa(fScore[end] - fScore[[2]int{0, 0}])
}
