package dynamicproblems

func rob(nums []int) int {
	dp := make(map[int]int, len(nums))

	var robRec func(int) int
	robRec = func(i int) int {
		if i < 0 {
			return 0
		}
		if i == 0 || i == 1 {
			return nums[i]
		}
		if value, ok := dp[i]; ok {
			return value
		}
		dp[i] = max(dp[i], nums[i]+robRec(i-2), nums[i]+robRec(i-3))
		return dp[i]
	}
	return max(robRec(len(nums)-1), robRec(len(nums)-2))
}

func robLinear(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]

	for i := 2; i < len(nums); i++ {
		if i==2 {
			dp[i] = nums[i] + dp[i-2]
			continue
		}
		dp[i] = max(nums[i]+dp[i-2], nums[i]+dp[i-3])
	}
	return max(dp[len(nums)-1], dp[len(nums)-2])
}
