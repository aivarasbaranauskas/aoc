package main

import (
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/go_helpers/o"
)

//go:embed input.txt
var inputHex []byte

func main() {
	inputBin := make([]byte, 0, len(inputHex)*4)
	for _, c := range inputHex {
		inputBin = append(
			inputBin,
			[]byte(fmt.Sprintf("%04b", o.ParseInt(string([]byte{c}), 16, 0)))...,
		)
	}

	p := ParsePacket(&inputBin)

	fmt.Println(VersionSum(p))
}

func VersionSum(p PacketI) int {
	if pL, ok := p.(*LiteralPacket); ok {
		return pL.version
	}
	pO := p.(*OperatorPacket)
	sum := pO.version
	for _, innerP := range pO.packets {
		sum += VersionSum(innerP)
	}
	return sum
}

type LiteralPacket struct {
	Packet
	value int
}

type OperatorPacket struct {
	Packet
	packets []PacketI
}

type Packet struct {
	PacketI
	version, typeId int
}

type PacketI interface{}

func ParsePacket(signal *[]byte) (p PacketI) {
	version := int(o.ParseInt(string((*signal)[:3]), 2, 0))
	*signal = (*signal)[3:]
	typeId := int(o.ParseInt(string((*signal)[:3]), 2, 0))
	*signal = (*signal)[3:]

	if typeId == 4 {
		hasNext := byte('1')
		var lit []byte
		for hasNext == '1' {
			hasNext = (*signal)[0]
			lit = append(lit, (*signal)[1:5]...)
			*signal = (*signal)[5:]
		}

		p = &LiteralPacket{
			Packet: Packet{version: version, typeId: typeId},
			value:  int(o.ParseInt(string(lit), 2, 0)),
		}
	} else {
		lengthTypeId := (*signal)[0]
		*signal = (*signal)[1:]

		pT := &OperatorPacket{
			Packet: Packet{version: version, typeId: typeId},
		}

		if lengthTypeId == '0' {
			subPacketsLength := int(o.ParseInt(string((*signal)[:15]), 2, 0))
			*signal = (*signal)[15:]
			end := len(*signal) - subPacketsLength
			for len(*signal) > end {
				pT.packets = append(pT.packets, ParsePacket(signal))
			}
		} else {
			subPacketsCount := int(o.ParseInt(string((*signal)[:11]), 2, 0))
			*signal = (*signal)[11:]

			for i := 0; i < subPacketsCount; i++ {
				pT.packets = append(pT.packets, ParsePacket(signal))
			}
		}
		p = pT
	}

	return p
}
