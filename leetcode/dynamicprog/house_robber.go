package dynamicproblems

// The `rob` function in Go uses dynamic programming to efficiently calculate the maximum amount that
// can be robbed from houses represented by an array of integers.
func rob(nums []int) int {
	dp := make(map[int]int, len(nums))

	var robRec func(int) int
	robRec = func(i int) int {
		if i < 0 {
			return 0
		}
		if value, ok := dp[i]; ok {
			return value
		}
		dp[i] = max(nums[i]+robRec(i-2), robRec(i-1))
		return dp[i]
	}
	return robRec(len(nums) - 1)
}

// The `robLinear` function calculates the maximum amount that can be robbed from houses in a linear
// manner without adjacent houses being robbed.
func robLinear(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i+1] = max(dp[i], nums[i]+dp[i-1])
	}
	return dp[len(nums)] // last element of dp
}

// The robEffective function efficiently calculates the maximum amount that can be robbed from a list
// of non-negative integers without adjacent elements being robbed.
func robEffective(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev1 := 0
	prev2 := 0
	for _, num := range nums {
		tmp := prev1
		prev1 = max(num+prev2, prev1)
		prev2 = tmp
	}
	return prev1 // last element of dp
}
