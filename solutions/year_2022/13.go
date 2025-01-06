package year_2022

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func init() {
	Solutions[13] = Day13{}
}

type Day13 struct{}

func (d Day13) Part1(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))
	var out int
	i := 1
	for r.Scan() {
		line := r.Text()
		if line == "" {
			continue
		}

		s1 := d.ParseLine(line)
		r.Scan()
		s2 := d.ParseLine(r.Text())

		if d.Compare(s1, s2) == 1 {
			out += i
		}
		i++
	}

	return strconv.Itoa(out)
}

func (d Day13) Part2(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))

	d1 := d.ParseLine("[[2]]")
	d2 := d.ParseLine("[[6]]")
	p := []*Day13Node{d1, d2}

	for r.Scan() {
		line := r.Text()
		if line == "" {
			continue
		}
		p = append(p, d.ParseLine(line))
	}

	sort.Slice(p, func(i, j int) bool {
		return d.Compare(p[i], p[j]) == 1
	})

	var d1i, d2i int
	for i := range p {
		if p[i] == d1 {
			d1i = i + 1
		}
		if p[i] == d2 {
			d2i = i + 1
		}
	}

	return strconv.Itoa(d1i * d2i)
}

func (d Day13) Compare(l, r *Day13Node) int {
	if l.IsList() && r.IsList() {
		i := 0
		for ; i < len(l.l) && i < len(r.l); i++ {
			c := d.Compare(l.l[i], r.l[i])
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
		return d.Compare(l, &Day13Node{l: []*Day13Node{r}})
	}

	return d.Compare(&Day13Node{l: []*Day13Node{l}}, r)
}

type Day13Node struct {
	v      *int
	l      []*Day13Node
	parent *Day13Node
}

func (n *Day13Node) IsList() bool {
	return n.v == nil
}

func (n *Day13Node) String() string {
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

func (Day13) ParseLine(line string) *Day13Node {
	s := &Day13Node{}

	for i := 1; i < len(line)-1; i++ {
		if line[i] == '[' {
			n := &Day13Node{parent: s}
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
		x := slices.Min(_slice.Filter([]int{x1, x2}, func(x int) bool {
			return x >= 0
		}))
		v := optimistic.Atoi(line[i : i+x])
		s.l = append(s.l, &Day13Node{v: &v})
		i = i + x - 1
	}

	return s
}
