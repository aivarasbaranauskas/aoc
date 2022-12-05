package main

import (
	"bufio"
	"embed"
	"fmt"
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

	r := bufio.NewScanner(f)
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

	fmt.Println(score)
}
