package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strings"
	"sync/atomic"
	"time"
)

var pr *uint64

//go:embed input.txt
var input string

func main() {
	spl := _slice.Map(strings.Split(input, "\n"), func(tin string) int {
		return optimistic.Atoi(strings.Split(tin, ": ")[1])
	})
	p1, p2 := spl[0], spl[1]
	p1--
	p2--

	pr = new(uint64)
	go func() {
		for {
			fmt.Printf("%v\n", atomic.LoadUint64(pr))
			time.Sleep(time.Second)
		}
	}()

	cache = make(map[entry][2]int)
	p1w, p2w := doTurn(p1, p2, 0, 0, 1)

	fmt.Println(p1w, p2w)
}

type entry struct {
	p1, p2, p1s, p2s, turn int
}

var cache map[entry][2]int

func doTurn(p1, p2, p1s, p2s, turn int) (p1w, p2w int) {
	if p1s >= 21 {
		atomic.AddUint64(pr, 1)
		return 1, 0
	}
	if p2s >= 21 {
		atomic.AddUint64(pr, 1)
		return 0, 1
	}

	e := entry{p1, p2, p1s, p2s, turn}
	if c, ok := cache[e]; ok {
		return c[0], c[1]
	}

	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			for c := 1; c <= 3; c++ {
				p1t, p2t := p1, p2
				if turn == 1 {
					p1t = (p1t + a + b + c) % 10
					p1wt, p2wt := doTurn(p1t, p2t, p1s+p1t+1, p2s, 2)
					p1w += p1wt
					p2w += p2wt
				} else {
					p2t = (p2t + a + b + c) % 10
					p1wt, p2wt := doTurn(p1t, p2t, p1s, p2s+p2t+1, 1)
					p1w += p1wt
					p2w += p2wt
				}
			}
		}
	}

	cache[e] = [2]int{p1w, p2w}

	return
}
