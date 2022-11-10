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
func searchRangeN(nums []int, target int) []int {
	// 寻找开始位置：为目标值的左侧边界
	start := findBound(nums,target)
	if start == len(nums) || nums[start] != target {
		return []int{-1,-1}
	}
	end := findBound(nums, target+1) -1
	return []int{start, end}
}
func findBound(nums []int, target int) int  {
	left, right := 0, len(nums) -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid -1
		} else {
			right = mid - 1
		}
	}
	return left
}