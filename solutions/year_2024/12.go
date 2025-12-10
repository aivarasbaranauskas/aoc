package year_2024

import (
	"bytes"
	"strconv"
)

func init() {
	Solutions[12] = Day12{}
}

type Day12 struct{}

func (day Day12) Part1(input []byte) string {
	m := bytes.Split(input, []byte("\n"))
	sum := 0

	for row := range m {
		for col := range m[row] {
			if m[row][col] < 'A' {
				// already visited
				continue
			}

			area, perimeter := day.walk(m, m[row][col], row, col)
			sum += area * perimeter
		}
	}

	return strconv.Itoa(sum)
}

func (day Day12) walk(m [][]byte, b byte, row, col int) (area int, perimeter int) {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[row]) {
		// outside
		return 0, 1
	}

	if m[row][col] != b {
		if m[row][col] == b-'A' {
			// already visited
			return 0, 0
		}

		// other region
		return 0, 1
	}

	area, perimeter = 1, 0
	m[row][col] -= 'A'

	a, p := day.walk(m, b, row+1, col)
	area += a
	perimeter += p
	a, p = day.walk(m, b, row-1, col)
	area += a
	perimeter += p
	a, p = day.walk(m, b, row, col+1)
	area += a
	perimeter += p
	a, p = day.walk(m, b, row, col-1)
	area += a
	perimeter += p

	return
}

func (day Day12) Part2(input []byte) string {
	m := bytes.Split(input, []byte("\n"))
	sum := 0

	for row := range m {
		for col := range m[row] {
			if m[row][col] < 'A' {
				// already visited
				continue
			}

			area, corners := day.walk2(m, m[row][col], row, col)
			sum += area * corners
		}
	}

	return strconv.Itoa(sum)
}

func (day Day12) walk2(m [][]byte, b byte, row, col int) (area int, corners int) {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[row]) {
		// outside
		return 0, 0
	}

	if m[row][col] != b {
		// already visited or other region
		return 0, 0
	}

	isPopulatedBySame := func(r, c int) bool {
		return r >= 0 && r < len(m) && c >= 0 && c < len(m[r]) && (m[r][c] == b || m[r][c] == b-'A')
	}

	isCorner := func(dr, dc int) int {
		dp := isPopulatedBySame(row+dr, col+dc)
		a1 := isPopulatedBySame(row, col+dc)
		a2 := isPopulatedBySame(row+dr, col)
		if ((!dp) && a1 == a2) || (dp && (!a1) && !a2) {
			return 1
		}
		return 0
	}

	m[row][col] -= 'A' // mark as visited
	area = 1
	corners = isCorner(1, 1) +
		isCorner(-1, -1) +
		isCorner(1, -1) +
		isCorner(-1, 1)

	a, c := day.walk2(m, b, row+1, col)
	area += a
	corners += c
	a, c = day.walk2(m, b, row-1, col)
	area += a
	corners += c
	a, c = day.walk2(m, b, row, col+1)
	area += a
	corners += c
	a, c = day.walk2(m, b, row, col-1)
	area += a
	corners += c

	return
}
