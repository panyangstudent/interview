package main


// 思路

func canCompleteCircuit(gas []int, cost []int) int {
	for i , n := 0 , len(gas); i< n ; {
		sumOfCost , sumOfGas, cnt := 0,0,0
		for cnt < n {
			j := (i + cnt) % n
			sumOfGas += gas[j]
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

func canCompleteCircuitN(gas []int, cost []int) int {
	for i, n:=0, len(gas); i < n; {
		sumoOfCast, sumOfGas, cut := 0,0,0
		for cut < n {
			j := (i+cut) % n
			sumOfGas += gas[j]
			sumoOfCast += cost[j]
			if sumoOfCast > sumOfGas {
				break
			}
			cut++
		}
		if cut == n {
			return i
		} else  {
			i += cut+1 // 从第一个无法到达的加油站开始检查
		}
	}
	return -1
}