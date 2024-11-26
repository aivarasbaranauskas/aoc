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
	var ct int

	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, ",")
		spl1 := strings.Split(spl[0], "-")
		spl2 := strings.Split(spl[1], "-")
		a := o.Atoi(spl1[0])
		b := o.Atoi(spl1[1])
		c := o.Atoi(spl2[0])
		d := o.Atoi(spl2[1])

		if (a <= c && c <= b) || (c <= a && a <= d) {
			ct++
		}
	}

	fmt.Println(ct)
}
