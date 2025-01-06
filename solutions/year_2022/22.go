package year_2022

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"slices"
	"strconv"
	"strings"
)

func init() {
	Solutions[22] = Day22{}
}

type Day22 struct{}

func (Day22) Part1(input []byte) string {
	var (
		m     [][]byte
		moves string
	)

	r := bufio.NewScanner(bytes.NewReader(input))
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
	w := slices.Max(_slice.Map(m, func(line []byte) int {
		return len(line)
	}))
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

	return strconv.Itoa(1000*(l+1) + 4*(c+1) + d)
}

func (Day22) Part2(input []byte) string {
	var (
		m     [][]byte
		moves string
	)

	r := bufio.NewScanner(bytes.NewReader(input))
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
	w := slices.Max(_slice.Map(m, func(line []byte) int {
		return len(line)
	}))
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

		for i := 0; i < ct; i++ {
			ln, cn, dn := l, c, d
			switch d {
			case 0:
				// right
				cn++
				if cn >= len(m[l]) || m[l][cn] == ' ' {
					switch l / 50 {
					case 0:
						// B -> D
						ln = 149 - l
						cn = 99
						dn = 2
					case 1:
						// C -> B
						ln = 49
						cn = 50 + l
						dn = 3
					case 2:
						// D -> B
						ln = 149 - l
						cn = 149
						dn = 2
					case 3:
						// F -> D
						ln = 149
						cn = l - 100
						dn = 3
					}
				}
			case 1:
				// down
				ln++
				if ln >= len(m) || m[ln][c] == ' ' {
					switch c / 50 {
					case 0:
						// F -> B
						ln = 0
						cn += 100
					case 1:
						// D -> F
						ln = 100 + c
						cn = 49
						dn = 2
					case 2:
						// B -> C
						ln = c - 50
						cn = 99
						dn = 2
					}
				}
			case 2:
				// left
				cn--
				if cn < 0 || m[l][cn] == ' ' {
					switch l / 50 {
					case 0:
						// A -> E
						ln = 149 - l
						cn = 0
						dn = 0
					case 1:
						// C -> E
						ln = 100
						cn = l - 50
						dn = 1
					case 2:
						// E -> A
						ln = 149 - l
						cn = 50
						dn = 0
					case 3:
						// F -> A
						ln = 0
						cn = l - 100
						dn = 1
					}
				}
			case 3:
				// up
				ln--
				if ln < 0 || m[ln][c] == ' ' {
					switch c / 50 {
					case 0:
						// E -> C
						ln = c + 50
						cn = 50
						dn = 0
					case 1:
						// A -> F
						ln = 100 + c
						cn = 0
						dn = 0
					case 2:
						// B -> F
						ln = 199
						cn = c - 100
					}
				}
			}
			if m[ln][cn] == '#' {
				break
			}
			l, c, d = ln, cn, dn
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

	return strconv.Itoa(1000*(l+1) + 4*(c+1) + d)
}
