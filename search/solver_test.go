package search

import (
	"testing"

	"sort"

	"github.com/Goita/go-goita/goita"
)

func TestSolve(t *testing.T) {
	board := goita.ParseBoard("11244556,12234569,11123378,11113457,s3,371,411,115,2p,3p,4p,145,252,3p,4p,124,2p")
	ret := Solve(board)
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Move.Hash() < ret[j].Move.Hash()
	})
	if len(ret) != 4 || ret[0].Score != -40 || ret[1].Score != -40 || ret[2].Score != -50 || ret[3].Score != -40 {
		t.Errorf("search.Solve() = %v, want [p:-40] [81:-40] [82:-50] [83:-40]", ret)
	}
	//[[p:-40 -> 3p,443,1p,2p,3p,415,1p,2p,3p,417] [81:-40 -> 381,413,1p,2p,3p,454,1p,2p,3p,417] [82:-50 -> 382,4p,1p,2p,311,413,1p,232,3p,4p,1p,264,3p,4p,1p,218] [83:-40 -> 383,4p,1p,2p,311,413,1p,234,3p,4p,1p,261,3p,415,1p,2p,3p,447]]
}

func TestHistory(t *testing.T) {
	board := goita.ParseBoard("11244556,12234569,11123378,11113457,s3,371,411,115,2p,3p,4p,145,252,3p,4p,124,2p")
	ret := Solve(board)
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Move.Hash() < ret[j].Move.Hash()
	})

	want := []string{
		"[p:-40 -> 3p,443,1p,2p,3p,415,1p,2p,3p,417]",
		"[81:-40 -> 381,413,1p,2p,3p,414,1p,2p,3p,457]",
		"[82:-50 -> 382,4p,1p,2p,311,413,1p,231,3p,414,1p,242,3p,4p,1p,268]",
		"[83:-40 -> 383,4p,1p,2p,311,413,1p,231,3p,414,1p,2p,3p,457]",
	}

	if len(ret) != 4 {
		t.Errorf("search.Solve() = %v, want [p:-40] [81:-40] [82:-50] [83:-40]", ret)
	}
	for i, r := range ret {
		if str := r.StringHistory(board.Turn); str != want[i] {
			t.Errorf("search.Solve() = %v, want %v", str, want[i])
		}
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
