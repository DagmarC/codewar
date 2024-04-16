package arrays

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConnectationArr(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{nums: []int{1}, want: []int{1, 1}},
		{nums: []int{1, 2, 1}, want: []int{1, 2, 1, 1, 2, 1}},
		{nums: []int{}, want: []int{}},
	}
	for _, tc := range testCases {
		if got := getConcatenation(tc.nums); !cmp.Equal(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}
