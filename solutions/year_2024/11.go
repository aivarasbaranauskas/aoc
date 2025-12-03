package year_2024

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[11] = Day11{}
}

type Day11 struct{}

func (day Day11) Part1(input []byte) string {
	return day.solve(input, 25)
}

func (day Day11) Part2(input []byte) string {
	return day.solve(input, 75)
}

func (day Day11) solve(input []byte, n int) string {
	pebbles := _slice.Map(bytes.Split(input, []byte(" ")), optimistic.AtoiBFast)
	mem := map[[2]int]int{}
	ct := 0

	for _, pebble := range pebbles {
		ct += day.expand(mem, pebble, n)
	}

	return strconv.Itoa(ct)
}

func (day Day11) expand(mem map[[2]int]int, pebble, n int) int {
	pow10s := [...]int{
		1,
		10,
		100,
		1000,
		10000,
		100000,
		1000000,
		10000000,
		100000000,
		1000000000,
		10000000000,
		100000000000,
		1000000000000,
		10000000000000,
		100000000000000,
	}

	if n == 0 {
		return 1
	}

	p := [2]int{pebble, n}
	if ct, ok := mem[p]; ok {
		return ct
	}

	var ct int
	if pebble == 0 {
		ct = day.expand(mem, 1, n-1)
	} else {
		decimalsCount := day.countDecimals(pebble)
		if decimalsCount%2 == 0 {
			pow10 := pow10s[decimalsCount/2]
			ct = day.expand(mem, pebble%pow10, n-1) + day.expand(mem, pebble/pow10, n-1)
		} else {
			ct = day.expand(mem, pebble*2024, n-1)
		}
	}

	mem[p] = ct
	return ct
}

func (day Day11) countDecimals(x int) (ct int) {
	for x > 0 {
		x /= 10
		ct++
	}
	return
}
