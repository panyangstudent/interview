package main
// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 输入: [-2,1,-3,4,-1,2,1,-5,4],
// 输出: 6
// 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

import (
	"fmt"
)

func main() {
	//	贪心算法
	arr := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println("value is ： ", maxSubArrayKMP(arr))

	//	动态规划	
}

// 贪心算法
func maxSubArr(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, currentSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}
		if maxSum < currentSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

// 动态规划
// 假设sum[i]为第i个元素结尾且和最大的连续子数组
// 假设对于元素i，所有以它前面的元素结尾的子数组的长度都已经求得
// 那么以第i个元素结尾且和最大的连续子数组实际上，要么是以第i-1个元素结尾且和最大的连续子数组加上这个元素，要么是只包含第i个元素
// 即sum[i] =  max(sum[j:i-1] + a[i],a[i])
// 可以通过判断sum[i-1]+a[i]是否大于a[i]来做选择，而这实际上等价于判断sum[i-1]是否大于0。
// 由于每次运算只需要前一次的结果，因此并不需要像普通的动态规划那样保留之前所有的计算结果，只需要保留上一次的即可，因此算法的时间和空间复杂度都很小
func maxSubArrayKMP(nums []int) int{
	max_sum := nums[0]
    for i := 1; i < len(nums); i++ {
		if nums[i - 1] > 0 {
			nums[i] += nums[i - 1]
		} 
		if nums[i] > max_sum {
			max_sum = nums[i]
		}    
	}
	return max_sum
}

func maxSubArr() {
	
}