package main

// 思想： 单调栈的思想，如果当前元素比栈顶的元素小，则入栈，如果大的话，则栈顶元素出栈，同时依次和栈顶元素进行对比
// 1. 当栈为空或新加入的元素小于栈顶元素（不会破坏递减的单调性）时，将元素入栈，作为左墙
// 2. 遇到破坏递增性的元素，把栈顶元素拿出来作为“底”（决定雨水的深度）（如果此时已经栈空了，则已经完成了匹配）。在
// 对比新的栈顶元素和新元素的高低l，取小的那个，与底求差，并与栈顶元素与新元素的坐标差-1（表示中间的空格数）相乘。加入到总雨水
// 3. 遍历数组，执行1，2步

import (
	"fmt"
)
func main() {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Printf("res is ",trap(height))
}

func trap(height []int) int {
	s := make([]int,0)
	var ret int
	for i, h :=range height {
		for len(s) > 0 && height[s[len(s)-1]] < h {
			top :=  s[len(s)-1]
			s = s[:len(s)-1]
			if len(s) == 0 {
				break
			}
			
			t :=  s[len(s)-1]
			W := i-t-1 
			H := min(height[t], h) - height[top]
			ret += W * H
		} 
		s = append(s, i)
	}
	return ret
}
func min(a, b int) int {
	if a < b {
		return a
	} 
	return b
}