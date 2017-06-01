package search

import (
	"fmt"
	"runtime"
	"sync/atomic"

	"github.com/Goita/go-goita/goita"
)

const (
	// MaxDepth is a limit of depth of goita
	MaxDepth = 50
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
	movesBuf []*goita.Move
	buf      goita.KomaArray
	//moveHashBuf goita.MoveHashArray
}

// Status holds search information
type Status struct {
	VisitedNode  uint64
	VisitedLeaf  uint64
	CutOffedNode uint64
	Routines     int32
	MaxRoutines  int32
}

func (em EvaluatedMove) String() string {
	return fmt.Sprintf("[%v:%v]", em.Move, em.Score)
}

// StringHistory make move+score+history string
func (em EvaluatedMove) StringHistory(startTurn int) string {
	return fmt.Sprintf("[%v:%v -> %v]", em.Move, em.Score, em.History.History(startTurn))
}

// evalScore returns the finish score
func evalScore(b *goita.Board) int {
	// must be finished by attack move
	return b.FinishMoveScore()
}

var negativeInf *EvaluatedMove

// Solve search the deal perfect
func Solve(board *goita.Board) []EvaluatedMove {
	return StartNegamax(board, MaxDepth, evalScore)
}

// StartNegamax manages negamax search goroutines
func StartNegamax(board *goita.Board, searchDepth int, eval evalFunc) []EvaluatedMove {
	negativeInf = &EvaluatedMove{Score: -999}

	moves := board.GetPossibleMoves()
	fmt.Printf("search begin on %v moves ", len(moves))
	fmt.Println(moves)
	evaledMoves := make([]EvaluatedMove, 0, len(moves))
	status := &Status{MaxRoutines: int32(runtime.NumCPU())}
	ch := make(chan *EvaluatedMove, len(moves))
	for _, move := range moves {
		go negamaxAsync(board.Copy(), board.Turn, move, 0, searchDepth, -999, 999, eval, status, ch)
		atomic.AddInt32(&status.Routines, 1)
	}
	for i := 0; i < len(moves); i++ {
		result := <-ch
		atomic.AddInt32(&status.Routines, -1)
		fmt.Println("search done!", result)
		evaledMoves = append(evaledMoves, *result)
	}

	fmt.Printf("%+v\n", status)

	return evaledMoves
}

// negamaxAsync run negamax search async and returns result to channel
func negamaxAsync(copyBoard *goita.Board, playerNo int, move *goita.Move, depth int, searchDepth int, alpha int, beta int, eval evalFunc, status *Status, ch chan *EvaluatedMove) {
	// just copy and allocate separeted memory to concurent task

	shared := make([]searchMemory, 0, searchDepth)
	bufBulk := make(goita.KomaArray, goita.FieldLength*searchDepth)
	movesBufBulk := make([]*goita.Move, 64*searchDepth)
	for i := 0; i < searchDepth; i++ {
		mem := searchMemory{
			buf:      bufBulk[i*goita.FieldLength : (i+1)*goita.FieldLength],
			movesBuf: movesBufBulk[i*64 : i*64],
		}
		shared = append(shared, mem)
	}
	evaledMove := negamax(copyBoard, playerNo, move, depth, searchDepth, alpha, beta, eval, status, &shared)
	evaledMove.Score = -evaledMove.Score
	evaledMove.Move = move
	ch <- evaledMove
}

