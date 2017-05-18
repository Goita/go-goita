package goita

import (
	"reflect"
	"testing"
)

func TestPlayer_pushKoma(t *testing.T) {
	type args struct {
		koma     Koma
		faceDown bool
	}
	tests := []struct {
		name   string
		fields *Player
		args   args
		want   *Player
	}{
		// TODO: Add test cases.
		{
			"11112233->1",
			&Player{
				hand:         ParseKomaArray("11112233"),
				field:        ParseKomaArray("00000000"),
				handCounter:  8,
				hiddenfield:  ParseKomaArray("00000000"),
				fieldCounter: 0,
			},
			args{Shi, false},
			&Player{
				hand:         ParseKomaArray("01112233"),
				field:        ParseKomaArray("10000000"),
				handCounter:  7,
				hiddenfield:  ParseKomaArray("10000000"),
				fieldCounter: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Player{
				hand:         tt.fields.hand,
				field:        tt.fields.field,
				hiddenfield:  tt.fields.hiddenfield,
				handCounter:  tt.fields.handCounter,
				fieldCounter: tt.fields.fieldCounter,
			}
			p.pushKoma(tt.args.koma, tt.args.faceDown)
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("Player.pushKoma() = %v, want %v", p, tt.want)
			}
		})
	}
}
