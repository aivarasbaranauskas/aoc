package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_map"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"log"
	"sort"
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

	r.Scan()
	chain := []byte(r.Text())
	r.Scan()

	transformations := map[byte]map[byte]byte{}

	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " -> ")
		from := spl[0][0]
		to := spl[0][1]
		n := []byte(spl[1])[0]
		if t, ok := transformations[from]; ok {
			t[to] = n
		} else {
			transformations[from] = map[byte]byte{to: n}
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < len(chain)-1; j++ {
			if n, ok := transformations[chain[j]][chain[j+1]]; ok {
				chain = append(chain[:j+1], append([]byte{n}, chain[j+1:]...)...)
				j++
			}
		}
	}

	cts := _map.Values(_slice.CountUnique(chain))
	sort.Ints(cts)

	fmt.Println(cts[len(cts)-1] - cts[0])
}
