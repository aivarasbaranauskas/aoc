package year_2022

import (
	"strconv"
)

func init() {
	Solutions[17] = Day17{
		figures: []Day17Figure{
			{
				bs: []uint8{
					0b1111,
				},
				w: 4, h: 1,
			},
			{
				bs: []uint8{
					0b010,
					0b111,
					0b010,
				},
				w: 3, h: 3,
			},
			{
				bs: []uint8{
					0b111,
					0b100,
					0b100,
				},
				w: 3, h: 3,
			},
			{
				bs: []uint8{
					0b1,
					0b1,
					0b1,
					0b1,
				},
				w: 1, h: 4,
			},
			{
				bs: []uint8{
					0b11,
					0b11,
				},
				w: 2, h: 2,
			},
		},
	}
}

type Day17 struct {
	figures []Day17Figure
}

func (d Day17) Part1(input []byte) string {
	ffbL := len(input)

	var m []uint8
	z := 0

	for n := 0; n < 2022; n++ {
		fig := d.figures[n%5]
		x, y := 2, len(m)+3

		for {
			ch := 1
			if input[z] == '<' {
				ch = -1
			}
			z = (z + 1) % ffbL

			xTmp := min(6, max(0, x+ch))
			if x != xTmp && xTmp+fig.w <= 7 && d.noCollision(xTmp, y, fig, m) {
				x = xTmp
			}

			if y > 0 && d.noCollision(x, y-1, fig, m) {
				y--
			} else {
				break
			}
		}

		if len(m) <= y+fig.h {
			// grow
			g := y + fig.h - len(m)
			for i := 0; i < g; i++ {
				m = append(m, 0)
			}
		}
		for yy, p := range fig.bs {
			m[y+yy] = m[y+yy] | (p << x)
		}
	}

	return strconv.Itoa(len(m))
}

func (d Day17) Part2(input []byte) string {
	ffbL := len(input)

	var m []uint8
	z := 0

	type www struct {
		w   [10]uint8
		fig int
		z   int
	}

	ws := map[www]int{} // n
	var hhs []int

	for n := 0; n < 1000000000000; n++ {
		fig := d.figures[n%5]
		x, y := 2, len(m)+3

		for {
			ch := 1
			if input[z] == '<' {
				ch = -1
			}
			z = (z + 1) % ffbL

			xTmp := min(6, max(0, x+ch))
			if x != xTmp && xTmp+fig.w <= 7 && d.noCollision(xTmp, y, fig, m) {
				x = xTmp
			}

			if y > 0 && d.noCollision(x, y-1, fig, m) {
				y--
			} else {
				break
			}
		}

		if len(m) <= y+fig.h {
			// grow
			g := y + fig.h - len(m)
			for i := 0; i < g; i++ {
				m = append(m, 0)
			}
		}
		for yy, p := range fig.bs {
			m[y+yy] = m[y+yy] | (p << x)
		}

		mL := len(m)
		if mL > 10 {
			w := www{
				w: [10]uint8{
					m[mL-10],
					m[mL-9],
					m[mL-8],
					m[mL-7],
					m[mL-6],
					m[mL-5],
					m[mL-4],
					m[mL-3],
					m[mL-2],
					m[mL-1],
				},
				fig: n % 5,
				z:   z,
			}
			if wN, ok := ws[w]; ok {
				cycleN := (1000000000000 - wN) / (n - wN)
				leftoverFigs := (1000000000000 - wN) % cycleN
				hH := hhs[wN] + (mL-hhs[wN])*cycleN + (hhs[wN+leftoverFigs] - hhs[wN]) - 1
				return strconv.Itoa(hH)
			}
			ws[w] = n
		}
		hhs = append(hhs, mL)
	}

	return strconv.Itoa(len(m))
}

type Day17Figure struct {
	bs []uint8
	w  int
	h  int
}

func (Day17) noCollision(x, y int, fig Day17Figure, m []uint8) bool {
	if y >= len(m) {
		return true
	}
	for i := 0; i < len(fig.bs) && y+i < len(m); i++ {
		if m[y+i]&(fig.bs[i]<<x) > 0 {
			return false
		}
	}

	return true
}
