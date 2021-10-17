package main

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