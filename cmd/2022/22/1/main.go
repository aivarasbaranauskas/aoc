package main

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"log"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var (
		m     [][]byte
		moves string
	)

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Text()
		if len(line) == 0 {
			r.Scan()
			moves = r.Text()
			break
		}
		m = append(m, []byte(line))
	}

	// pad map
	w := _num.Max(_slice.Map(m, func(line []byte) int {
		return len(line)
	})...)
	for i := range m {
		if len(m[i]) < w {
			m[i] = append(m[i], bytes.Repeat([]byte(" "), w-len(m[i]))...)
		}
	}

	l, c, d := 0, bytes.IndexAny(m[0], ".#"), 0

	for len(moves) > 0 {
		next := strings.IndexAny(moves, "LR")
		var ct int
		if next == -1 {
			ct = optimistic.Atoi(moves)
		} else {
			ct = optimistic.Atoi(moves[:next])
		}

	MoveLoop:
		for i := 0; i < ct; i++ {
			switch d {
			case 0:
				// right
				cn := c + 1
				if cn >= len(m[l]) || m[l][cn] == ' ' {
					cn = 0
				}
				for m[l][cn] == ' ' {
					cn++
				}
				if m[l][cn] == '#' {
					break MoveLoop
				}
				c = cn
			case 1:
				// down
				ln := l + 1
				if ln >= len(m) || m[ln][c] == ' ' {
					ln = 0
				}
				for m[ln][c] == ' ' {
					ln++
				}
				if m[ln][c] == '#' {
					break MoveLoop
				}
				l = ln
			case 2:
				// left
				cn := c - 1
				if cn < 0 || m[l][cn] == ' ' {
					cn = len(m[l]) - 1
				}
				for m[l][cn] == ' ' {
					cn--
				}
				if m[l][cn] == '#' {
					break MoveLoop
				}
				c = cn
			case 3:
				// up
				ln := l - 1
				if ln < 0 || m[ln][c] == ' ' {
					ln = len(m) - 1
				}
				for m[ln][c] == ' ' {
					ln--
				}
				if m[ln][c] == '#' {
					break MoveLoop
				}
				l = ln
			}
		}

		if next >= 0 {
			if moves[next] == 'R' {
				d = (d + 1) % 4
			} else {
				d--
				if d == -1 {
					d = 3
				}
			}
			moves = moves[next+1:]
		} else {
			moves = ""
		}
	}

	fmt.Println(1000*(l+1) + 4*(c+1) + d)
}
