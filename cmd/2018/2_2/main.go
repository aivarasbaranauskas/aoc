package main

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func stringDiffer(str1 string, str2 string) (int, string) {
	diffCounter := 0
	solution := ""
	for i := range str1 {
		if str1[i] == str2[i] {
			solution += string(str1[i])
		} else {
			diffCounter++
		}
	}
	return diffCounter, solution
}

func main() {
	file, _ := inputData.ReadFile("input.txt")
	answer2 := 0
	answer3 := 0
	// Solution to first part
	for _, line := range strings.Split(string(file), "\n") {
		abcCounter := make(map[rune]int)
		val2 := false
		val3 := false
		for _, char := range line {
			if _, exists := abcCounter[char]; exists {
				abcCounter[char]++
			} else {
				abcCounter[char] = 1
			}
		}
		for _, value := range abcCounter {
			if value == 2 && val2 == false {
				answer2++
				val2 = true
			} else if value == 3 && val3 == false {
				answer3++
				val3 = true
			}
		}
		fmt.Println(val2, val3)
	}
	fmt.Println(answer2 * answer3)
	// Solution to first part
	var abc []string
	diff := 0
	solution := ""
	for _, line := range strings.Split(string(file), "\n") {
		abc = append(abc, line)
	}
	for i := range abc {
		for j := range abc {
			diff, solution = stringDiffer(abc[i], abc[j])
			if diff == 1 {
				fmt.Println(solution)
			}
		}
	}
}
