package dynamicproblems

import "testing"

var tests = []struct {
	coins []int
	sum   int
	want  int
}{
	{coins: []int{1, 2, 5}, sum: 11, want: 3}, // 11 = 5 + 5 + 1
	{coins: []int{2}, sum: 3, want: -1},
	{coins: []int{1}, sum: 0, want: 0},
	{coins: []int{1, 2, 3}, sum: 6, want: 2},
	{coins: []int{2, 5, 10, 1}, sum: 27, want: 4},
	{coins: []int{3, 5, 1}, sum: 10, want: 2},

}

func TestCoinsChangeMin(t *testing.T) {
	for _, tc := range tests {
		var got int
		if got = coinChange(tc.coins, tc.sum); got != tc.want {
			t.Errorf("DP Coin Change Fewest path: got %d want %d", got, tc.want)
		}
	}
}
