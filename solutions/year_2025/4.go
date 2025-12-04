package year_2025

import (
	"bytes"
	"strconv"
)

func init() {
	Solutions[4] = Day4{}
}

type Day4 struct{}

func (day Day4) Part1(input []byte) string {
	rowLength := bytes.IndexByte(input, '\n') + 1
	rowCount := len(input)/rowLength + 1

	ct := 0

	for i := 0; i < len(input); i++ {
		if input[i] != '@' {
			continue
		}

		if day.canRemove(input, rowLength, rowCount, i) {
			ct++
		}
	}

	return strconv.Itoa(ct)
}

func (day Day4) Part2(input []byte) string {
	rowLength := bytes.IndexByte(input, '\n') + 1
	rowCount := len(input)/rowLength + 1

	ct := 0
	prevCt := -1

	for prevCt != ct {
		prevCt = ct

		for i := 0; i < len(input); i++ {
			if input[i] != '@' {
				continue
			}

			if day.canRemove(input, rowLength, rowCount, i) {
				input[i] = '.'
				ct++
			}
		}
	}

	return strconv.Itoa(ct)
}

func (Day4) canRemove(input []byte, rowLength, rowCount, i int) bool {
	row := i / rowLength
	col := i % rowLength

	adjRolls := 0

	if col > 0 && input[row*rowLength+col-1] == '@' {
		adjRolls += 1
	}
	if col+2 < rowLength && input[row*rowLength+col+1] == '@' {
		adjRolls += 1
	}
	if row > 0 && input[(row-1)*rowLength+col] == '@' {
		adjRolls += 1
	}
	if row+1 < rowCount && input[(row+1)*rowLength+col] == '@' {
		adjRolls += 1
	}

	if col > 0 && row > 0 && input[(row-1)*rowLength+col-1] == '@' {
		adjRolls += 1
	}
	if col+2 < rowLength && row > 0 && input[(row-1)*rowLength+col+1] == '@' {
		adjRolls += 1
	}
	if col > 0 && row+1 < rowCount && input[(row+1)*rowLength+col-1] == '@' {
		adjRolls += 1
	}
	if col+2 < rowLength && row+1 < rowCount && input[(row+1)*rowLength+col+1] == '@' {
		adjRolls += 1
	}

	return adjRolls < 4
}
