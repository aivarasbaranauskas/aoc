package main

import (
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"io"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)
	fullFileB, err := io.ReadAll(f)
	_a.CheckErr(err)

	for i := range fullFileB {
		if i >= 14 {
			if _set.FromSlice(fullFileB[i-14:i]).Len() == 14 {
				fmt.Println(i)
				break
			}
		}
	}
}
