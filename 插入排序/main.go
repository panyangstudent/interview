package main
// 平均复杂度 O(n^2) 且稳定
// 设计思路： 将数组划分成两个部分，第一部分是有序的，第二部分是无序的
// 每次都从无序部分取一个元素，将这个元素插入到有序部分，保持有序部分的有序性质
// 直到无序部分为空
import (
	"fmt"
)
func main()  {
	sortArr := []int64{2,33,4,55,6,77,34,25,67,87,24,12,23,45,68,60,1,2,3}
	insert(sortArr)
}

func insert(sortArr []int64)  {
	arrLen := len(sortArr)
	for i := 0; i < arrLen-1; i++ {
		for j := i+1; j > 0; j-- {
			if sortArr[j] < sortArr[j-1] {
				sortArr[j-1], sortArr[j] = sortArr[j], sortArr[j-1]
			}
		}
	}
	fmt.Println("insert %v",sortArr)
	return
}
