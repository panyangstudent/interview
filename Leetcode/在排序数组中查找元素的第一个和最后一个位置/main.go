package main


// 时间复杂度为o(log N)的查找算法为二分查找算法
func searchRange(nums []int, target int) []int {
	left, right := 0 , len(nums)-1
	for left <= right {
		if nums[left] < target  {
			left++
		} else if nums[right] > target {
			right--
		} else {
			return []int{left, right}
		}
	}
	return []int{-1,-1}
}