package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/go_helpers/_num"
	"github.com/aivarasbaranauskas/aoc/go_helpers/o"
	"log"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var points [][3]int

	r := bufio.NewScanner(f)
	for r.Scan() {
		spl := strings.Split(r.Text(), ",")
		points = append(
			points,
			[3]int{
				o.Atoi(spl[0]),
				o.Atoi(spl[1]),
				o.Atoi(spl[2]),
			},
		)
	}

	sides := 0
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			if manhattan(points[i], points[j]) == 1 {
				sides++
			}
		}
	}

	fmt.Println(len(points)*6 - sides*2)
}

func manhattan(a, b [3]int) int {
	return _num.Abs(a[0]-b[0]) + _num.Abs(a[1]-b[1]) + _num.Abs(a[2]-b[2])
}
