package main

import (
	"fmt"
)

// 贪心算法
func canJUmp(nums []int) bool {
	if len(nums) <=0 {
		return false
	}
	dp := make([]bool,0)
	dp[0] = true
	for i := 1;i<len(nums);i++ {
		for j := i-1; j>=0;j-- {
			if dp[j] && nums[j] + j >= i {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}