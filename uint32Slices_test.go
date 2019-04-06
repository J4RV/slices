package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_Uint32FastShuffle(t *testing.T) {
	shuffles := [][]uint32{}
	for i := 0; i < 8; i++ {
		or := []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Uint32FastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Uint32Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.Uint32FastShuffle(nil)
}

func Test_Uint32SecureShuffle(t *testing.T) {
	shuffles := [][]uint32{}
	for i := 0; i < 8; i++ {
		or := []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		slices.Uint32SecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Uint32Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
	// check that nil does not panic
	slices.Uint32SecureShuffle(nil)
}

func Test_Uint32Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []uint32
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []uint32{}, true},
		{"Right nil, left empty", []uint32{}, nil, true},
		{"Left nil, right not empty", nil, []uint32{4294967295}, false},
		{"Right nil, left not empty", []uint32{4294967295}, nil, false},

		{"Equals empty", []uint32{}, []uint32{}, true},
		{"Equals 1", []uint32{4294967295}, []uint32{4294967295}, true},
		{"Equals 2", []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{"Not equals 1", []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 4294967295}, false},
		{"Not equals 2", []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint32{}, false},
		{"Not equals 3", []uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []uint32{4294967295, 4294967295}, false},
	}
	for _, tc := range tcs {
		got := slices.Uint32Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
