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

	r := bufio.NewScanner(f)
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
	fmt.Println(out)
}
