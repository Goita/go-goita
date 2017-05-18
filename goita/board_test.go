package goita

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseBoard(t *testing.T) {
	type args struct {
		historyString string
	}
	tests := []struct {
		name string
		args args
		// want is input historyString
	}{
		// TODO: Add test cases.
		{"initial", args{"12345678,12345679,11112345,11112345,s1"}},
		{"end of deal", args{"22221678,11111345,11345679,11345345,s1,112,2p,3p,4p,162,2p,3p,4p,172,2p,3p,4p,128"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseBoard(tt.args.historyString).String(); !reflect.DeepEqual(got, tt.args.historyString) {
				t.Errorf("ParseBoard() = %v, want %v", got, tt.args.historyString)
			}
		})
	}
}

func TestBoard_GetPossibleMoves(t *testing.T) {
	tests := []struct {
		name  string
		board string
		want  string
	}{
		// TODO: Add test cases.
		{"gon-ou finish", "12345678,12345679,11112345,11112345,s1,113,2p,3p,431,1p,2p,315,4p,156,267,3p,4p,174,242,3p,4p", "p,28"},
		{"end of deal", "12345678,12345679,11112345,11112345,s1,113,2p,3p,431,1p,2p,315,4p,156,267,3p,4p,174,242,3p,4p,128", ""},
		{"king's double-up finish", "12667789,12345543,11112345,11112345,s1,116,2p,3p,4p,126,2p,3p,4p,177,2p,3p,4p", "88"},
		{"finish with yaku", "22235567,12345679,11133448,11111145,s1", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := ParseBoard(tt.board)
			got := b.GetPossibleMoves()
			ret := make([]string, 0)
			for _, v := range got {
				ret = append(ret, v.OpenString())
			}
			moves := strings.Join(ret, ",")
			if len(moves) != len(tt.want) {
				t.Errorf("Board.GetPossibleMoves() = %v, want %v", moves, tt.want)
			}
		})
	}
}
