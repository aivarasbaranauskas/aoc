package main

import (
	"fmt"
)

func main() {
	fmt.Println(theThing([14]int{9, 2, 9, 2, 8, 9, 1, 4, 9, 9, 9, 9, 9, 1})) // Part 1
	fmt.Println(theThing([14]int{9, 1, 8, 1, 1, 2, 1, 1, 6, 1, 1, 9, 8, 1})) // Part 2
}

func theThing(ws [14]int) bool {
	var x, y, z, w, i int
	w = ws[i]
	i++
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 1
	x = x + 12
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 4
	y = y * x
	z = z + y
	w = ws[i]
	i++
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 1
	x = x + 15
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 11
	y = y * x
	z = z + y
	w = ws[i]
	i++
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 1
	x = x + 11
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 7
	y = y * x
	z = z + y
	w = ws[i]
	i++
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -14
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 2
	y = y * x
	z = z + y
	w = ws[i]
	i++
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 1
	x = x + 12
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 11
	y = y * x
	z = z + y
	w = ws[i]
	i++
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -10
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 13
	y = y * x
	z = z + y
	w = ws[i]
	i++
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 1
	x = x + 11
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 9
	y = y * x
	z = z + y
	w = ws[i]
	i++
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 1
	x = x + 13
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 12
	y = y * x
	z = z + y
	w = ws[i]
	i++
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -7
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 6
	y = y * x
	z = z + y
	w = ws[i]
	i++
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 1
	x = x + 10
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 2
	y = y * x
	z = z + y
	w = ws[i]
	i++
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -2
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 11
	y = y * x
	z = z + y
	w = ws[i]
	i++
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -1
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 12
	y = y * x
	z = z + y
	w = ws[i]
	i++
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -4
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 3
	y = y * x
	z = z + y
	w = ws[i]
	i++
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -12
	if x == w {
		x = 0
	} else {
		x = 1
	}
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 13
	y = y * x
	z = z + y

	return z == 0
}

/**
ws[3] == ws[2] - 7 ->  1=8-7
ws[4] == ws[5] - 1 ->  1=2-1
ws[7] == ws[8] - 5 ->  1=6-5
ws[9] == ws[10]		-> 1=1
ws[6] == ws[11] - 8 -> 1=9-8
ws[1] == ws[12] - 7 -> 1=8-7
ws[13] == ws[0] - 8 -> 1=9-1

92928914999991 <- CORRECT (for pt. 1)

91811211611981 <- CORRECT (for pt. 2)
*/
