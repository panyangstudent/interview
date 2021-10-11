package main

import (
	"fmt"
)

// 快慢指针
func hasCycle(head *ListNode) bool {
	if  head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// 哈希表

func hasCycle(head *ListNode) bool {
	if  head == nil || head.Next == nil {
		return false
	}
	tempMap := make(map[*ListNode]bool)
	for head != nil [
		if _,ok := tempMap[head];ok {
			return true
		}
		tempMap[head] = true
		head = head.Next
	]
	return false
}