package year_2015

import (
	"bytes"
	"strconv"
	"strings"
)

func init() {
	Solutions[19] = Day19{}
}

type Day19 struct{}

func (day Day19) Part1(input []byte) string {
	m, s := day.parse(input)
	mem := map[string]struct{}{}

	for from, toA := range m {
		i := 0
		for {
			iF := strings.Index(s[i:], from)
			if iF == -1 {
				break
			}
			for _, to := range toA {
				newS := s[:i+iF] + to + s[i+iF+len(from):]
				mem[newS] = struct{}{}
			}
			i += iF + 1
		}
	}

	return strconv.Itoa(len(mem))
}

func (day Day19) Part2(input []byte) string {
	_, s := day.parse(input)
	ct := 0

	for _, v := range s {
		if 'A' <= v && v <= 'Z' {
			ct++
		}
	}

	rn := strings.Count(s, "Rn")
	ar := strings.Count(s, "Ar")
	y := strings.Count(s, "Y")

	return strconv.Itoa(ct - rn - ar - 2*y - 1)
}

func (day Day19) parse(input []byte) (map[string][]string, string) {
	spl := bytes.Split(input, []byte("\n\n"))
	m := make(map[string][]string)
	for line := range bytes.Lines(spl[0]) {
		line = bytes.TrimSpace(line)
		spaceI := bytes.IndexByte(line, ' ')

		from := string(line[:spaceI])
		to := string(line[spaceI+4:])

		if mi, ok := m[from]; ok {
			m[from] = append(mi, to)
		} else {
			m[from] = []string{to}
		}
	}

	return m, string(bytes.TrimSpace(spl[1]))
}
