package main

import (
	"bufio"
	"embed"
	"fmt"

	"github.com/aivarasbaranauskas/aoc/internal/_a"
)

//go:embed input.txt
var inputData embed.FS

func readInput(file string) [][]byte {
	f, err := inputData.Open(file)
	_a.CheckErr(err)
	scanner := bufio.NewScanner(f)

	var res [][]byte
	for scanner.Scan() {
		res = append(res, scanner.Bytes())
	}

	return res
}

func countLetters(ID string) (hasDoubles bool, hasTriples bool) {
	hasDoubles = false
	hasTriples = false

	buckets := make(map[rune]int)

	for _, letter := range ID {
		if _, ok := buckets[letter]; ok {
			buckets[letter]++
		} else {
			buckets[letter] = 1
		}
	}

	for _, ct := range buckets {
		if ct == 2 {
			hasDoubles = true
		}
		if ct == 3 {
			hasTriples = true
		}
	}

	return hasDoubles, hasTriples
}

func main() {
	data := readInput("input.txt")

	doubles := 0
	triples := 0
	for _, line := range data {
		hasDoubles, hasTriples := countLetters(string(line[:]))
		if hasDoubles {
			doubles++
		}
		if hasTriples {
			triples++
		}
		fmt.Println(hasDoubles, hasTriples)
	}

	fmt.Println(len(data))
	fmt.Println("doubles:", doubles, "triples:", triples)
	fmt.Println("Checksum:", doubles*triples)
}
