package main



// 动态规划
// dp[i][j] 表示到达i,j位置的最小和，dp[i][j] = min(dp[i-1][j], dp[i][j-1])
// 初始 dp[0][0] = nums[0][0] 
func minPathSum(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = grid[0][0]
     for i := 1; i < n; i++ {
        dp[i][0] = dp[i - 1][0] + grid[i][0]
    }
    for j := 1; j < m; j++ {
        dp[0][j] = dp[0][j - 1] + grid[0][j]
    }
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
            dp[i][j] = min(dp[i-1][j], dp[i][j-1])
			dp[i][j] += grid[i][j]
		}
	}
	return dp[n-1][m-1]
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b 
}