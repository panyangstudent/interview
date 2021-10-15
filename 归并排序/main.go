package main

// 归并排序 复杂度，O(nlogn)， 非原地排序，比较耗费空间
// 算法思想：如果要排序一个数组，我们先把数组从中间分成前后两个部分，
// 然后对前后两个部分分别排序，再将排序好的两部分合并在一起，这样整个数组就有序了
import (
	"fmt"
)
func main()  {
	sortArr := []int64{1,23,4,5,67,45,23,11,34,56,23,12,14}
	res := mergerSort(sortArr)
	fmt.Println(res)
}

func mergerSort(arr []int64) []int64 {
    if len(arr) <= 1 {
        return arr
    }
    m := len(arr) /2 
    left := mergerSort(arr[:m])
    right := mergerSort(arr[m:])
    return merge(left, right)
}
func merge(left []int64, right []int64) []int64{
    lindex , rindex := 0,0
    llen := len(left)
    rLen := len(right)
    res := make([]int64,0)
    for lindex <llen && rindex < rLen {
        if left[lindex] < right[rindex] {
            res = append(res, left[lindex])
            lindex++
        } else {
            res = append(res, right[rindex])
            rindex++
        }
    }
    if rindex < rLen {
        res = append(res,right[rindex:]...)
    }
    if lindex < llen {
        res = append(res,left[lindex:]...)
    }
    return res
}