package year_2024

import (
	"strconv"
)

func init() {
	Solutions[9] = Day9{}
}

type Day9 struct{}

func (Day9) Part1(input []byte) string {
	for i := range input {
		input[i] -= '0'
	}

	i, j := 0, (len(input)-1)/2
	cf, cj := 0, 0
	s := 0

	for i < j {
		for range input[i*2] {
			s += i * cf
			cf++
		}
		if i+1 < len(input) {
			for range input[i*2+1] {
				for cj == int(input[j*2]) {
					j--
					cj = 0
				}
				s += j * cf
				cj++
				cf++
			}
		}
		i++
	}

	for cj < int(input[j*2]) {
		s += j * cf
		cj++
		cf++
	}

	return strconv.Itoa(s)
}

func (Day9) Part2(input []byte) string {
	l := (len(input) + 1) / 2
	diskMap0 := make([]int, l)
	diskMap1 := make([]int, l)
	diskMap2 := make([]int, l)
	for i := range l {
		diskMap0[i] = int(input[i*2] - '0')
		diskMap2[i] = i
		if i+1 < l {
			diskMap1[i] = int(input[i*2+1] - '0')
		}
	}

	for i := l - 1; i >= 0; i-- {
		v := diskMap0[i]
		for j := 0; j < i; j++ {
			if diskMap1[j] < v {
				continue
			}

			tmp0 := diskMap0[i]
			tmp1 := diskMap1[i]
			tmp2 := diskMap2[i]

			copy(diskMap0[j+2:], diskMap0[j+1:i])
			copy(diskMap1[j+2:], diskMap1[j+1:i])
			copy(diskMap2[j+2:], diskMap2[j+1:i])

			diskMap0[j+1] = tmp0
			diskMap1[j+1] = diskMap1[j] - tmp0
			diskMap2[j+1] = tmp2
			diskMap1[j] = 0
			diskMap1[i] += tmp0 + tmp1

			i++
			break
		}
	}

	s := 0
	c := 0
	for i := range l {
		for range diskMap0[i] {
			s += c * diskMap2[i]
			c++
		}
		c += diskMap1[i]
	}

	return strconv.Itoa(s)
}
