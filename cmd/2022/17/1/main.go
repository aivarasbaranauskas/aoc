package main

import (
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"io/ioutil"
	"log"
)

type Figure struct {
	bs []uint8
	w  int
	h  int
}

var figures = []Figure{
	{
		bs: []uint8{
			0b1111,
		},
		w: 4, h: 1,
	},
	{
		bs: []uint8{
			0b010,
			0b111,
			0b010,
		},
		w: 3, h: 3,
	},
	{
		bs: []uint8{
			0b111,
			0b100,
			0b100,
		},
		w: 3, h: 3,
	},
	{
		bs: []uint8{
			0b1,
			0b1,
			0b1,
			0b1,
		},
		w: 1, h: 4,
	},
	{
		bs: []uint8{
			0b11,
			0b11,
		},
		w: 2, h: 2,
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
	ffbL := len(ffb)

	var m []uint8
	z := 0

	for n := 0; n < 2022; n++ {
		fig := figures[n%5]
		x, y := 2, len(m)+3

		for {
			ch := 1
			if ffb[z] == '<' {
				ch = -1
			}
			z = (z + 1) % ffbL

			xTmp := clamp(x+ch, 0, 6)
			if x != xTmp && xTmp+fig.w <= 7 && noCollision(xTmp, y, fig, m) {
				x = xTmp
			}

			if y > 0 && noCollision(x, y-1, fig, m) {
				y--
			} else {
				break
			}
		}

		if len(m) <= y+fig.h {
			// grow
			g := y + fig.h - len(m)
			for i := 0; i < g; i++ {
				m = append(m, 0)
			}
		}
		for yy, p := range fig.bs {
			m[y+yy] = m[y+yy] | (p << x)
		}
		//draw(m)
	}

	fmt.Println(len(m))
}

func draw(m []uint8) {
	for i := len(m) - 1; i >= 0; i-- {
		for j := 0; j < 7; j++ {
			if m[i]&(1<<j) == uint8(1<<j) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func noCollision(x, y int, fig Figure, m []uint8) bool {
	if y >= len(m) {
		return true
	}
	for i := 0; i < len(fig.bs) && y+i < len(m); i++ {
		if m[y+i]&(fig.bs[i]<<x) > 0 {
			return false
		}
	}

	return true
}

func clamp(x, min, max int) int {
	return _num.Min(max, _num.Max(min, x))
}
