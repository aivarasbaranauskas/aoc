package year_2021

import (
	"bytes"
	"errors"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"io"
	"log"
	"strconv"
	"strings"
)

func init() {
	Solutions[4] = Day4{}
}

type Day4 struct{}

func (d Day4) Part1(input []byte) string {
	numbers, boards := d.loadData(input)

	for _, number := range numbers {
		for _, board := range boards {
			if board.Mark(number) {
				return strconv.Itoa(board.SumUnmarked() * number)
			}
		}
	}

	return "not found"
}

func (d Day4) Part2(input []byte) string {
	numbers, boards := d.loadData(input)

	var lastWinner *Day4Board
	var lastWinnerNumber int

outer:
	for _, number := range numbers {
		ct := len(boards)
		for i := 0; i < ct; i++ {
			board := boards[i]
			if board.Mark(number) {
				lastWinner = board
				lastWinnerNumber = number

				if len(boards) == 1 {
					break outer
				}
				boards[i] = boards[ct-1]
				boards = boards[:ct-1]
				ct--
				i--
			}
		}
	}

	return strconv.Itoa(lastWinner.SumUnmarked() * lastWinnerNumber)
}

func (Day4) loadData(input []byte) ([]int, []*Day4Board) {
	r := optimistic.NewReader(bytes.NewReader(input))
	numbersLine := r.ReadStringLine()
	_ = r.ReadStringLine() // Empty line

	numbersS := strings.Split(numbersLine, ",")
	numbers := make([]int, len(numbersS))
	for i := range numbersS {
		numbers[i] = optimistic.Atoi(numbersS[i])
	}

	var boards []*Day4Board

	for {
		board := Day4Board{}
		var err error
		for i := 0; i < 5; i++ {
			var line []byte
			line, err = r.Reader.ReadBytes('\n')
			line = bytes.Trim(line, "\n")
			for j := 0; j < 5; j++ {
				b := j * 3
				x := line[b : b+2]
				board.Numbers[i][j] = optimistic.Atoi(string(bytes.TrimSpace(x)))
			}
		}
		boards = append(boards, &board)
		if err != nil && errors.Is(err, io.EOF) {
			if errors.Is(err, io.EOF) {
				break
			} else {
				log.Fatalln(err)
			}
		}
		_, _ = r.Reader.ReadString('\n')
	}

	return numbers, boards
}

type Day4Board struct {
	Numbers [5][5]int
	marks   [5][5]bool
}

func (b *Day4Board) Mark(number int) bool {
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

func (b *Day4Board) SumUnmarked() (sum int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.marks[i][j] {
				sum += b.Numbers[i][j]
			}
		}
	}
	return
}
