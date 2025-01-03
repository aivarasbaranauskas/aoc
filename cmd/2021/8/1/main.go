package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	var count int

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " | ")
		spl = strings.Split(spl[1], " ")

		for _, s := range spl {
			l := len(s)
			switch l {
			case 2, 4, 3, 7:
				count++
			default:
				// none
			}
		}
	}

	fmt.Println(count)
}
