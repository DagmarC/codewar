package slidingwindow

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSlidingWindow(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		want []int
	}{
		{nums: []int{1}, k: 1, want: []int{1}},
		{nums: []int{1, -1}, k: 1, want: []int{1, -1}},
		{nums: []int{1,3,1,2,0,5}, k: 3, want: []int{3,3,2,5}},
		{nums: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, want: []int{3, 3, 5, 5, 6, 7}},
	}
	for _, tc := range testCases {
		if got := maxSlidingWindow(tc.nums, tc.k); !cmp.Equal(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}
