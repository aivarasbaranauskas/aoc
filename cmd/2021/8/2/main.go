package main

import (
	"bufio"
	"embed"
	"fmt"
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

	var sum int

	r := bufio.NewScanner(f)
	for r.Scan() {
		line := r.Text()
		spl := strings.Split(line, " | ")
		outputSignals := strings.Split(spl[1], " ")
		spl = append(strings.Split(spl[0], " "), outputSignals...)

		var patterns [10]uint8
		p069 := make(map[uint8]struct{})
		p235 := make(map[uint8]struct{})

		for _, s := range spl {
			pattern := signalsBin(s)

			switch len(s) {
			case 2:
				patterns[1] = pattern
			case 3:
				patterns[7] = pattern
			case 4:
				patterns[4] = pattern
			case 5:
				p235[pattern] = struct{}{}
			case 6:
				p069[pattern] = struct{}{}
			case 7:
				patterns[8] = pattern
			default:
				log.Fatal(s)
			}
		}

		if len(p235) != 3 {
			log.Fatal("can't find 235")
		}

		// find 3
		for p := range p235 {
			if patterns[1]&p == patterns[1] {
				patterns[3] = p
				delete(p235, p)
				break
			}
		}
		if patterns[3] == 0 {
			log.Fatal("3 not found")
		}

		// find 5
		for p := range p235 {
			if countBitsSet(patterns[4]&p) == 3 {
				patterns[5] = p
				delete(p235, p)
				break
			}
		}
		if patterns[5] == 0 {
			log.Fatal("5 not found")
		}

		// left with 2
		for p := range p235 {
			patterns[2] = p
		}
		if patterns[2] == 0 {
			log.Fatal("2 not found")
		}

		if len(p069) != 3 {
			log.Fatal("can't find p069")
		}

		// find 6
		for p := range p069 {
			if countBitsSet(patterns[1]&p) == 1 {
				patterns[6] = p
				delete(p069, p)
				break
			}
		}
		if patterns[6] == 0 {
			log.Fatal("6 not found")
		}

		// find 9
		for p := range p069 {
			if countBitsSet(patterns[4]&p) == 4 {
				patterns[9] = p
				delete(p069, p)
				break
			}
		}
		if patterns[9] == 0 {
			log.Fatal("9 not found")
		}

		// left with 0
		for p := range p069 {
			patterns[0] = p
		}
		if patterns[0] == 0 {
			log.Fatal("0 not found")
		}

		patternsMap := make(map[uint8]int)
		for i, p := range patterns {
			patternsMap[p] = i
		}

		output := patternsMap[signalsBin(outputSignals[0])]*1000 + patternsMap[signalsBin(outputSignals[1])]*100 + patternsMap[signalsBin(outputSignals[2])]*10 + patternsMap[signalsBin(outputSignals[3])]
		sum += output
	}

	fmt.Println(sum)
}

func countBitsSet(input uint8) (ct int) {
	for i := 0; i < 8; i++ {
		if input&(1<<i) == 1<<i {
			ct++
		}
	}
	return
}

func signalsBin(input string) uint8 {
	var pattern uint8
	for _, c := range []byte(input) {
		switch c {
		case 'a':
			pattern |= 1
		case 'b':
			pattern |= 1 << 2
		case 'c':
			pattern |= 1 << 3
		case 'd':
			pattern |= 1 << 4
		case 'e':
			pattern |= 1 << 5
		case 'f':
			pattern |= 1 << 6
		case 'g':
			pattern |= 1 << 7
		}
	}
	return pattern
}
