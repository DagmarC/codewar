package dynamicprog

import (
	"reflect"
	"testing"
)

func TestBestSum(t *testing.T) {
	tests := []struct {
		target   int
		arr      []int
		expected []int
	}{
		{7, []int{5, 3, 4, 7}, []int{7}},
		{8, []int{2, 3, 5}, []int{5, 3}},
		{8, []int{1, 4, 5}, []int{4, 4}},
		{100, []int{1, 2, 5, 25}, []int{25, 25, 25, 25}},
	}

	for _, test := range tests {
		result := bestSum(test.target, test.arr, nil)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("bestSum(%d, %v) = %v; want %v", test.target, test.arr, result, test.expected)
		}
	}
}
