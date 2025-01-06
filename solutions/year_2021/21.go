package year_2021

import (
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[21] = Day21{}
}

type Day21 struct{}

func (Day21) Part1(input []byte) string {
	spl := _slice.Map(strings.Split(string(input), "\n"), func(tin string) int {
		return optimistic.Atoi(strings.Split(tin, ": ")[1])
	})
	p1, p2 := spl[0], spl[1]
	p1--
	p2--
	p1s, p2s := 0, 0
	i := 0
	ct := 0
	turn := 1
	for p1s < 1000 && p2s < 1000 {
		if turn == 1 {
			p1 += i + 1
			i = (i + 1) % 100
			p1 += i + 1
			i = (i + 1) % 100
			p1 += i + 1
			i = (i + 1) % 100
			p1 = p1 % 10
			p1s += p1 + 1
			turn = 2
		} else {
			p2 += i + 1
			i = (i + 1) % 100
			p2 += i + 1
			i = (i + 1) % 100
			p2 += i + 1
			i = (i + 1) % 100
			p2 = p2 % 10
			p2s += p2 + 1
			turn = 1
		}
		ct += 3
	}

	return strconv.Itoa(min(p1s, p2s) * ct)
}

func (d Day21) Part2(input []byte) string {
	spl := _slice.Map(strings.Split(string(input), "\n"), func(tin string) int {
		return optimistic.Atoi(strings.Split(tin, ": ")[1])
	})
	p1, p2 := spl[0], spl[1]
	p1--
	p2--

	p1w, p2w := d.doTurn(make(map[day21Entry][2]int), p1, p2, 0, 0, 1)

	return strconv.Itoa(max(p1w, p2w))
}

type day21Entry struct {
	p1, p2, p1s, p2s, turn int
}

func (d Day21) doTurn(cache map[day21Entry][2]int, p1, p2, p1s, p2s, turn int) (p1w, p2w int) {
	if p1s >= 21 {
		return 1, 0
	}
	if p2s >= 21 {
		return 0, 1
	}

	e := day21Entry{p1, p2, p1s, p2s, turn}
	if c, ok := cache[e]; ok {
		return c[0], c[1]
	}

	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			for c := 1; c <= 3; c++ {
				p1t, p2t := p1, p2
				if turn == 1 {
					p1t = (p1t + a + b + c) % 10
					p1wt, p2wt := d.doTurn(cache, p1t, p2t, p1s+p1t+1, p2s, 2)
					p1w += p1wt
					p2w += p2wt
				} else {
					p2t = (p2t + a + b + c) % 10
					p1wt, p2wt := d.doTurn(cache, p1t, p2t, p1s, p2s+p2t+1, 1)
					p1w += p1wt
					p2w += p2wt
				}
			}
		}
	}

	cache[e] = [2]int{p1w, p2w}

	return
}
