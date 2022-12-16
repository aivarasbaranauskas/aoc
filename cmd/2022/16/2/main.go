package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"log"
	"math"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

var (
	rates   map[string]int
	tunnels map[string][]string
	pathL   map[[2]string]int
)

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	rates = make(map[string]int)
	tunnels = make(map[string][]string)
	pathL = make(map[[2]string]int)

	r := bufio.NewScanner(f)
	for r.Scan() {
		spl := strings.Split(r.Text(), "; ")
		spl1 := strings.Split(spl[0], " ")
		name := spl1[1]
		rate := optimistic.Atoi(strings.Split(spl1[4], "=")[1])
		if rate > 0 {
			rates[name] = rate
		}

		spl2 := strings.Split(spl[1], " ")
		var tt []string
		for i := 4; i < len(spl2); i++ {
			tt = append(tt, strings.Trim(spl2[i], ","))
		}
		tunnels[name] = tt
	}

	fmt.Println(walk(26, 26, 0, "AA", "AA", _set.Set[string]{}))
}

func walk(t, tE, released int, current, currentE string, open _set.Set[string]) int {
	if (t == 0 && tE == 0) || open.Len() == len(rates) {
		return released
	}

	maxRel := released

	for next, rate := range rates {
		if open.Has(next) {
			continue
		}

		if t > tE {
			l, ok := pathL[[2]string{current, next}]
			if !ok {
				l = findPath(current, next, nil)
				pathL[[2]string{current, next}] = l
				pathL[[2]string{next, current}] = l
			}
			timeLeft := t - l - 1
			if timeLeft >= 0 {
				open.Add(next)
				rel := walk(timeLeft, tE, released+rate*timeLeft, next, currentE, open)
				if maxRel < rel {
					maxRel = rel
				}
				open.Remove(next)
			}
		} else {
			l, ok := pathL[[2]string{currentE, next}]
			if !ok {
				l = findPath(currentE, next, nil)
				pathL[[2]string{currentE, next}] = l
				pathL[[2]string{next, currentE}] = l
			}
			timeLeft := tE - l - 1
			if timeLeft >= 0 {
				open.Add(next)
				rel := walk(t, timeLeft, released+rate*timeLeft, current, next, open)
				if maxRel < rel {
					maxRel = rel
				}
				open.Remove(next)
			}
		}
	}

	return maxRel
}

func findPath(from, to string, visited []string) int {
	min := math.MaxInt - 100
L:
	for _, next := range tunnels[from] {
		if next == to {
			min = 0
			break
		}

		for _, v := range visited {
			if next == v {
				continue L
			}
		}

		min = _num.Min(
			min,
			findPath(next, to, append(visited, from)),
		)
	}
	return min + 1
}
