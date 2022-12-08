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

	shortcuts := [40]map[[2]byte]map[byte]int{}
	shortcuts[0] = map[[2]byte]map[byte]int{}
	for from, t1 := range transformations {
		for to, n := range t1 {
			shortcuts[0][[2]byte{from, to}] = _slice.CountUnique([]byte{from, to, n})
		}
	}

	for i := 1; i < 40; i++ {
		shortcuts[i] = map[[2]byte]map[byte]int{}
		for from, t1 := range transformations {
			for to, n := range t1 {
				p1, ok := shortcuts[i-1][[2]byte{from, n}]
				if !ok {
					p1 = _slice.CountUnique([]byte{from, n})
				}

				p2, ok := shortcuts[i-1][[2]byte{n, to}]
				if !ok {
					p2 = _slice.CountUnique([]byte{n, to})
				}

				p := mapSum(p1, p2)
				p[n]--

				shortcuts[i][[2]byte{from, to}] = p
			}
		}
	}

	ctsMap := map[byte]int{}
	for j := 0; j < len(chain)-1; j++ {
		p, ok := shortcuts[39][[2]byte{chain[j], chain[j+1]}]
		if !ok {
			p = _slice.CountUnique([]byte{chain[j], chain[j+1]})
		}
		ctsMap = mapSum(ctsMap, p)

		if j != 0 {
			ctsMap[chain[j]]--
		}
	}

	cts := _map.Values(ctsMap)
	sort.Ints(cts)

	fmt.Println(cts[len(cts)-1] - cts[0])
}

func mapSum(a, b map[byte]int) map[byte]int {
	s := _map.Duplicate(a)
	for k, v := range b {
		s[k] += v
	}
	return s
}
