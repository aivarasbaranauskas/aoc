package year_2015

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[7] = Day7{}
}

type Day7 struct{}

type Day7Ins struct {
	// l AND r // l OR r // l LSHIFT v // l RSHIFT v // NOT l // assign v
	// 0 - assign
	// 1 - NOT
	// 2 - AND
	// 3 - OR
	// 4 - LSHIFT
	// 5 - RSHIFT
	op   uint8
	v    uint16
	l, r string
}

func (day Day7) Part1(input []byte) string {
	circuit := day.parseCircuit(input)

	result := day.calculate(circuit, map[string]uint16{}, "a")

	return strconv.Itoa(int(result))
}

func (day Day7) Part2(input []byte) string {
	circuit := day.parseCircuit(input)

	result1 := day.calculate(circuit, map[string]uint16{}, "a")

	b := circuit["b"]
	b.v = result1
	circuit["b"] = b

	result := day.calculate(circuit, map[string]uint16{}, "a")

	return strconv.Itoa(int(result))
}

func (day Day7) calculate(circuit map[string]Day7Ins, mem map[string]uint16, wire string) uint16 {
	if v, ok := mem[wire]; ok {
		return v
	}

	instruction := circuit[wire]

	var v uint16
	switch instruction.op {
	case 0:
		if instruction.l != "" {
			v = day.calculate(circuit, mem, instruction.l)
		} else {
			v = instruction.v
		}
	case 1:
		v = ^day.calculate(circuit, mem, instruction.l)
	case 2:
		var l uint16
		if instruction.l == "" {
			l = instruction.v
		} else {
			l = day.calculate(circuit, mem, instruction.l)
		}
		v = l & day.calculate(circuit, mem, instruction.r)
	case 3:
		v = day.calculate(circuit, mem, instruction.l) | day.calculate(circuit, mem, instruction.r)
	case 4:
		v = day.calculate(circuit, mem, instruction.l) << instruction.v
	case 5:
		v = day.calculate(circuit, mem, instruction.l) >> instruction.v
	}

	mem[wire] = v

	return v
}

func (day Day7) parseCircuit(input []byte) map[string]Day7Ins {
	circuit := map[string]Day7Ins{}

	for line := range bytes.Lines(input) {
		spl := bytes.Split(bytes.TrimSpace(line), []byte(" -> "))
		var instruction Day7Ins
		if bytes.HasPrefix(spl[0], []byte("NOT")) {
			instruction.op = 1
			instruction.l = string(spl[0][4:])
		} else if bytes.Index(spl[0], []byte("AND")) >= 0 {
			instruction.op = 2
			spl2 := bytes.Split(spl[0], []byte(" AND "))
			if spl2[0][0] >= '0' && spl2[0][0] <= '9' {
				instruction.v = uint16(optimistic.AtoiB(spl2[0]))
			} else {
				instruction.l = string(spl2[0])
			}

			instruction.r = string(spl2[1])
		} else if bytes.Index(spl[0], []byte("OR")) >= 0 {
			instruction.op = 3
			spl2 := bytes.Split(spl[0], []byte(" OR "))
			instruction.l = string(spl2[0])
			instruction.r = string(spl2[1])
		} else if bytes.Index(spl[0], []byte("LSHIFT")) >= 0 {
			instruction.op = 4
			spl2 := bytes.Split(spl[0], []byte(" LSHIFT "))
			instruction.l = string(spl2[0])
			instruction.v = uint16(optimistic.AtoiB(spl2[1]))
		} else if bytes.Index(spl[0], []byte("RSHIFT")) >= 0 {
			instruction.op = 5
			spl2 := bytes.Split(spl[0], []byte(" RSHIFT "))
			instruction.l = string(spl2[0])
			instruction.v = uint16(optimistic.AtoiB(spl2[1]))
		} else {
			instruction.op = 0
			if spl[0][0] >= '0' && spl[0][0] <= '9' {
				instruction.v = uint16(optimistic.AtoiB(spl[0]))
			} else {
				instruction.l = string(spl[0])
			}
		}

		circuit[string(spl[1])] = instruction
	}
	return circuit
}
