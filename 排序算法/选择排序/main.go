package main
//选择排序,平均复杂度 O(n^2) 且稳定
// 算法思想： 选择排序提高了冒泡排序的性能，它每遍历一遍列表只交换一次数据，
// 即遍历时找到最大的项，和冒泡排序一样，第一次遍历后，最大的数据项就已经归位，
// 第二次遍历使次最大项归位，这个过程持续进行，一共需要n-1次遍历来排好n个数据，
// 因为最后一个数据必须在第n-1n-1次遍历之后才能归位


import (
	"fmt"
)
func main()  {
	sortArr := []int64{1,23,4,5,67,45,23,11,34,56,23,12,14}
	SelectSort(sortArr)
}

func SelectSort(sortArr []int64) {
	arrlen := len(sortArr)
	for i := 0; i < arrlen; i++ {
		minIndex := i
		for j := i+1 ; j < arrlen; j++ {
			if sortArr[minIndex] > sortArr[j] {
				minIndex = j
			}
		}
		sortArr[minIndex], sortArr[i] = sortArr[i], sortArr[minIndex]
	}
	fmt.Println("SelectSort result is : %v",sortArr)
	return
}
