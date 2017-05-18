package goita

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strings"
	"unsafe"
)

// Koma represents koma of goita
type Koma byte

// KomaArray has 8 koma (including empty)
type KomaArray []Koma

// GetByte returns the raw byte value
func (koma *Koma) GetByte() byte {
	return byte(*koma)
}

// GetBytes returns the raw byte array
func (k *KomaArray) GetBytes() []byte {
	b := *(*[]byte)(unsafe.Pointer(k))
	return b
}

// Empty (0x0) represents no koma is here
const Empty Koma = 0x0

// Hidden (0xf) is face-down koma
const Hidden Koma = 0xf

// Shi (0x1)
const Shi Koma = 0x1

// Gon (0x2)
const Gon Koma = 0x2

// Bakko (0x3)
const Bakko Koma = 0x3

// Gin (0x4)
const Gin Koma = 0x4

// Kin (0x5)
const Kin Koma = 0x5

// Kaku (0x6)
const Kaku Koma = 0x6

// Hisha (0x7)
const Hisha Koma = 0x7

// Ou (0x8)
const Ou Koma = 0x8

// ParseKoma converts string to Koma byte value
func ParseKoma(str string) Koma {
	switch str {
	case "0":
		return Empty
	case "1":
		return Shi
	case "2":
		return Gon
	case "3":
		return Bakko
	case "4":
		return Gin
	case "5":
		return Kin
	case "6":
		return Kaku
	case "7":
		return Hisha
	case "8", "9":
		return Ou
	case "x":
		return Hidden
	default:
		panic(fmt.Sprintf("Invalid koma string value %v was given", str))
	}
}

// GetScore returns the koma finnish score
func (koma Koma) GetScore() int {
	if koma == Hidden {
		panic("cannot get the score of Hidden")
	}
	if koma == Empty {
		panic("cannot get the score of Empty")
	}
	return int(math.Floor(float64(koma/2.0))*10 + 10)
}

// IsKing returns true if the koma is Ou or Gyoku
func (koma Koma) IsKing() bool {
	return koma == Ou
}

// IsShi returns true if the koma is Shi
func (koma Koma) IsShi() bool {
	return koma == Shi
}

// IsEmpty returns true if the koma is Empty
func (koma Koma) IsEmpty() bool {
	return koma == Empty
}

// IsHidden returns true if the koma is Hidden
func (koma Koma) IsHidden() bool {
	return koma == Hidden
}

// CanBlock returns true if the koma can block the target koma
func (koma Koma) CanBlock(target Koma) bool {
	if koma.IsKing() {
		if target.IsShi() || target == Gon {
			return false
		}
		return true
	}
	return koma == target
}

func (koma Koma) String() string {
	if koma.IsHidden() {
		return "x"
	}
	return fmt.Sprintf("%x", byte(koma))
}

// ParseKomaArray creates array of koma from string
func ParseKomaArray(handOrField string) KomaArray {
	if len(handOrField) > FieldLength {
		panic(fmt.Sprintf("hand_or_field must be <= 8 length string. but %v was given", handOrField))
	}
	arr := strings.Split(handOrField, "")
	ret := make(KomaArray, 0, FieldLength)
	for _, v := range arr {
		ret = append(ret, ParseKoma(v))
	}
	return ret
}

// NewKomaArrayFromBytes creates array of koma from bytes
func NewKomaArrayFromBytes(handOrField []byte) *KomaArray {
	if len(handOrField) > FieldLength {
		panic(fmt.Sprintf("hand_or_field must be <= 8 length string. but %v was given", handOrField))
	}
	return (*KomaArray)(unsafe.Pointer(&handOrField))
}

// GetUnique gets distinct koma
func (k *KomaArray) GetUnique() KomaArray {
	unqMap := make(map[Koma]bool)
	unq := make(KomaArray, 0, FieldLength)
	for _, v := range *k {
		if v.IsEmpty() || v.IsHidden() {
			continue
		}
		if _, ok := unqMap[v]; !ok {
			unqMap[v] = true
			unq = append(unq, v)
		}
	}
	return unq
}

// Index returns the index of the first koma in array, or -1 for nothing found
func (k *KomaArray) Index(koma Koma) int {
	return bytes.IndexByte(k.GetBytes(), koma.GetByte())
}

// Contains detects the koma is in the hand
func (k *KomaArray) Contains(koma Koma) bool {
	return bytes.IndexByte(k.GetBytes(), koma.GetByte()) >= 0
}

// Count count up the koma in the hand
func (k *KomaArray) Count(koma Koma) int {
	return bytes.Count(k.GetBytes(), []byte{koma.GetByte()})
}

// Implements sort interface

// Len returns length of KomaArray
func (k *KomaArray) Len() int {
	return len(*k)
}

// Less returns true if array[i] is less than array[j]
func (k *KomaArray) Less(i, j int) bool {
	return (*k)[i] < (*k)[j]
}

// More returns true if array[i] is more than array[j]. it's for descending sort.
func (k *KomaArray) More(i, j int) bool {
	return (*k)[i] > (*k)[j]
}

// Swap changes the place of 2 items by given indexes
func (k *KomaArray) Swap(i, j int) {
	(*k)[i], (*k)[j] = (*k)[j], (*k)[i]
}

// SortDesc makes the array order sorted descending
func (k *KomaArray) SortDesc() {
	sort.Slice(k, k.More)
}

// Search returns the index of koma
func (k *KomaArray) Search(koma Koma) int {
	return sort.Search(len(*k), func(i int) bool { return (*k)[i] <= koma })
}

func (k *KomaArray) String() string {
	// []byte append solution is good for unknown length string concatination
	str := make([]byte, 0, FieldLength*10)
	for _, v := range *k {
		str = append(str, v.String()...)
	}
	return string(str)
}
