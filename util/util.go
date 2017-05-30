package util

import (
	"math/rand"
	"time"
)

var rngSource = rand.NewSource(time.Now().UnixNano())
var rng = rand.New(rngSource)

// GetRNG returns a Random Number Generator
func GetRNG() *rand.Rand {
	return rng
}

// ShiftTurn shifts the turn by the offset number
func ShiftTurn(turn int, offset int) int {
	if offset < 0 {
		return (turn + (4 + offset%4)) % 4
	}
	return (turn + (offset % 4)) % 4
}

// GetNextTurn returns the next turn number
func GetNextTurn(turn int) int {
	return ShiftTurn(turn, 1)
}

// GetPreviousTurn returns the previous turn number
func GetPreviousTurn(turn int) int {
	return ShiftTurn(turn, -1)
}

// IsSameTeam checks no1 and no2 player number belongs to the same team
func IsSameTeam(no1 int, no2 int) bool {
	return (no1+no2)%2 == 0
}
