// Code generated by yasupGen; DO NOT EDIT.

package yasup

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

var zeroValueUint64 uint64

//Uint64Insert will append elem at the position i. Might return ErrIndexOutOfBounds.
func Uint64Insert(sl *[]uint64, elem uint64, i int) error {
	if i < 0 || i > len(*sl) {
		return ErrIndexOutOfBounds
	}
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
	return nil
}

//Uint64Delete delete the element at the position i. Might return ErrIndexOutOfBounds.
func Uint64Delete(sl *[]uint64, i int) error {
	if i < 0 || i >= len(*sl) {
		return ErrIndexOutOfBounds
	}
	*sl = append((*sl)[:i], (*sl)[i+1:]...)
	return nil
}

//Uint64Contains will return true if elem is present in the slice and false otherwise.
func Uint64Contains(sl []uint64, elem uint64) bool {
	for i := range sl {
		if sl[i] == elem {
			return true
		}
	}
	return false
}

//Uint64Index returns the index of the first instance of elem, or -1 if elem is not present.
func Uint64Index(sl []uint64, elem uint64) int {
	for i := range sl {
		if sl[i] == elem {
			return i
		}
	}
	return -1
}

//Uint64Push is equivalent to Uint64Insert with index len(*sl)
func Uint64Push(sl *[]uint64, elem uint64) {
	Uint64Insert(sl, elem, len(*sl))
}

//Uint64FrontPush is equivalent to Uint64Insert with index 0
func Uint64FrontPush(sl *[]uint64, elem uint64) {
	Uint64Insert(sl, elem, 0)
}

//Uint64Pop is equivalent to getting and removing the last element of the slice. Might return ErrEmptySlice.
func Uint64Pop(sl *[]uint64) (uint64, error) {
	if len(*sl) == 0 {
		return zeroValueUint64, ErrEmptySlice
	}
	last := len(*sl) - 1
	ret := (*sl)[last]
	Uint64Delete(sl, last)
	return ret, nil
}

//Uint64Pop is equivalent to getting and removing the first element of the slice. Might return ErrEmptySlice.
func Uint64FrontPop(sl *[]uint64) (uint64, error) {
	if len(*sl) == 0 {
		return zeroValueUint64, ErrEmptySlice
	}
	ret := (*sl)[0]
	Uint64Delete(sl, 0)
	return ret, nil
}

//Uint64Equals compares two uint64 slices. Returns true if their elements are equal.
func Uint64Equals(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//Uint64FastShuffle will randomly swap the uint64 elements of a slice using math/rand (fast but not cryptographycally secure).
func Uint64FastShuffle(sp []uint64) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Uint64SecureShuffle will randomly swap the uint64 elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func Uint64SecureShuffle(sp []uint64) error {
	var i int64
	size := int64(len(sp)) - 1
	for i = 0; i < size+1; i++ {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(size))
		if err != nil {
			return err
		}
		randI := bigRandI.Int64()
		sp[size-i], sp[randI] = sp[randI], sp[size-i]
	}
	return nil
}
