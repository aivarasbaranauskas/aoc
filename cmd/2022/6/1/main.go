package main

import (
	"embed"
	"fmt"
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
