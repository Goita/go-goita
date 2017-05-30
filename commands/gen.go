package commands

import (
	"flag"
	"fmt"

	"github.com/Goita/go-goita/goita"
)

// CmdGen generate a random history string
var CmdGen = &Command{
	Name: "gen",
	Run: func(args []string) {
		fmt.Println("gen cmd run with args = ", args)

		// turn := flag.Int("t", 0, "random play turns")
		dealer := flag.Int("d", 0, "dealer number")
		flag.CommandLine.Parse(args)

		board := goita.NewBoardRandom(*dealer)
		fmt.Println(board.String())
	},
}
