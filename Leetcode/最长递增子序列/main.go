// 问题： 给定一个无序的整数数组，找到其中最长的上升子序列的长度	
// 输入：[10,9,2,5,3,7,101,18]
// 解释：最长的上升子序列是[2，3，7，101]，长度为4

// 思路：我们设计动态规划算法，需要一个dp数组，我们可以假设dp[0....i-1]，都已经被算出来了
// ，然后问自己：怎么通过这些结果算出dp[i]

// 我们定义dp数组的含义，dp[i]表示以nums[i]这个数字结尾的最长递增子序列的长度

/**
int res = 0 
for (i := 0; i<len(dp);i++){
	res = math.max(res, dp[i])
}
return res
**/

package main

import (
	"fmt"
)
func main() {
	nums := []int{10,9,2,5,3,7,101,18}
	fmt.Println("value is :", 	maxChildStr(nums))
	return
}

func maxChildStr(nums []int) int {
	dp := make([]int, len(nums),len(nums))
	max := 1
	 for i := 0; i < len(nums); i++ {
		 dp[i] = 1
		for j := i-1; j>=0 ; j -- {
			if nums[i] > nums[j] {
				dp[i] = funcMax(dp[j] + 1, dp[i])
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	 } 
	 return max
}


func funcMax (int1 int, int2 int) int {
	if int1 > int2 {
		return int1
	}
	return int2
}

func maxChildStr(nums []int) {
	dp := make([]int, len(nums))
	max := 1

	for i := 0; i< len(nums);i++ {
		dp[i] = 1 
		for j = i-1; j >= 0;j-- {
			if num[i] > num[j] {
				dp[i] = findMax(dp[j]+1,dp[i])
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}