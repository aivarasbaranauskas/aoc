package year_2023

import (
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_map"
	"strconv"
	"strings"
)

func init() {
	Solutions[20] = Day20{}
}

type Day20 struct{}

func (Day20) Part1(input []byte) string {
	lines := strings.Split(string(input), "\n")

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

	return strconv.Itoa(ctLow * ctHigh)
}

func (Day20) Part2(input []byte) string {
	lines := strings.Split(string(input), "\n")

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

	flipped := map[string]bool{}
	pushButton := func() {
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

	for i, v := range flippedAt {
		flippedAt[i] = (v-1)/2 + 1
	}
	return strconv.Itoa(_a.LCM(_map.Values(flippedAt)...))
}

type Transfer struct {
	source, dest string
	signal       bool
}
