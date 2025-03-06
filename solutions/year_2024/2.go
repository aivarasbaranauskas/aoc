package year_2024

import (
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
)

func init() {
	Solutions[2] = Day2{}
}

type Day2 struct{}

func (day Day2) Part1(input []byte) string {
	return day.solve(input, day.isRowOk)
}

func (day Day2) Part2(input []byte) string {
	return day.solve(input, day.isRowOk2)
}

func (day Day2) solve(input []byte, checkF func([]int) bool) string {
	ct := 0

	var numbers [8]int
	checkRow := func(row []byte) {
		b := 0
		i := 0
		for e, v := range row {
			if v == ' ' {
				numbers[i] = optimistic.AtoiBFast(row[b:e])
				b = e + 1
				i++
			}
		}
		numbers[i] = optimistic.AtoiBFast(row[b:])
		i++
		if checkF(numbers[:i]) {
			ct++
		}
	}

	b := 0
	for e, v := range input {
		if v == '\n' {
			checkRow(input[b:e])
			b = e + 1
		}
	}
	checkRow(input[b:])

	// 1 allocation
	return strconv.Itoa(ct)
}

func (day Day2) isRowOk(row []int) bool {
	direction := row[0] < row[1]
	for i := range len(row) - 1 {
		diff := _num.Abs(row[i] - row[i+1])
		if diff < 1 || diff > 3 || direction != (row[i] < row[i+1]) {
			return false
		}
	}
	return true
}

func (day Day2) isRowOk2(row []int) bool {
	if day.isRowOk(row) {
		return true
	}

	for i := range row {
		tmp := row[i]
		copy(row[i:], row[i+1:])
		if day.isRowOk(row[:len(row)-1]) {
			return true
		}
		copy(row[i+1:], row[i:])
		row[i] = tmp
	}
	return false
}
