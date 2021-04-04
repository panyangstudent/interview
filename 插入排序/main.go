package main
import (
	"fmt"
)
func main()  {
	sortArr := []int64{2,33,4,55,6,77,34,25,67,87,24,12,23,45,68,60}
	insert(sortArr)
}

func insert(sortArr []int64)  {
	var temp int64
	arrLen := len(sortArr)
	for i := 1; i < arrLen; i++ {
		for j := i-1 ; j >= 0 ; j-- {
			if sortArr[j] < sortArr[i]{
				temp =  sortArr[i]
				sortArr[i] =  sortArr[j]
				sortArr[j] =  temp
			}
		}
	}
	fmt.Println("insert %v",sortArr)
	return
}
