package year_2022

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"math"
	"strconv"
	"strings"
)

func init() {
	Solutions[16] = Day16{}
}

type Day16 struct{}

func (d Day16) Part1(input []byte) string {
	rates := make(map[string]int)
	tunnels := make(map[string][]string)
	pathL := make(map[[2]string]int)

	r := bufio.NewScanner(bytes.NewReader(input))
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

	return strconv.Itoa(d.walk(30, 0, "AA", _set.Set[string]{}, rates, pathL, tunnels))
}

func (d Day16) Part2(input []byte) string {
	rates := make(map[string]int)
	tunnels := make(map[string][]string)
	pathL := make(map[[2]string]int)

	r := bufio.NewScanner(bytes.NewReader(input))
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

	return strconv.Itoa(d.walk2(26, 26, 0, "AA", "AA", _set.Set[string]{}, rates, pathL, tunnels))
}

func (d Day16) walk(
	t int,
	released int,
	current string,
	open _set.Set[string],
	rates map[string]int,
	pathL map[[2]string]int,
	tunnels map[string][]string,
) int {
	if t == 0 || open.Len() == len(rates) {
		return released
	}

	maxRel := released

	for next, rate := range rates {
		if open.Has(next) {
			continue
		}

		l, ok := pathL[[2]string{current, next}]
		if !ok {
			l = d.findPath(current, next, nil, tunnels)
			pathL[[2]string{current, next}] = l
			pathL[[2]string{next, current}] = l
		}
		timeLeft := t - l - 1
		if timeLeft >= 0 {
			open.Add(next)
			maxRel = max(maxRel, d.walk(timeLeft, released+rate*timeLeft, next, open, rates, pathL, tunnels))
			open.Remove(next)
		}
	}

	return maxRel
}

func (d Day16) findPath(from, to string, visited []string, tunnels map[string][]string) int {
	minVal := math.MaxInt - 100
L:
	for _, next := range tunnels[from] {
		if next == to {
			minVal = 0
			break
		}

		for _, v := range visited {
			if next == v {
				continue L
			}
		}

		minVal = min(
			minVal,
			d.findPath(next, to, append(visited, from), tunnels),
		)
	}
	return minVal + 1
}

func (d Day16) walk2(
	t, tE, released int,
	current, currentE string,
	open _set.Set[string],
	rates map[string]int,
	pathL map[[2]string]int,
	tunnels map[string][]string,
) int {
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
				l = d.findPath(current, next, nil, tunnels)
				pathL[[2]string{current, next}] = l
				pathL[[2]string{next, current}] = l
			}
			timeLeft := t - l - 1
			if timeLeft >= 0 {
				open.Add(next)
				rel := d.walk2(timeLeft, tE, released+rate*timeLeft, next, currentE, open, rates, pathL, tunnels)
				if maxRel < rel {
					maxRel = rel
				}
				open.Remove(next)
			}
		} else {
			l, ok := pathL[[2]string{currentE, next}]
			if !ok {
				l = d.findPath(currentE, next, nil, tunnels)
				pathL[[2]string{currentE, next}] = l
				pathL[[2]string{next, currentE}] = l
			}
			timeLeft := tE - l - 1
			if timeLeft >= 0 {
				open.Add(next)
				rel := d.walk2(t, timeLeft, released+rate*timeLeft, current, next, open, rates, pathL, tunnels)
				if maxRel < rel {
					maxRel = rel
				}
				open.Remove(next)
			}
		}
	}

	return maxRel
}
