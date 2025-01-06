package year_2021

import (
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_num"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"slices"
	"strconv"
)

func init() {
	Solutions[16] = Day16{}
}

type Day16 struct{}

func (d Day16) Part1(input []byte) string {
	return strconv.Itoa(d.VersionSum(d.readPackets(input)))
}

func (d Day16) Part2(input []byte) string {
	return strconv.Itoa(d.Calc(d.readPackets(input)))
}

func (d Day16) readPackets(input []byte) Day16PacketI {
	inputBin := make([]byte, 0, len(input)*4)
	for _, c := range input {
		inputBin = append(
			inputBin,
			[]byte(fmt.Sprintf("%04b", optimistic.ParseInt(string([]byte{c}), 16, 0)))...,
		)
	}

	p := Day16ParsePacket(&inputBin)
	return p
}

func (d Day16) VersionSum(p Day16PacketI) int {
	if pL, ok := p.(*Day16LiteralPacket); ok {
		return pL.version
	}
	pO := p.(*Day16OperatorPacket)
	sum := pO.version
	for _, innerP := range pO.packets {
		sum += d.VersionSum(innerP)
	}
	return sum
}

func (d Day16) Calc(p Day16PacketI) int {
	if pL, ok := p.(*Day16LiteralPacket); ok {
		return pL.value
	}
	pO := p.(*Day16OperatorPacket)

	values := _slice.Map(pO.packets, d.Calc)

	switch pO.typeId {
	case 0:
		return _num.Sum(values...)
	case 1:
		return _num.Product(values...)
	case 2:
		return slices.Min(values)
	case 3:
		return slices.Max(values)
	case 5:
		if values[0] > values[1] {
			return 1
		}
		return 0
	case 6:
		if values[0] < values[1] {
			return 1
		}
		return 0
	case 7:
		if values[0] == values[1] {
			return 1
		}
		return 0
	}
	return -1
}

type Day16LiteralPacket struct {
	Day16Packet
	value int
}

type Day16OperatorPacket struct {
	Day16Packet
	packets []Day16PacketI
}

type Day16Packet struct {
	Day16PacketI
	version, typeId int
}

type Day16PacketI interface{}

func Day16ParsePacket(signal *[]byte) (p Day16PacketI) {
	version := int(optimistic.ParseInt(string((*signal)[:3]), 2, 0))
	*signal = (*signal)[3:]
	typeId := int(optimistic.ParseInt(string((*signal)[:3]), 2, 0))
	*signal = (*signal)[3:]

	if typeId == 4 {
		hasNext := byte('1')
		var lit []byte
		for hasNext == '1' {
			hasNext = (*signal)[0]
			lit = append(lit, (*signal)[1:5]...)
			*signal = (*signal)[5:]
		}

		p = &Day16LiteralPacket{
			Day16Packet: Day16Packet{version: version, typeId: typeId},
			value:       int(optimistic.ParseInt(string(lit), 2, 0)),
		}
	} else {
		lengthTypeId := (*signal)[0]
		*signal = (*signal)[1:]

		pT := &Day16OperatorPacket{
			Day16Packet: Day16Packet{version: version, typeId: typeId},
		}

		if lengthTypeId == '0' {
			subPacketsLength := int(optimistic.ParseInt(string((*signal)[:15]), 2, 0))
			*signal = (*signal)[15:]
			end := len(*signal) - subPacketsLength
			for len(*signal) > end {
				pT.packets = append(pT.packets, Day16ParsePacket(signal))
			}
		} else {
			subPacketsCount := int(optimistic.ParseInt(string((*signal)[:11]), 2, 0))
			*signal = (*signal)[11:]

			for i := 0; i < subPacketsCount; i++ {
				pT.packets = append(pT.packets, Day16ParsePacket(signal))
			}
		}
		p = pT
	}

	return p
}
