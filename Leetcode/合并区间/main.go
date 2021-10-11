package main
import (
	"fmt"
)

// 思路：先排序
// 之后在取两个指针left， right， 这两个指针left指向新的返回，right指向老的数组结构。之后在判断，合并
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool{
		return intervals[a][0] < intervals[b][0]
	})
	length := len(intervals)
	right, left := 1,0
	newIntervals := make([][]int,0)
	newIntervals = append(newIntervals, intervals[0])
	for right < length {
		if newIntervals[left][1] >= intervals[right][0] {
			if newIntervals[left][1] < intervals[right][1] {
				newIntervals[left][1] = intervals[right][1]
			}
		} else {
			newIntervals = append(newIntervals,intervals[right])
			left++ 
		}
		right++
	}
	return newIntervals
}
