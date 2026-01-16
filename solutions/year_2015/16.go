package year_2015

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[16] = Day16{
		expected: Day16S{
			children:    _a.Ptr(3),
			cats:        _a.Ptr(7),
			samoyeds:    _a.Ptr(2),
			pomeranians: _a.Ptr(3),
			akitas:      _a.Ptr(0),
			vizslas:     _a.Ptr(0),
			goldfish:    _a.Ptr(5),
			trees:       _a.Ptr(3),
			cars:        _a.Ptr(2),
			perfumes:    _a.Ptr(1),
		},
	}
}

type Day16 struct {
	expected Day16S
}

func (day Day16) Part1(input []byte) string {
	m := day.parse(input)

	for i, v := range m {
		if v.children != nil && *(v.children) != *(day.expected.children) {
			continue
		}
		if v.cats != nil && *(v.cats) != *(day.expected.cats) {
			continue
		}
		if v.samoyeds != nil && *(v.samoyeds) != *(day.expected.samoyeds) {
			continue
		}
		if v.pomeranians != nil && *(v.pomeranians) != *(day.expected.pomeranians) {
			continue
		}
		if v.akitas != nil && *(v.akitas) != *(day.expected.akitas) {
			continue
		}
		if v.vizslas != nil && *(v.vizslas) != *(day.expected.vizslas) {
			continue
		}
		if v.goldfish != nil && *(v.goldfish) != *(day.expected.goldfish) {
			continue
		}
		if v.trees != nil && *(v.trees) != *(day.expected.trees) {
			continue
		}
		if v.cars != nil && *(v.cars) != *(day.expected.cars) {
			continue
		}
		if v.perfumes != nil && *(v.perfumes) != *(day.expected.perfumes) {
			continue
		}
		return strconv.Itoa(i + 1)
	}

	return ""
}

func (day Day16) Part2(input []byte) string {
	m := day.parse(input)

	for i, v := range m {
		if v.children != nil && *(v.children) != *(day.expected.children) {
			continue
		}
		if v.cats != nil && *(v.cats) <= *(day.expected.cats) {
			continue
		}
		if v.samoyeds != nil && *(v.samoyeds) != *(day.expected.samoyeds) {
			continue
		}
		if v.pomeranians != nil && *(v.pomeranians) >= *(day.expected.pomeranians) {
			continue
		}
		if v.akitas != nil && *(v.akitas) != *(day.expected.akitas) {
			continue
		}
		if v.vizslas != nil && *(v.vizslas) != *(day.expected.vizslas) {
			continue
		}
		if v.goldfish != nil && *(v.goldfish) >= *(day.expected.goldfish) {
			continue
		}
		if v.trees != nil && *(v.trees) <= *(day.expected.trees) {
			continue
		}
		if v.cars != nil && *(v.cars) != *(day.expected.cars) {
			continue
		}
		if v.perfumes != nil && *(v.perfumes) != *(day.expected.perfumes) {
			continue
		}
		return strconv.Itoa(i + 1)
	}

	return ""
}

type Day16S struct {
	children, cats, samoyeds, pomeranians, akitas, vizslas, goldfish, trees, cars, perfumes *int
}

func (Day16) parse(input []byte) []Day16S {
	n := 500
	m := make([]Day16S, n)
	for line := range bytes.Lines(input) {
		line = bytes.TrimSpace(line)
		i := bytes.IndexByte(line, ':')
		sueN := optimistic.AtoiBFast(line[4:i]) - 1

		spl := bytes.Split(line[i+2:], []byte(", "))
		for _, v := range spl {
			i = bytes.IndexByte(v, ':')
			vv := optimistic.AtoiBFast(v[i+2:])
			switch string(v[:i]) {
			case "children":
				m[sueN].children = &vv
			case "cats":
				m[sueN].cats = &vv
			case "samoyeds":
				m[sueN].samoyeds = &vv
			case "pomeranians":
				m[sueN].pomeranians = &vv
			case "akitas":
				m[sueN].akitas = &vv
			case "vizslas":
				m[sueN].vizslas = &vv
			case "goldfish":
				m[sueN].goldfish = &vv
			case "trees":
				m[sueN].trees = &vv
			case "cars":
				m[sueN].cars = &vv
			case "perfumes":
				m[sueN].perfumes = &vv
			}
		}
	}

	return m
}
