package year_2021

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_map"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"sort"
	"strconv"
	"strings"
)

func init() {
	Solutions[14] = Day14{}
}

type Day14 struct{}

func (Day14) Part1(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))

	r.Scan()
	chain := []byte(r.Text())
	r.Scan()

	transformations := map[byte]map[byte]byte{}

	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " -> ")
		from := spl[0][0]
		to := spl[0][1]
		n := []byte(spl[1])[0]
		if t, ok := transformations[from]; ok {
			t[to] = n
		} else {
			transformations[from] = map[byte]byte{to: n}
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < len(chain)-1; j++ {
			if n, ok := transformations[chain[j]][chain[j+1]]; ok {
				chain = append(chain[:j+1], append([]byte{n}, chain[j+1:]...)...)
				j++
			}
		}
	}

	cts := _map.Values(_slice.CountUnique(chain))
	sort.Ints(cts)

	return strconv.Itoa(cts[len(cts)-1] - cts[0])
}

func (Day14) Part2(input []byte) string {
	r := bufio.NewScanner(bytes.NewReader(input))

	r.Scan()
	chain := []byte(r.Text())
	r.Scan()

	transformations := map[[2]byte][2][2]byte{}

	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " -> ")
		from := spl[0][0]
		to := spl[0][1]
		n := []byte(spl[1])[0]
		transformations[[2]byte{from, to}] = [2][2]byte{{from, n}, {n, to}}
	}

	pairs := map[[2]byte]int{}
	for i := 0; i < len(chain)-1; i++ {
		pair := [2]byte{chain[i], chain[i+1]}
		if _, ok := pairs[pair]; !ok {
			pairs[pair] = 0
		}
		pairs[pair]++
	}

	for i := 0; i < 40; i++ {
		tmp := map[[2]byte]int{}
		for pair, n := range pairs {
			if newPairs, ok := transformations[pair]; ok {
				if _, ok := tmp[newPairs[0]]; !ok {
					tmp[newPairs[0]] = 0
				}
				tmp[newPairs[0]] += n

				if _, ok := tmp[newPairs[1]]; !ok {
					tmp[newPairs[1]] = 0
				}
				tmp[newPairs[1]] += n
			} else {
				if _, ok := tmp[pair]; !ok {
					tmp[pair] = 0
				}
				tmp[pair] += n
			}
		}
		pairs = tmp
	}

	ctsMap := map[byte]int{
		chain[len(chain)-1]: 1,
	}
	for pair, ct := range pairs {
		c := pair[0]
		if _, ok := ctsMap[c]; !ok {
			ctsMap[c] = 0
		}
		ctsMap[c] += ct
	}

	cts := _map.Values(ctsMap)
	sort.Ints(cts)

	return strconv.Itoa(cts[len(cts)-1] - cts[0])
}
