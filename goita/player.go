package goita

import (
	"bytes"
)

// FieldLength is field and hand length. the value is 8
const FieldLength int = 8

// Player has 8 koma in total in the hand or/and on the field
type Player struct {
	hand         KomaArray
	field        KomaArray
	hiddenfield  KomaArray
	handCounter  int
	fieldCounter int
}

// NewPlayer create a new player data
func NewPlayer(hand *KomaArray) *Player {
	p := new(Player)
	p.init(hand)
	return p
}

func (p *Player) init(hand *KomaArray) {
	p.hand = *hand
	p.handCounter = len(p.hand)
	p.field = make(KomaArray, FieldLength)
	p.hiddenfield = make(KomaArray, FieldLength)
	p.fieldCounter = 0
}

func (p *Player) pushKoma(koma Koma, faceDown bool) {
	if koma == Empty || koma == Hidden {
		panic("cannot put Empty neither Hidden")
	}
	i := bytes.IndexByte(p.hand.GetBytes(), koma.GetByte())
	p.hand[i] = Empty
	p.handCounter--
	if faceDown {
		p.field[p.fieldCounter] = Hidden
	} else {
		p.field[p.fieldCounter] = koma
	}

	p.hiddenfield[p.fieldCounter] = koma
	p.fieldCounter++
}

func (p *Player) popKoma() {
	if p.fieldCounter == 0 {
		return
	}
	removingIndex := p.fieldCounter - 1
	koma := p.hiddenfield[removingIndex]
	p.field[removingIndex] = Empty
	p.hiddenfield[removingIndex] = Empty
	p.fieldCounter--
	i := p.hand.Index(Empty)
	p.hand[i] = koma
	p.handCounter++
}

// GetHidden creates a field list. the list represents only hidden(face-down) koma and hide the others.
func (p *Player) GetHidden() KomaArray {
	diff := make(KomaArray, FieldLength)
	for i := 0; i < FieldLength; i++ {
		if p.field[i] == p.hiddenfield[i] {
			diff[i] = p.hiddenfield[i]
		}
	}
	return diff
}
