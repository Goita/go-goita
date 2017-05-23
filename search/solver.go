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

type searchMemory struct {
	movesBuf    []*goita.Move
	mapBuf      []goita.Koma
	buf         goita.KomaArray
	moveHashBuf goita.MoveHashArray
}

func (em EvaluatedMove) String() string {
	return fmt.Sprintf("[%v:%v -> %v]", em.Move.OpenString(), em.Score, em.History)
}

// Solve search the deal perfect
func Solve(board *goita.Board) []EvaluatedMove {
	moves := board.GetPossibleMoves()
	evaledMoves := make([]EvaluatedMove, 0, len(moves))
	ch := make(chan *EvaluatedMove, len(moves))
	for _, move := range moves {
		go StartAlphaBetaSearch(board, move, func(b *goita.Board) int {
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

// StartAlphaBetaSearch run alpha-beta search
func StartAlphaBetaSearch(board *goita.Board, move *goita.Move, eval evalFunc, ch chan *EvaluatedMove) {
	copyBoard := board.Copy()
	min := &EvaluatedMove{Score: -999}
	max := &EvaluatedMove{Score: 999}
	shared := make([]searchMemory, 0, 50)
	for i := 0; i < 50; i++ {
		mem := searchMemory{
			buf:         make(goita.KomaArray, 0, goita.FieldLength),
			mapBuf:      make([]goita.Koma, 10, 10),
			movesBuf:    make([]*goita.Move, 0, 64),
			moveHashBuf: make(goita.MoveHashArray, 0, 50)}
		shared = append(shared, mem)
	}
	evaledMove := alphaBetaSearch(copyBoard, board.Turn, eval, move, min, max, 0, &shared)
	evaledMove.Move = move
	ch <- evaledMove
}

func alphaBetaSearch(board *goita.Board, playerNo int, eval evalFunc, move *goita.Move, min *EvaluatedMove, max *EvaluatedMove, depth int, shared *[]searchMemory) *EvaluatedMove {
	board.PlayMove(move)

	// fmt.Print(move.OpenString())
	if board.Finish {
		score := eval(board)
		if !util.IsSameTeam(playerNo, board.LastAttacker()) {
			score *= -1
		}
		history := board.SubHistory(board.MoveHistoryIndex-depth, board.MoveHistoryIndex+1, (*shared)[depth].moveHashBuf)
		board.UndoMove()
		return &EvaluatedMove{Score: score, History: history}
	}

	moves := board.PossibleMoves((*shared)[depth].mapBuf, (*shared)[depth].buf, (*shared)[depth].movesBuf)
	var v *EvaluatedMove
	if util.IsSameTeam(playerNo, board.Turn) {
		v = min
		for _, move := range moves {
			t := alphaBetaSearch(board, playerNo, eval, move, v, max, depth+1, shared)
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
			t := alphaBetaSearch(board, playerNo, eval, move, min, v, depth+1, shared)
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
