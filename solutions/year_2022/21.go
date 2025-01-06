package year_2022

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[21] = Day21{}
}

type Day21 struct{}

func (Day21) Part1(input []byte) string {
	nums := map[string]int{}
	acts := map[string]Day21Action{}

	r := bufio.NewScanner(bytes.NewReader(input))
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " ")
		if len(spl) == 2 {
			nums[spl[0][:len(spl[0])-1]] = optimistic.Atoi(spl[1])
		} else {
			acts[spl[0][:len(spl[0])-1]] = Day21Action{
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

	return strconv.Itoa(nums["root"])
}

type Day21Action struct {
	a, b, sign string
}

func (Day21) Part2(input []byte) string {
	nums := map[string]int{}
	acts := map[string]Day21Action{}

	r := bufio.NewScanner(bytes.NewReader(input))
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " ")
		if len(spl) == 2 {
			nums[spl[0][:len(spl[0])-1]] = optimistic.Atoi(spl[1])
		} else {
			acts[spl[0][:len(spl[0])-1]] = Day21Action{
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
	actsNew := map[string]Day21Action{}
	for name, act := range acts {
		if _, ok := nums[act.b]; ok {
			switch act.sign {
			case "+":
				actsNew[act.a] = Day21Action{
					a:    name,
					b:    act.b,
					sign: "-",
				}
			case "-":
				actsNew[act.a] = Day21Action{
					a:    name,
					b:    act.b,
					sign: "+",
				}
			case "*":
				actsNew[act.a] = Day21Action{
					a:    name,
					b:    act.b,
					sign: "/",
				}
			case "/":
				actsNew[act.a] = Day21Action{
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
			actsNew[act.b] = Day21Action{
				a:    name,
				b:    act.a,
				sign: "-",
			}
		case "-":
			actsNew[act.b] = Day21Action{
				a:    act.a,
				b:    name,
				sign: "-",
			}
		case "*":
			actsNew[act.b] = Day21Action{
				a:    name,
				b:    act.a,
				sign: "/",
			}
		case "/":
			actsNew[act.b] = Day21Action{
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

	return strconv.Itoa(nums["humn"])
}
