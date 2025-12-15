package year_2015

import "strconv"

func init() {
	Solutions[3] = Day3{}
}

type Day3 struct{}

func (Day3) Part1(input []byte) string {
	mem := map[[2]int]struct{}{}
	x, y, ct := 0, 0, 0
	for _, v := range input {
		switch v {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		if _, ok := mem[[2]int{x, y}]; !ok {
			ct++
			mem[[2]int{x, y}] = struct{}{}
		}
	}

	return strconv.Itoa(ct)
}

func (Day3) Part2(input []byte) string {
	mem := map[[2]int]struct{}{}
	sx, sy, rx, ry, ct := 0, 0, 0, 0, 0
	robo := false
	for _, v := range input {
		var x, y *int
		if robo {
			x, y = &rx, &ry
		} else {
			x, y = &sx, &sy
		}
		robo = !robo

		switch v {
		case '^':
			(*y)++
		case 'v':
			(*y)--
		case '>':
			(*x)++
		case '<':
			(*x)--
		}
		if _, ok := mem[[2]int{(*x), (*y)}]; !ok {
			ct++
			mem[[2]int{(*x), (*y)}] = struct{}{}
		}
	}

	return strconv.Itoa(ct)
}
