package main

import (
	"fmt"
)

func main() {
	merge([]int{1,2,3,0,0,0}, 3,[]int{2,5,6},3)
}

// 实现时间复杂度为噢O(m+n)
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