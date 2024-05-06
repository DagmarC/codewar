package dynamicprog

// InitDPTable will create the 2D table that is usually used for dynamic programming purposes
func InitDPTable(m, n, value int) [][]int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		// Fill in the table with value if it is not the default 0
		if value != 0 {
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
	}
	return dp
}