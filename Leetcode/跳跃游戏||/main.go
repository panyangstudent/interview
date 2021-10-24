package main

import (
	"fmt"
)

func jump(nums []int) int {
	steps, maxPos, end := 0, 0 ,0 
	for i := 0; i< len(nums); i++ {
		maxPos = max(maxPos, i + nums[i])
		if i == end {
			end = maxPos
			steps++
		}
	}
	return steps
}
func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}