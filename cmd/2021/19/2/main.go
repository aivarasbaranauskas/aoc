package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_map"
	"github.com/aivarasbaranauskas/aoc/internal/_matrix"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	scanners := [][][3]int{{}}
	n := 0

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Text()
		if line == "" {
			scanners = append(scanners, [][3]int{})
			n++
			continue
		}
		if strings.HasPrefix(line, "---") {
			continue
		}
		spl := strings.Split(line, ",")
		scanners[n] = append(scanners[n], [3]int{
			optimistic.Atoi(spl[0]),
			optimistic.Atoi(spl[1]),
			optimistic.Atoi(spl[2]),
		})
	}

	rotations := GetRotations()

	all := _set.FromSlice(scanners[0])
	scannersPos := [][3]int{{0, 0, 0}}

	q := _a.Queue[[][][3]int]{}
	for j := 1; j <= n; j++ {
		var tmp [][][3]int
		for _, rotation := range rotations {
			tmp2 := _slice.Duplicate(scanners[j])
			for i := range tmp2 {
				t := _matrix.Multiply([][]int{tmp2[i][:]}, rotation)
				tmp2[i] = [3]int{t[0][0], t[0][1], t[0][2]}
			}
			tmp = append(tmp, tmp2)
		}
		q.Enqueue(tmp)
	}

Loop:
	for !q.Empty() {
		rots := q.Dequeue()

		for _, rotated := range rots {
			var distances [][3]int
			ddd := map[[3]int][][2][3]int{}
			for _, a := range all.ToSlice() {
				for _, b := range rotated {
					d := [3]int{
						a[0] - b[0],
						a[1] - b[1],
						a[2] - b[2],
					}
					distances = append(distances, d)

					if _, ok := ddd[d]; ok {
						ddd[d] = append(ddd[d], [2][3]int{a, b})
					} else {
						ddd[d] = [][2][3]int{{a, b}}
					}
				}
			}

			ud := _slice.CountUnique(distances)
			d, ct := _map.Max(ud)
			if ct >= 12 {
				for _, b := range rotated {
					all.Add([3]int{
						b[0] + d[0],
						b[1] + d[1],
						b[2] + d[2],
					})
				}

				scannersPos = append(scannersPos, d)
				continue Loop
			}
		}

		q.Enqueue(rots)
	}

	var m int
	for i := 0; i < len(scannersPos)-1; i++ {
		for j := i + 1; j < len(scannersPos); j++ {
			d := _num.Abs(scannersPos[i][0]-scannersPos[j][0]) + _num.Abs(scannersPos[i][1]-scannersPos[j][1]) + _num.Abs(scannersPos[i][2]-scannersPos[j][2])
			m = max(m, d)
		}
	}

	fmt.Println(m)
}

func GetRotations() [][][]int {

	rotationAroundZ := [][][]int{
		{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
		{
			{0, -1, 0},
			{1, 0, 0},
			{0, 0, 1},
		},
		{
			{-1, 0, 0},
			{0, -1, 0},
			{0, 0, 1},
		},
		{
			{0, 1, 0},
			{-1, 0, 0},
			{0, 0, 1},
		},
	}
	zRotations := [][][]int{
		{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
		{
			{1, 0, 0},
			{0, 0, -1},
			{0, 1, 0},
		},
		{
			{1, 0, 0},
			{0, -1, 0},
			{0, 0, -1},
		},
		{
			{1, 0, 0},
			{0, 0, 1},
			{0, -1, 0},
		},
		{
			{0, 0, 1},
			{0, 1, 0},
			{-1, 0, 0},
		},
		{
			{0, 0, -1},
			{0, 1, 0},
			{1, 0, 0},
		},
	}

	var allRotations [][][]int
	for _, r1 := range rotationAroundZ {
		for _, r2 := range zRotations {
			allRotations = append(allRotations, _matrix.Multiply(r1, r2))
		}
	}

	return allRotations
}
