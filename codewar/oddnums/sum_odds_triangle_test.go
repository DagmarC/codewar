package oddnums

import "testing"

func TestSumOdssTriangle(t *testing.T) {
	tc := []struct{
		n int
		expected int
	} {
		{
			n: 1,
			expected: 1,
		},
		{
			n: 2,
			expected: 8,
		},
		{
			n: 13,
			expected: 2197,
		},
		{
			n: 42,
			expected: 74088,
		},
		{
			n: 86,
			expected: 636056,
		},
		{
			n: 93,
			expected: 804357,
		},
	}

	for _, test := range tc {
		if got := RowSumOddNumbers(test.n); got != test.expected {
			t.Errorf("got %d, expected %d", got, test.expected)
		}
	}

}