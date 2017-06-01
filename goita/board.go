/*
Package goita implements the components on the goita game.
such as places of koma on a game board, history of play.

*/
package goita

import (
	"strconv"
	"strings"

	"github.com/Goita/go-goita/util"
)

const moveHistoryCapacity = 49 // 3 moves * 4 player * 4 (max pass count) + 1 (last move)
const attackerLogCapacity = 13 // 3moves * 4 player + 1 (last move)

// Board represents the whole component of the goita game
type Board struct {
	Players          []*Player
	Turn             int
	MoveHistory      []*Move
	MoveHistoryLen   int
	MoveHistoryIndex int
	LastAttackMove   *Move
	AttackerLog      []int
	AttackMoveLog    []*Move
	KingUsed         int
	Finish           bool // true when finished
	InitialHands     []string
	Dealer           int
	InitialShiCounts []int // less than 5 will be 0
}

// NewBoard creates board instance with hands and Dealer
func NewBoard(dealer int, hands []KomaArray) *Board {
	b := &Board{}
	b.initWithInitialStateData(dealer, hands)
	return b
}

// NewBoardRandom creates board instance with random hands
func NewBoardRandom(dealer int) *Board {
	b := &Board{}
	ring := KomaRing()
	Shuffle(ring)
	hands := make([]KomaArray, 4)
	for i := 0; i < 4; i++ {
		hands[i] = ring[i*FieldLength : (i+1)*FieldLength]
		hands[i].Sort()
	}
	b.initWithInitialStateData(dealer, hands)
	return b
}

// ParseBoard creates board instance from HistoryString
func ParseBoard(historyString string) *Board {
	b := &Board{}
	b.initWithHistoryString(historyString)
	return b
}

// Copy duplicate board
func (b *Board) Copy() *Board {
	str := b.String()
	return ParseBoard(str)
}

func (b *Board) initBase() {
	b.Players = make([]*Player, 4)
	b.MoveHistory = make([]*Move, moveHistoryCapacity)
	b.MoveHistoryLen = 0
	b.MoveHistoryIndex = -1
	b.LastAttackMove = nil
	b.AttackerLog = make([]int, 0, attackerLogCapacity)
	b.AttackMoveLog = make([]*Move, 0, attackerLogCapacity)
	b.KingUsed = 0
	b.Finish = false
	b.InitialHands = make([]string, 4)
	b.InitialShiCounts = make([]int, 4)
}

func (b *Board) initWithInitialStateData(Dealer int, hands []KomaArray) {
	b.initBase()
	b.Dealer = Dealer
	for i, v := range hands {
		b.InitialHands[i] = v.String()
		p := NewPlayer(v)
		b.Players[i] = p
	}
	b.initYaku()
}

func (b *Board) initWithHistoryString(history string) {
	b.initBase()
	state := strings.Split(history, ",")
	if len(state) < 5 {
		panic("history must contain 4 initial hands and Dealer info")
	}

	for i := 0; i < 4; i++ {
		b.InitialHands[i] = state[i]
		k := ParseKomaArray(state[i])
		p := NewPlayer(k)
		b.Players[i] = p
	}
	Dealerinfo := strings.Split(state[4], "")
	noStr := Dealerinfo[1]
	no, err := strconv.ParseInt(noStr, 10, 4)
	if err != nil {
		panic(err)
	}
	b.Dealer = int(no) - 1
	b.Turn = b.Dealer

	for i := 5; i < len(state); i++ {
		m, _ := ParseMove(state[i])
		b.PlayMove(m)
	}

	b.initYaku()
}

func (b *Board) initYaku() {
	for i, p := range b.Players {
		shiCount := p.hand.Count(Shi)
		if shiCount >= 5 {
			b.InitialShiCounts[i] = shiCount
		}
	}
	goshi := -1
	for i, c := range b.InitialShiCounts {
		if c > 5 {
			b.Finish = true
		} else if c == 5 {
			if goshi < 0 {
				goshi = i
			} else {
				b.Finish = true
			}
		}
	}
}

