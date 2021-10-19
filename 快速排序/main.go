package main
//快速排序 ,原地排序 ，平均复杂度 O(nlogn),最差为O(n^2) 不稳定
// 算法思想： 通过分治法的思想，从数组中选取一个基准元素pivot，
// 把这个元素中小于pivot的移动到左边，把大于povit的移动到右边，
// 然后在分别对左右两边的数组进行排序
import (
	"fmt"
	"strings"
)

func main()  {
	sortArr := []int64{1,23,4,5,67,45,23,11,34,56,23,12,14}
	quickSort(sortArr,0,int64(len(sortArr)))
	fmt.Println("result is %v",sortArr)
	fmt.Println(strings.Replace("https:/io","/","\\\\/",-1))
}
// func quickSort(arr []int64, left int64, right int64) {
// 	if left < right {
// 		privote := arr[left]
// 		j := left 
// 		for i := left; i< right; i++ {
// 			if arr[i] < privote {
// 				j++
// 				arr[j] ,arr[i] = arr[i],arr[j]
// 			}
// 		}
// 		arr[j], arr[left] = arr[left],arr[j]
// 		quickSort(arr,left, j)
// 		quickSort(arr,j+1,right)
// 	}
// }


func quickSort(arr []int64, left int64, right int64){
	if left < right {
		privote := arr[left]
		j := left 
		for i := left;i<right;i++ {
			if arr[i] < privote {
				j++
				arr[i],arr[j] = arr[j], arr[i]
			}
		}
		arr[left], arr[j] = arr[j], arr[left]
		quickSort(arr, left, j)
		quickSort(arr, j+1, right)
	}
}

