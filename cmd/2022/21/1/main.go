package main

import (
	"bufio"
	"embed"
	"fmt"
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

	nums := map[string]int{}
	acts := map[string]Action{}

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " ")
		if len(spl) == 2 {
			nums[spl[0][:len(spl[0])-1]] = optimistic.Atoi(spl[1])
		} else {
			acts[spl[0][:len(spl[0])-1]] = Action{
				a:    spl[1],
				b:    spl[3],
				sign: spl[2],
			}
		}
	}

	for len(acts) > 0 {
		for name, act := range acts {
			aN, ok := nums[act.a]
			if !ok {
				continue
			}
			bN, ok := nums[act.b]
			if !ok {
				continue
			}
			switch act.sign {
			case "+":
				nums[name] = aN + bN
			case "-":
				nums[name] = aN - bN
			case "*":
				nums[name] = aN * bN
			case "/":
				nums[name] = aN / bN
			}
			delete(acts, name)
		}
	}

	fmt.Println(nums["root"])
}

type Action struct {
	a, b, sign string
}
