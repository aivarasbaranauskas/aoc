package year_2015

import (
	"strconv"
)

func init() {
	Solutions[10] = Day10{}
}

type Day10 struct{}

func (day Day10) Part1(input []byte) string {
	return day.run(input, 40)
}

func (day Day10) run(input []byte, n int) string {
	tmp := make([]byte, 0, len(input))

	for range n {
		c := 0
		tmp = tmp[:0]
		for i := 1; i < len(input); i++ {
			if input[i] == input[c] {
				continue
			}
			tmp = append(tmp, '0'+byte(i-c))
			tmp = append(tmp, input[c])
			c = i
		}

		tmp = append(tmp, '0'+byte(len(input)-c))
		tmp = append(tmp, input[c])

		input, tmp = tmp, input
	}

	return strconv.Itoa(len(input))
}

func (day Day10) Part2(input []byte) string {
	return day.run(input, 50)
}
