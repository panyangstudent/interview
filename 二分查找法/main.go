package main

import (
	"fmt"
)

// 二分查找法通过对折的方式查找一个数据，条件必须是一个有序的数组，数组的底层是顺序链表
func main()  {
	searchArr := []int{2,4,6,12,23,24,25,33,34,45,55,60,67,68,77,87}
	res := BinarySearch(4,searchArr)
	fmt.Println("search target is  %v",res)
}
func BinarySearch(target int, searchArr []int) int {
	return 0
}