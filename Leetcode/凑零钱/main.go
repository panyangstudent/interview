package main

import (
	"math"
)

// 动规五部曲分析如下：
// 1. 确定dp数组以及下标含义
//		dp[j]:凑足总额为j所需钱币的最少个数为dp[j]
// 2. 确定递推公式
//		得到dp[j]（考虑coins[i]）,只有一个来源，dp[j-conis[i]](没有考虑conis[i])，凑足总额为j-conis[i]的最少个数为dp[j-conis[i]]，那么只需要加上一个conis[i]即dp[j-conis[i]]+1就是dp[j]
//		所以dp[j]要取所有dp[j-coins]+1中最小的。
//		递推公式就是：dp[j] = min(dp[j-coins[i]]+1,dp[j])
// 3. dp数组如何初始化
//		首先凑足总金额为0所需钱币个数一定为0，那么dp[0] = 0 
// 		考虑到递推公式的的特性，dp[j]必须初始化一个最大值，否则就会在min(dp[j - coins[i]] + 1, dp[j]），比较过程中被初始值覆盖。
// 4. 确定遍历顺序
//		本题求铅笔的最小个数，那么钱币有顺序和没有顺序都不会影响钱币的最小个数。
//		所以本题强调的不是集合的排列或组合
//		* 如果求组合数就是外层for循环遍历物品，内层for遍历背包。
//		* 如果求排列数就是外层for遍历背包，内层for循环遍历物品。


// 版本一, 先遍历物品,再遍历背包
func coinChange(coins []int, amount int) int {
	dp := make([]int,amount + 1)
	// 初始化dp数组
	dp[0] = 0
	// 初始化为math.MaxInt32
	for i:=1; i<= amount;i++{
		dp[i] = math.MaxInt32
	}
	//遍历物品
	for i:= 0;i<len(coins);i++ {
		// 遍历背包
		for j := coins[i]; j<=amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			} 
		} 
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 版本二，先遍历背包，在遍历物品
func coinChangev2(conis []int, amount int) int {
	dp := make([]int,amount+1)
	dp[0] = 0 
	for j:=1;j <= amount; j++ {
		dp[j] = math.MaxInt32
		for i:=0; i< len(conis); i++ {
			if j >=conis[i] && dp[j-conis[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = min(dp[j], dp[j-conis[i]] + 1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}