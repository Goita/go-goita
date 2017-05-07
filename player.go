package goita

// Player has 8 koma in total in the hand or/and on the field
type Player struct {
	hand         []Koma
	field        []Koma
	hiddenfield  []Koma
	handCounter  int
	fieldCounter int
}
