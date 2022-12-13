package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"log"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	r := bufio.NewScanner(f)
	var out int
	i := 1
	for r.Scan() {
		line := r.Text()
		if line == "" {
			continue
		}

		s1 := ParseLine(line)
		r.Scan()
		s2 := ParseLine(r.Text())

		fmt.Println(s1)
		fmt.Println(s2)
		if Compare(s1, s2) == 1 {
			out += i
		}
		i++
	}
	fmt.Println(out)
}

func Compare(l, r *Node) int {
	if l.IsList() && r.IsList() {
		i := 0
		for ; i < len(l.l) && i < len(r.l); i++ {
			c := Compare(l.l[i], r.l[i])
			if c == 0 {
				continue
			}
			return c
		}

		if i == len(l.l) {
			if i == len(r.l) {
				return 0
			}
			return 1
		}
		return -1
	}

	if !l.IsList() && !r.IsList() {
		if *l.v < *r.v {
			return 1
		}
		if *l.v == *r.v {
			return 0
		}
		return -1
	}

	if l.IsList() {
		return Compare(l, &Node{l: []*Node{r}})
	}

	return Compare(&Node{l: []*Node{l}}, r)
}

type Node struct {
	v      *int
	l      []*Node
	parent *Node
}

func (n *Node) IsList() bool {
	return n.v == nil
}

func (n *Node) String() string {
	b := strings.Builder{}
	b.WriteString("[")
	if len(n.l) > 0 {
		b.WriteString(n.l[0].String())
		for i := 1; i < len(n.l); i++ {
			b.WriteString(",")
			b.WriteString(n.l[i].String())
		}
	} else if n.v != nil {
		b.WriteString(fmt.Sprintf("%v", *n.v))
	}
	b.WriteString("]")
	return b.String()
}

func ParseLine(line string) *Node {
	s := &Node{}

	for i := 1; i < len(line)-1; i++ {
		if line[i] == '[' {
			n := &Node{parent: s}
			s.l = append(s.l, n)
			s = n
			continue
		}
		if line[i] == ']' {
			s = s.parent
			continue
		}
		if line[i] == ',' {
			continue
		}

		x1 := strings.IndexByte(line[i:], ',')
		x2 := strings.IndexByte(line[i:], ']')
		x := _num.Min(_slice.Filter([]int{x1, x2}, func(x int) bool {
			return x >= 0
		})...)
		s.l = append(s.l, &Node{v: ip(optimistic.Atoi(line[i : i+x]))})
		i = i + x - 1
	}

	return s
}

func ip(i int) *int {
	return &i
}
