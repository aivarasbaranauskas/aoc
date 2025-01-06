package year_2019

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
)

func init() {
	Solutions[24] = Day24{}
}

type Day24 struct{}

func (Day24) Part1(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var f [][]bool
	for r.Scan() {
		line := r.Text()
		fieldLine := make([]bool, len(line))
		for i := range line {
			if line[i] == '#' {
				fieldLine[i] = true
			}
		}
		f = append(f, fieldLine)

		if len(fieldLine) != 5 {
			log.Fatal("not 5 cols in row", len(f))
		}
	}

	if len(f) != 5 {
		log.Fatal("not 5 rows")
	}

	field := fieldT1{f: f}
	lastStates := map[uint32]struct{}{}

	for {
		enc := field.encode()
		if _, ok := lastStates[enc]; ok {
			break
		}
		lastStates[enc] = struct{}{}
		field.next()
	}

	return strconv.Itoa(int(field.encode()))
}

func (Day24) Part2(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var f fieldT
	for r.Scan() {
		line := r.Text()
		fieldLine := make([]bool, len(line))
		for i := range line {
			if line[i] == '#' {
				fieldLine[i] = true
			}
		}
		f = append(f, fieldLine)

		if len(fieldLine) != 5 {
			log.Fatal("not 5 cols in row", len(f))
		}
	}

	if len(f) != 5 {
		log.Fatal("not 5 rows")
	}

	field := &fieldW{
		f: f,
	}
	field.in = &fieldW{f: emptyField(), out: field}
	field.out = &fieldW{f: emptyField(), in: field}

	for i := 0; i < 200; i++ {
		field.next()
	}

	return strconv.Itoa(field.count())
}

type fieldT1 struct {
	f [][]bool
}

func (f *fieldT1) next() {
	tmp := [][]bool{
		make([]bool, 5),
		make([]bool, 5),
		make([]bool, 5),
		make([]bool, 5),
		make([]bool, 5),
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			var s int
			//left
			if j > 0 && f.f[i][j-1] {
				s++
			}
			//right
			if j < 4 && f.f[i][j+1] {
				s++
			}
			//up
			if i > 0 && f.f[i-1][j] {
				s++
			}
			//down
			if i < 4 && f.f[i+1][j] {
				s++
			}

			tmp[i][j] = (!f.f[i][j] && (s == 1 || s == 2)) || (f.f[i][j] && s == 1)
		}
	}

	f.f = tmp
}

func (f *fieldT1) encode() (out uint32) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if f.f[i][j] {
				out |= 1 << (i*5 + j)
			}
		}
	}
	return
}

// part 2

type fieldT [][]bool
type fieldW struct {
	f   fieldT
	in  *fieldW
	out *fieldW
}

func (f *fieldW) next() {
	next := f.nextField()
	f.in.nextIn()
	f.out.nextOut()
	f.f = next
}

func (f *fieldW) nextIn() {
	if f.in == nil {
		f.in = &fieldW{f: emptyField(), out: f}
		f.f = f.nextField()
	} else {
		next := f.nextField()
		f.in.nextIn()
		f.f = next
	}
}

func (f *fieldW) nextOut() {
	if f.out == nil {
		f.out = &fieldW{f: emptyField(), in: f}
		f.f = f.nextField()
	} else {
		next := f.nextField()
		f.out.nextOut()
		f.f = next
	}
}

func (f *fieldW) nextField() fieldT {
	newF := emptyField()

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				// not interested in middle as it is _inception_
				continue
			}

			var s int
			//left
			if i == 2 && j == 3 {
				for k := 0; k < 5; k++ {
					if f.in.f[k][4] {
						s++
					}
				}
			} else if (j > 0 && f.f[i][j-1]) || (j == 0 && f.out.f[2][1]) {
				s++
			}
			//right
			if i == 2 && j == 1 {
				for k := 0; k < 5; k++ {
					if f.in.f[k][0] {
						s++
					}
				}
			} else if (j < 4 && f.f[i][j+1]) || (j == 4 && f.out.f[2][3]) {
				s++
			}
			//up
			if i == 3 && j == 2 {
				for k := 0; k < 5; k++ {
					if f.in.f[4][k] {
						s++
					}
				}
			} else if (i > 0 && f.f[i-1][j]) || (i == 0 && f.out.f[1][2]) {
				s++
			}
			//down
			if i == 1 && j == 2 {
				for k := 0; k < 5; k++ {
					if f.in.f[0][k] {
						s++
					}
				}
			} else if (i < 4 && f.f[i+1][j]) || (i == 4 && f.out.f[3][2]) {
				s++
			}

			newF[i][j] = (!f.f[i][j] && (s == 1 || s == 2)) || (f.f[i][j] && s == 1)
		}
	}

	return newF
}

func (f fieldT) encode() (out uint32) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if f[i][j] {
				out |= 1 << (i*5 + j)
			}
		}
	}
	return
}

func (f fieldT) count() (out int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if f[i][j] {
				out++
			}
		}
	}
	return
}

func (f *fieldW) count() int {
	return f.f.count() + f.in.countIn() + f.out.countOut()
}

func (f *fieldW) countIn() int {
	if f == nil {
		return 0
	}
	return f.f.count() + f.in.countIn()
}

func (f *fieldW) countOut() int {
	if f == nil {
		return 0
	}
	return f.f.count() + f.out.countOut()
}

func emptyField() fieldT {
	return [][]bool{
		make([]bool, 5),
		make([]bool, 5),
		make([]bool, 5),
		make([]bool, 5),
		make([]bool, 5),
	}
}
