package year_2015

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func init() {
	Solutions[4] = Day4{}
}

type Day4 struct{}

func (day Day4) Part1(input []byte) string {
	return day.calculate(input, 5)
}

func (day Day4) Part2(input []byte) string {
	return day.calculate(input, 6)
}

func (day Day4) calculate(input []byte, numZeros int) string {
	expectedPrefix := bytes.Repeat([]byte{'0'}, numZeros)

	inputExpanded := make([]byte, len(input), len(input)+10)
	copy(inputExpanded, input)

	h := md5.New()
	sum := make([]byte, 0, md5.Size)
	encoded := make([]byte, md5.Size*2)

	for i := int64(1); ; i++ {
		appended := strconv.AppendInt(inputExpanded, i, 10)
		sum = sum[:0]
		h.Reset()
		h.Write(appended)
		hex.Encode(encoded, h.Sum(sum))

		if bytes.HasPrefix(encoded, expectedPrefix) {
			return strconv.Itoa(int(i))
		}
	}
}
