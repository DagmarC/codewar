package tribonacci

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTribonacci(t *testing.T) {
	tc := []struct {
		signature [3]float64
		n         int
		exoected  []float64
	}{
		{
			signature: [3]float64{1, 1, 1},
			n:         10,
			exoected:  []float64{1, 1, 1, 3, 5, 9, 17, 31, 57, 105},
		},
		{
			signature: [3]float64{0, 0, 1},
			n:         10,
			exoected:  []float64{0, 0, 1, 1, 2, 4, 7, 13, 24, 44},
		},
		{
			signature: [3]float64{3, 2, 1},
			n:         10,
			exoected:  []float64{3, 2, 1, 6, 9, 16, 31, 56, 103, 190},
		},
		{
			signature: [3]float64{1, 1, 1},
			n:         1,
			exoected:  []float64{1},
		},
		{
			signature: [3]float64{1, 200, 100},
			n:         0,
			exoected:  []float64{},
		},
	}

	for _, test := range tc {
		if res := Tribonacci(test.signature, test.n); !cmp.Equal(res, test.exoected) {
			t.Errorf("got %v, expected %v with n=%d", res, test.exoected, test.n)
		}
	}
}
