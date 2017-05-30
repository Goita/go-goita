package main

import (
	"fmt"
	"os"

	"github.com/Goita/go-goita/commands"
)

func main() {
	var commands = []*commands.Command{
		commands.CmdSolve,
		commands.CmdGen,
	}

	if len(os.Args) < 2 {
		fmt.Println("usage")
		return
	}
	subcmd := os.Args[1]

	for _, c := range commands {
		if subcmd == c.Name {
			c.Run(os.Args[2:])
		}
	}

	fmt.Println(os.Args)

	os.Exit(0)
}
