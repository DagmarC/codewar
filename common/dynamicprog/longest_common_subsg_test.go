package dynamicprog

import (
	"fmt"
	"testing"
)

func TestLCS(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want string
	}{
		{s1: "AGGTAB", s2: "GXTXAYB", want: "GTAB"},
		{s1: "AB", s2: "A", want: "A"},
		{s1: "CBB", s2: "BCB", want: "CB"},
		{s1: "", s2: "BCB", want: ""},
		{s1: "AAAA", s2: "AJA", want: "AA"},

	}
	var got string
	for _, tc := range tests {
		if got = LCS(tc.s1, tc.s2); got != tc.want {
			t.Errorf("ERROR: s1=%s and s2=%s -> got %s, want %s\n", tc.s1, tc.s2, got, tc.want)
		}
		fmt.Printf("s1=%10s\ts2=%10s\t\tLCS=%10s\n", tc.s1, tc.s2, got)
	}
}
