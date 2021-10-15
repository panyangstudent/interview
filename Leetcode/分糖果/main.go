package main

import (
	"fmt"
)

func candy(ratings []int) int {
	count := 0
	mapArr := make(map[int]int)
	// 右规则
	for i:=0;i< len(ratings);i++ {
		mapArr[i] = 1
		if i - 1 >= 0 && ratings[i] > ratings[i-1] && mapArr[i] <= mapArr[i-1]{
			mapArr[i] = mapArr[i-1] + 1
		}
	}
	// 左规则
	for i:= len(ratings) -2;i>=0;i-- {
		if ratings[i] > ratings[i+1] && mapArr[i] <= mapArr[i+1] {
			mapArr[i] = mapArr[i+1] + 1
		} 
	}

	for _, Val := range  mapArr {
		count = count + Val
	}
	return count
}
