package arrays

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindPoisendDuration(t *testing.T) {
	testCases := []struct {
		timeSeries []int
		duration int
		want int
	}{
		{timeSeries: []int{1, 4}, duration: 2, want: 4},
		{timeSeries: []int{1, 2}, duration: 2, want: 3},
		{timeSeries: []int{1, 2, 4, 9}, duration: 3, want: 9}, 
		{timeSeries: []int{}, duration: 1, want: 0},
	}
	for _, tc := range testCases {
		if got := findPoisonedDuration(tc.timeSeries, tc.duration); !cmp.Equal(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}

	// 1 2 3 4 5 6 9 10 11
}