package main

import (
	"fmt"
)
// 递归
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}
// 迭代
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy

	for head != nil && head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = head
		prev.Next = next

		prev = head
		head = head.Next
	}
	return dummy.Next
}
