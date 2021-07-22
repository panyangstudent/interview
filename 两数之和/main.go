package main
import (
	"fmt"
)
func main() {
    nums := []int{ 10,7,1, 3, 4, 6}
    target := 17
    newArr := twoSum(nums, target)
    fmt.Println(newArr)
}
func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int,0)
    for key1, value1 := range nums {
       value2 := target - value1
       if value3 ,ok := numsMap[value2];ok{
            return []int{key1, value3}
       }
       numsMap[value1] = key1
    }
    return []int{}
}

// 详细解说： https://juejin.cn/post/6937313943320952869