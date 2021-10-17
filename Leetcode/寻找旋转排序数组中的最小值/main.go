package main

// 二分查找
func findMin(nums []int) int {

	low, hight := 0, len(nums)-1
	for low < hight {
		priv := (low + hight) / 2  
		if nums[priv] < nums[hight] {
			hight = priv
		} else {
			low = priv + 1
		}
	}
	return nums[low]
}