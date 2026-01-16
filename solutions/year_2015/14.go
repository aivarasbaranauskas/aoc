package year_2015

import (
	"bytes"
	"slices"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[14] = Day14{
		t: 2503,
	}
}

type Day14 struct {
	t int
}

func (day Day14) Part1(input []byte) string {
	m := day.parse(input)
	t := day.t
	maxD := 0

	for _, v := range m {
		maxD = max(maxD, day.posAtTime(v, t))
	}

	return strconv.Itoa(maxD)
}

func (day Day14) Part2(input []byte) string {
	m := day.parse(input)
	scores := make([]int, len(m))
	poses := make([]int, len(m))

	for t := 1; t <= day.t; t++ {
		for i, v := range m {
			poses[i] = day.posAtTime(v, t)
		}
		maxPos := slices.Max(poses)
		for i, v := range poses {
			if v == maxPos {
				scores[i]++
			}
		}
	}

	return strconv.Itoa(slices.Max(scores))
}

func (day Day14) posAtTime(v Day14S, t int) int {
	tInterval := v.time + v.rest
	sPerInterval := v.time * v.speed
	d := (t / tInterval) * sPerInterval
	leftT := t % tInterval

	if leftT >= v.time {
		d += sPerInterval
	} else {
		d += leftT * v.speed
	}
	return d
}

type Day14S struct {
	speed, time, rest int
}

func (day Day14) parse(input []byte) []Day14S {
	n := bytes.Count(input, []byte{'\n'}) + 1
	m := make([]Day14S, n)
	for i := 0; i < n; i++ {
		lineEnd := bytes.Index(input, []byte{'\n'})
		var line []byte
		if lineEnd == -1 {
			line = input
			input = nil
		} else {
			line = input[:lineEnd]
			input = input[lineEnd+1:]
		}

		flyIdx := bytes.Index(line, []byte("fly"))
		kmsIdx := bytes.Index(line, []byte("km/s"))
		secondsIdx := bytes.Index(line, []byte("seconds,"))

		speed := line[flyIdx+4 : kmsIdx-1]
		t := line[kmsIdx+9 : secondsIdx-1]
		rest := line[secondsIdx+32 : len(line)-9]

		m[i].speed = optimistic.AtoiBFast(speed)
		m[i].time = optimistic.AtoiBFast(t)
		m[i].rest = optimistic.AtoiBFast(rest)
	}

	return m
}
