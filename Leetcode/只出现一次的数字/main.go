package main

// 位运算
func singleNumber(nums []int) int {
	flag := 0
	for _, num := range nums {
		flag ^= num
	}
	return flag
}