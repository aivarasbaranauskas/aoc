package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

//go:embed input.txt
var input []byte

func main() {
	fmt.Println("part 1:", p1())
	fmt.Println("part 2:", p2())
}

func p1() (sm int) {
	for _, s := range bytes.Split(input, []byte(",")) {
		sm += int(HASH(s))
	}
	return
}

type Lens struct {
	Label       string
	FocalLength int
}

func p2() int {
	seq := bytes.Split(input, []byte(","))
	boxes := make([][]Lens, 256)

OpLoop:
	for _, s := range seq {
		opI := bytes.IndexAny(s, "=-")
		label := string(s[:opI])
		h := HASH(s[:opI])

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

	return s
}

func HASH(s []byte) (hash uint8) {
	for _, c := range s {
		// overflow!
		hash = (hash + c) * 17
	}
	return hash
}
