package main

// 动态规划问题解题思路：
// 1. 将问题分解成最优子问题
// 2. 用递归方式将问题表述成最优子问题的解
// 3. 自底向上的将递归转化成迭代问题
// 打家劫舍问题思路：	
// 定义状态：dp[i]:盗窃到第i个房子所能获取到的最高金额
// 状态转移方程：dp[i] = max(dp[i-2] + nums[i],dp[i-1])
// 边界条件：dp[0] = nums[0],dp[1] = max(dp[0],dp[1])
import (
	"fmt"
)
func main() {
	fmt.Println("max is : ",rob([]int{1,2,3,4,5,6}))
}
func rob(nums []int) int {
	// prevMax:记录偷盗第i间房子之前，盗取的最大金额
	// currmax ： 记录偷盗第i间房子(包括第i间房子)，偷盗到的最大金额
    prevMax, currMax := 0, 0
    for i := 0; i < len(nums); i++ {
        temp := currMax
        currMax, prevMax = max(prevMax + nums[i], currMax), temp
    }

    return currMax
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}