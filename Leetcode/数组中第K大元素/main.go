package main
 import (
	"fmt"
 )
// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// 堆+分治

func main(){
	num := []int{3,2,1,5,6,4}
	fmt.Printf("findKthLargest nums Is %v  value Is %v",num,	findKthLargest(num,1))
}

func findKthLargest(nums []int, k int) int {
	// 构建
	for i := len(nums)/2-1 ;i>=0 ;i-- {
		heapSort(nums, i, len(nums))
	}
	// 重建
	for i := len(nums)-1 ; i>= 0 ; i-- {
		if nums[i] > nums[0] {
			continue
		}

		nums[0] ,nums[i] = nums[i], nums[0]
		heapSort(nums, 0, i)
	}
	return nums[len(nums)-k]
}


func heapSort(nums []int,pos int, length int) {
	for {
		child := pos * 2 + 1
		if child >= length - 1 {
			break
		}
		if nums[child] < nums[child +1] {
			child ++ 
		}

		if nums[pos] < nums[child] {
			nums[pos], nums[child] = nums[child], nums[pos]
			pos = child
		} else {
			break
		}
 	}

}
