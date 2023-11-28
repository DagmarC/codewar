package anagrams

import (
	"math/big"
	"testing"
)

func TestListPosition(t *testing.T) {
	tcs := []struct {
		w    string
		want *big.Int
	}{
		{w: "AA", want: big.NewInt(1)},
		{w: "BAAA", want: big.NewInt(4)},
		{w: "ABAB", want: big.NewInt(2)},
		{w: "QUESTION", want: big.NewInt(24572)},
		{w: "BOOKKEEPER", want: big.NewInt(10743)},
	}

	for _, tc := range tcs {
		if got := ListPosition(tc.w); got.Int64() != tc.want.Int64() {
			t.Errorf("w: %s: got %d, want %d", tc.w, got, tc.want)
		}
	}
}
