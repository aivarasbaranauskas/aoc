package _matrix

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
