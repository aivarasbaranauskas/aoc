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

	fmt.Println(score)
}
