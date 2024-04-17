package dynamicproblems

// import (
// 	"testing"

// 	"github.com/google/go-cmp/cmp"
// )

// // n == s1.length == s2.length
// // 1 <= n, x <= 500

// func TestStrEqlMinOperations(t *testing.T) {

// 	testCases := []struct {
// 		s1   string
// 		s2   string
// 		x    int
// 		want int
// 	}{
// 		{s1: "1100011000", s2: "0101001010", x: 2, want: 4},
// 		{s1: "10110", s2: "00011", x: 4, want: -1}, // impossible
// 	}
// 	for _, tc := range testCases {
// 		if got := minOperations(tc.s1, tc.s2, tc.x); !cmp.Equal(got, tc.want) {
// 			t.Errorf("got %v, want %v", got, tc.want)
// 		}
// 	}
// }

// //1100011000
// //0101001010
// //1001010010
// //-----------------> Index i=0, 1st op: 1
// //0000011000
// //0101001010
// //0101010010
// //-----------------> Index i=1, 1st op: 1
// //0110011000
// //0101001010
// //0011010010
// //-----------------> Index i=2, 1st op: 1
// //0101011000
// //0101001010
// //0000010010
// //-----------------> i=5, j=9, 2nd op: 2
