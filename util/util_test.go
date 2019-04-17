package util

import (
	"testing"
)

func TestShiftTurn(t *testing.T) {
	type args struct {
		turn   int
		offset int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0+2", args{0, 2}, 2},
		{"0+2", args{3, 2}, 1},
		{"0-2", args{0, 2}, 2},
		{"0-2", args{3, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShiftTurn(tt.args.turn, tt.args.offset); got != tt.want {
				t.Errorf("ShiftTurn() = %v, want %v", got, tt.want)
			}
		})
	}
}
