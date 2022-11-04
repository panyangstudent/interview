package main

import "sort"

/*
给定一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，请你判断一个人是否能够参加这里面的全部会议。

统计当前每个会议时间的开始时间和结束时间，并且进行排序，如果最小的开始时间小于最小的结束时间，则将会议室数量+1，如果最小的开始时间大于最小的结束时间则将会议室的数量-1

 */


func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i:=1; i<len(intervals); i++{
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}