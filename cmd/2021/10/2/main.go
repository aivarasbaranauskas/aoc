package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"sort"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	r := bufio.NewScanner(f)
	var scores []int
OuterLoop:
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, "")
		stack := _a.Stack[string]{}
		for _, s := range spl {
			if stack.Empty() || s == "(" || s == "[" || s == "{" || s == "<" {
				stack.Push(s)
				continue
			}
			last := stack.Pop()
			if last == "(" && s == ")" || last == "[" && s == "]" || last == "{" && s == "}" || last == "<" && s == ">" {
				continue
			}

			continue OuterLoop
		}

		if stack.Empty() {
			continue
		}

		var score int
		for !stack.Empty() {
			score *= 5
			switch stack.Pop() {
			case "(":
				score += 1
			case "[":
				score += 2
			case "{":
				score += 3
			case "<":
				score += 4
			}
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
