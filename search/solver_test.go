package search

import (
	"testing"

	"fmt"

	"github.com/Goita/go-goita/goita"
)

func TestSolve(t *testing.T) {
	b := goita.ParseBoard("11244556,12234569,11123378,11113457,s3,371,411,115,2p,3p,4p,145,252,3p,4p,124")
	ret := Solve(b)

	fmt.Println(ret)
}
