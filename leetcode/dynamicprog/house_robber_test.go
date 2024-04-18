package dynamicproblems

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

// Example 1:

// Input: nums = [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
// Total amount you can rob = 1 + 3 = 4.
// Example 2:

// Input: nums = [2,7,9,3,1]
// Output: 12
// Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
// Total amount you can rob = 2 + 9 + 1 = 12.

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRob(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{input: []int{1, 2, 3, 1}, want: 4},
		{input: []int{2, 7, 9, 3, 1}, want: 12},
		{input: []int{2, 1, 1, 2}, want: 4},
		{input: []int{226, 174, 214, 16, 218, 48, 153, 131, 128, 17, 157, 142, 88, 43, 37, 157, 43, 221, 191, 68, 206, 23, 225, 82, 54, 118, 111, 46, 80, 49, 245, 63, 25, 194, 72, 80, 143, 55, 209, 18, 55, 122, 65, 66, 177, 101, 63, 201, 172, 130, 103, 225, 142, 46, 86, 185, 62, 138, 212, 192, 125, 77, 223, 188, 99, 228, 90, 25, 193, 211, 84, 239, 119, 234, 85, 83, 123, 120, 131, 203, 219, 10, 82, 35, 120, 180, 249, 106, 37, 169, 225, 54, 103, 55, 166, 124}, want: 7102},
	}
	for _, tc := range tests {
		var got int
		if got = robEffective(tc.input); got != tc.want {
			t.Errorf("House Robber: got %d want %d", got, tc.want)
		}
	}
}

// Benchmark function for rob function
func BenchmarkRob(b *testing.B) {
	// Initialize random number generator with a fixed seed
	// Initialize random number generator
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)

	// Generate random test cases
	for i := 0; i < b.N; i++ {
		n := r.Intn(1000) + 1 // Generate a random number between 1 and 1000 for array length
		nums := generateRandomArray(n)
		b.Run(fmt.Sprintf("Length: %d", n), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				rob(nums)
			}
		})
	}
}

// Benchmark function for robLinear function
func BenchmarkRobLinear(b *testing.B) {
	// Initialize random number generator with a fixed seed
	// Initialize random number generator
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)

	// Generate random test cases
	for i := 0; i < b.N; i++ {
		n := r.Intn(1000) + 1 // Generate a random number between 1 and 1000 for array length
		nums := generateRandomArray(n)
		b.Run(fmt.Sprintf("Length: %d", n), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				robLinear(nums)
			}
		})
	}
}

// Helper function to generate a random array of integers
func generateRandomArray(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(1000) // Generate random integers between 0 and 999
	}
	return arr
}

