package goita

import (
	"fmt"
	"strings"
)

// Hand holds komas in hand. The index corresponds to koma value and each item holds the counter of the Koma.
type Hand []uint

// ParseHand creates array of koma from string
func ParseHand(hand string) Hand {
	if len(hand) > FieldLength {
		panic(fmt.Sprintf("hand must be <= 8 length string. but %v was given", hand))
	}
	arr := strings.Split(hand, "")
	ret := make(Hand, 9, 9)
	for _, v := range arr {
		koma := ParseKoma(v)
		if koma == Empty || koma == Hidden {
			continue
		}
		ret[koma]++
	}
	return ret
}

// GetUnique gets distinct koma, excluding 0:Empty and f:Hidden
func (h Hand) GetUnique() KomaArray {
	// koma range 0-8 ()
	unq := make(KomaArray, 9)
	count := h.Unique(unq)
	return unq[:count]
}

// Unique gets distinct koma, excluding 0:Empty and f:Hidden (no memory allocation)
func (h Hand) Unique(buf KomaArray) (count int) {
	count = 0
	for i := 0; i < len(h); i++ {
		if h[i] == 0 {
			continue
		}
		buf[count] = Koma(i)
		count++
	}
	return count
}

// Count returns the count of the koma in hand
func (h Hand) Count(koma Koma) int {
	return int(h[koma])
}

// Contains returns true if the koma in hand
func (h Hand) Contains(koma Koma) bool {
	return h[koma] > 0
}

// Array converts Hand to KomaArray
func (h Hand) Array() KomaArray {
	arr := make(KomaArray, 0, FieldLength)
	for i := 1; i < 9; i++ {
		for c := 0; c < int(h[i]); c++ {
			arr = append(arr, Koma(i))
		}
	}
	return arr
}

func (h Hand) String() string {
	str := make([]byte, 0, FieldLength)
	for i := 1; i < 9; i++ {
		for c := 0; c < int(h[i]); c++ {
			str = append(str, Koma(i).String()...)
		}
	}
	return string(str)
}
