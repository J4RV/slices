// Code generated by yasupGen; DO NOT EDIT.

package yasup_test

import (
	yasup "github.com/j4rv/yasup"
	"testing"
)

func Test_IntInsert(t *testing.T) {
	type testCase struct {
		name     string
		slice    []int
		insertAt int
	}
	base := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tcs := []testCase{
		{"First", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0},
		{"Middle", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, len(base) / 2},
		{"Last", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, len(base)},
		{"Empty slice", []int{}, 0},
		{"Nil slice", nil, 0},
	}
	for _, tc := range tcs {
		yasup.IntInsert(&tc.slice, -2147483648, tc.insertAt)
		if tc.slice[tc.insertAt] != (-2147483648) {
			t.Error(tc)
		}
	}
}

func Test_IntFastShuffle(t *testing.T) {
	shuffles := [][]int{}
	for i := 0; i < 8; i++ {
		or := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		yasup.IntFastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i >= j {
				continue
			}
			if yasup.IntEquals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	yasup.IntFastShuffle(nil)
}

func Test_IntSecureShuffle(t *testing.T) {
	shuffles := [][]int{}
	for i := 0; i < 8; i++ {
		or := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		yasup.IntSecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i >= j {
				continue
			}
			if yasup.IntEquals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	yasup.IntSecureShuffle(nil)
}

func Test_IntEquals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []int
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []int{}, true},
		{"Right nil, left empty", []int{}, nil, true},
		{"Left nil, right not empty", nil, []int{-2147483648}, false},
		{"Right nil, left not empty", []int{-2147483648}, nil, false},

		{"Equals empty", []int{}, []int{}, true},
		{"Equals 1", []int{-2147483648}, []int{-2147483648}, true},
		{"Equals 2", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -2147483648}, false},
		{"Not equals 2", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{}, false},
		{"Not equals 3", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{-2147483648, -2147483648}, false},
	}
	for _, tc := range tcs {
		got := yasup.IntEquals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc)
		}
	}
}
