package main

import (
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"io"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)
	fullFileB, err := io.ReadAll(f)
	_a.CheckErr(err)

	var o []byte
	for i, c := range fullFileB {
		o = append(o, c)
		if i >= 3 && o[i] != o[i-1] && o[i] != o[i-2] && o[i] != o[i-3] && o[i-1] != o[i-2] && o[i-1] != o[i-3] && o[i-2] != o[i-3] {
			break
		}
	}

	fmt.Println(string(o))
	fmt.Println(len(o))
}
