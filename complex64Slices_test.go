package slices_test

import "testing"
import "github.com/j4rv/slices"

func Test_Complex64FastShuffle(t *testing.T) {
	shuffles := [][]complex64{}
	for i := 0; i < 8; i++ {
		or := []complex64{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}
		slices.Complex64FastShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Complex64Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
}

func Test_Complex64SecureShuffle(t *testing.T) {
	shuffles := [][]complex64{}
	for i := 0; i < 8; i++ {
		or := []complex64{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}
		slices.Complex64SecureShuffle(or)
		shuffles = append(shuffles, or)
	}
	for i := range shuffles {
		for j := range shuffles {
			if i == j {
				continue
			}
			if slices.Complex64Equals(shuffles[i], shuffles[j]) {
				// If there is any collision in 8 shuffles, the Shuffle function is probably broken
				t.Fail()
			}
		}
	}
}

func Test_Complex64Equals(t *testing.T) {
	type TestCase struct {
		name string
		a, b []complex64
		exp  bool
	}
	tcs := []TestCase{
		// nil checks
		{"Equals nil", nil, nil, true},
		// golang treats empty and nil slices as the same thing in most cases, we'll do the same
		{"Left nil, right empty", nil, []complex64{}, true},
		{"Right nil, left empty", []complex64{}, nil, true},
		{"Left nil, right not empty", nil, []complex64{-5.64 + 15.82i}, false},
		{"Right nil, left not empty", []complex64{-5.64 + 15.82i}, nil, false},

		{"Equals empty", []complex64{}, []complex64{}, true},
		{"Equals 1", []complex64{-5.64 + 15.82i}, []complex64{-5.64 + 15.82i}, true},
		{"Equals 2", []complex64{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}, []complex64{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}, true},
		{"Not equals 1", []complex64{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}, []complex64{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i, -5.64 + 15.82i}, false},
		{"Not equals 2", []complex64{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}, []complex64{}, false},
		{"Not equals 3", []complex64{0 + 0i, 1 + 2i, -2 + 7.5i, 3 + 42.1i, 4 - 74.6i, -5 + 4i, 6 - 88i, 7 - 0i, 8 + 100i, 9 + 99i}, []complex64{-5.64 + 15.82i, -5.64 + 15.82i}, false},
	}
	for _, tc := range tcs {
		got := slices.Complex64Equals(tc.a, tc.b)
		if got != tc.exp {
			t.Error(tc.name)
		}
	}
}
