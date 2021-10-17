package main
// 左右乘法表
func productExceptSelf(nums []int) []int {
    dp1 := make([]int, len(nums))
    dp2 := make([]int, len(nums))
    dp3 := make([]int, len(nums))
    dp1[0] = 1
    dp1[1] = nums[0]
    dp2[len(nums)-1] = 1
    dp2[len(nums)-2] = nums[len(nums)-1]
    for i := 2; i< len(nums);i++ {
        dp1[i] = dp1[i-1] * nums[i-1]
    }
    for j := len(nums)-3; j >= 0; j-- {
        dp2[j] = dp2[j + 1] * nums[j + 1]
    }
    for i := 0 ;i < len(nums); i++ {
        dp3[i] = dp1[i] * dp2[i]
    }
    return dp3
}