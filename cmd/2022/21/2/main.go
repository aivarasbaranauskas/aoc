package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strings"
)

//go:embed input.txt
var inputData embed.FS

func main() {
	f, err := inputData.Open("input.txt")
	_a.CheckErr(err)

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

	delete(nums, "humn")

	do := true
	for do {
		do = false

		for name, act := range acts {
			aN, ok := nums[act.a]
			if !ok {
				continue
			}
			bN, ok := nums[act.b]
			if !ok {
				continue
			}

			do = true
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

	// change root to num
	if a, ok := nums[acts["root"].a]; ok {
		nums[acts["root"].b] = a
	} else {
		// naively assume
		nums[acts["root"].a] = nums[acts["root"].b]
	}
	delete(acts, "root")

	// flip all acts
	actsNew := map[string]Action{}
	for name, act := range acts {
		if _, ok := nums[act.b]; ok {
			switch act.sign {
			case "+":
				actsNew[act.a] = Action{
					a:    name,
					b:    act.b,
					sign: "-",
				}
			case "-":
				actsNew[act.a] = Action{
					a:    name,
					b:    act.b,
					sign: "+",
				}
			case "*":
				actsNew[act.a] = Action{
					a:    name,
					b:    act.b,
					sign: "/",
				}
			case "/":
				actsNew[act.a] = Action{
					a:    name,
					b:    act.b,
					sign: "*",
				}
			}
			// n = H + b -> H = n - b
			// n = H - b -> H = n + b
			// n = H * b -> H = n / b
			// n = H / b -> H = n * b
			continue
		}

		switch act.sign {
		case "+":
			actsNew[act.b] = Action{
				a:    name,
				b:    act.a,
				sign: "-",
			}
		case "-":
			actsNew[act.b] = Action{
				a:    act.a,
				b:    name,
				sign: "-",
			}
		case "*":
			actsNew[act.b] = Action{
				a:    name,
				b:    act.a,
				sign: "/",
			}
		case "/":
			actsNew[act.b] = Action{
				a:    act.a,
				b:    name,
				sign: "/",
			}
		}
		// n = a + H -> H = n - a
		// n = a - H -> H = a - n
		// n = a * H -> H = n / a
		// n = a / H -> H = a / n
	}

	acts = actsNew

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

	fmt.Println(nums["humn"])
}

type Action struct {
	a, b, sign string
}