// LastAttacker returns the number of last attacker
func (b *Board) LastAttacker() int {
	if len(b.AttackerLog) == 0 {
		return b.Dealer
	}
	return b.AttackerLog[len(b.AttackerLog)-1]
}

// PlayMove apply move to board
func (b *Board) PlayMove(move *Move) (ok bool) {
	// face-down check
	if b.LastAttacker() == b.Turn {
		move.FaceDown = true
	}

	b.MoveHistoryIndex++
	b.MoveHistory[b.MoveHistoryIndex] = move
	b.MoveHistoryLen = b.MoveHistoryIndex + 1

	finished := false
	if !move.IsPass() {
		p := b.Players[b.Turn]
		p.pushKoma(move.Block, move.FaceDown)
		p.pushKoma(move.Attack, false)

		if (move.Block.IsKing() && !move.FaceDown) || move.Attack.IsKing() {
			b.KingUsed++
		}
		b.LastAttackMove = move
		b.AttackerLog = append(b.AttackerLog, b.Turn)
		b.AttackMoveLog = append(b.AttackMoveLog, move)
		finished = p.fieldCounter == FieldLength
	}

	b.Turn = util.GetNextTurn(b.Turn)
	if finished {
		// TODO: calc score or something
		b.Finish = true
	}

	return true
}

// RedoMove turn back undo. it can redo to the latest move.
func (b *Board) RedoMove() (ok bool) {
	if b.MoveHistoryLen <= b.MoveHistoryIndex+1 {
		return false
	}
	move := b.MoveHistory[b.MoveHistoryIndex+1]
	b.PlayMove(move)
	return true
}

// UndoMove undo the last move. it can undo to the beginning of the deal.
func (b *Board) UndoMove() (ok bool) {
	if b.MoveHistoryIndex < 0 {
		return false
	}
	b.Turn = util.GetPreviousTurn(b.Turn)
	b.Finish = false

	move := b.MoveHistory[b.MoveHistoryIndex]
	b.MoveHistoryIndex--

	if !move.IsPass() {
		p := b.Players[b.Turn]
		p.popKoma()
		p.popKoma()

		if (move.Block.IsKing() && !move.FaceDown) || move.Attack.IsKing() {
			b.KingUsed--
		}

		b.AttackerLog = b.AttackerLog[:len(b.AttackerLog)-1]
		b.AttackMoveLog = b.AttackMoveLog[:len(b.AttackMoveLog)-1]
		if len(b.AttackMoveLog) == 0 {
			b.LastAttackMove = nil
		} else {
			b.LastAttackMove = b.AttackMoveLog[len(b.AttackMoveLog)-1]
		}
	}
	return true
}

// PossibleMoves returns a list of possible moves (no memory allocation)
func (b *Board) PossibleMoves(buf KomaArray, moves []*Move) []*Move {
	if b.Finish {
		return []*Move{}
	}
	hand := b.Players[b.Turn].hand
	uniqueCount := hand.Unique(buf)
	uniqueHand := buf[:uniqueCount]
	fieldCounter := b.Players[b.Turn].fieldCounter
	moves = moves[:0]
	if b.LastAttackMove == nil || b.Turn == b.AttackerLog[len(b.AttackerLog)-1] {
		// Face-Down move
		for _, faceDown := range uniqueHand {
			for _, attack := range uniqueHand {
				if faceDown == attack && hand.Count(faceDown) < 2 {
					continue
				}
				// Ou(王) as attack koma rule
				if fieldCounter < 6 && attack.IsKing() {
					if hand.Count(Ou) < 2 && b.KingUsed == 0 {
						continue
					}
				}
				moves = append(moves, NewFaceDownMove(faceDown, attack))
			}
		}
	} else {
		// Match move
		moves = append(moves, NewPassMove())
		block := b.LastAttackMove.Attack
		if hand.Contains(block) {
			for _, attack := range uniqueHand {
				if block == attack && hand.Count(block) < 2 {
					continue
				}
				// Ou(王) as attack koma rule
				if fieldCounter < 6 && attack.IsKing() {
					if hand.Count(Ou) < 2 && b.KingUsed == 0 {
						continue
					}
				}
				moves = append(moves, NewMatchMove(block, attack))
			}
		}
		if hand.Contains(Ou) && Ou.CanBlock(block) {
			for _, attack := range uniqueHand {
				if attack.IsKing() && hand.Count(attack) < 2 {
					continue
				}
				moves = append(moves, NewMatchMove(Ou, attack))
			}
		}
	}

	return moves
}

