package main

import (
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"io/ioutil"
	"log"
)

var figures = [][][2]int{
	{
		{0, 0}, {0, 1}, {0, 2}, {0, 3},
	},
	{
		{0, 1},
		{1, 0}, {1, 1}, {1, 2},
		{2, 1},
	},
	{
		{2, 2},
		{1, 2},
		{0, 0}, {0, 1}, {0, 2},
	},
	{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
	},
	{
		{0, 0}, {0, 1},
		{1, 0}, {1, 1},
	},
}

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	ffb, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	var m [][7]bool
	z := 0

	for n := 0; n < 2022; n++ {
		fig := figures[n%5]
		x, y := 2, len(m)+3

		for {
			ch := 1
			if ffb[z] == '<' {
				ch = -1
			}
			z = (z + 1) % len(ffb)

			xTmp := clamp(x+ch, 0, 6)
			if x != xTmp && xInBounds(xTmp, fig) && noCollision(xTmp, y, fig, m) {
				x = xTmp
			}

			if y > 0 && noCollision(x, y-1, fig, m) {
				y--
			} else {
				break
			}
		}

		for _, p := range fig {
			if len(m) <= y+p[0] {
				// grow
				g := y + p[0] + 1 - len(m)
				for i := 0; i < g; i++ {
					m = append(m, [7]bool{})
				}
			}
			m[y+p[0]][x+p[1]] = true
		}
		//draw(m)
	}

	//for _, d := range ffb {
	//	ch := 1
	//	if d == '<' {
	//		ch = -1
	//	}
	//
	//	xTmp := clamp(x+ch, 0, 6)
	//	if x != xTmp && xInBounds(xTmp, fig) && noCollision(xTmp, y, fig, m) {
	//		x = xTmp
	//	}
	//
	//	if y > 0 && noCollision(x, y-1, fig, m) {
	//		y--
	//		continue
	//	}
	//
	//	for _, p := range fig {
	//		if len(m) <= y+p[0] {
	//			// grow
	//			g := y + p[0] + 1 - len(m)
	//			for i := 0; i < g; i++ {
	//				m = append(m, [7]bool{})
	//			}
	//		}
	//		m[y+p[0]][x+p[1]] = true
	//	}
	//	figId = (figId + 1) % 5
	//	fig = figures[figId]
	//	x = 2
	//	y = len(m) + 3
	//}

	fmt.Println(len(m))
}

func draw(m [][7]bool) {
	for i := len(m) - 1; i >= 0; i-- {
		for j := 0; j < 7; j++ {
			if m[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func xInBounds(x int, fig [][2]int) bool {
	for _, p := range fig {
		if p[1]+x >= 7 {
			return false
		}
	}
	return true
}

func noCollision(x, y int, fig [][2]int, m [][7]bool) bool {
	for _, p := range fig {
		if len(m) > y+p[0] && m[y+p[0]][x+p[1]] {
			return false
		}
	}

	return true
}

func clamp(x, min, max int) int {
	return _num.Min(max, _num.Max(min, x))
}
