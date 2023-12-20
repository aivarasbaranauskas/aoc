package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_map"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	memCon := make(map[string]map[string]bool)
	memFlip := make(map[string]bool)
	modulesDest := make(map[string][]string)
	modulesTypes := make(map[string]byte)

	var start []string

	for _, line := range lines {
		spl := strings.Split(line, " -> ")
		if spl[0] == "broadcaster" {
			start = strings.Split(spl[1], ", ")
			continue
		}

		_type := spl[0][0]
		name := spl[0][1:]
		modulesTypes[name] = _type
		modulesDest[name] = strings.Split(spl[1], ", ")

		if _type == '%' {
			memFlip[name] = false
		}
		if _type == '&' {
			memCon[name] = make(map[string]bool)
		}
	}

	for name, dests := range modulesDest {
		for _, dest := range dests {
			if modulesTypes[dest] == '&' {
				memCon[dest][name] = false
			}
		}
	}

	ctLow := 0
	ctHigh := 0
	flipped := map[string]bool{}
	pushButton := func() {
		ctLow++
		q := _a.Queue[Transfer]{}
		for _, name := range start {
			q.Enqueue(Transfer{
				source: "broadcast",
				dest:   name,
				signal: false,
			})
		}

		for !q.Empty() {
			s := q.Dequeue()

			if s.dest == "dd" && s.signal {
				flipped[s.source] = true
			}

			if s.signal {
				ctHigh++
			} else {
				ctLow++
			}

			switch modulesTypes[s.dest] {
			case '%':
				if !s.signal {
					newSignal := !memFlip[s.dest]
					memFlip[s.dest] = newSignal
					for _, dest := range modulesDest[s.dest] {
						q.Enqueue(Transfer{
							source: s.dest,
							dest:   dest,
							signal: newSignal,
						})
					}
				}
			case '&':
				memCon[s.dest][s.source] = s.signal
				allHigh := true
				for _, st := range memCon[s.dest] {
					allHigh = allHigh && st
				}
				for _, dest := range modulesDest[s.dest] {
					q.Enqueue(Transfer{
						source: s.dest,
						dest:   dest,
						signal: !allHigh,
					})
				}
			}
		}
	}

	for i := 0; i < 1000; i++ {
		pushButton()
	}

	fmt.Println("part 1:", ctLow, ctHigh, ctLow*ctHigh)

	// reset
	for name := range memFlip {
		memFlip[name] = false
	}
	for name := range memCon {
		for name2 := range memCon[name] {
			memCon[name][name2] = false
		}
	}

	flippedAt := make(map[string]int)
	for i := 0; len(flippedAt) < len(memCon["dd"]); i++ {
		pushButton()
		i++
		if len(flipped) > 0 {
			for n := range flipped {
				if _, ok := flippedAt[n]; !ok {
					flippedAt[n] = i
				}
			}
		}
		flipped = map[string]bool{}
	}
	fmt.Println("part 2:", _a.LCM(_map.Values(flippedAt)...))
}

type Transfer struct {
	source, dest string
	signal       bool
}
