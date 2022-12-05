package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
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

	var energyLevels [][]int

	r := bufio.NewScanner(f)
	for r.Scan() {
		energyLevels = append(energyLevels,
			_slice.Map(
				strings.Split(r.Text(), ""),
				optimistic.Atoi,
			),
		)
	}

	steps := 100
	flashes := 0
	checkQueue := _a.Queue[[2]int]{}
	for i := 0; i < steps; i++ {
		// increase
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				energyLevels[x][y]++
			}
		}

		//flash
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				if energyLevels[x][y] > 9 {
					flashes++
					energyLevels[x][y] = -1

					for a := -1; a <= 1; a++ {
						for b := -1; b <= 1; b++ {
							nx := x + a
							ny := y + b
							if 0 <= nx && nx < 10 && 0 <= ny && ny < 10 && energyLevels[nx][ny] >= 0 {
								energyLevels[nx][ny]++
								checkQueue.Enqueue([2]int{nx, ny})
							}
						}
					}
				}
			}
		}

		for !checkQueue.Empty() {
			tmp := checkQueue.Dequeue()
			x, y := tmp[0], tmp[1]
			if energyLevels[x][y] > 9 {
				flashes++
				energyLevels[x][y] = -1

				for a := -1; a <= 1; a++ {
					for b := -1; b <= 1; b++ {
						nx := x + a
						ny := y + b
						if 0 <= nx && nx < 10 && 0 <= ny && ny < 10 && energyLevels[nx][ny] >= 0 {
							energyLevels[nx][ny]++
							checkQueue.Enqueue([2]int{nx, ny})
						}
					}
				}
			}
		}

		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				if energyLevels[x][y] < 0 {
					energyLevels[x][y] = 0
				}
			}
		}
	}

	fmt.Println(flashes)
}
