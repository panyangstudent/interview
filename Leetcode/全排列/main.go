package main 


func permute(nums []int) [][]int {
	if len(nums) == 1 {
        return [][]int{nums}
    }

    res := [][]int{}

    for i, num := range nums {
        // 把num从 nums 拿出去 得到tmp
        tmp := make([]int, len(nums)-1)
        copy(tmp[0:], nums[0:i])
        copy(tmp[i:], nums[i+1:])

        // sub 是把num 拿出去后，数组中剩余数据的全排列
        sub := permute(tmp)
        for _, s := range sub {
            res = append(res, append(s, num))
        }
    }
    return res
}

/*
全排列一个不含重复数字的数组
 */
func permute1(nums []int) [][]int {
    if len(nums) == 1 {
        return  [][]int{nums}
    }
    res := make([][]int, 0)
    for i, num := range nums {
        temp := make([]int, len(nums)-1)
        // 剪枝，将该位置的元素去除
        copy(nums[0:i], temp)
        copy(nums[i+1:],temp)
        // 剩余元素进行全排列
        sub := permute1(temp)
        // 循环添加当前元素，到每个已全排列的数组中
        for _, ints := range sub {
            res = append(res,append(ints, num))
        }
    }
    return res
}