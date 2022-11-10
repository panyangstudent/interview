package main
import (
	"sort"
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
func mergeN(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	left, right := 0,1
	res = append(res, intervals[left])
	for right < len(intervals) {
		if res[left][1] >= intervals[right][0] {
			if res[left][1] < intervals[right][1] {
				res[left][1] = intervals[right][1]
			}
		} else {
			res = append(res, intervals[right])
			left++
		}
		right++
	}
	return res
}