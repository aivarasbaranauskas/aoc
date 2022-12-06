package main

import (
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"io/ioutil"
	"log"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fullFileB, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	for i := range fullFileB {
		if i >= 14 {
			if _set.FromSlice(fullFileB[i-14:i]).Len() == 14 {
				fmt.Println(i)
				break
			}
		}
	}
}
