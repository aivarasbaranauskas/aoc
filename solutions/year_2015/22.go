package year_2015

import (
	"bytes"
	"strconv"

	"github.com/aivarasbaranauskas/aoc/internal/optimistic"
)

func init() {
	Solutions[22] = Day22{
		magicMissileCost: 53,
		drainCost:        73,
		shieldCost:       113,
		poisonCost:       173,
		rechargeCost:     229,
	}
}

type Day22 struct {
	magicMissileCost, drainCost, shieldCost, poisonCost, rechargeCost int
}

func (day Day22) Part1(input []byte) string {
	return day.solve(input, false)
}

func (day Day22) Part2(input []byte) string {
	return day.solve(input, true)
}

func (day Day22) solve(input []byte, hardMode bool) string {
	bossHitPoints, bossDamage := day.parse(input)

	type State struct {
		hitPoints, bossHitPoints,
		mana, manaSpent,
		shieldLeft, poisonLeft, rechargeLeft int
	}

	var minManaSpent *int
	var bruteForce func(
		isPlayerTurn bool,
		hitPoints,
		bossHitPoints,
		mana,
		manaSpent,
		shieldLeft,
		poisonLeft,
		rechargeLeft int,
	)

	bruteForce = func(
		isPlayerTurn bool,
		hitPoints,
		bossHitPoints,
		mana,
		manaSpent,
		shieldLeft,
		poisonLeft,
		rechargeLeft int,
	) {
		if minManaSpent != nil && *minManaSpent < manaSpent {
			// no need to continue
			return
		}

		if hardMode && isPlayerTurn {
			hitPoints--
			if hitPoints <= 0 {
				// dead
				return
			}
		}

		// apply effects
		if shieldLeft > 0 {
			shieldLeft--
		}
		if poisonLeft > 0 {
			bossHitPoints -= 3
			poisonLeft--
		}
		if rechargeLeft > 0 {
			mana += 101
			rechargeLeft--
		}

		// check if boss not dead after effects
		if bossHitPoints <= 0 {
			if minManaSpent == nil || *minManaSpent > manaSpent {
				minManaSpent = &manaSpent
			}
			return
		}

		if !isPlayerTurn {
			// boss move
			activeBossDamage := bossDamage
			if shieldLeft > 0 {
				activeBossDamage -= 7
				if activeBossDamage <= 0 {
					activeBossDamage = 1
				}
			}
			hitPoints -= activeBossDamage
			// check if not dead
			if hitPoints <= 0 {
				// dead
				return
			}
			// call player turn
			bruteForce(true, hitPoints, bossHitPoints, mana, manaSpent, shieldLeft, poisonLeft, rechargeLeft)
			return
		}

		// try to do all options

		// magic missile
		if mana > day.magicMissileCost {
			newBossHitPoints := bossHitPoints - 4
			newManaSpent := manaSpent + day.magicMissileCost
			if newBossHitPoints <= 0 {
				if minManaSpent == nil || *minManaSpent > newManaSpent {
					minManaSpent = &newManaSpent
				}
				return
			}
			newMana := mana - day.magicMissileCost
			bruteForce(false, hitPoints, newBossHitPoints, newMana, newManaSpent, shieldLeft, poisonLeft, rechargeLeft)
		}

		// drain
		if mana > day.drainCost {
			newBossHitPoints := bossHitPoints - 2
			newManaSpent := manaSpent + day.drainCost
			if newBossHitPoints <= 0 {
				if minManaSpent == nil || *minManaSpent > newManaSpent {
					minManaSpent = &newManaSpent
				}
				return
			}
			newMana := mana - day.drainCost
			newHitPoints := hitPoints + 2
			bruteForce(false, newHitPoints, newBossHitPoints, newMana, newManaSpent, shieldLeft, poisonLeft, rechargeLeft)
		}

		// shield
		if mana > day.shieldCost {
			newManaSpent := manaSpent + day.shieldCost
			newMana := mana - day.shieldCost
			bruteForce(false, hitPoints, bossHitPoints, newMana, newManaSpent, 6, poisonLeft, rechargeLeft)
		}

		// poison
		if mana > day.poisonCost {
			newManaSpent := manaSpent + day.poisonCost
			newMana := mana - day.poisonCost
			bruteForce(false, hitPoints, bossHitPoints, newMana, newManaSpent, shieldLeft, 6, rechargeLeft)
		}

		// recharge
		if mana > day.rechargeCost {
			newManaSpent := manaSpent + day.rechargeCost
			newMana := mana - day.rechargeCost
			bruteForce(false, hitPoints, bossHitPoints, newMana, newManaSpent, shieldLeft, poisonLeft, 5)
		}
	}

	bruteForce(true, 50, bossHitPoints, 500, 0, 0, 0, 0)

	if minManaSpent == nil {
		return "no min found"
	}

	return strconv.Itoa(*minManaSpent)
}

func (Day22) parse(input []byte) (bossHitPoints, bossDamage int) {
	spl := bytes.Split(input, []byte("\n"))
	bossHitPoints = optimistic.AtoiBFast(spl[0][12:])
	bossDamage = optimistic.AtoiBFast(spl[1][8:])
	return
}
