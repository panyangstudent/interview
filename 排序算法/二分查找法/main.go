package main

import (
	"fmt"
)

// 二分查找法通过对折的方式查找一个数据，条件必须是一个有序的数组，数组的底层是顺序链表
func main()  {
	searchArr := []int{1, 3}
	res := BinarySearch(searchArr,0, len(searchArr)-1,3)
	fmt.Printf("找到了，下标为：%v \n", res)
}
func BinarySearch(arr []int, leftIndex int, rightIndex int, findVal int) int {
	res := -1
	// 判断 leftindex是否大于rightindex
	if leftIndex > rightIndex {
		fmt.Printf("没找到")
		return res
	}
	// 先找打中间节点
	middle := (leftIndex + rightIndex) / 2 
	if arr[middle] > findVal {
		// 要查找的数，范围应该在leftindex 和 middle之间
		res = BinarySearch(arr, leftIndex, middle-1, findVal)
	} else if arr[middle] < findVal {
		res = BinarySearch(arr, middle+1, rightIndex, findVal)
	} else {
		res = middle
	}
	return res
}