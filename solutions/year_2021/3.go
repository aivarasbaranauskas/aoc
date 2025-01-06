package year_2021

import (
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[3] = Day3{}
}

type Day3 struct{}

func (Day3) Part1(input []byte) string {
	lines := strings.Split(string(input), "\n")
	m := make([]uint64, len(lines))
	for i, line := range lines {
		m[i] = uint64(optimistic.ParseInt(line, 2, 64))
	}

	gamma := 0
	l := 12

	onesC := make([]int, l)
	zeroesC := make([]int, l)
	for _, ip := range m {
		for p := 0; p < l; p++ {
			f := uint64(1 << p)
			if ip&f == f {
				onesC[p]++
			} else {
				zeroesC[p]++
			}
		}
	}

	for i := 0; i < l; i++ {
		if onesC[i] > zeroesC[i] {
			gamma += 1 << i
		}
	}

	epsilon := ^gamma & 0b111111111111
	powerConsumption := gamma * epsilon
	return strconv.Itoa(powerConsumption)
}

func (Day3) Part2(input []byte) string {
	lines := strings.Split(string(input), "\n")
	m := make([]uint64, len(lines))
	for i, line := range lines {
		m[i] = uint64(optimistic.ParseInt(line, 2, 64))
	}

	getRating := func(input []uint64, l int, checkF func(onesC, zeroesC int) bool) uint64 {
		for p := l - 1; p >= 0; p-- {
			if len(input) == 1 {
				break
			}

			mask := uint64(1 << p)
			onesC := 0
			zeroesC := 0
			for _, ip := range input {
				if ip&mask == mask {
					onesC++
				} else {
					zeroesC++
				}
			}

			doPut1 := checkF(onesC, zeroesC)

			var newI []uint64
			for i := range input {
				if (input[i]&mask == mask) == doPut1 {
					newI = append(newI, input[i])
				}
			}
			input = newI
		}

		if len(input) > 1 {
			panic(fmt.Sprint(input))
		}

		return input[0]
	}

	oxygen := getRating(_slice.Duplicate(m), 12, func(onesC, zeroesC int) bool { return onesC >= zeroesC })
	co2 := getRating(_slice.Duplicate(m), 12, func(onesC, zeroesC int) bool { return onesC < zeroesC })

	lifeSupportRating := oxygen * co2

	return strconv.Itoa(int(lifeSupportRating))
}