func negamax(board *goita.Board, playerNo int, move *goita.Move, depth int, searchDepth int, alpha int, beta int, eval evalFunc, status *Status, shared *[]searchMemory) *EvaluatedMove {
	// normal negamax search
	board.PlayMove(move)

	// fmt.Print(move.OpenString())
	if board.Finish {
		atomic.AddUint64(&status.VisitedLeaf, 1)
		score := -eval(board) // it's the opponent turn. score will be negative for lost player.
		history := board.SubHistory(board.MoveHistoryIndex-depth, board.MoveHistoryIndex+1)
		board.UndoMove()
		return &EvaluatedMove{Score: score, History: history}
	}

	atomic.AddUint64(&status.VisitedNode, 1)

	moves := board.PossibleMoves((*shared)[depth].buf, (*shared)[depth].movesBuf)

	// order moves
	// p := board.Players[board.Turn]
	// sameteam := util.IsSameTeam(board.LastAttacker(), board.Turn)
	// handcount := p.GetHandCount()
	// sort.Slice(moves, func(i, j int) bool {
	// 	return OrderSimple(moves[i], handcount, sameteam) < OrderSimple(moves[j], handcount, sameteam)
	// })

	var best *EvaluatedMove
	best = negativeInf

	// start new search routine if total running routine is less than cpu count
	// it's effective when more than 2 possible moves
	if len(moves) > 2 && atomic.LoadInt32(&status.Routines) < status.MaxRoutines {
		atomic.AddInt32(&status.Routines, int32(len(moves)))
		ch := make(chan *EvaluatedMove, len(moves))
		for _, nextMove := range moves {
			go negamaxAsync(board.Copy(), playerNo, nextMove, depth+1, searchDepth, -beta, -alpha, eval, status, ch)
			// fmt.Println("start new search!")
		}
		for j := 0; j < len(moves); j++ {
			//var v *EvaluatedMove
			v := <-ch
			atomic.AddInt32(&status.Routines, -1)
			v.Score = -v.Score
			if v.Score > best.Score {
				best = v
			}
			if v.Score > alpha {
				alpha = v.Score
			}
			if alpha >= beta {
				atomic.AddUint64(&status.CutOffedNode, 1)
				//TODO : cancel all goroutine started from this node
				break // beta cut-off
			}
		}
		//atomic.AddInt32(&status.Routines, int32(-len(moves)))
	} else {
		for _, next := range moves {
			v := negamax(board, playerNo, next, depth+1, searchDepth, -beta, -alpha, eval, status, shared)
			v.Score = -v.Score
			if v.Score > best.Score {
				best = v
			}
			if v.Score > alpha {
				alpha = v.Score
			}
			if alpha >= beta {
				atomic.AddUint64(&status.CutOffedNode, 1)
				break // beta cut-off
			}
		}
	}

	board.UndoMove()
	return best
}

// SolveSimple search the deal perfect
func SolveSimple(board *goita.Board) []EvaluatedMove {
	moves := board.GetPossibleMoves()
	evaledMoves := make([]EvaluatedMove, 0, len(moves))
	ch := make(chan *EvaluatedMove, len(moves))
	for _, move := range moves {
		go StartNegamaxSimple(board, move, evalScore, MaxDepth, ch)
	}
	for i := 0; i < len(moves); i++ {
		result := <-ch
		// fmt.Println("search done!")
		evaledMoves = append(evaledMoves, *result)
	}

	return evaledMoves
}

// StartNegamaxSimple run negamax search
func StartNegamaxSimple(board *goita.Board, move *goita.Move, eval evalFunc, searchDepth int, ch chan *EvaluatedMove) {
	copyBoard := board.Copy()
	negativeInf = &EvaluatedMove{Score: -999}
	shared := make([]searchMemory, 0, searchDepth)
	bufBulk := make(goita.KomaArray, goita.FieldLength*searchDepth)
	movesBufBulk := make([]*goita.Move, 64*searchDepth)
	for i := 0; i < searchDepth; i++ {
		mem := searchMemory{
			buf:      bufBulk[i*goita.FieldLength : (i+1)*goita.FieldLength],
			movesBuf: movesBufBulk[i*64 : i*64],
		}
		shared = append(shared, mem)
	}
	evaledMove := negamaxSimple(copyBoard, board.Turn, eval, move, -999, 999, 0, &shared)
	evaledMove.Score = -evaledMove.Score
	evaledMove.Move = move
	ch <- evaledMove
}

func negamaxSimple(board *goita.Board, playerNo int, eval evalFunc, move *goita.Move, alpha int, beta int, depth int, shared *[]searchMemory) *EvaluatedMove {
	board.PlayMove(move)

	// fmt.Print(move.OpenString())
	if board.Finish {
		score := -eval(board)
		history := board.SubHistory(board.MoveHistoryIndex-depth, board.MoveHistoryIndex+1)
		board.UndoMove()
		return &EvaluatedMove{Score: score, History: history}
	}

	moves := board.PossibleMoves((*shared)[depth].buf, (*shared)[depth].movesBuf)
	var best *EvaluatedMove
	best = negativeInf
	for _, move := range moves {
		v := negamaxSimple(board, playerNo, eval, move, -beta, -alpha, depth+1, shared)
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
