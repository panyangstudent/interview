
/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
示例 1：
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

示例 2：
输入：nums = [0,1,0,3,2,3]
输出：4

示例 3：
输入：nums = [7,7,7,7,7,7,7]
输出：1
*/

// 确定状态：dp[n] = i 表示当前位置元素的最长子序列的长度
// 状态转移方程： dp[n] = max(dp[n-1], dp[n-1] + 1) 如果当前元素比dp[n-1]代表的位置元素要大， 那么对比dp[n-1]和dp[n-1]+1的大小
package main
import (
	"fmt"
)
func main() {
	fmt.Print(lengthOfLIS([]int{123,2,3,4,5,2}))
}
func lengthOfLIS(nums []int) int {
    dp := make([]int,len(nums),len(nums))
	max := 1
	for i:=0;i<len(nums);i++ {
        dp[i] = 1 
		for j:= i-1; j>=0; j-- {
			if nums[i] > nums[j] {
				dp[i] = Funcmax(dp[i], dp[j]+1)
			}
		} 
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}
func Funcmax(num1 int, num2 int) int{
	if num1 > num2 {
		return num1
	}
	return num2
}