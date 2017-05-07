package goita

import (
	"fmt"
	"math"
)

// Koma represents koma of goita
type Koma byte

// GetValue returns the raw byte value
func (koma Koma) GetValue() byte {
	return byte(koma)
}

// Empty (0x0) represents no koma is here
const Empty Koma = 0x0

// Hidden (0xf) is face-down koma
const Hidden Koma = 0xf

// Shi (0x1)
const Shi Koma = 0x1

// Gon (0x2)
const Gon Koma = 0x2

// Bakko (0x3)
const Bakko Koma = 0x3

// Gin (0x4)
const Gin Koma = 0x4

// Kin (0x5)
const Kin Koma = 0x5

// Kaku (0x6)
const Kaku Koma = 0x6

// Hisha (0x7)
const Hisha Koma = 0x7

// Ou (0x8)
const Ou Koma = 0x8

// NewKomaFromStr converts string to Koma byte value
func NewKomaFromStr(str string) Koma {
	switch str {
	case "0":
		return Empty
	case "1":
		return Shi
	case "2":
		return Gon
	case "3":
		return Bakko
	case "4":
		return Gin
	case "5":
		return Kin
	case "6":
		return Kaku
	case "7":
		return Hisha
	case "8":
		return Ou
	case "x":
		return Hidden
	default:
		panic(fmt.Sprintf("Invalid koma string value %v was given", str))
	}
}

// GetScore returns the koma finnish score
func (koma Koma) GetScore() int {
	if koma == Hidden {
		panic("cannot get the score of Hidden")
	}
	if koma == Empty {
		panic("cannot get the score of Empty")
	}
	return int(math.Floor(float64(koma/2.0))*10 + 10)
}

// IsKing returns true if the koma is Ou or Gyoku
func (koma Koma) IsKing() bool {
	return koma == Ou
}

// IsShi returns true if the koma is Shi
func (koma Koma) IsShi() bool {
	return koma == Shi
}

// IsEmpty returns true if the koma is Empty
func (koma Koma) IsEmpty() bool {
	return koma == Empty
}

// IsHidden returns true if the koma is Hidden
func (koma Koma) IsHidden() bool {
	return koma == Hidden
}

// CanBlock returns true if the koma can block the target koma
func (koma Koma) CanBlock(target Koma) bool {
	if koma.IsKing() {
		if target.IsShi() || target == Gon {
			return false
		}
		return true
	}
	return koma == target
}

func (koma Koma) String() string {
	if koma.IsHidden() {
		return "x"
	}
	return fmt.Sprintf("%x", byte(koma))
}
