package main

import (
	"fmt"
)

func main() {
	merge([]int{1,2,3,0,0,0}, 3,[]int{2,5,6},3)
}

// 实现时间复杂度为O(m+n)
func merge(nums1 []int, m int, nums2 []int, n int)  {
	newnums := make([]int,0)
	left ,right := 0,0
	for i := 0; i< m+n ;i++ {
		if left >= m {
			newnums = append(newnums,nums2[right])
			right++  
			continue
		}
		if right >= n {
			newnums = append(newnums,nums1[left])
			left++  
			continue
		} 
		if nums1[left] < nums2[right] {
			newnums = append(newnums,nums1[left])
			left++
		} else {
			newnums = append(newnums,nums2[right])
			right++
		}
	}
	copy(nums1, newnums)
	fmt.Printf("merge value is ", nums1)
}

func mergeN(nums1 []int, m int, nums2 []int, n int)  {
	newNums := make([]int, m+n)
	left1, left2 , num , i := 0,0,0, 0
	for left1 < m || left2 < n {
		if left1 < m {
			num = nums1[left1]
		}
		if left2 < n {
			num = min(num, nums2[left2])
		}
		newNums[i] = num
		i++
	}
	copy(nums1, newNums)
	return
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}