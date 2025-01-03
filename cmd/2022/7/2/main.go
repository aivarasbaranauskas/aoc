package main

import (
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"io"
	"log"
	"sort"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

	fullFileB, err := io.ReadAll(f)
	_a.CheckErr(err)

	fullFile := string(fullFileB)
	fileSpl := strings.Split(fullFile, "$")

	root := NewDir("/", nil)
	var current *Dir
	dirs := []*Dir{root}

	for x, part := range fileSpl {
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
						newD := NewDir(fName, current)
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

	var dirSizes []int
	for _, d := range dirs {
		dirSizes = append(dirSizes, d.Size())
	}
	sort.Ints(dirSizes)

	total := 70000000
	wantFree := 30000000
	isFree := total - root.Size()
	needToFree := wantFree - isFree
	x := sort.SearchInts(dirSizes, needToFree)

	fmt.Println(x, dirSizes[x])
}

type Dir struct {
	name   string
	parent *Dir
	dirs   map[string]*Dir
	files  map[string]int
}

func NewDir(name string, parent *Dir) *Dir {
	return &Dir{
		name:   name,
		parent: parent,
		dirs:   map[string]*Dir{},
		files:  map[string]int{},
	}
}

func (d *Dir) Size() (s int) {
	for _, fileSize := range d.files {
		s += fileSize
	}

	for _, innerD := range d.dirs {
		s += innerD.Size()
	}
	return
}
