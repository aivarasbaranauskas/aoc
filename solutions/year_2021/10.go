package year_2021

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"sort"
	"strconv"
	"strings"
)

func init() {
	Solutions[10] = Day10{}
}

type Day10 struct{}

func (Day10) Part1(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var out int
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

			switch s {
			case ")":
				out += 3
			case "]":
				out += 57
			case "}":
				out += 1197
			case ">":
				out += 25137
			}
		}
	}
	return strconv.Itoa(out)
}

func (Day10) Part2(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
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
	return strconv.Itoa(scores[len(scores)/2])
}
