package goita

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Goita/go-goita/util"
)

// KomaArray has 8 koma (including empty)
type KomaArray []Koma

// KomaRing returns a set of all koma
func KomaRing() KomaArray {
	ring := make(KomaArray, 0, 32)
	s := "11111111112222333344445555667789"
	for _, v := range strings.Split(s, "") {
		ring = append(ring, ParseKoma(v))
	}
	return ring
}

// Shuffle randomize the provided array order with Fisher-Yates algorithm
func Shuffle(array KomaArray) {
	rng := util.GetRNG()
	for i := len(array) - 1; i > 0; i = i - 1 {
		j := rng.Intn(i)
		array[i], array[j] = array[j], array[i]
	}
}

// ParseKomaArray creates array of koma from string
func ParseKomaArray(komas string) KomaArray {
	if len(komas) > FieldLength {
		panic(fmt.Sprintf("hand_or_field must be <= 8 length string. but %v was given", komas))
	}
	arr := strings.Split(komas, "")
	ret := make(KomaArray, 0, FieldLength)
	for _, v := range arr {
		ret = append(ret, ParseKoma(v))
	}
	return ret
}

// NewKomaArrayFromBytes creates array of koma from bytes
func NewKomaArrayFromBytes(komas []byte) KomaArray {
	l := len(komas)
	if l > FieldLength {
		panic(fmt.Sprintf("hand_or_field must be <= 8 length string. but %v was given", komas))
	}
	arr := make(KomaArray, l, FieldLength)
	for i, b := range komas {
		arr[i] = Koma(b)
	}
	return arr
}

// GetUnique gets distinct koma, excluding 0:Empty and f:Hidden
func (k KomaArray) GetUnique() KomaArray {
	// koma range 0-8 ()
	unqMap := make([]Koma, 9)
	unq := make(KomaArray, 0, FieldLength)
	return k.Unique(unqMap, unq)
}

// Unique gets distinct koma, excluding 0:Empty and f:Hidden (no memory allocation)
func (k KomaArray) Unique(mapBuf []Koma, buf KomaArray) KomaArray {
	unq := buf[:0]
	for i := 1; i < 9; i++ {
		mapBuf[i] = 0
	}
	for _, v := range k {
		if v.IsEmpty() || v.IsHidden() {
			continue
		}
		mapBuf[v] = v
	}
	for _, v := range mapBuf {
		if v == 0 {
			continue
		}
		unq = append(unq, v)
	}
	return unq
}

// Index returns the index of the first koma in array, or -1 for nothing found
func (k KomaArray) Index(koma Koma) int {
	for i, b := range k {
		if b == koma {
			return i
		}
	}
	return -1
}

// Contains detects the koma is in the hand
func (k KomaArray) Contains(koma Koma) bool {
	return k.Index(koma) >= 0
}

// Count count up the koma in the hand
func (k KomaArray) Count(koma Koma) int {
	c := 0
	for _, b := range k {
		if b == koma {
			c++
		}
	}
	return c
}

// Implements sort interface

// Len returns length of KomaArray
func (k KomaArray) Len() int {
	return len(k)
}

// Less returns true if array[i] is less than array[j]
func (k KomaArray) Less(i, j int) bool {
	return k[i] < k[j]
}

// More returns true if array[i] is more than array[j]. it's for descending sort.
func (k KomaArray) More(i, j int) bool {
	return k[i] > k[j]
}

// Swap changes the place of 2 items by given indexes
func (k KomaArray) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

// Sort makes the array order sorted ascending
func (k KomaArray) Sort() {
	sort.Slice(k, func(i, j int) bool { return k[i] < k[j] })
}

// SortDesc makes the array order sorted descending
func (k KomaArray) SortDesc() {
	sort.Slice(k, func(i, j int) bool { return k[i] > k[j] })
}

// Search returns the index of koma
func (k KomaArray) Search(koma Koma) int {
	return sort.Search(len(k), func(i int) bool { return k[i] <= koma })
}

// Hand converts KomaArray to Hand
func (k KomaArray) Hand() Hand {
	// lazy way
	return ParseHand(k.String())
}

func (k KomaArray) String() string {
	// []byte append solution is good for unknown length string concatination
	str := make([]byte, len(k))
	for i, v := range k {
		if v == Hidden {
			str[i] = 'x'
		} else {
			str[i] = byte(v) + '0'
		}
	}
	return string(str)
}

// StringOld is older way of Stringer method
func (k KomaArray) StringOld() string {
	// []byte append solution is good for unknown length string concatination
	str := make([]byte, 0, FieldLength*10)
	for _, v := range k {
		str = append(str, v.String()...)
	}
	return string(str)
}
