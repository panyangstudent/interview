package main

import (
	"fmt"
)

// 思路： 
//	动态规划： 我们用f(i,j)表示从左上角走到(i,j)的路径数量，其中i和j的范围是[0,m)和[0,n）。由于我们每一步只能从向下或者向右移动一步，所以
//		我们只能从（i-1， j）， 或者（i， j-1）走过过来。因此动态转移方程是：dp[i][j] = dp[i-1][j] + dp[i][j-1]， 初始状态为dp[0][0] = 1
// 		为了方便我们将矩形的边和宽，也就是dp[0][j] 和 dp[i][0]初始化成 1， 这样在计算dp[0][j]时就是前面的dp[0][j-1]的累加
func uniquePaths(m, n int) int {
	dp := make([][]int, m) 
	for i := 0 ;i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	} 
	for j := 0;j < n ;j++ {
		dp[0][j] = 1
	}
	for i := 1; i< m; i++  {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		} 
	}
	return dp[m-1][n-1]
}
func ()  {
	
}