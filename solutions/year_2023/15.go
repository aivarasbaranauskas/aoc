package year_2023

import (
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
)

func init() {
	Solutions[15] = Day15{}
}

type Day15 struct{}

func (d Day15) Part1(input []byte) string {
	sm := 0
	for _, s := range bytes.Split(input, []byte(",")) {
		sm += int(d.HASH(s))
	}
	return strconv.Itoa(sm)
}

func (d Day15) Part2(input []byte) string {
	seq := bytes.Split(input, []byte(","))
	boxes := make([][]Lens, 256)

OpLoop:
	for _, s := range seq {
		opI := bytes.IndexAny(s, "=-")
		label := string(s[:opI])
		h := d.HASH(s[:opI])

		if s[opI] == '=' {
			focalLength := optimistic.Atoi(string(s[opI+1:]))

			for i, lens := range boxes[h] {
				if lens.Label == label {
					boxes[h][i].FocalLength = focalLength
					continue OpLoop
				}
			}

			boxes[h] = append(boxes[h], Lens{
				Label:       label,
				FocalLength: focalLength,
			})
		} else {
			for i, lens := range boxes[h] {
				if lens.Label == label {
					boxes[h], _ = _slice.Remove(boxes[h], i)
					continue OpLoop
				}
			}
		}
	}

	s := 0
	for i, box := range boxes {
		for j, lens := range box {
			s += (i + 1) * (j + 1) * lens.FocalLength
		}
	}

	return strconv.Itoa(s)
}

type Lens struct {
	Label       string
	FocalLength int
}

func (Day15) HASH(s []byte) (hash uint8) {
	for _, c := range s {
		// overflow!
		hash = (hash + c) * 17
	}
	return hash
}
