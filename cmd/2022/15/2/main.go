package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"os"
	"sort"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	var sb [][2][2]int
	n := 0

	r := bufio.NewScanner(f)
	for r.Scan() {
		spl := strings.Split(r.Text(), " ")
		sb = append(sb, [2][2]int{
			{
				gI(spl[2][:len(spl[2])-1]),
				gI(spl[3][:len(spl[3])-1]),
			},
			{
				gI(spl[8][:len(spl[8])-1]),
				gI(spl[9][:len(spl[9])]),
			},
		})
		n++
	}

	sort.Slice(sb, func(i, j int) bool {
		return sb[i][0][1] < sb[j][0][1]
	})

	var dBs []int
	for i := 0; i < n; i++ {
		dBs = append(dBs, _num.Abs(sb[i][0][0]-sb[i][1][0])+_num.Abs(sb[i][0][1]-sb[i][1][1]))
	}

	// search field
	iR := 4000000

	for y := 0; y <= iR; y++ {
		var ranges [][2]int
		for i := 0; i < n; i++ {
			if sb[i][0][1]-dBs[i] > y || y > sb[i][0][1]+dBs[i] {
				continue
			}

			mx := dBs[i] - _num.Abs(y-sb[i][0][1])
			ranges = append(ranges, [2]int{
				max(sb[i][0][0]-mx, 0),
				min(sb[i][0][0]+mx, iR),
			})
		}

		sort.Slice(ranges, func(i, j int) bool {
			if ranges[i][0] == ranges[j][0] {
				return ranges[i][1] < ranges[j][1]
			}
			return ranges[i][0] < ranges[j][0]
		})

		rr := ranges[0]
		if rr[0] > 0 {
			fmt.Println(ranges)
			x := 0
			fmt.Println(x*4000000 + y)
			os.Exit(0)
		}
		for i := 1; i < len(ranges)-1; i++ {
			if rr[1] < ranges[i][0] {
				fmt.Println(ranges)
				x := rr[1] + 1
				fmt.Println(x*4000000 + y)
				os.Exit(0)
			} else {
				rr[1] = max(rr[1], ranges[i][1])
			}
		}
	}

	fmt.Println("not found")
}

func gI(s string) int {
	spl := strings.Split(s, "=")
	return optimistic.Atoi(spl[1])
}
