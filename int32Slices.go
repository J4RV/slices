package slices

import (
	crypto "crypto/rand"
	"math/big"
	"math/rand"
)

//Int32Insert will append elem at the position i
func Int32Insert(elem int32, sl *[]int32, i int) {
	*sl = append(*sl, elem)
	copy((*sl)[i+1:], (*sl)[i:])
	(*sl)[i] = elem
}

//Int32FastShuffle will randomly swap the int32 elements of a slice using math/rand (fast but not cryptographycally secure).
func Int32FastShuffle(sp []int32) {
	rand.Shuffle(len(sp), func(i, j int) {
		sp[i], sp[j] = sp[j], sp[i]
	})
}

//Int32SecureShuffle will randomly swap the int32 elements of a slice using crypto/rand (resource intensive but cryptographycally secure).
func Int32SecureShuffle(sp []int32) {
	for i := int64(len(sp) - 1); i >= 0; i-- {
		bigRandI, err := crypto.Int(crypto.Reader, big.NewInt(i+1))
		if err != nil {
			panic(err)
		}
		randI := bigRandI.Int64()
		sp[i], sp[randI] = sp[randI], sp[i]
	}
}

//Int32Equals compares two int32 slices. Returns true if their elements are equal.
func Int32Equals(a, b []int32) bool {
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
