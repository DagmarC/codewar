package dynamicprog

// ################################################
// ##############  RECURSION  #####################
// ################################################

func countCoinsRec(coins []int, sum int) int {
	// exactly 1 solution - do not include any coin
	if sum == 0 {
		return 1
	}
	// no solution exists
	if sum < 0 {
		return 0
	}
	// no coins left and sum > 0 -> no solution exists
	if len(coins) == 0 {
		return 0
	}
	last := coins[len(coins)-1]
	// count is sum of solutions (i)
	// including last -> coins[n-1], excluding coins[n-1]
	return countCoinsRec(coins, sum-last) + countCoinsRec(coins[:len(coins)-1], sum)
}

// ################### DP #########################
// ##############  MEMOIZATION  ###################
// ################################################

// countCoinsMemoization using dynamic programming memoization to store partial results in table to prevent repeating of calculation
func countCoinsMemoization(coins []int, sum int) int {
	dp := InitDPTable(len(coins)+1, sum+1, -1)
	return countCoinsMemoRec(coins, sum, len(coins), &dp)
}

func countCoinsMemoRec(coins []int, sum int, n int, dp *[][]int) int {
	// base case
	if sum == 0 {
		(*dp)[n][sum] = 1
		return (*dp)[n][sum]
	}
	// sum is less than 0 or no coins have left
	if sum < 0 || n == 0 {
		return 0
	}
	// if the subproblem has already been calculated then simply return it
	if el := (*dp)[n][sum]; el != -1 {
		return el
	}
	// Recursion: including last -> coins[n-1], excluding coins[n-1]
	(*dp)[n][sum] = countCoinsMemoRec(coins, sum-coins[n-1], n, dp) + countCoinsMemoRec(coins, sum, n-1, dp)
	return (*dp)[n][sum]
}

// ################### DP #########################
// ##############  TABULATION  ####################
// ################### 2D #########################

func countCoinsTabulation(coins []int, sum int) int {
	n := len(coins)
	// n is the number of coin denominations and sum i the target sum
	dp := InitDPTable(n+1, sum+1, 0)

	// Base case where the target sum = 0, only 1 way to make change - do not select any coin
	dp[0][0] = 1
	for i := 1; i < n+1; i++ {
		for j := 0; j < sum+1; j++ {
			// Add the number of ways to make change without using current coin
			dp[i][j] += dp[i-1][j]

			// check if the current sum (j) after using (inserting) the current coin will still be postivie or 0
			// j represents the current sum j values will be from 1 to sum included
			if j-coins[i-1] >= 0 {
				// Add the number of ways to make change using the current coin
				dp[i][j] += dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[n][sum]
}

// ################### DP #########################
// ##############  TABULATION  ####################
// ################### 1D #########################

func countCoinsTabulation1D(coins []int, sum int) int {
	// dp[i] represents the number of ways to make change for partial sum i (sum = i)
	dp := make([]int, sum+1)

	// Base case where the target sum = 0 (do not select any coin)
	dp[0] = 1

	// Pick all coins 1 by 1 and update dp[] values until index j reaches the wanted sum
	for _, coin := range coins {
		for j := coin; j < sum+1; j++ {
			dp[j] += dp[j-coin]
		}
	}
	return dp[sum]
}
