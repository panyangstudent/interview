package main

// 动态规划类题目
// 不过这里借助了贪心算法的思想，一次遍历找出最小的数值（最小的数值有可能不更新）
// 然后利用当前的数值减去最小的数值，对比差距与上一次的差距，谁大谁小，如果大的直接替换
// 最终计算出对应的差距

import (
	"fmt"
)

func  main() {
	nums := []int{7,1,5,3,6,4}
	fmt.Printf("max prices is ",maxProfix(nums))
}


func maxProfix(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min , gap := prices[0],0
	for i := 1; i<len(prices);i++ {
		if prices[i] < min {
			min = prices[i]
		}	
		temp := prices[i] - min 
		if temp > gap {
			gap = temp
		}
	}
	return gap
}