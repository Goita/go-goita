package search

import (
	"fmt"

	"github.com/Goita/go-goita/goita"
	"github.com/Goita/go-goita/util"
)

// EvaluatedMove is Move data and score
type EvaluatedMove struct {
	Move    goita.Move
	Score   int
	History string
}

func (em EvaluatedMove) String() string {
	return fmt.Sprintf("[%v:%v]", em.Move.OpenString(), em.Score)
}

type evalFunc func(*goita.Board) int

// Solve search the deal perfect
func Solve(board *goita.Board) []EvaluatedMove {
	evaledMoves := make([]EvaluatedMove, 0)
	moves := board.GetPossibleMoves()
	// ch := make(chan int)
	var score int
	for _, move := range moves {
		score = StartAlphaBetaSearch(board, move, func(b *goita.Board) int {
			return b.Score()
		})
		evaledMoves = append(evaledMoves, EvaluatedMove{Score: score})
	}
	for i := 0; i < len(moves); i++ {
		//score := <-ch

	}

	return evaledMoves
}

// StartAlphaBetaSearch run alpha-beta search
func StartAlphaBetaSearch(board *goita.Board, move *goita.Move, eval evalFunc) int {
	copyBoard := board.Copy()
	// min := &EvaluatedMove{Score: -999}
	// max := &EvaluatedMove{Score: 999}
	score := alphaBetaSearch(copyBoard, board.Turn, func(b *goita.Board) int {
		return b.Score()
	}, move, -999, 999)
	return score
}

func alphaBetaSearch(board *goita.Board, playerNo int, eval evalFunc, move *goita.Move, min int, max int) int {
	board.PlayMove(move)
	defer board.UndoMove()

	fmt.Print(move.OpenString())
	if board.Finish {
		score := eval(board)
		if !util.IsSameTeam(playerNo, board.LastAttacker()) {
			score *= -1
		}
		// history := board.String()
		return score // &EvaluatedMove{Move: *move, Score: score, History: history}
	}

	moves := board.GetPossibleMoves()
	var v int
	if util.IsSameTeam(playerNo, board.Turn) {
		v = min
		for _, move := range moves {
			t := alphaBetaSearch(board, playerNo, eval, move, v, max)
			if t > v {
				v = t
			}
			if v > max {
				fmt.Println("->cut (max)")
				return max
			}
		}
	} else {
		v = max
		for _, move := range moves {
			t := alphaBetaSearch(board, playerNo, eval, move, min, v)
			if t < v {
				v = t
			}
			if v < min {
				fmt.Println("->cut (min)")
				return min
			}
		}
	}

	fmt.Println("->undo")
	return v
}
