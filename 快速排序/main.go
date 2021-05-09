package main
//快速排序 ,原地排序 ，平均复杂度 O(nlogn),最差为）O(n^2) 不稳定
// 算法思想： 通过分治法的思想，从数组中选取一个基准元素pivot，
// 把这个元素中小于pivot的移动到左边，把大于povit的移动到右边，
// 然后在分别对左右两边的数组进行排序
import (
	"fmt"
)

func main()  {
	sortArr := []int{1,23,4,5,67,45,23,11,34,56,23,12,14}
	QuickSort(&sortArr,0,len(sortArr)-1)
	fmt.Println("result is %v",sortArr)

}
// func quickSort(sortArr *[]int, l int ,h int)  {
// 	if l < h {
// 		position := partition(sortArr,l ,h)
// 		quickSort(sortArr,l, position-1)
// 		quickSort(sortArr,position+1, h)
// 	}
// }
// func partition(sortArr *[]int, l int ,h int)(left int){
// 	if l > h {
// 		return 0
// 	}
// 	fmt.Println("partition is %v ，%v", l, h)
// 	value := (*sortArr)[l]
// 	l++
// 	for l < h {
// 		if (*sortArr)[l] > value {
// 			(*sortArr)[l], (*sortArr)[h] = (*sortArr)[h], (*sortArr)[l]
// 			h--
// 		}

// 		if (*sortArr)[h] < value {
// 			(*sortArr)[h], (*sortArr)[l] = (*sortArr)[l], (*sortArr)[h]
// 			l++
// 		}
// 	}
// 	(*sortArr)[l] = value
// 	return l
// }

//划分
func partition(arr *[]int,left int,right int)int{
	privot:=(*arr)[right]
	i:=left-1
	for j:=left;j<right;j++{
		if (*arr)[j]<privot{
			i++
			(*arr)[i],(*arr)[j] = (*arr)[j],(*arr)[i]
		}
	}
	(*arr)[i+1],(*arr)[right] = (*arr)[right],(*arr)[i+1]
	return i+1
}
//递归
func QuickSort(arr *[]int,left int,right int){
	if left>= right{
		return
	}
	privot:=partition(arr,left,right)
	QuickSort(arr,left,privot-1)
	QuickSort(arr,privot+1,right)
}