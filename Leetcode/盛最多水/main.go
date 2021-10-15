package  main

// 思路：两条之间形成的区域总是会受到其中较短那条的限制。此外两条线段距离越远等到的面积越大
import (
	"fmt"
)
func main() {
	height := []int{4,3,2,1,4}
	fmt.Printf("max area is ",maxArea(height))
}
func maxArea(height []int) int {
    start, end := 0,len(height) -1
	maxArea := 0 
	for start < end {
		// 面积 = 底边 * 高。指针往中间移动，底边长变小，只有短木板的高度变高了才能有最大面积
		// 因此移动短木板的指针，这样原来的长木板才可能变成移动后的短木板
		if height[start] < height[end] {
			maxArea = max(height[start] * (end - start) , maxArea)
			// 短木板的指针向中间移动
			start++
		} else {
			maxArea = max(height[end] * (end - start), maxArea)
	        // 短木板的指针向中间移动		
			end--
		}
	}
	return maxArea
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}