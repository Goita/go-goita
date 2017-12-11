package search

import (
	"fmt"
	"testing"

	"github.com/Goita/go-goita/goita"
)

func TestSolve(t *testing.T) {
	board := goita.ParseBoard("11244556,12234569,11123378,11113457,s3,371,411,115,2p,3p,4p,145,252,3p,4p,124,2p")
	ret := Solve(board)
	moves := board.GetPossibleMoves()
	results := make([]*EvaluatedMove, 0, len(moves))
	for r := range ret {
		results = append(results, r)
		fmt.Printf("move:[%v] score:[%v] %v\n\n", r.Move, r.Score, r.History.History(board.Turn))
	}
	if len(results) != 4 || results[0].Score != -40 || results[1].Score != -40 || results[2].Score != -50 || results[3].Score != -40 {
		t.Errorf("search.Solve() = %v, want [p:-40] [81:-40] [82:-50] [83:-40]", ret)
	}
	//[[p:-40 -> 3p,443,1p,2p,3p,415,1p,2p,3p,417] [81:-40 -> 381,413,1p,2p,3p,454,1p,2p,3p,417] [82:-50 -> 382,4p,1p,2p,311,413,1p,232,3p,4p,1p,264,3p,4p,1p,218] [83:-40 -> 383,4p,1p,2p,311,413,1p,234,3p,4p,1p,261,3p,415,1p,2p,3p,447]]
}

func BenchmarkSolve(b *testing.B) {
	board := goita.ParseBoard("11244556,12234569,11123378,11113457,s3,371,411,115,2p,3p,4p,145,252,3p,4p,124,2p")
	for i := 0; i < b.N; i++ {
		Solve(board)
	}
	// fmt.Println(ret)
}
