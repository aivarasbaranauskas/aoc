package year_2022

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

func init() {
	Solutions[2] = Day2{}
}

type Day2 struct{}

func (Day2) Part1(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var score int
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " ")

		switch spl[1] {
		case "X":
			score += 1
		case "Y":
			score += 2
		case "Z":
			score += 3
		}

		switch line {
		case "A X", "B Y", "C Z":
			score += 3
		case "A Z", "B X", "C Y":
			score += 0
		case "A Y", "B Z", "C X":
			score += 6
		}
	}

	return strconv.Itoa(score)
}

func (Day2) Part2(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var score int
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " ")

		switch spl[1] {
		case "X":
			score += 0

			switch spl[0] {
			case "A":
				score += 3 //Z
			case "B":
				score += 1 //X
			case "C":
				score += 2 //Y
			}
		case "Y":
			score += 3

			switch spl[0] {
			case "A":
				score += 1 //X
			case "B":
				score += 2 //Y
			case "C":
				score += 3 //Z
			}
		case "Z":
			score += 6

			switch spl[0] {
			case "A":
				score += 2 //Y
			case "B":
				score += 3 //Z
			case "C":
				score += 1 //X
			}
		}
	}

	return strconv.Itoa(score)
}
