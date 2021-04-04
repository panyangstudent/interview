package main
// 选择排序的思想
// 两次循环，第一层循环固定一个位置上的元素，用该元素和后面其他位置上的元素比较大小，
// 如果大于/或者小于，就交换该位置上的元素， 然后继续用该元素和后续的元素进行比较。
// 最后生成一个升序或者降序的数组
import (
	"fmt"
)
func main()  {
	sortArr := []int64{2,33,4,55,6,77,34,25,67,87,24,12,23,45,68,90}
	selectSort(sortArr)
}

func selectSort(sortArr []int64,)  {
	var temp int64
	sortLen := len(sortArr)
	for j := 0; j< sortLen; j++ {
		for i := j + 1; i< sortLen; i++ {
			if  sortArr[j] < sortArr[i] {
				temp =  sortArr[i]
				sortArr[i] =  sortArr[j]
				sortArr[j] =  temp
			}
		}
	}
	fmt.Println("sortArr : %v",sortArr)  
	return
} 