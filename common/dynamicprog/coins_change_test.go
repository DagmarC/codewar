package dynamicprog

import (
	"fmt"
	"testing"
)

var tests = []struct {
	coins []int
	sum   int
	want  int
}{
	{coins: []int{1, 2, 3}, sum: 4, want: 4},
	{coins: []int{2, 5, 3, 6}, sum: 10, want: 5},
}

func TestCoinsChange(t *testing.T) {
	for _, tc := range tests {
		var got int
		if got = countCoinsRec(tc.coins, tc.sum); got != tc.want {
			t.Errorf("Recursove: got %d want %d", got, tc.want)
		}
		if got = countCoinsMemoization(tc.coins, tc.sum); got != tc.want {
			t.Errorf("DP Memoization: got %d want %d", got, tc.want)
		}
		if got = countCoinsTabulation(tc.coins, tc.sum); got != tc.want {
			t.Errorf("DP Tabulation 2D: got %d want %d", got, tc.want)
		}
		if got = countCoinsTabulation1D(tc.coins, tc.sum); got != tc.want {
			t.Errorf("DP Tabulation 1D: got %d want %d", got, tc.want)
		}
		fmt.Printf("Coins %v,sum=%d,\tchanges %d\n", tc.coins, tc.sum, got)
	}
}
