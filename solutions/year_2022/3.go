package year_2022

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"strconv"
	"strings"
)

func init() {
	Solutions[3] = Day3{}
}

type Day3 struct{}

func (Day3) Part1(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var sum int
	for r.Scan() {
		line := r.Text()
		items := []byte(strings.TrimSpace(line))
		l := len(items)
		intr := _slice.Intersect(items[:l/2], items[l/2:])
		item := intr[0]
		if int(item) >= int('a') {
			sum += int(item) - int('a') + 1
		} else {
			sum += int(item) - int('A') + 27
		}
	}

	return strconv.Itoa(sum)
}

func (Day3) Part2(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var sum int
	var group [][]byte
	for r.Scan() {
		group = append(group, []byte(strings.TrimSpace(r.Text())))
		if len(group) < 3 {
			continue
		}

		intr := _slice.Intersect(group...)
		item := intr[0]
		if int(item) >= int('a') {
			sum += int(item) - int('a') + 1
		} else {
			sum += int(item) - int('A') + 27
		}
		group = group[:0]
	}

	return strconv.Itoa(sum)
}
