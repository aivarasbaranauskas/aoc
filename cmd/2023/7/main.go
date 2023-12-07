package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/aivarasbaranauskas/aoc/internal/_map"
	"github.com/aivarasbaranauskas/aoc/internal/_slice"
	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
	"sort"
)

//go:embed input.txt
var input []byte

type Hand struct {
	cards []int
	_type int
	bid   int
}

func main() {
	p1()
	p2()
}

func p2() {
	lines := bytes.Split(input, []byte("\n"))

	hands := _slice.Map(lines, func(line []byte) Hand {
		spl := bytes.Split(line, []byte(" "))
		var cards []int
		for _, c := range spl[0] {
			if c >= '2' && c <= '9' {
				cards = append(cards, int(c-'0'))
			} else {
				switch c {
				case 'A':
					cards = append(cards, 14)
				case 'K':
					cards = append(cards, 13)
				case 'Q':
					cards = append(cards, 12)
				case 'J':
					cards = append(cards, 1)
				case 'T':
					cards = append(cards, 10)
				}
			}
		}

		grouped := _slice.CountUnique(cards)
		jCt := 0 // J count
		if ct, ok := grouped[1]; ok {
			jCt = ct
			delete(grouped, 1)
		}
		counted := _slice.CountUnique(_map.Values(grouped))

		_type := 1 // high card

		if _, ok := counted[5]; ok {
			_type = 7 // five of a kind
		} else if _, ok := counted[4]; ok {
			if jCt == 1 {
				_type = 7 // five of a kind with J
			} else {
				_type = 6 // four of a kind
			}
		} else if _, ok := counted[3]; ok {
			if jCt == 2 {
				_type = 7 // five of a kind with 2 Js
			} else if jCt == 1 {
				_type = 6 // four of a kind with J
			} else {
				if _, ok := counted[2]; ok || jCt == 1 {
					_type = 5 // full house (with 3 and 2 or 3 and 1 and J)
				} else {
					_type = 4 // three of a kind
				}
			}
		} else if ct, ok := counted[2]; ok {
			if ct == 2 {
				if jCt == 1 {
					_type = 5 // full house with 2 and 2 and J
				} else {
					_type = 3 // two pair
				}
			} else {
				if jCt == 3 {
					_type = 7 // five of kind with 2 and 3J
				} else if jCt == 2 {
					_type = 6 // four of kind with 2 and 2J
				} else if jCt == 1 {
					_type = 4 // three of kind with 2 and J
				} else {
					_type = 2 // one pair
				}
			}
		} else {
			switch jCt {
			case 5:
				_type = 7 // five of a kind
			case 4:
				_type = 7 // five of a kind 4J and 1
			case 3:
				_type = 6 // four of a kind 3J and 1
			case 2:
				_type = 4 // three of kind 2J and 1
			case 1:
				_type = 2 // one pair J and 1
			}
		}

		fmt.Println(string(line), cards, _type)

		return Hand{
			cards: cards,
			_type: _type,
			bid:   optimistic.Atoi(string(spl[1])),
		}
	})

	sort.Slice(hands, func(i, j int) bool {
		if hands[i]._type == hands[j]._type {
			for k := 0; k < 5; k++ {
				if hands[i].cards[k] == hands[j].cards[k] {
					continue
				}

				return hands[i].cards[k] > hands[j].cards[k]
			}
		}

		return hands[i]._type > hands[j]._type
	})

	s := 0

	for i := 0; i < len(hands); i++ {
		s += hands[i].bid * (len(hands) - i)
	}

	fmt.Println("part 2:", s)
}

func p1() {
	lines := bytes.Split(input, []byte("\n"))

	hands := _slice.Map(lines, func(line []byte) Hand {
		spl := bytes.Split(line, []byte(" "))
		var cards []int
		for _, c := range spl[0] {
			if c >= '2' && c <= '9' {
				cards = append(cards, int(c-'0'))
			} else {
				switch c {
				case 'A':
					cards = append(cards, 14)
				case 'K':
					cards = append(cards, 13)
				case 'Q':
					cards = append(cards, 12)
				case 'J':
					cards = append(cards, 11)
				case 'T':
					cards = append(cards, 10)
				}
			}
		}

		grouped := _slice.CountUnique(_map.Values(_slice.CountUnique(cards)))

		_type := 1 // high card

		if _, ok := grouped[5]; ok {
			_type = 7 // five of a kind
		} else if _, ok := grouped[4]; ok {
			_type = 6 // four of a kind
		} else if _, ok := grouped[3]; ok {
			if _, ok := grouped[2]; ok {
				_type = 5 // full house
			} else {
				_type = 4 // three of a kind
			}
		} else if ct, ok := grouped[2]; ok {
			if ct == 2 {
				_type = 3 // two pair
			} else {
				_type = 2 // one pair
			}
		}

		return Hand{
			cards: cards,
			_type: _type,
			bid:   optimistic.Atoi(string(spl[1])),
		}
	})

	sort.Slice(hands, func(i, j int) bool {
		if hands[i]._type == hands[j]._type {
			for k := 0; k < 5; k++ {
				if hands[i].cards[k] == hands[j].cards[k] {
					continue
				}

				return hands[i].cards[k] > hands[j].cards[k]
			}
		}

		return hands[i]._type > hands[j]._type
	})

	s := 0

	for i := 0; i < len(hands); i++ {
		s += hands[i].bid * (len(hands) - i)
	}

	fmt.Println("part 1:", s)
}
