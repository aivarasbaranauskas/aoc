package year_2015

import (
	"maps"
	"strconv"
	"time"

	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/_set"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[20] = Day20{}
}

type Day20 struct{}

func (Day20) Part1(input []byte) string {
	num := optimistic.AtoiBFast(input)

	mem := map[int]_set.Set[int]{}
	s := 0
	i := 1

	p := _a.NewProgress(time.Second, 0)
	for ; s < num; i++ {
		memT := _set.FromSlice([]int{1, i})
		for a := 2; a < i; a++ {
			if i%a != 0 {
				continue
			}
			b := i / a
			memT.Union(mem[a])
			if a != b {
				memT.Union(mem[b])
			}
		}
		mem[i] = memT
		s = 0
		for v := range maps.Keys(memT) {
			s += v
		}
		s *= 10
		p.Inc()
	}
	p.Stop()

	return strconv.Itoa(i - 1)
}

func (Day20) Part2(input []byte) string {
	num := optimistic.AtoiBFast(input)

	mem := map[int]_set.Set[int]{}
	finishedMem := _set.New[int]()
	usedTracker := map[int]int{}
	s := 0
	i := 1

	p := _a.NewProgress(time.Second, 0)
	for ; s < num; i++ {
		memT := _set.FromSlice([]int{1, i})
		for a := 2; a < i; a++ {
			if i%a != 0 {
				continue
			}
			b := i / a
			memT.Union(mem[a])
			if a != b {
				memT.Union(mem[b])
			}
		}
		memT.Difference(finishedMem)
		mem[i] = memT
		s = 0
		for v := range maps.Keys(memT) {
			s += v

			if _, ok := usedTracker[v]; ok {
				usedTracker[v]++
				if usedTracker[v] == 50 {
					delete(usedTracker, v)
					finishedMem.Add(v)
				}
			} else {
				usedTracker[v] = 1
			}
		}
		s *= 11
		p.Inc()
	}
	p.Stop()

	return strconv.Itoa(i - 1)
}
