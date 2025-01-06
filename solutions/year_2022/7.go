package year_2022

import (
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"log"
	"sort"
	"strconv"
	"strings"
)

func init() {
	Solutions[7] = Day7{}
}

type Day7 struct{}

func (d Day7) Part1(input []byte) string {
	dirs, _ := d.parseData(input)
	var sum int
	for _, dir := range dirs {
		size := dir.Size()
		if size <= 100000 {
			sum += size
		}
	}
	return strconv.Itoa(sum)
}

func (d Day7) Part2(input []byte) string {
	dirs, root := d.parseData(input)
	var dirSizes []int
	for _, dir := range dirs {
		dirSizes = append(dirSizes, dir.Size())
	}
	sort.Ints(dirSizes)

	total := 70000000
	wantFree := 30000000
	isFree := total - root.Size()
	needToFree := wantFree - isFree
	x := sort.SearchInts(dirSizes, needToFree)

	return strconv.Itoa(dirSizes[x])
}

func (Day7) parseData(input []byte) ([]*Day7Dir, *Day7Dir) {
	inputSpl := strings.Split(string(input), "$")

	root := NewDay7Dir("/", nil)
	var current *Day7Dir
	dirs := []*Day7Dir{root}

	for x, part := range inputSpl {
		if part == "" {
			//start of file...
			continue
		}

		spl := strings.Split(strings.TrimSpace(part), "\n")
		if len(spl) == 1 {
			splLine := strings.Split(strings.TrimSpace(spl[0]), " ")
			command := strings.TrimSpace(splLine[0])

			if command == "cd" {
				dest := strings.TrimSpace(splLine[1])
				if dest == "/" {
					current = root
				} else if dest == ".." {
					current = current.parent
				} else {
					current = current.dirs[dest]
				}
			} else if command == "ls" {
				// empty ls
			} else {
				log.Fatalf("wtf, %v, '%v'\n", x, part)
			}
		} else {
			command := strings.TrimSpace(spl[0])

			if command == "ls" {
				for i := 1; i < len(spl); i++ {
					splLine := strings.Split(spl[i], " ")
					fName := splLine[1]
					if splLine[0] == "dir" {
						newD := NewDay7Dir(fName, current)
						current.dirs[fName] = newD
						dirs = append(dirs, newD)
					} else {
						current.files[fName] = optimistic.Atoi(splLine[0])
					}
				}
			} else {
				log.Fatalf("wtf, %v, '%v'\n", x, part)
			}
		}
	}

	return dirs, root
}

type Day7Dir struct {
	name   string
	parent *Day7Dir
	dirs   map[string]*Day7Dir
	files  map[string]int
}

func NewDay7Dir(name string, parent *Day7Dir) *Day7Dir {
	return &Day7Dir{
		name:   name,
		parent: parent,
		dirs:   map[string]*Day7Dir{},
		files:  map[string]int{},
	}
}

func (d *Day7Dir) Size() (s int) {
	for _, fileSize := range d.files {
		s += fileSize
	}

	for _, innerD := range d.dirs {
		s += innerD.Size()
	}
	return
}
