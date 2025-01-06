package year_2021

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"strconv"
)

func init() {
	Solutions[20] = Day20{}
}

type Day20 struct{}

func (d Day20) Part1(input []byte) string {
	return d.solve(input, 2)
}

func (d Day20) Part2(input []byte) string {
	return d.solve(input, 50)
}

func (Day20) solve(input []byte, iterations int) string {
	r := bufio.NewScanner(bytes.NewReader(input))
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

	for i := 0; i < iterations; i++ {
		newImage := make([][]bool, len(image)+4)
		for x := range newImage {
			newImage[x] = make([]bool, len(image[0])+4)
		}

		for x := range newImage {
			for y := range newImage[x] {
				var idx int
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
			}
		}
	}

	return strconv.Itoa(totalLit)
}
