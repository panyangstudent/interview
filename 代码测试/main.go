package main

import "fmt"

/*
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。
 */
func main()  {
	nums := []int{2,7,11,15}
	target := 9
	twoSum(nums, target)
	fmt.Println(twoSum(nums, target))
}
func twoSum(nums []int, target int) []int {
	keyValue := make(map[int]int, 0)
	resp := make([]int, 0)
	for i, num := range nums {
		subValue := target - num
		if _, ok := keyValue[subValue]; ok {
			return append(resp, i, keyValue[subValue])
		}
		keyValue[subValue] = i
	}
	return resp
}