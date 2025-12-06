package year_2025

import (
	"bytes"
	"strconv"
)

func init() {
	Solutions[6] = Day6{}
}

type Day6 struct{}

func (Day6) Part1(input []byte) string {
	lines := bytes.Split(input, []byte{'\n'})

	tmpLine := bytes.TrimSpace(lines[0])
	colCount := 1
	for i := range len(tmpLine) - 1 {
		if tmpLine[i] == ' ' && tmpLine[i+1] != ' ' {
			colCount++
		}
	}

	m := make([][]int, len(lines)-1)
	for i, line := range lines[:len(lines)-1] {
		m[i] = make([]int, colCount)
		lineI := 0
		for ; lineI < len(line) && line[lineI] == ' '; lineI++ {
		}

		for j := 0; j < colCount; j++ {
			for ; lineI < len(line) && line[lineI] != ' '; lineI++ {
				m[i][j] = m[i][j]*10 + int(line[lineI]-'0')
			}
			for ; lineI < len(line) && line[lineI] == ' '; lineI++ {
			}
		}
	}

	sum := 0
	signs := bytes.ReplaceAll(lines[len(lines)-1], []byte{' '}, []byte{})
	for i, sign := range signs {
		if sign == '+' {
			for j := 0; j < len(lines)-1; j++ {
				sum += m[j][i]
			}
		} else {
			rez := m[0][i]
			for j := 1; j < len(lines)-1; j++ {
				rez *= m[j][i]
			}
			sum += rez
		}
	}

	return strconv.Itoa(sum)
}

func (Day6) Part2(input []byte) string {
	lines := bytes.Split(input, []byte{'\n'})
	sum := 0

	for i := 0; i < len(lines[0]); {
		if lines[len(lines)-1][i] == '+' {
			lines[len(lines)-1][i] = ' '
			for ; i < len(lines[0])-1 && lines[len(lines)-1][i] == ' '; i++ {
				num := 0

				for j := 0; j < len(lines)-1; j++ {
					if lines[j][i] != ' ' {
						num = num*10 + int(lines[j][i]-'0')
					}
				}
				sum += num
			}
		} else {
			lines[len(lines)-1][i] = ' '
			mul := 1
			for ; i < len(lines[0]) && lines[len(lines)-1][i] == ' '; i++ {
				num := 0

				for j := 0; j < len(lines)-1; j++ {
					if lines[j][i] != ' ' {
						num = num*10 + int(lines[j][i]-'0')
					}
				}

				if num != 0 {
					mul *= num
				}
			}
			sum += mul
		}
	}

	return strconv.Itoa(sum)
}
