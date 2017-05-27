package goita

import "testing"

// go test ./goita -bench moveHash -benchmem -benchtime 1s
func Benchmark_moveHash(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for block := 0; block < 9; block++ {
			for attack := 0; attack < 9; attack++ {
				moveHash(Koma(block), Koma(attack), true)
				moveHash(Koma(block), Koma(attack), false)
			}
		}
	}
}
