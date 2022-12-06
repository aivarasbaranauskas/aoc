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
		if i >= 14 {
			x := make(map[byte]struct{})
			for j := 0; j < 14; j++ {
				x[o[i-j]] = struct{}{}
			}
			if len(x) == 14 {
				break
			}
		}
	}

	fmt.Println(string(o))
	fmt.Println(len(o))
}
