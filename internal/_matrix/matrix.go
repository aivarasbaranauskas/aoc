package _matrix

import (
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"math"
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

func Minor[T any](m [][]T, i, j int) [][]T {
	m2 := make([][]T, 0, len(m)-1)
	for ii, row := range m {
		if i == ii {
			continue
		}
		m2 = append(m2, append(row[:j], row[:j+1]...))
	}
	return m2
}

func Determinant[T _num.SignedNumeric](m [][]T) T {
	//base case for 2x2 matrix
	if len(m) == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	}

	d := T(0)
	for c := range m {
		cc := T(math.Pow(-1, float64(c)))
		d += (cc) * m[0][c] * Determinant(Minor(m, 0, c))
	}
	return d
}

func Inverse[T _num.SignedNumeric](m [][]T) ([][]T, bool) {
	d := Determinant(m)
	if d == 0 {
		return nil, false
	}

	//special case for 2x2 matrix:
	if len(m) == 2 {
		return [][]T{
			{m[1][1] / d, -1 * m[0][1] / d},
			{-1 * m[1][0] / d, m[0][0] / d},
		}, true
	}

	// find matrix of cofactors
	var cofactors [][]T
	for r := range m {
		var cofactorRow []T
		for c := range m {
			minor := Minor(m, r, c)
			cc := T(math.Pow(-1, float64(r+c)))
			cofactorRow = append(cofactorRow, cc*Determinant(minor))
		}
		cofactors = append(cofactors, cofactorRow)
	}
	cofactors = Transpose(cofactors)
	for r := range cofactors {
		for c := range cofactors {
			cofactors[r][c] = cofactors[r][c] / d
		}
	}

	return cofactors, true
}
