package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/go_helpers/o"
	"log"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	r := bufio.NewScanner(f)
	x := 1
	c := 1
	out := 0
	for r.Scan() {
		line := r.Text()
		if line == "noop" {
			//here
			if (c+20)%40 == 0 {
				out += c * x
			}
			c++
			continue
		}

		//here
		if (c+20)%40 == 0 {
			out += c * x
		}

		c++
		//here
		if (c+20)%40 == 0 {
			out += c * x
		}
		c++

		spl := strings.Split(line, " ")
		x += o.Atoi(spl[1])
	}
	fmt.Println(out)
}
