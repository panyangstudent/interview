package main

import (
	"fmt"
)
// 	置换，将给定的数组恢复成下面的形式	
// 	如果数组中含有x ∈ [1,N], 那么恢复之后，数组的第x-1个元素为x
//	在恢复之后，数组应当有[1,2,3,4,5.....N]的形式，但是其中有若干位置上的数是错误的，每一个错误的位置就代表了一个缺失的正数。以[3, 4, -1, 1]
//	为例恢复后的数组应该是[1,-1,3,4]，我们就可以认为缺失的是2

//	那么我们如何将数组恢复呢？ 我们可以对数组进行遍历，对遍历到的数，x = nums[i], 如果 x ∈ [1,N]，我们就知道 x 应当出现在数组中的 x−1 的位置
//	因此交换nums[i] 和 nums[x-1]， 在交换完之后nums[i]可能还是属于[1,N]， 那我们继续交换
func firstMissingPositive(nums []int) int {
    n := len(nums)
    for i := 0; i < n; i++ {
        for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
            nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
        }
    }
    for i := 0; i < n; i++ {
        if nums[i] != i + 1 {
            return i + 1
        }
    }
    return n + 1
}