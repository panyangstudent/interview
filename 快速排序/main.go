package main
//快速排序 ,原地排序 ，平均复杂度 O(nlogn),最差为O(n^2) 不稳定
// 算法思想： 通过分治法的思想，从数组中选取一个基准元素pivot，
// 把这个元素中小于pivot的移动到左边，把大于povit的移动到右边，
// 然后在分别对左右两边的数组进行排序
import (
	"fmt"
)

func main()  {
	sortArr := []int{1,23,4,5,67,45,23,11,34,56,23,12,14}
	headSort(sortArr)
	fmt.Println("result is %v",sortArr)

}
// func quickSort(arr []int, left, right int)  {
// 	if left < right{
// 	   pivot := arr[left]
// 	   j := left
// 	   for i := left; i < right; i++ {
// 		  if arr[i] < pivot {
// 			 j++
// 			 arr[j], arr[i] = arr[i], arr[j]
// 		  }
// 	   }
// 	   arr[left], arr[j] = arr[j], arr[left]
// 	   quickSort(arr, left, j)
// 	   quickSort(arr, j+1, right)
// 	}
//  }



//  func quickSort(arr []int64, left int64, right int64) {
// 	if left < right {
// 		privot := arr[left]
// 		j := left 
// 		for i := left; i<right ; i++ {
// 			if arr[i] < privot {
// 				j++
// 				arr[j],arr[i] = arr[i], arr[j]
// 			}
// 		}
// 		arr[left],arr[j] = arr[j], arr[left]
// 		quickSort(arr,left, j)
// 		quickSort(arr, j+1,right)
// 	}
// } 



func quickSort(num []int64, left int64, right int64) {
	if left < right {
		privot := num[left] 
		j := left 
		for i := left ; i< right ; i++ {
			if num[i] < privot {
				j++ 
				num[j], num[i] = num[i], num[j]
			}
		}
		num[j], num[left] = num[left], num[j]
		quickSort(num, j+1, right)
		quickSort(num, left, j)
	}
}

func headSort (num []int) {
	// 先从第一个非叶子结点构造堆
	for i := len(num)/2 - 1; i >= 0 ; i-- {
		adjestHeap(num, i, len(num))
	} 
	// 再次调整堆
	for i:= len(num) - 1 ; i >= 0 ; i-- {
		num[0], num[i] = num[i],num[0]
		adjestHeap(num, 0, i)
	}
}

func adjestHeap(num []int, pos int,length int) {
	for {
		// 计算孩子节点
		child := pos * 2 + 1
		if child >= length-1 {
			return
		}

		// 孩子节点最大的
		if num[child + 1] > num[child] {
			child ++ 
		}

		if num[pos] < num[child] {
			num[pos], num[child] = num[child], num[pos]
			pos = child
		} else {
			break
		}
	}
}