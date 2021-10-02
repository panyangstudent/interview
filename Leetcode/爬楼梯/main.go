package main

import (
	"fmt"
)
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 注意：给定 n 是一个正整数。

// base case n = 1 n =0 => dp[0] = 0,dp[1] = 1 dp[2] = dp[1] + 1
// dp[i] = dp[i-1] + dp[i-2] 因为dp[i]这一位置，可以上一步到达，也可以上两步到达
func climbStairs(n int) int {
	left, right,cur := 0,0,1
	for i := 1; i <= n ;i++ {
        left = right
        right = cur
		cur = left + right 
	}
	return cur
}