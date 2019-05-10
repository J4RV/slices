// Code generated by yasupGen; DO NOT EDIT.

package yasup_test

import (
	yasup "github.com/j4rv/yasup"
	"testing"
)

func Test_BoolInsert(t *testing.T) {
	type testCase struct {
		name     string
		slice    []bool
		insertAt int
	}
	base := []bool{true, false, true, false, true, false, true, false, true, false, true, false}
	tcs := []testCase{
		{"First", []bool{true, false, true, false, true, false, true, false, true, false, true, false}, 0},
		{"Middle", []bool{true, false, true, false, true, false, true, false, true, false, true, false}, len(base) / 2},
		{"Last", []bool{true, false, true, false, true, false, true, false, true, false, true, false}, len(base)},
		{"Empty slice", []bool{}, 0},
		{"Nil slice", nil, 0},
	}
	for _, tc := range tcs {
		yasup.BoolInsert(&tc.slice, false, tc.insertAt)
		if tc.slice[tc.insertAt] != (false) {
			t.Error(tc)
		}
	}
}

func Test_BoolFastShuffle(t *testing.T) {
	shuffles := [][]bool{}
	for i := 0; i < 8; i++ {
		or := []bool{true, false, true, false, true, false, true, false, true, false, true, false}
		yasup.BoolFastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i >= j {
				continue
			}
			if yasup.BoolEquals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	yasup.BoolFastShuffle(nil)
}

func Test_BoolSecureShuffle(t *testing.T) {
	shuffles := [][]bool{}
	for i := 0; i < 8; i++ {
		or := []bool{true, false, true, false, true, false, true, false, true, false, true, false}
		yasup.BoolSecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i >= j {
				continue
			}
			if yasup.BoolEquals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	yasup.BoolSecureShuffle(nil)
}

func Test_BoolEquals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []bool
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []bool{}, true},
		{"Right nil, left empty", []bool{}, nil, true},
		{"Left nil, right not empty", nil, []bool{false}, false},
		{"Right nil, left not empty", []bool{false}, nil, false},

		{"Equals empty", []bool{}, []bool{}, true},
		{"Equals 1", []bool{false}, []bool{false}, true},
		{"Equals 2", []bool{true, false, true, false, true, false, true, false, true, false, true, false}, []bool{true, false, true, false, true, false, true, false, true, false, true, false}, true},
		{"Not equals 1", []bool{true, false, true, false, true, false, true, false, true, false, true, false}, []bool{true, false, true, false, true, false, true, false, true, false, true, false, false}, false},
		{"Not equals 2", []bool{true, false, true, false, true, false, true, false, true, false, true, false}, []bool{}, false},
		{"Not equals 3", []bool{true, false, true, false, true, false, true, false, true, false, true, false}, []bool{false, false}, false},
	}
	for _, tc := range tcs {
		got := yasup.BoolEquals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc)
		}
	}
}
