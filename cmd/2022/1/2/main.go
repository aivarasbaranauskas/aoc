package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"log"
	"sort"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var (
		x int
		s []int
	)

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Text()
		if line == "" {
			s = append(s, x)
			x = 0
		} else {
			x += optimistic.Atoi(line)
		}
	}
	s = append(s, x)

	sort.Ints(s)
	l := len(s)

	fmt.Println("Max3:", s[l-3]+s[l-2]+s[l-1])
	fmt.Println(s[l-3:])
	fmt.Println(s)
}
