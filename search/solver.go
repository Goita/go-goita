package search

import (
	"fmt"

	"github.com/Goita/go-goita/goita"
	"github.com/Goita/go-goita/util"
)

type evalFunc func(*goita.Board) int

// EvaluatedMove is Move data and score
type EvaluatedMove struct {
	Move    *goita.Move
	Score   int
	History goita.MoveHashArray
}

// LinkedMove is considerable as a solution for reducing memory allocation of EvaluatedMove
type LinkedMove struct {
	Move *goita.Move
	next *LinkedMove
}

type searchMemory struct {
	movesBuf    []*goita.Move
	mapBuf      []goita.Koma
	buf         goita.KomaArray
	moveHashBuf goita.MoveHashArray
}

func (em EvaluatedMove) String() string {
	return fmt.Sprintf("[%v:%v]", em.Move.OpenString(), em.Score)
}

// StringHistory make move+score+history string
func (em EvaluatedMove) StringHistory(startTurn int) string {
	return fmt.Sprintf("[%v:%v -> %v]", em.Move.OpenString(), em.Score, em.History.History(startTurn))
}

// Solve search the deal perfect
func Solve(board *goita.Board) []EvaluatedMove {
	moves := board.GetPossibleMoves()
	evaledMoves := make([]EvaluatedMove, 0, len(moves))
	ch := make(chan *EvaluatedMove, len(moves))
	for _, move := range moves {
		go StartNegamax(board, move, func(b *goita.Board) int {
			return b.Score()
		}, ch)
	}
	for i := 0; i < len(moves); i++ {
		result := <-ch
		// fmt.Println("search done!")
		evaledMoves = append(evaledMoves, *result)
	}

	return evaledMoves
}

var positiveInf *EvaluatedMove
var negativeInf *EvaluatedMove

// StartNegamax run negamax search
func StartNegamax(board *goita.Board, move *goita.Move, eval evalFunc, ch chan *EvaluatedMove) {
	copyBoard := board.Copy()
	negativeInf = &EvaluatedMove{Score: -999}
	shared := make([]searchMemory, 0, 50)
	for i := 0; i < 50; i++ {
		mem := searchMemory{
			buf:         make(goita.KomaArray, 0, goita.FieldLength),
			mapBuf:      make([]goita.Koma, 10, 10),
			movesBuf:    make([]*goita.Move, 0, 64),
			moveHashBuf: make(goita.MoveHashArray, 0, 50)}
		shared = append(shared, mem)
	}
	evaledMove := negamax(copyBoard, board.Turn, eval, move, -999, 999, 0, &shared)
	evaledMove.Score = -evaledMove.Score
	evaledMove.Move = move
	ch <- evaledMove
}

func negamax(board *goita.Board, playerNo int, eval evalFunc, move *goita.Move, alpha int, beta int, depth int, shared *[]searchMemory) *EvaluatedMove {
	board.PlayMove(move)

	// fmt.Print(move.OpenString())
	if board.Finish {
		score := -eval(board)
		history := board.SubHistory(board.MoveHistoryIndex-depth, board.MoveHistoryIndex+1)
		board.UndoMove()
		return &EvaluatedMove{Score: score, History: history}
	}

	moves := board.PossibleMoves((*shared)[depth].mapBuf, (*shared)[depth].buf, (*shared)[depth].movesBuf)
	var best *EvaluatedMove
	best = negativeInf
	for _, move := range moves {
		v := negamax(board, playerNo, eval, move, -beta, -alpha, depth+1, shared)
		v.Score = -v.Score
		if v.Score > best.Score {
			best = v
		}
		if v.Score > alpha {
			alpha = v.Score
		}
		if alpha >= beta {
			break // beta cut-off
		}
	}

	board.UndoMove()
	return best
}

// StartAlphaBeta run alpha-beta search
func StartAlphaBeta(board *goita.Board, move *goita.Move, eval evalFunc, ch chan *EvaluatedMove) {
	copyBoard := board.Copy()
	min := negativeInf
	max := positiveInf
	shared := make([]searchMemory, 0, 50)
	for i := 0; i < 50; i++ {
		mem := searchMemory{
			buf:         make(goita.KomaArray, 0, goita.FieldLength),
			mapBuf:      make([]goita.Koma, 10, 10),
			movesBuf:    make([]*goita.Move, 0, 64),
			moveHashBuf: make(goita.MoveHashArray, 0, 50)}
		shared = append(shared, mem)
	}
	evaledMove := alphaBeta(copyBoard, board.Turn, eval, move, min, max, 0, &shared)
	evaledMove.Move = move
	ch <- evaledMove
}

func alphaBeta(board *goita.Board, playerNo int, eval evalFunc, move *goita.Move, min *EvaluatedMove, max *EvaluatedMove, depth int, shared *[]searchMemory) *EvaluatedMove {
	board.PlayMove(move)

	// fmt.Print(move.OpenString())
	if board.Finish {
		score := eval(board)
		if !util.IsSameTeam(playerNo, board.LastAttacker()) {
			score *= -1
		}
		history := board.SubHistory(board.MoveHistoryIndex-depth, board.MoveHistoryIndex+1)
		board.UndoMove()
		return &EvaluatedMove{Score: score, History: history}
	}

	moves := board.PossibleMoves((*shared)[depth].mapBuf, (*shared)[depth].buf, (*shared)[depth].movesBuf)
	var v *EvaluatedMove
	if util.IsSameTeam(playerNo, board.Turn) {
		v = min
		for _, move := range moves {
			t := alphaBeta(board, playerNo, eval, move, v, max, depth+1, shared)
			if t.Score > v.Score {
				v = t
			}
			if v.Score > max.Score {
				// fmt.Println("->cut (max)")
				board.UndoMove()
				return max
			}
		}
	} else {
		v = max
		for _, move := range moves {
			t := alphaBeta(board, playerNo, eval, move, min, v, depth+1, shared)
			if t.Score < v.Score {
				v = t
			}
			if v.Score < min.Score {
				// fmt.Println("->cut (min)")
				board.UndoMove()
				return min
			}
		}
	}

	board.UndoMove()
	return v
}
