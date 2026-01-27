package year_2015

import "strconv"

func init() {
	Solutions[18] = Day18{}
}

type Day18 struct{}

func (day Day18) Part1(input []byte) string {
	m := day.parse(input)

	for range 100 {
		m = day.nextState(m)
	}

	ct := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if m[i][j] {
				ct++
			}
		}
	}

	return strconv.Itoa(ct)
}

func (day Day18) Part2(input []byte) string {
	m := day.parse(input)
	m[0][0] = true
	m[0][99] = true
	m[99][0] = true
	m[99][99] = true

	for range 100 {
		m = day.nextState(m)
		m[0][0] = true
		m[0][99] = true
		m[99][0] = true
		m[99][99] = true
	}

	ct := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if m[i][j] {
				ct++
			}
		}
	}

	return strconv.Itoa(ct)
}

func (day Day18) nextState(m [100][100]bool) [100][100]bool {
	var m2 [100][100]bool

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			ct := 0

			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					if di == 0 && dj == 0 {
						continue
					}

					ti := i + di
					tj := j + dj

					if 0 <= ti && ti < 100 && 0 <= tj && tj < 100 && m[ti][tj] {
						ct++
					}
				}
			}

			m2[i][j] = ct == 3 || (ct == 2 && m[i][j])
		}
	}

	return m2
}

func (day Day18) parse(input []byte) [100][100]bool {
	var m [100][100]bool

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			m[i][j] = input[i*101+j] == '#'
		}
	}

	return m
}
