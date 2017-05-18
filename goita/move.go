package goita

import (
	"strings"
)

// Move represents block and attack move, or pass move.
// pass move's block and attack are empty.
type Move struct {
	block    Koma
	attack   Koma
	faceDown bool
}

// NewMatchMove creates match move
func NewMatchMove(block, attack Koma) *Move {
	return &Move{block, attack, false}
}

// NewFaceDownMove creates face-down move
func NewFaceDownMove(block, attack Koma) *Move {
	return &Move{block, attack, true}
}

// NewPassMove creates pass move
func NewPassMove() *Move {
	return &Move{Empty, Empty, false}
}

// ParseMove parse
func ParseMove(m string) (*Move, bool) {
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

// String returns the string representation for block(may be hidden) and attack, or for pass
func (m Move) String() string {
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
