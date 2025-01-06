package year_2022

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[10] = Day10{}
}

type Day10 struct{}

func (Day10) Part1(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	x := 1
	c := 1
	out := 0
	for r.Scan() {
		line := r.Text()
		if line == "noop" {
			//here
			if (c+20)%40 == 0 {
				out += c * x
			}
			c++
			continue
		}

		//here
		if (c+20)%40 == 0 {
			out += c * x
		}

		c++
		//here
		if (c+20)%40 == 0 {
			out += c * x
		}
		c++

		spl := strings.Split(line, " ")
		x += optimistic.Atoi(spl[1])
	}
	return strconv.Itoa(out)
}

func (Day10) Part2(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	x := 1
	c := 1
	var crt [240]bool
	for r.Scan() {
		line := r.Text()
		if line == "noop" {
			//here
			if x-1 <= c%40-1 && c%40-1 <= x+1 {
				crt[c-1] = true
			}
			c++
			continue
		}

		//here
		if x-1 <= c%40-1 && c%40-1 <= x+1 {
			crt[c-1] = true
		}

		c++
		//here
		if x-1 <= c%40-1 && c%40-1 <= x+1 {
			crt[c-1] = true
		}
		c++

		spl := strings.Split(line, " ")
		x += optimistic.Atoi(spl[1])
	}
	s := strings.Builder{}
	for i := 0; i <= 5; i++ {
		s.WriteByte('\n')
		for j := 0; j < 40; j++ {
			if crt[i*40+j] {
				s.WriteByte('#')
			} else {
				s.WriteByte(' ')
			}
		}
	}
	return s.String()
}
