package main

import (
	"fmt"
	"os"
	"time"

	"flag"

	"github.com/Goita/go-goita/goita"
	"github.com/Goita/go-goita/search"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no history was input")
		return
	}
	fmt.Println(os.Args)

	history := flag.String("h", "11244556,12234569,11123378,11113457,s3,371,411,115,2p,3p,4p,145,252,3p,4p,124,2p", "goita history string")
	flag.Parse()

	board := goita.ParseBoard(*history)
	start := time.Now()
	ret := search.Solve(board)
	elapsed := time.Since(start)
	fmt.Println(ret)
	for _, r := range ret {
		fmt.Printf("move:[%v] score:[%v] %v\n\n", r.Move, r.Score, r.History.History(board.Turn))
	}
	fmt.Printf("execution time: %s\n", elapsed)
}
