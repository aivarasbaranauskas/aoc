package main

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"io"
	"log"
	"os"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	numbers, boards := loadData()

	for _, number := range numbers {
		for _, board := range boards {
			if board.Mark(number) {
				fmt.Println("Score:", board.SumUnmarked()*number)
				os.Exit(0)
			}
		}
	}
}

type Board struct {
	Numbers [5][5]int
	marks   [5][5]bool
}

func (b *Board) Mark(number int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.Numbers[i][j] == number {
				b.marks[i][j] = true

				bingo := true
				for k := 0; k < 5; k++ {
					bingo = bingo && b.marks[i][k]
				}
				if bingo {
					return true
				}

				bingo = true
				for k := 0; k < 5; k++ {
					bingo = bingo && b.marks[k][j]
				}
				if bingo {
					return true
				}
			}
		}
	}

	return false
}

func (b *Board) SumUnmarked() (sum int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.marks[i][j] {
				sum += b.Numbers[i][j]
			}
		}
	}
	return
}

func loadData() ([]int, []*Board) {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	r := optimistic.NewReader(f)
	numbersLine := r.ReadStringLine()
	_ = r.ReadStringLine() // Empty line

	numbersS := strings.Split(numbersLine, ",")
	numbers := make([]int, len(numbersS))
	for i := range numbersS {
		numbers[i] = optimistic.Atoi(numbersS[i])
	}

	var boards []*Board

	for {
		board := Board{}
		for i := 0; i < 5; i++ {
			line := r.ReadBytesLine()
			for j := 0; j < 5; j++ {
				b := j * 3
				x := line[b : b+2]
				board.Numbers[i][j] = optimistic.Atoi(string(bytes.TrimSpace(x)))
			}
		}
		boards = append(boards, &board)

		_, err = r.Reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				log.Fatalln(err)
			}
		}
	}

	return numbers, boards
}
