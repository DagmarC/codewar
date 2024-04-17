package dynamicproblems

import (
	"math"
)

// The `coinChange` function in Go calculates the minimum number of coins needed to make up a given sum
// using a dynamic programming approach.
func coinChange(coins []int, sum int) int {
	dp := make([]int, sum+1)
	for i := 1; i <= sum; i++ {
		dp[i] = math.MaxInt32
	}
	for _, coin := range coins {
		for currSum := 1; currSum < sum+1; currSum++ {
			if currSum >= coin {
				dp[currSum] = min(dp[currSum], 1+dp[currSum-coin])
			}
		}
	}
	if dp[sum] == math.MaxInt32 {
		return -1
	}
	return dp[sum]
}
