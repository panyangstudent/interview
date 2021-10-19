package main

// 归并排序 复杂度，O(nlogn)， 非原地排序，比较耗费空间
// 算法思想：如果要排序一个数组，我们先把数组从中间分成前后两个部分，
// 然后对前后两个部分分别排序，再将排序好的两部分合并在一起，这样整个数组就有序了
import (
	"fmt"
)
func main()  {
	sortArr := []int{1,23,4,5,67,45,23,11,34,56,23,12,14}
	res := mergeSort(sortArr)
	fmt.Println(res)
}

func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    mid := len(arr) / 2 
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])
    return merge(left, right)
}
func merge(left ,right []int) []int {
    rindex, lindex := 0,0
    llen, rlen := len(left), len(right)
    res := make([]int, 0)
    for rindex < rlen && lindex < llen {
        if left[lindex] < right[rindex] {
            res = append(res, left[lindex])
            lindex++
        } else {
            res = append(res, right[rindex])
            rindex++
        }
    }
    if lindex == llen {
        res = append(res, right[rindex:]...)
    }
    if rindex == rlen {
        res = append(res, left[lindex:]...)
    }
    return res
}