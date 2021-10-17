package main

import (
	"fmt"
)

//给定一个会议时间安排的数组，每个会议时间都会包括开始和结束的时间 [[s1,e1],[s2,e2],...] (si < ei)，
// 为避免会议冲突，同时要考虑充分利用会议室资源，请你计算至少需要多少间会议室，才能满足这些会议安排。
// 示例 1:
// 输入: [[0, 30],[5, 10],[15, 20]]
// 输出: 2
// 示例 2:

// 输入: [[7,10],[2,4]]
// 输出: 1


// 思路先将会议时间按照开始时间进行排序
// 用最小堆存放会议的结束时间 
// 遍历会议curInterval：
//		1. 如果curInterval开始时间大于最小堆的堆顶， 即已申请的会议室最小结束时间， 则不需要申请会议室，将堆顶替换为当前遍历的会议结束时间 重新堆排序
// 		2. 如果curInterval开始时间小于最小堆的堆顶即 已申请会议室的最小结束时间，则需要申请一个新的会议室，将当前会议的结束时间加入堆中 重新排序

func minMeetingRooms(intervals [][]int) int {
	// 按照会议的开始时间来排序
	sort.Slice(intervals , func(i, j int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	// 堆用来存放已申请会议室的会议结束时间
	heap := make([]int, 0)
	for _, interval := range intervals {
		if len(heap) == 0 {
			heap = append(heap, interval[1])
		} else {
			if interval[0]  >= heap[0] { // 如果要开的会议的开始时间大于已申请会议室的结束时间， 不需要申请会议室， 并替换堆顶
				heap[0]  = interval[1]
			} else { // 如果要开的会议的开始时间小于已申请会议室的最小结束时间 需要申请会议室， 并将会议结束时间加入堆
				heap = append(heap, interval[1])
			}
			minMeetingRoomsBuildHeap(heap)
		}
	}
	return len(heap) 
}
// 重建堆排序数组
func minMeetingRoomsBuildHeap(nums []int) {
	size := len(nums) 
	for i := size/2 - 1; i>=0 ; i-- {
		minMeetingRoomsHeapBody(nums, size, i)
	} 

	for i := size -1 ; i > 0 ; i-- {
		if nums[i] > nums[0]{
			continue
		}
		nums[i] , nums[0] = nums[0], nums[i]
		minMeetingRoomsHeapBody(nums, 0, i)
	} 

}

func minMeetingRoomsHeapBody(nums []int, size, curRoot int) {
	for {
		child := curRoot * 2 + 1
		if child >= size-1 {
			return
		}

		if nums[child] < nums[child + 1] {
			child++
		}

		if nums[child] > nums[curRoot] {
			nums[child] , nums[curRoot] = nums[curRoot], nums[child]
			curRoot = child
		} else {
			break
		}
	}
}