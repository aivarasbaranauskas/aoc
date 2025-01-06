package year_2022

import (
	"bufio"
	"bytes"
	"math"
)

func init() {
	Solutions[25] = Day25{}
}

type Day25 struct{}

func (d Day25) Part1(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var sum int
	for r.Scan() {
		line := r.Text()
		sum += d.decode([]byte(line))
	}
	return string(d.encode(sum))
}

func (Day25) decode(n []byte) (x int) {
	l := len(n)
	for i, c := range n {
		switch c {
		case '=':
			x -= int(math.Pow(5, float64(l-i-1))) * 2
		case '-':
			x -= int(math.Pow(5, float64(l-i-1)))
		case '1':
			x += int(math.Pow(5, float64(l-i-1)))
		case '2':
			x += int(math.Pow(5, float64(l-i-1))) * 2
		}
	}
	return
}

func (Day25) encode(n int) (x []byte) {
	var carry int
	for n > 0 {
		n += carry
		carry = 0
		c := n % 5
		n = n / 5
		if c < 3 {
			x = append([]byte{byte('0' + c)}, x...)
		} else {
			carry = 1
			if c == 3 {
				x = append([]byte{'='}, x...)
			} else {
				x = append([]byte{'-'}, x...)
			}
		}
	}
	return
}

func (Day25) Part2(_ []byte) string {
	return ""
}
