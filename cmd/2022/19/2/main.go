package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	var bl []Blueprint

	r := bufio.NewScanner(f)
	for r.Scan() {
		spl := strings.Split(r.Text(), " ")
		bl = append(
			bl,
			Blueprint{
				oreRobot:           optimistic.Atoi(spl[6]),
				clayRobot:          optimistic.Atoi(spl[12]),
				obsidianRobotOre:   optimistic.Atoi(spl[18]),
				obsidianRobotClay:  optimistic.Atoi(spl[21]),
				geodeRobotOre:      optimistic.Atoi(spl[27]),
				geodeRobotObsidian: optimistic.Atoi(spl[30]),
				maxOreC: max(
					optimistic.Atoi(spl[6]),
					optimistic.Atoi(spl[12]),
					optimistic.Atoi(spl[18]),
					optimistic.Atoi(spl[27]),
				),
			},
		)
	}

	out := 1
	for i, b := range bl {
		s := simulate(state{OreR: 1}, 32, b, make(map[state]int))
		out *= s
		fmt.Println(i, s)
	}
	fmt.Println(out)
}

type Blueprint struct {
	oreRobot                            int
	clayRobot                           int
	obsidianRobotOre, obsidianRobotClay int
	geodeRobotOre, geodeRobotObsidian   int
	maxOreC                             int
}

type state struct {
	Ore, Clay, Obsidian, Geode     int
	OreR, ClayR, ObsidianR, GeodeR int
	T                              int
}

func simulate(s state, t int, b Blueprint, cache map[state]int) int {
	if s.T == t {
		return s.Geode
	}

	s.OreR = min(s.OreR, b.maxOreC)
	s.ClayR = min(s.ClayR, b.obsidianRobotClay)
	s.ObsidianR = min(s.ObsidianR, b.geodeRobotObsidian)
	s.Ore = min(s.Ore, (t-s.T)*b.maxOreC-s.OreR*(t-s.T-1))
	s.Clay = min(s.Clay, (t-s.T)*b.obsidianRobotClay-s.ClayR*(t-s.T-1))
	s.Obsidian = min(s.Obsidian, (t-s.T)*b.geodeRobotObsidian-s.ObsidianR*(t-s.T-1))
	if v, ok := cache[s]; ok {
		return v
	}

	maxVal := 0

	if s.Ore >= b.oreRobot && b.maxOreC > s.OreR {
		maxVal = max(
			maxVal,
			simulate(
				state{
					Ore:       s.Ore + s.OreR - b.oreRobot,
					Clay:      s.Clay + s.ClayR,
					Obsidian:  s.Obsidian + s.ObsidianR,
					Geode:     s.Geode + s.GeodeR,
					OreR:      s.OreR + 1,
					ClayR:     s.ClayR,
					ObsidianR: s.ObsidianR,
					GeodeR:    s.GeodeR,
					T:         s.T + 1,
				},
				t,
				b,
				cache,
			),
		)
	}

	if s.Ore >= b.clayRobot && b.obsidianRobotClay > s.ClayR {
		maxVal = max(
			maxVal,
			simulate(
				state{
					Ore:       s.Ore + s.OreR - b.clayRobot,
					Clay:      s.Clay + s.ClayR,
					Obsidian:  s.Obsidian + s.ObsidianR,
					Geode:     s.Geode + s.GeodeR,
					OreR:      s.OreR,
					ClayR:     s.ClayR + 1,
					ObsidianR: s.ObsidianR,
					GeodeR:    s.GeodeR,
					T:         s.T + 1,
				},
				t,
				b,
				cache,
			),
		)
	}

	if s.Ore >= b.obsidianRobotOre && s.Clay >= b.obsidianRobotClay && b.geodeRobotObsidian > s.ObsidianR {
		maxVal = max(
			maxVal,
			simulate(
				state{
					Ore:       s.Ore + s.OreR - b.obsidianRobotOre,
					Clay:      s.Clay + s.ClayR - b.obsidianRobotClay,
					Obsidian:  s.Obsidian + s.ObsidianR,
					Geode:     s.Geode + s.GeodeR,
					OreR:      s.OreR,
					ClayR:     s.ClayR,
					ObsidianR: s.ObsidianR + 1,
					GeodeR:    s.GeodeR,
					T:         s.T + 1,
				},
				t,
				b,
				cache,
			),
		)
	}

	if s.Ore >= b.geodeRobotOre && s.Obsidian >= b.geodeRobotObsidian {
		maxVal = max(
			maxVal,
			simulate(
				state{
					Ore:       s.Ore + s.OreR - b.geodeRobotOre,
					Clay:      s.Clay + s.ClayR,
					Obsidian:  s.Obsidian + s.ObsidianR - b.geodeRobotObsidian,
					Geode:     s.Geode + s.GeodeR,
					OreR:      s.OreR,
					ClayR:     s.ClayR,
					ObsidianR: s.ObsidianR,
					GeodeR:    s.GeodeR + 1,
					T:         s.T + 1,
				},
				t,
				b,
				cache,
			),
		)
	}

	maxVal = max(
		maxVal,
		simulate(
			state{
				Ore:       s.Ore + s.OreR,
				Clay:      s.Clay + s.ClayR,
				Obsidian:  s.Obsidian + s.ObsidianR,
				Geode:     s.Geode + s.GeodeR,
				OreR:      s.OreR,
				ClayR:     s.ClayR,
				ObsidianR: s.ObsidianR,
				GeodeR:    s.GeodeR,
				T:         s.T + 1,
			},
			t,
			b,
			cache,
		),
	)

	cache[s] = maxVal

	return maxVal
}
