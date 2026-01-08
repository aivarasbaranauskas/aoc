package year_2015

import (
	"bytes"
	"iter"
)

func init() {
	Solutions[11] = Day11{}
}

type Day11 struct{}

func (day Day11) Part1(input []byte) string {
	for pass := range day.iteratePasswords(input) {
		if day.isPasswordValid(pass) {
			return string(pass)
		}
	}
	panic("not found")
}

func (day Day11) Part2(input []byte) string {
	ct := 0
	for pass := range day.iteratePasswords(input) {
		if day.isPasswordValid(pass) {
			ct++
			if ct == 2 {
				return string(pass)
			}
		}
	}
	panic("not found")
}

func (day Day11) iteratePasswords(pass []byte) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		for {
			for i := len(pass) - 1; i >= 0; i-- {
				if pass[i] != 'z' {
					pass[i]++
					break
				}
				pass[i] = 'a'
			}

			if !yield(pass) {
				break
			}
		}
	}
}

func (day Day11) isPasswordValid(pass []byte) bool {
	if bytes.IndexByte(pass, 'i') != -1 {
		return false
	}
	if bytes.IndexByte(pass, 'o') != -1 {
		return false
	}
	if bytes.IndexByte(pass, 'l') != -1 {
		return false
	}

	has3 := false
	for i := 0; i < len(pass)-3; i++ {
		has3 = has3 || (pass[i]+1 == pass[i+1] && pass[i+1]+1 == pass[i+2])
	}
	if !has3 {
		return false
	}

	pairs := 0
	for i := 0; i < len(pass)-1; i++ {
		if pass[i] == pass[i+1] {
			pairs++
			i++
		}
	}

	return pairs > 1
}
