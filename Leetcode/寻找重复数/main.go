package main

import "sort"

func findDuplicate(nums []int) int {
    sort.Ints(nums)
    fest := nums[0]
    for i := 1 ; i< len(nums); i++ {
        if fest ^ nums[i] == 0 {
            return fest
        }
        fest = nums[i]
    }
    return fest
}