package main

import "fmt"

/*
解题思路：
1. 两个正序数组，设置两个游走指针，
 */
func main() {
	findMedianSortedArrays([]int{1,2},[]int{3,4})
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		resp          = 1.0
		lenNums1      = len(nums1)
		lenNums2      = len(nums2)
		sumLen        = lenNums1 + lenNums2
		modelRemIndex = sumLen % 2
		Nums1Index    = 0
		Nums2Index    = 0
		index         = 0
		nums3         []int
	)
	for index < sumLen/2+1 {
		index++
		if Nums1Index >= lenNums1 {
			nums3 = append(nums3, nums2[Nums2Index])
			Nums2Index ++
			continue
		}
		if Nums2Index >= lenNums2 {
			nums3 = append(nums3, nums1[Nums1Index])
			Nums1Index ++
			continue
		}
		if nums1[Nums1Index] <= nums2[Nums2Index] {
			nums3 = append(nums3, nums1[Nums1Index])
			Nums1Index++
		} else {
			nums3 = append(nums3, nums2[Nums2Index])
			Nums2Index++
		}
	}
	if modelRemIndex == 0 {
		resp = float64((nums3[index-2] + nums3[index-1]))/ 2.0
	} else {
		resp = float64(nums3[index-1])
	}
	fmt.Printf("index : %v sumLen/2+1 ： %v nums3: %v resp : %v",index,sumLen/2+1,nums3, resp)

	return resp
}