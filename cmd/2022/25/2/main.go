package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"math"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	r := bufio.NewScanner(f)
	var sum int
	for r.Scan() {
		line := r.Text()
		sum += decode([]byte(line))
	}
	fmt.Println(sum)
	fmt.Println(string(encode(sum)))
}

func decode(n []byte) (x int) {
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

func encode(n int) (x []byte) {
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
