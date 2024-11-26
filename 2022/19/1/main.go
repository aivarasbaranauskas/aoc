package main

import (
	"bufio"
	"embed"
	"encoding/json"
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

	var bl []Blueprint

	r := bufio.NewScanner(f)
	for r.Scan() {
		spl := strings.Split(r.Text(), " ")
		bl = append(
			bl,
			Blueprint{
				oreRobot:           o.Atoi(spl[6]),
				clayRobot:          o.Atoi(spl[12]),
				obsidianRobotOre:   o.Atoi(spl[18]),
				obsidianRobotClay:  o.Atoi(spl[21]),
				geodeRobotOre:      o.Atoi(spl[27]),
				geodeRobotObsidian: o.Atoi(spl[30]),
			},
		)
	}

	var out int
	for i, b := range bl {
		s := simulate(state{OreR: 1}, 24, b, make(map[state]int))
		out += (i + 1) * s
		fmt.Println(i, s)
	}
	fmt.Println(out)
}

type Blueprint struct {
	oreRobot                            int
	clayRobot                           int
	obsidianRobotOre, obsidianRobotClay int
	geodeRobotOre, geodeRobotObsidian   int
}

type state struct {
	Ore, Clay, Obsidian, Geode     int
	OreR, ClayR, ObsidianR, GeodeR int
	T                              int
}

func (s *state) String() string {
	v, _ := json.Marshal(s)
	return string(v)
}

func simulate(s state, t int, b Blueprint, cache map[state]int) int {
	if s.T == t {
		return s.Geode
	}

	if v, ok := cache[s]; ok {
		return v
	}

	maxV := 0

	if s.Ore >= b.oreRobot {
		maxV = max(
			maxV,
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

	if s.Ore >= b.clayRobot {
		maxV = max(
			maxV,
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

	if s.Ore >= b.obsidianRobotOre && s.Clay >= b.obsidianRobotClay {
		maxV = max(
			maxV,
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
		maxV = max(
			maxV,
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

	maxV = max(
		maxV,
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

	cache[s] = maxV

	return maxV
}
