package year_2022

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[4] = Day4{}
}

type Day4 struct{}

func (Day4) Part1(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var ct int

	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, ",")
		spl1 := strings.Split(spl[0], "-")
		spl2 := strings.Split(spl[1], "-")
		a := optimistic.Atoi(spl1[0])
		b := optimistic.Atoi(spl1[1])
		c := optimistic.Atoi(spl2[0])
		d := optimistic.Atoi(spl2[1])

		if (a <= c && b >= d) || (c <= a && d >= b) {
			ct++
		}
	}

	return strconv.Itoa(ct)
}

func (Day4) Part2(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var ct int

	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, ",")
		spl1 := strings.Split(spl[0], "-")
		spl2 := strings.Split(spl[1], "-")
		a := optimistic.Atoi(spl1[0])
		b := optimistic.Atoi(spl1[1])
		c := optimistic.Atoi(spl2[0])
		d := optimistic.Atoi(spl2[1])

		if (a <= c && c <= b) || (c <= a && a <= d) {
			ct++
		}
	}

	return strconv.Itoa(ct)
}
