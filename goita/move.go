package goita

import (
	"strings"

	"github.com/Goita/go-goita/util"
)

// Move represents block and attack move, or pass move.
// pass move's block and attack are empty.
type Move struct {
	block    Koma
	attack   Koma
	faceDown bool
}

// MoveHashArray have hashes of move
type MoveHashArray []uint

var moveMap []*Move

// NewMatchMove creates match move
func NewMatchMove(block, attack Koma) *Move {
	if moveMap == nil {
		moveMap = getMoveMap()
	}
	return moveMap[moveHash(block, attack, false)]
	//return &Move{block, attack, false}
}

// NewFaceDownMove creates face-down move
func NewFaceDownMove(block, attack Koma) *Move {
	if moveMap == nil {
		moveMap = getMoveMap()
	}
	return moveMap[moveHash(block, attack, true)]
}

// NewPassMove creates pass move
func NewPassMove() *Move {
	if moveMap == nil {
		moveMap = getMoveMap()
	}
	return moveMap[moveHash(Empty, Empty, false)]
}

// ParseMove parse
func ParseMove(m string) (move *Move, ok bool) {
	elem := strings.Split(m, "")
	if elem[1] == "p" {
		return NewPassMove(), true
	}
	b := ParseKoma(elem[1])
	a := ParseKoma(elem[2])
	return NewMatchMove(b, a), true
}

// IsPass returns true if move is pass
func (m *Move) IsPass() bool {
	return m.block == Empty
}

// Hash returns identical hash key
func (m *Move) Hash() uint {
	return moveHash(m.block, m.attack, m.faceDown)
}

// MoveHash calculate identical hash key
func moveHash(block Koma, attack Koma, faceDown bool) uint {
	// block 0-8 : 4 bit
	// attack 0-8 : 4 bit
	// facedown flag 0-1 : 1 bit
	// total: 9bit
	if faceDown {
		return 1<<8 | uint(block)<<4 | uint(attack)
	}
	return uint(block)<<4 | uint(attack)
}

func getMoveMap() []*Move {
	mapLen := 1 << 9
	mmap := make([]*Move, mapLen, mapLen)

	for b := 0; b < 9; b++ {
		for a := 0; a < 9; a++ {
			key1 := moveHash(Koma(b), Koma(a), true)
			mmap[key1] = &Move{Koma(b), Koma(a), true}
			key2 := moveHash(Koma(b), Koma(a), false)
			mmap[key2] = &Move{Koma(b), Koma(a), false}
		}
	}
	return mmap
}

// StringHidden returns the string representation for block(may be hidden) and attack, or for pass
func (m *Move) StringHidden() string {
	if m.IsPass() {
		return "p"
	}
	if m.faceDown {
		return Hidden.String() + m.attack.String()
	}
	return m.block.String() + m.attack.String()
}

// String returns the string representation for opened block and attack, or for pass
func (m *Move) String() string {
	if m.IsPass() {
		return "p"
	}
	return m.block.String() + m.attack.String()
}

// History converts move hash array into history string
func (a MoveHashArray) History(startTurn int) string {
	buf := make([]byte, 0, len(a)*4)
	turn := startTurn
	for i, m := range a {
		if i > 0 {
			buf = append(buf, ',')
		}
		buf = append(buf, '1'+byte(turn))
		buf = append(buf, moveMap[m].String()...)
		turn = util.GetNextTurn(turn)
	}
	return string(buf)
}
