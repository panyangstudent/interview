package main


// 动态规划
// 状态定义： dp[i][0] : 表示第i天交易完后手里没有股票的最大利润，dp[i][1]表示第i天交易完后手里持有股票的最大利润
// 	虑到dp[i][0] 如果这一天交易完后手里没有股票，那么可能的转移状态为前一天已经没有股票，即 dp[i−1][0]，
//		或者前一天结束的时候手里持有一支股票，即 dp[i−1][1]，这时候我们要将其卖出，并获得 prices[i] 的收益。
//		状态转移方程为：p[i][0]=max{dp[i−1][0],dp[i−1][1]+prices[i]}
//	再来考虑 dp[i][1]，按照同样的方式考虑转移状态，那么可能的转移状态为前一天已经持有一支股票，即 
//		dp[i−1][1]，或者前一天结束时还没有股票，即 dp[i−1][0]，这时候我们要将其买入，并减少 prices[i] 的收益
//		状态转移方程为：dp[i][1]=max{dp[i−1][1],dp[i−1][0]−prices[i]}
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i :=1 ; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][0] - prices[i], dp[i-1][1])
	}
	return dp[n-1][0]
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}