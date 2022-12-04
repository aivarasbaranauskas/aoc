package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/a"
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
	var sum int
	var group [][]byte
	for r.Scan() {
		group = append(group, []byte(strings.TrimSpace(r.Text())))
		if len(group) < 3 {
			continue
		}

		intr := a.Intersect(group...)
		item := intr[0]
		if int(item) >= int('a') {
			sum += int(item) - int('a') + 1
		} else {
			sum += int(item) - int('A') + 27
		}
		group = group[:0]
	}

	fmt.Println(sum)
}
