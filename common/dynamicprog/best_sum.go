package dynamicprog

func bestSum(target int, arr []int, dp map[int][]int) []int {
	if dp == nil {
		dp = make(map[int][]int, 0)
	}
	if _, ok := dp[target]; ok {
		return dp[target]
	}

	if target == 0 {
		return  []int{}
	}
	if target < 0 {
		return nil
	}

	for _, n := range arr {
		partialResult := bestSum(target - n, arr, dp)
		if partialResult != nil {
			partialResult = append(partialResult, n)
			if _, ok := dp[target]; !ok || len(dp[target]) > len(partialResult) {
				dp[target] = partialResult
			}
		}
	}
	return dp[target]
}