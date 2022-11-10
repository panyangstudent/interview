package main

import (
	"container/heap"
)

// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i:=0; i< k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
 	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(x interface{}) {
    *h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func topKFrequentN(nums []int, k int) []int {
	m := make(map[int]int, 0)
	ret := make([]int, 0)
	for _, num := range nums {
		m[num]++
	}
	buckets := make([][]int, len(nums)+1)
	for num, v := range m {
		if len(buckets[v]) <=0 {
			buckets[v] = make([]int, 0)
		}
		buckets[v] = append(buckets[v], num)
	}
	for i:= len(buckets) -1; i>=0;i-- {
		if len(buckets[i]) > 0 {
			ret = append(ret, buckets[i]...)
			if len(ret) >= k {
				return ret[:k]
			}
		}
	}
	return  ret
}