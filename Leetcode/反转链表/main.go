package main

import (
	"fmt"
)
type ListNode struct {
	Val int
	Next *ListNode
}
func  main() {
	
}

// 传统迭代方法
func reverseList(head *ListNode) *ListNode {
	res := &ListNode{}
	for {
		if head == nil {
            break
		}
        tempNode := head 
		head = tempNode.Next
		tempNode.Next = res.Next
		res.Next = tempNode
	}
	return res.Next
}

// 传统递归方法
func recursionReverseList(head *ListNode) *ListNode {
	if head.Next == nil || head == nil {
		return head
	}
	newHead := recursionReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
    return newHead
}