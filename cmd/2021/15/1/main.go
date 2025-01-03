package main

import (
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_matrix"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"math"
	"sort"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	m := _matrix.ReadIntMatrix(f)
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
			if neighborFScore, ok := fScore[neighbor]; !ok || (ok && tentativeFScore < neighborFScore) {
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
	//fmt.Println(totalPath)
	//fmt.Println(items)
	fmt.Println(fScore[end] - fScore[[2]int{0, 0}])
}
