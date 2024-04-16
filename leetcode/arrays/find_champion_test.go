package arrays

import "testing"

func TestFindChampion(t *testing.T) {
	tests := []struct{
		grid [][]int
		want int
	}{
		{grid: [][]int{{0, 1}, {0, 0}}, want: 0},
		{grid: [][]int{{0, 0, 1}, {1, 0, 1}, {0,0,0}}, want: 1},
	}

	for _, tc := range tests {
		if got := findChampion(tc.grid); got != tc.want {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}	