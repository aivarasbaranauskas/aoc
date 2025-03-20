package year_2024

import (
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
)

func init() {
	Solutions[7] = Day7{}
}

type Day7 struct{}

func (day Day7) Part1(input []byte) string {
	s := 0

	values := make([]int, 1)

	for i := 0; i < len(input); i++ {
		values = values[:1]
		i2 := bytes.IndexByte(input[i:], ':')
		testValue := optimistic.AtoiBFast(input[i : i+i2])
		values[0] = testValue

		eol := bytes.IndexByte(input[i:], '\n')
		if eol == -1 {
			eol = len(input) - i
		}

		e := eol
		for i3 := eol - 1; i3 > i2+2; i3-- {
			if input[i+i3] == ' ' {
				value := optimistic.AtoiBFast(input[i+i3+1 : i+e])

				vl := len(values)
				for j := 0; j < vl; j++ {
					if values[j]%value == 0 {
						values = append(values, values[j]/value)
					}
					values[j] -= value
				}
				e = i3
			}
		}

		lastValue := optimistic.AtoiBFast(input[i+i2+2 : i+e])
		for _, value := range values {
			if value == lastValue {
				s += testValue
				break
			}
		}

		i += eol
	}

	return strconv.Itoa(s)
}

func (day Day7) Part2(input []byte) string {
	s := 0

	values := make([]int, 1)

	for i := 0; i < len(input); i++ {
		values = values[:1]
		i2 := bytes.IndexByte(input[i:], ':')
		testValue := optimistic.AtoiBFast(input[i : i+i2])
		values[0] = testValue

		eol := bytes.IndexByte(input[i:], '\n')
		if eol == -1 {
			eol = len(input) - i
		}

		e := eol
		for i3 := eol - 1; i3 > i2+2; i3-- {
			if input[i+i3] == ' ' {
				value := optimistic.AtoiBFast(input[i+i3+1 : i+e])

				vl := len(values)
				for j := 0; j < vl; j++ {
					valuePower10 := day.get10Power(value) * 10
					if values[j]%valuePower10 == value {
						values = append(values, values[j]/valuePower10)
					}
					if values[j]%value == 0 {
						values = append(values, values[j]/value)
					}
					values[j] -= value
				}
				e = i3
			}
		}

		lastValue := optimistic.AtoiBFast(input[i+i2+2 : i+e])
		for _, value := range values {
			if value == lastValue {
				s += testValue
				break
			}
		}

		i += eol
	}

	return strconv.Itoa(s)
}

func (day Day7) get10Power(value int) int {
	power := 1
	for value > 9 {
		value /= 10
		power *= 10
	}
	return power
}
