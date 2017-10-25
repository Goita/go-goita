package goita

import "testing"

func TestKomaArray_String(t *testing.T) {
	tests := []struct {
		name string
		k    KomaArray
		want string
	}{
		// TODO: Add test cases.
		{"12345678", ParseKomaArray("12345678"), "12345678"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.String(); got != tt.want {
				t.Errorf("KomaArray.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKomaArray_StringOld(t *testing.T) {
	tests := []struct {
		name string
		k    KomaArray
		want string
	}{
		// TODO: Add test cases.
		{"12345678", ParseKomaArray("12345678"), "12345678"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.StringOld(); got != tt.want {
				t.Errorf("KomaArray.StringOld() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_String(b *testing.B) {
	arr := ParseKomaArray("12345678")
	for i := 0; i < b.N; i++ {
		_ = arr.String()
	}
}

func Benchmark_StringOld(b *testing.B) {
	arr := ParseKomaArray("12345678")
	for i := 0; i < b.N; i++ {
		_ = arr.StringOld()
	}
}
