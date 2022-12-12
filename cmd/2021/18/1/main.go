package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"math"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var nums []*Node

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := []byte(r.Text())
		num := &Node{}

		for i := 1; i < len(line)-1; i++ {
			if line[i] == '[' {
				num.ln = &Node{parent: num}
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
					num.rn = &Node{parent: num}
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

	acc := nums[0]
	nums = nums[1:]

	for _, num := range nums {
		acc = &Node{ln: acc, rn: num}
		acc.ln.parent = acc
		acc.rn.parent = acc
		e := true
		//fmt.Println(acc)
		for e {
			e = acc.Explode(0)
			if e {
				//fmt.Println("E:", acc)
				continue
			}
			if e = acc.Split(); e {
				//fmt.Println("S:", acc)
			}
		}
		//fmt.Println(acc)
		//fmt.Println()
	}

	fmt.Println(acc)
	fmt.Println(acc.Magnitude())
}

type Node struct {
	l, r   int
	ln, rn *Node
	parent *Node
}

func (n *Node) String() string {
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

func (n *Node) Explode(d int) bool {
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

func (n *Node) Magnitude() int {
	l, r := n.l, n.r
	if n.ln != nil {
		l = n.ln.Magnitude()
	}
	if n.rn != nil {
		r = n.rn.Magnitude()
	}
	return 3*l + 2*r
}

func (n *Node) Split() (s bool) {
	if n.l >= 10 {
		//fmt.Println("s:", n.l, n)
		n.ln = &Node{
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
		n.rn = &Node{
			l:      int(math.Floor(float64(n.r) / 2)),
			r:      int(math.Ceil(float64(n.r) / 2)),
			parent: n,
		}
		n.r = 0
		return true
	}
	return n.rn != nil && n.rn.Split()
}
