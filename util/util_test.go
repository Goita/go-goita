package util

import (
	"bytes"
	"testing"
)

func TestShuffle(t *testing.T) {
	sumIndex := 0
	loop := 1000
	for n := 0; n < loop; n++ {
		arr := make([]byte, 100)
		for i := range arr {
			arr[i] = byte(i)
		}
		Shuffle(arr)
		sumIndex += bytes.IndexByte(arr, byte(0))
	}
	if indicator := sumIndex / loop; indicator < 45 && indicator > 55 {
		t.Error("Shuffle seems to be not working propery")
	}
}

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
