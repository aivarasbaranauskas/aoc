package year_2021

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func init() {
	Solutions[18] = Day18{}
}

type Day18 struct{}

func (d Day18) Part1(input []byte) string {
	nums := d.readData(input)

	acc := nums[0]
	nums = nums[1:]

	for _, num := range nums {
		acc = &Day18Node{ln: acc, rn: num}
		acc.ln.parent = acc
		acc.rn.parent = acc
		e := true
		for e {
			e = acc.Explode(0)
			if e {
				continue
			}
			e = acc.Split()
		}
	}

	return strconv.Itoa(acc.Magnitude())
}

func (d Day18) Part2(input []byte) string {
	nums := d.readData(input)
	var maxVal int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			a := &Day18Node{ln: nums[i].Clone(), rn: nums[j].Clone()}
			a.ln.parent = a
			a.rn.parent = a

			e := true
			for e {
				e = a.Explode(0)
				if e {
					continue
				}
				if e = a.Split(); e {
				}
			}

			maxVal = max(maxVal, a.Magnitude())
		}
	}

	return strconv.Itoa(maxVal)
}

func (d Day18) readData(input []byte) []*Day18Node {
	var nums []*Day18Node
	r := bufio.NewScanner(bytes.NewReader(input))
	for r.Scan() {
		line := []byte(r.Text())
		num := &Day18Node{}

		for i := 1; i < len(line)-1; i++ {
			if line[i] == '[' {
				num.ln = &Day18Node{parent: num}
				num = num.ln
				continue
			}
			if line[i] == ']' {
				num = num.parent
				continue
			}
			if line[i] == ',' {
				i++
				if line[i] == '[' {
					num.rn = &Day18Node{parent: num}
					num = num.rn
					continue
				}

				num.r = int(line[i] - '0')
				continue
			}
			num.l = int(line[i] - '0')
		}
		nums = append(nums, num)
	}

	return nums
}

type Day18Node struct {
	l, r   int
	ln, rn *Day18Node
	parent *Day18Node
}

func (n *Day18Node) String() string {
	b := strings.Builder{}
	b.WriteString("[")
	if n.ln != nil {
		b.WriteString(n.ln.String())
	} else {
		b.WriteString(fmt.Sprintf("%v", n.l))
	}
	b.WriteString(",")
	if n.rn != nil {
		b.WriteString(n.rn.String())
	} else {
		b.WriteString(fmt.Sprintf("%v", n.r))
	}
	b.WriteString("]")
	return b.String()
}

func (n *Day18Node) Clone() *Day18Node {
	c := &Day18Node{
		l: n.l,
		r: n.r,
	}
	if n.ln != nil {
		c.ln = n.ln.Clone()
		c.ln.parent = c
	}
	if n.rn != nil {
		c.rn = n.rn.Clone()
		c.rn.parent = c
	}
	return c
}

func (n *Day18Node) Explode(d int) bool {
	d++
	if n.ln == nil && n.rn == nil {
		if d <= 4 {
			return false
		}
		//fmt.Println("e:", n)
		// add left
		p, c := n, n.parent
		for c != nil {
			if c.ln == nil {
				c.l += n.l
				break
			}
			if c.ln != p {
				c = c.ln
				for c.rn != nil {
					c = c.rn
				}
				c.r += n.l
				break
			}
			p, c = c, c.parent
		}

		// add right
		p, c = n, n.parent
		for c != nil {
			if c.rn == nil {
				c.r += n.r
				break
			}
			if c.rn != p {
				c = c.rn
				for c.ln != nil {
					c = c.ln
				}
				c.l += n.r
				break
			}
			p, c = c, c.parent
		}

		//remove node
		if n.parent == nil {
			panic("wth")
		}
		p = n.parent
		if p.ln == n {
			p.ln = nil
		} else {
			p.rn = nil
		}

		return true
	}
	if n.ln != nil {
		if boom := n.ln.Explode(d); boom {
			return true
		}
	}
	if n.rn != nil {
		return n.rn.Explode(d)
	}
	return false
}

func (n *Day18Node) Magnitude() int {
	l, r := n.l, n.r
	if n.ln != nil {
		l = n.ln.Magnitude()
	}
	if n.rn != nil {
		r = n.rn.Magnitude()
	}
	return 3*l + 2*r
}

func (n *Day18Node) Split() (s bool) {
	if n.l >= 10 {
		//fmt.Println("s:", n.l, n)
		n.ln = &Day18Node{
			l:      int(math.Floor(float64(n.l) / 2)),
			r:      int(math.Ceil(float64(n.l) / 2)),
			parent: n,
		}
		n.l = 0
		return true
	}

	if n.ln != nil && n.ln.Split() {
		return true
	}

	if n.r >= 10 {
		//fmt.Println("s:", n.r, n)
		n.rn = &Day18Node{
			l:      int(math.Floor(float64(n.r) / 2)),
			r:      int(math.Ceil(float64(n.r) / 2)),
			parent: n,
		}
		n.r = 0
		return true
	}
	return n.rn != nil && n.rn.Split()
}
