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
// func mergerSort(sortArr []int) []int {
// 	len1 := len(sortArr)
// 	if len1 == 1 {
// 		return sortArr // 切割到最后只剩下一个元素
// 	}
// 	m := len1/2
// 	left := mergerSort(sortArr[:m])
// 	right := mergerSort(sortArr[m:])
// 	return merge(left, right)
// }
// //把两个有序切片合并成一个有序切片
// func merge(l []int, r []int) []int{
//     lLen := len(l)
//     rLen := len(r)
//     res := make([]int, 0)

//     lIndex,rIndex := 0,0 //两个切片的下标，插入一个数据，下标加一
//     for lIndex<lLen && rIndex<rLen {
//         if l[lIndex] > r[rIndex] {
//             res = append(res, r[rIndex])
//             rIndex++
//         }else{
//             res = append(res, l[lIndex])
//             lIndex++
//         }
//     }
//     if lIndex < lLen { //左边的还有剩余元素
//         res = append(res, l[lIndex:]...)
//     }
//     if rIndex < rLen {
//         res = append(res, r[rIndex:]...)
//     }
//     return res
// }

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