// GetPossibleMoves returns a list of possible moves
func (b *Board) GetPossibleMoves() []*Move {
	buf := make(KomaArray, FieldLength)
	moves := make([]*Move, 0, 64)
	return b.PossibleMoves(buf, moves)
}

// IsEnd returns true if the deal is finished (this is slow)
func (b *Board) IsEnd() bool {
	if b.Finish {
		return true
	}
	for _, p := range b.Players {
		if p.fieldCounter == FieldLength {
			return true
		}
	}
	return false
}

// IsGoshi returns true if only one player has goshi
func (b *Board) IsGoshi() bool {
	goshi := -1
	for i, v := range b.InitialShiCounts {
		if v == 5 {
			if goshi >= 0 {
				return false
			}
			goshi = i
		}
	}
	return false
}

// HasYaku returns true if there is a yaku
func (b *Board) HasYaku() bool {
	goshi := -1
	for i, v := range b.InitialShiCounts {
		if v >= 6 {
			return true
		} else if v >= 5 {
			if goshi >= 0 {
				if util.IsSameTeam(i, goshi) {
					return true
				}
			} else {
				goshi = i
			}
		}
	}
	return false
}

// WonPlayerNo returns the won player number, or -1 if deal is not end.
func (b *Board) WonPlayerNo() int {
	if b.Finish {
		for i, p := range b.Players {
			if p.fieldCounter == FieldLength {
				return i
			}
		}
	}
	return -1
}

// Score returns the score value of finish state
func (b *Board) Score() int {
	if !b.Finish {
		return 0
	}

	// yaku ? or finish move
	// b.HasYaku() //TODO: yaku score
	s := b.LastAttackMove.Attack.GetScore()
	if b.LastAttackMove.FaceDown && b.LastAttackMove.Block == b.LastAttackMove.Attack {
		s = s * 2
	}
	return s
}

// FinishMoveScore returns the score value of finish attack. it may cause panic when the board is not finished.
func (b *Board) FinishMoveScore() int {
	if b.LastAttackMove.FaceDown && b.LastAttackMove.Block == b.LastAttackMove.Attack {
		return b.LastAttackMove.Attack.GetScore() * 2
	}
	return b.LastAttackMove.Attack.GetScore()
}

func (b *Board) String() string {
	// 00000000,00000000,00000000,00000000,s1(38chars) + ,100,2p,3p,4p (13char) * 49 = 675byte
	bufLen := 38 + 13*len(b.AttackMoveLog)
	buf := make([]byte, 0, bufLen)
	for _, v := range b.InitialHands {
		buf = append(buf, v...)
		buf = append(buf, ',')
	}
	buf = append(buf, 's')
	buf = append(buf, '1'+byte(b.Dealer))
	turn := b.Dealer
	for i := 0; i < b.MoveHistoryLen; i++ {
		m := b.MoveHistory[i]
		buf = append(buf, ',')
		buf = append(buf, '1'+byte(turn))
		buf = append(buf, m.String()...)
		turn = util.GetNextTurn(turn)
	}
	return string(buf)
}

// SubHistory returns a part of history
func (b *Board) SubHistory(start int, end int) MoveHashArray {
	moves := make(MoveHashArray, 0, end-start+1)
	for i := start; i < end; i++ {
		m := b.MoveHistory[i]
		moves = append(moves, m.Hash())
	}
	return moves
}
