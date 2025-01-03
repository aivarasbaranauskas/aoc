package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	r := bufio.NewScanner(f)
	var sum int
	for r.Scan() {
		line := r.Text()
		items := []byte(strings.TrimSpace(line))
		l := len(items)
		intr := _slice.Intersect(items[:l/2], items[l/2:])
		item := intr[0]
		if int(item) >= int('a') {
			sum += int(item) - int('a') + 1
		} else {
			sum += int(item) - int('A') + 27
		}
	}

	fmt.Println(sum)
}
