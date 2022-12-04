package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
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
		a := optimistic.Atoi(spl1[0])
		b := optimistic.Atoi(spl1[1])
		c := optimistic.Atoi(spl2[0])
		d := optimistic.Atoi(spl2[1])

		if (a <= c && c <= b) || (c <= a && a <= d) {
			ct++
		}
	}

	fmt.Println(ct)
}
