package main

import (
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func main() {
	m := bytes.Split(input, []byte("\n"))

	lineL := len(m[0])

	lS, cS := 0, 0
	for m[lS][cS] != 'S' {
		cS++
		lS += cS / lineL
		cS %= lineL
	}

	fmt.Println("S:", lS, cS)

	tryGo := func(c, l, pc, pl int) (int, bool, [][]int) {
		mS := make([][]int, len(m))
		for i := range mS {
			mS[i] = make([]int, len(m[i]))
			for j := range mS[i] {
				mS[i][j] = -1
			}
		}
		mS[pl][pc] = 0

		ln := 1
		for l != lS || c != cS {
			mS[l][c] = ln
			switch m[l][c] {
			case '-':
				if l != pl {
					return 0, false, nil
				}
				if pc < c {
					pc = c
					c++
				} else {
					pc = c
					c--
				}
			case '|':
				if pc != c {
					return 0, false, nil
				}
				if pl < l {
					pl = l
					l++
				} else {
					pl = l
					l--
				}
			case '7':
				if pc > c || pl < l {
					return 0, false, nil
				}
				if pc < c {
					pl = l
					pc = c
					l++
				} else {
					pl = l
					pc = c
					c--
				}
			case 'J':
				if pc > c || pl > l {
					return 0, false, nil
				}
				if pl < l {
					pl = l
					pc = c
					c--
				} else {
					pl = l
					pc = c
					l--
				}
			case 'L':
				if pc < c || pl > l {
					return 0, false, nil
				}
				if pl < l {
					pl = l
					pc = c
					c++
				} else {
					pl = l
					pc = c
					l--
				}
			case 'F':
				if pc < c || pl < l {
					return 0, false, nil
				}
				if pc > c {
					pl = l
					pc = c
					l++
				} else {
					pl = l
					pc = c
					c++
				}
			}

			if l < 0 || c < 0 || l >= len(m) || c >= len(m[0]) || m[l][c] == '.' {
				return 0, false, nil
			}

			ln++
		}

		return ln, true, mS
	}

TheLoop:
	for _, a := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		c, l := cS+a[0], lS+a[1]
		if l < 0 || c < 0 || l >= len(m) || c >= len(m[0]) || m[l][c] == '.' {
			continue
		}
		ln, ok, mS := tryGo(c, l, cS, lS)
		if !ok {
			continue
		}
		fmt.Println("part 1:", ln/2)

		// replace S with correct
		var cB, lB, mx int
		for _, b := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			if a == b {
				continue
			}
			cBt, lBt := cS+b[0], lS+b[1]
			if mS[lB][cB] > mx {
				cB, lB = cBt, lBt
				mx = mS[lB][cB]
			}
		}

		up := (c == cS && l == lS-1) || (cB == cS && lB == lS-1)
		down := (c == cS && l == lS+1) || (cB == cS && lB == lS+1)
		left := (c == cS-1 && l == lS) || (cB == cS-1 && lB == lS)
		right := (c == cS+1 && l == lS) || (cB == cS+1 && lB == lS)

		if up && down {
			m[lS][cS] = '|'
		} else if left && right {
			m[lS][cS] = '-'
		} else if left && down {
			m[lS][cS] = '7'
		} else if left && up {
			m[lS][cS] = 'J'
		} else if right && up {
			m[lS][cS] = 'L'
		} else if right && down {
			m[lS][cS] = 'F'
		}

		// color insides
		for i := range mS {
			in := false
			down := false
			up := false
			for j := range mS[i] {
				if mS[i][j] == -1 {
					if in {
						mS[i][j] = -2
					}
					continue
				}

				switch m[i][j] {
				case '|':
					in = !in
				case 'F':
					up = true
				case 'L':
					down = true
				case '7':
					if up {
						up = false
					} else if down {
						down = false
						in = !in
					}
				case 'J':
					if down {
						down = false
					} else if up {
						up = false
						in = !in
					}
				}
			}
		}

		// count

		ct := 0
		for i := range mS {
			for j := range mS[i] {
				if mS[i][j] == -2 {
					ct++
				}
			}
		}

		//for i := range mS {
		//	for j := range mS[i] {
		//		if mS[i][j] == -2 {
		//			fmt.Print("O")
		//		} else if mS[i][j] == -1 {
		//			fmt.Print(".")
		//		} else {
		//			fmt.Print("+")
		//		}
		//	}
		//	fmt.Println()
		//}
		fmt.Println("part 2:", ct)

		break TheLoop
	}
}
