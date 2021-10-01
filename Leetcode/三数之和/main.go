package main

// 排序 + 双指针

import (
	"fmt"
)
func main()  {
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Printf("value is : ",threeSum(nums))
}

func threeSum(nums []int) [][]int {
	// 排序
	ans := make([][]int,0)
	sort.Ints(nums)
	for i := 0; i< len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 初始化双指针
		right := len(nums) - 1
		target := -1 * nums[i]
		for second := i+1;second < len(nums);second++ {
			// 枚举不能重复
			if second >  i + 1 && nums[second] == nums[second - 1] {
				continue
			}
			// 需要保证right一直在secon的右侧
			for second < right && nums[second] + nums[right] > target {
				right-- 
			}
			// 元素重叠了，停止
			if second == right {
				break
			}
			// 判断是否相等
			if nums[second] + nums[right] == target {
				ans = append(ans, []int{nums[i], nums[second], nums[right]})
			}

		}
	}
	return ans
}
// func quickSort(arr []int, left int, right int) {
// 	if left < right {
// 		privote := arr[left]
// 		j := left 
// 		for i := left; i< right; i++ {
// 			if arr[i]<privote {
// 				j++
// 				arr[j] ,arr[i] = arr[i],arr[j]
// 			}
// 		}
// 		arr[j], arr[left] = arr[left],arr[j]
// 		quickSort(arr,left, j)
// 		quickSort(arr,j+1,right)
// 	}
// }