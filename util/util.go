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

// Shuffle randomize the provided array order with Fisher-Yates algorithm
func Shuffle(array []byte) {
	for i := len(array) - 1; i > 0; i = i - 1 {
		j := rng.Intn(i)
		array[i], array[j] = array[j], array[i]
	}
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
