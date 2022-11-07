package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput(file string) [][]byte {
	f, err := os.Open(file)
	check(err)
	scanner := bufio.NewScanner(f)

	res := []([]byte){}
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
