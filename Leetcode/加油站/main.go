package main

import (
	"fmt"
)

// 思路

func canCompleteCircuit(gas []int, cost []int) int {
	for i , n := 0 , len(gas); i< n ; {
		sumOfCost , sumOfGas, cnt := 0,0,0
		for cnt < n {
			j := (i + cnt) % n 
			sumOfCost += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1 // 从第一个无法到达的加油站开始检查
		}
	}
	return -1 
}
