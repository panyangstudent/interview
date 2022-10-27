package main

// 贪心算法
func canJUmp(nums []int) bool {
	if len(nums) <=0 {
		return false
	}
	dp := make([]bool,len(nums))
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

// 动态规划
func canJUmpNew(nums []int) bool {
	n, m := len(nums), 0
	for i, num := range nums {
		if i <= m {
			m = max(m, i + num)
			if m > n-1 {
				return true
			} else {
				return false
			}
		}
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}