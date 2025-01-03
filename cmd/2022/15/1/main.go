package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	var s, b [][2]int
	n := 0

	r := bufio.NewScanner(f)
	for r.Scan() {
		spl := strings.Split(r.Text(), " ")
		s = append(s, [2]int{
			gI(spl[2][:len(spl[2])-1]),
			gI(spl[3][:len(spl[3])-1]),
		})
		b = append(b, [2]int{
			gI(spl[8][:len(spl[8])-1]),
			gI(spl[9][:len(spl[9])]),
		})
		n++
	}

	// inspected row
	iR := 2000000
	ps := _set.New[[2]int]()
	for i := 0; i < n; i++ {
		d := _num.Abs(s[i][0]-b[i][0]) + _num.Abs(s[i][1]-b[i][1])
		mx := d - _num.Abs(iR-s[i][1])
		for k := -mx; k <= mx; k++ {
			ps.Add([2]int{
				s[i][0] + k,
				iR,
			})
		}
	}

	for i := 0; i < n; i++ {
		ps.Remove(b[i])
	}

	fmt.Println(ps.Len())
}

func gI(s string) int {
	spl := strings.Split(s, "=")
	return optimistic.Atoi(spl[1])
}
