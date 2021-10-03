package main


import (
	"fmt"
)

// 解题思路：a+b=b+a
// for 内循环在两个 链表（a+b和b+a）都结束后就自动结束了。
// 如果链表有相交，那么会在中途相等，返回相交节点；
// 如果链表不相交，那么最后会 nil == nil，返回 nil；
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a,b := headA,headB
	for a!=b {
		if  a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}	
// 时间复杂度为O(n), 空间复杂度为O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	
}	