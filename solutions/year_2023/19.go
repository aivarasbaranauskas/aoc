package year_2023

import (
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"strconv"
	"strings"
)

func init() {
	Solutions[19] = Day19{}
}

type Day19 struct{}

func (Day19) Part1(input []byte) string {
	parts := strings.Split(string(input), "\n\n")
	workflowsRaw := make(map[string]string)
	for _, line := range strings.Split(parts[0], "\n") {
		i := strings.Index(line, "{")
		workflowsRaw[line[:i]] = line[i:]
	}

	workflows := map[string]NodeI{
		"A": &EndNode{true},
		"R": &EndNode{false},
	}
	mainWorkflow := buildWorkflow(&workflows, workflowsRaw, "in")

	p1 := 0
	for _, line := range strings.Split(parts[1], "\n") {
		item := make(map[string]int)
		for _, s := range strings.Split(line[1:len(line)-1], ",") {
			spl := strings.Split(s, "=")
			item[spl[0]] = optimistic.Atoi(spl[1])
		}
		if mainWorkflow.Eval(item) {
			for _, v := range item {
				p1 += v
			}
		}
	}

	return strconv.Itoa(p1)
}

func (Day19) Part2(input []byte) string {
	workflowsRaw := make(map[string]string)
	for _, line := range strings.Split(strings.Split(string(input), "\n\n")[0], "\n") {
		i := strings.Index(line, "{")
		workflowsRaw[line[:i]] = line[i:]
	}

	workflows := map[string]NodeI{
		"A": &EndNode{true},
		"R": &EndNode{false},
	}
	mainWorkflow := buildWorkflow(&workflows, workflowsRaw, "in")

	ranges := map[string]*[2]int{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}

	possibleRanges := mainWorkflow.GetPossible(ranges)
	p2 := 0

	for _, r := range possibleRanges {
		s := 1
		for _, rr := range r {
			s *= rr[1] - rr[0] + 1
		}
		p2 += s
	}

	return strconv.Itoa(p2)
}

func buildWorkflow(workflows *map[string]NodeI, workflowsRaw map[string]string, current string) NodeI {
	if wf, ok := (*workflows)[current]; ok {
		return wf
	}

	wfRaw, ok := workflowsRaw[current]
	if !ok {
		panic("AAAAAAAAAAAAAAA")
	}

	var wf NodeI
	conds := strings.Split(wfRaw[1:len(wfRaw)-1], ",")
	for i := len(conds) - 1; i >= 0; i-- {
		cond := conds[i]
		si := strings.IndexAny(cond, "<>")
		if si == -1 {
			wf = buildWorkflow(workflows, workflowsRaw, cond)
			continue
		}

		spl := strings.Split(cond, ":")
		wf = &WorkflowNode{
			variable:  spl[0][:si],
			sign:      spl[0][si : si+1],
			c:         optimistic.Atoi(spl[0][si+1:]),
			nextTrue:  buildWorkflow(workflows, workflowsRaw, spl[1]),
			nextFalse: wf,
		}
	}

	(*workflows)[current] = wf

	return wf
}

type WorkflowNode struct {
	variable            string
	sign                string
	c                   int
	nextTrue, nextFalse NodeI
}

func (w WorkflowNode) Eval(m map[string]int) bool {
	v, ok := m[w.variable]
	if !ok {
		panic("AJAJAY")
	}

	var match bool
	switch w.sign {
	case "<":
		match = v < w.c
	case ">":
		match = v > w.c
	default:
		panic(fmt.Sprintf("unknown sign: %v", w.sign))
	}

	if match {
		return w.nextTrue.Eval(m)
	}

	return w.nextFalse.Eval(m)
}

func (w WorkflowNode) GetPossible(ranges map[string]*[2]int) []map[string][2]int {
	var a, b []map[string][2]int
	switch w.sign {
	case "<":
		// ///////////////////c.......................
		// .............[0.........1].................
		if ranges[w.variable][0] < w.c {
			tmp := ranges[w.variable][1]
			ranges[w.variable][1] = w.c - 1
			a = w.nextTrue.GetPossible(ranges)
			ranges[w.variable][1] = tmp
		}

		if ranges[w.variable][1] >= w.c {
			tmp := ranges[w.variable][0]
			ranges[w.variable][0] = w.c
			b = w.nextFalse.GetPossible(ranges)
			ranges[w.variable][0] = tmp
		}
	case ">":
		// ...................c///////////////////////
		// .............[0.........1].................
		if ranges[w.variable][1] > w.c {
			tmp := ranges[w.variable][0]
			ranges[w.variable][0] = w.c + 1
			a = w.nextTrue.GetPossible(ranges)
			ranges[w.variable][0] = tmp
		}

		if ranges[w.variable][0] <= w.c {
			tmp := ranges[w.variable][1]
			ranges[w.variable][1] = w.c
			b = w.nextFalse.GetPossible(ranges)
			ranges[w.variable][1] = tmp
		}
	default:
		panic(fmt.Sprintf("unknown sign: %v", w.sign))
	}

	return append(a, b...)
}

type EndNode struct {
	accepted bool
}

func (e EndNode) Eval(_ map[string]int) bool {
	return e.accepted
}

func (e EndNode) GetPossible(ranges map[string]*[2]int) []map[string][2]int {
	if e.accepted {
		return []map[string][2]int{
			{
				"x": {ranges["x"][0], ranges["x"][1]},
				"m": {ranges["m"][0], ranges["m"][1]},
				"a": {ranges["a"][0], ranges["a"][1]},
				"s": {ranges["s"][0], ranges["s"][1]},
			},
		}
	}
	return nil
}

type NodeI interface {
	Eval(map[string]int) bool
	GetPossible(map[string]*[2]int) []map[string][2]int
}
