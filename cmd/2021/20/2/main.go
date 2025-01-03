package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	r := bufio.NewScanner(f)
	r.Scan()
	algo := _slice.Map(
		r.Bytes(),
		func(x byte) bool { return x == '#' },
	)
	r.Scan()

	var image [][]bool
	for r.Scan() {
		image = append(
			image,
			_slice.Map(
				[]byte(r.Text()),
				func(x byte) bool { return x == '#' },
			),
		)
	}

	for i := 0; i < 50; i++ {
		newImage := make([][]bool, len(image)+4)
		for x := range newImage {
			newImage[x] = make([]bool, len(image[0])+4)
		}

		for x := range newImage {
			for y := range newImage[x] {
				var idx int
				if x == 2 && y == len(newImage[x])-1 {
					fmt.Println(idx)
				}
				for xi := 0; xi < 3; xi++ {
					for yi := 0; yi < 3; yi++ {
						xxi := x + xi - 2
						yyi := y + yi - 2

						if 0 <= xxi && xxi < len(image) && 0 <= yyi && yyi < len(image[0]) {
							if image[xxi][yyi] {
								idx |= 1 << (8 - (xi*3 + yi))
							}
						} else {
							if i%2 == 1 {
								idx |= 1 << (8 - (xi*3 + yi))
							}
						}
					}
				}

				newImage[x][y] = algo[idx]
			}
		}
		image = newImage
	}

	var totalLit int
	for x := range image {
		for _, b := range image[x] {
			if b {
				totalLit++
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println(totalLit)
}
