package _matrix

import (
	"github.com/aivarasbaranauskas/aoc/internal/_num"
)

func Transpose[T any](m [][]T) [][]T {
	l := len(m)
	w := len(m[0])
	mT := make([][]T, w)
	for i := 0; i < w; i++ {
		mT[i] = make([]T, l)
		for j := 0; j < l; j++ {
			mT[i][j] = m[j][i]
		}
	}
	return mT
}

func Multiply[T _num.Numeric](a, b [][]T) [][]T {
	if len(a[0]) != len(b) {
		panic("mismatched matrices")
	}

	p := make([][]T, len(a))
	for i := range p {
		p[i] = make([]T, len(b[0]))
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(b); k++ {
				p[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return p
}
