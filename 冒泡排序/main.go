package main
// 冒泡排序的基本思想是将数组中的每个相邻元素进行两两比较，按照小元素在前（或大元素在前）的原则确定是否进行交换。
// 这样每一轮执行之后，最大（或最小）的元素就会被交换到了最后一位。
// 完成一轮之后，我们可以再从头进行第二轮的比较，直到倒数第二位（因为最后一位已经是被排序好的了）时结束。
// 这样一轮之后，第二大（或第二小）的元素就会被交换到了倒数第二位。
// 同样的过程会依次进行，直到所有元素都被排列成预期的顺序为止
import (
	"fmt"
)
func main()  {
	sortArr := []int64{2,33,4,55,6,77,34,25,67,87,24,12,23,45,68,90}
	BuSort(sortArr)
}

func BuSort(sortArr []int64)  {
	var temp int64
	arrLen := len(sortArr)
	for i := arrLen; i > 0; i-- {
		for j := 0; j <  i-1; j++ {
			if sortArr[j] < sortArr[j+1] {
				temp =  sortArr[j]
				sortArr[j] =  sortArr[j+1]
				sortArr[j+1] =  temp
			}
		}
	}
	fmt.Println("BuSort %v",sortArr)
	return
}