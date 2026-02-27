package year_2015

import (
	"bytes"
	"math"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[21] = Day21{
		weapons: []Day21Item{
			{
				name:   "Dagger",
				cost:   8,
				damage: 4,
			},
			{
				name:   "Shortsword",
				cost:   10,
				damage: 5,
			},
			{
				name:   "Warhammer",
				cost:   25,
				damage: 6,
			},
			{
				name:   "Longsword",
				cost:   40,
				damage: 7,
			},
			{
				name:   "Greataxe",
				cost:   74,
				damage: 8,
			},
		},
		armors: []Day21Item{
			{
				name:  "Leather",
				cost:  13,
				armor: 1,
			},
			{
				name:  "Chainmail",
				cost:  31,
				armor: 2,
			},
			{
				name:  "Splintmail",
				cost:  53,
				armor: 3,
			},
			{
				name:  "Bandedmail",
				cost:  75,
				armor: 4,
			},
			{
				name:  "Platemail",
				cost:  102,
				armor: 5,
			},
		},
		rings: []Day21Item{
			{
				name:   "Damage +1",
				cost:   25,
				damage: 1,
			},
			{
				name:   "Damage +2",
				cost:   50,
				damage: 2,
			},
			{
				name:   "Damage +3",
				cost:   100,
				damage: 3,
			},
			{
				name:  "Defense +1",
				cost:  20,
				armor: 1,
			},
			{
				name:  "Defense +2",
				cost:  40,
				armor: 2,
			},
			{
				name:  "Defense +4",
				cost:  80,
				armor: 3,
			},
		},
	}
}

type Day21Item struct {
	name                string
	cost, damage, armor int
}

type Day21Stats struct {
	hitPoints, damage, armor int
}

type Day21 struct {
	weapons []Day21Item
	armors  []Day21Item
	rings   []Day21Item
}

func (day Day21) Part1(input []byte) string {
	boss := day.parse(input)
	minPrice := math.MaxInt

	checkItemsCombo := func(items ...Day21Item) {
		player, cost := day.calculateStats(items)
		player.hitPoints = 100
		if day.simulate(player, boss) && cost < minPrice {
			minPrice = cost
		}
	}

	day.withAllItemsCombos(checkItemsCombo)

	return strconv.Itoa(minPrice)
}

func (day Day21) Part2(input []byte) string {
	boss := day.parse(input)
	maxPrice := 0

	checkItemsCombo := func(items ...Day21Item) {
		player, cost := day.calculateStats(items)
		player.hitPoints = 100
		if !day.simulate(player, boss) && cost > maxPrice {
			maxPrice = cost
		}
	}

	day.withAllItemsCombos(checkItemsCombo)

	return strconv.Itoa(maxPrice)
}

func (day Day21) withAllItemsCombos(f func(items ...Day21Item)) {
	for _, weapon := range day.weapons {
		for _, armor := range day.armors {
			for i, ring1 := range day.rings[:len(day.rings)-1] {
				for _, ring2 := range day.rings[i+1:] {
					f(weapon)
					f(weapon, armor)
					f(weapon, armor, ring1)
					f(weapon, ring1)
					f(weapon, armor, ring1, ring2)
					f(weapon, ring1, ring2)
				}
			}
		}
	}
}

func (day Day21) calculateStats(items []Day21Item) (stats Day21Stats, cost int) {
	for _, item := range items {
		stats.damage += item.damage
		stats.armor += item.armor
		cost += item.cost
	}
	return
}

func (day Day21) simulate(player, boss Day21Stats) bool {
	playerDamage := max(1, player.damage-boss.armor)
	bossDamage := max(1, boss.damage-player.armor)

	playerTurnsToKillBoss := boss.hitPoints / playerDamage
	if boss.hitPoints%playerDamage > 0 {
		playerTurnsToKillBoss++
	}

	bossTurnsToKillPlayer := player.hitPoints / bossDamage
	if player.hitPoints%bossDamage > 0 {
		bossTurnsToKillPlayer++
	}

	return playerTurnsToKillBoss <= bossTurnsToKillPlayer
}

func (day Day21) parse(input []byte) Day21Stats {
	spl := bytes.Split(input, []byte("\n"))

	return Day21Stats{
		hitPoints: optimistic.AtoiBFast(spl[0][12:]),
		damage:    optimistic.AtoiBFast(spl[1][8:]),
		armor:     optimistic.AtoiBFast(spl[2][7:]),
	}
}
