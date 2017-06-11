package commands

import (
	"flag"
	"fmt"
	"time"

	"github.com/Goita/go-goita/goita"
	"github.com/Goita/go-goita/search"
)

// CmdSolve run solve all moves
var CmdSolve = &Command{
	Name: "solve",
	Run: func(args []string) {
		fmt.Println("solve cmd run with args = ", args)

		history := flag.String("h", "11244556,12234569,11123378,11113457,s3,371,411,115,2p,3p,4p,145,252,3p,4p,124,2p", "goita history string")
		flag.CommandLine.Parse(args)

		board := goita.ParseBoard(*history)
		moves := board.GetPossibleMoves()
		fmt.Printf("search begin on %v moves ", len(moves))
		fmt.Println(moves)

		results := make([]*search.EvaluatedMove, 0, len(moves))
		start := time.Now()
		ret := search.Solve(board)
		for r := range ret {
			results = append(results, r)
			fmt.Printf("move:[%v] score:[%v] %v\n\n", r.Move, r.Score, r.History.History(board.Turn))
		}
		elapsed := time.Since(start)
		fmt.Println(results)
		fmt.Printf("execution time: %s\n", elapsed)
	},
}
