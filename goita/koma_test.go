package goita

import (
	"reflect"
	"testing"
)

func TestKomaArray_GetUnique(t *testing.T) {
	tests := []struct {
		name string
		k    KomaArray
		want KomaArray
	}{
		// TODO: Add test cases.
		{"11112234", ParseKomaArray("11112234"), ParseKomaArray("1234")},
		{"12345678", ParseKomaArray("12345678"), ParseKomaArray("12345678")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.GetUnique(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KomaArray.GetUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseKoma(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want Koma
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseKoma(tt.args.str); got != tt.want {
				t.Errorf("ParseKoma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKoma_String(t *testing.T) {
	tests := []struct {
		name string
		koma Koma
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.koma.String(); got != tt.want {
				t.Errorf("Koma.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKomaArray_String(t *testing.T) {
	tests := []struct {
		name string
		k    KomaArray
		want string
	}{
		// TODO: Add test cases.
		{"11112345", ParseKomaArray("11112345"), "11112345"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.String(); got != tt.want {
				t.Errorf("KomaArray.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
