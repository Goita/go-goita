package search

import "github.com/Goita/go-goita/goita"

// OrderSimple calculate score with simple conditions
func OrderSimple(move *goita.Move, handCounter int, partnerAttack bool) (score int) {
	score = 0

	if partnerAttack {
		return -30
	}

	if handCounter < 6 {
		if move.Attack == move.Block {
			return -40
		}
	}

	if handCounter == 0 {
		if move.IsPass() {
			score -= 5
		} else if move.FaceDown {
			if move.Block.IsKing() {
				score -= 50
			} else if move.Block == goita.Kaku || move.Block == goita.Hisha {
				score -= 30
			} else if move.Block == goita.Gon {
				score -= 20
			} else if move.Block == goita.Gin || move.Block == goita.Kin {
				score -= 10
			}
		} else {
			if move.Attack == goita.Shi {
				score -= 10
			} else if move.Attack == goita.Kaku || move.Attack == goita.Hisha {
				score += 20
			} else if move.Attack.IsKing() {
				score -= 50
			}
		}
	} else if handCounter == 2 {
		if move.IsPass() {
			score -= 15
		} else if move.FaceDown {
			if move.Block.IsKing() {
				score -= 50
			} else if move.Block == goita.Kaku || move.Block == goita.Hisha {
				score -= 30
			} else if move.Block == goita.Gon {
				score -= 20
			} else if move.Block == goita.Gin || move.Block == goita.Kin {
				score -= 10
			}

		}
		if move.Attack.IsKing() {
			score += 40
		} else if move.Attack == goita.Kaku || move.Attack == goita.Hisha {
			score += 30
		}
	} else if handCounter == 4 {
		if move.IsPass() {
			score -= 15
		} else if move.FaceDown {
			if move.Block.IsKing() {
				score -= 50
			} else if move.Block == goita.Kaku || move.Block == goita.Hisha {
				score -= 30
			}
		}
		if move.Attack.IsKing() {
			score += 40
		} else if move.Attack == goita.Kaku || move.Attack == goita.Hisha {
			score += 30
		}
	} else if handCounter == 6 {
		if move.IsPass() {
			score = -25
		} else if move.FaceDown && move.Block == move.Attack {
			score = move.Attack.GetScore() * 2
		} else {
			score = move.Attack.GetScore()
		}
	}

	return score
}
