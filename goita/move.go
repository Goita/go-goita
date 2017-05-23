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
type MoveHashArray []byte

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
func (m *Move) Hash() byte {
	return moveHash(m.block, m.attack, m.faceDown)
}

// MoveHash calculate identical hash key
func moveHash(block Koma, attack Koma, faceDown bool) byte {
	// block 0-8
	// attack 0-8
	// facedown flag 1 bit
	if faceDown {
		return 100 + byte(block)*10 + byte(attack)
	}
	return byte(block)*10 + byte(attack)
}

func getMoveMap() []*Move {
	mmap := make([]*Move, 256, 256)

	for b := 0; b < 10; b++ {
		for a := 0; a < 10; a++ {
			key1 := moveHash(Koma(b), Koma(a), true)
			mmap[key1] = &Move{Koma(b), Koma(a), true}
			key2 := moveHash(Koma(b), Koma(a), false)
			mmap[key2] = &Move{Koma(b), Koma(a), false}
		}
	}
	return mmap
}

// String returns the string representation for block(may be hidden) and attack, or for pass
func (m *Move) String() string {
	if m.IsPass() {
		return "p"
	}
	if m.faceDown {
		return Hidden.String() + m.attack.String()
	}
	return m.block.String() + m.attack.String()
}

// OpenString returns the string representation for opened block and attack, or for pass
func (m *Move) OpenString() string {
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
		buf = append(buf, moveMap[m].OpenString()...)
		turn = util.GetNextTurn(turn)
	}
	return string(buf)
}
