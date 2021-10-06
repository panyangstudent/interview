package main

import (
	"fmt"
)

// 构建一个头节点
// 根据该节点从前往后推进， 然后反转
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummyNode := &ListNode{Val: -1}
    dummyNode.Next = head
    pre := dummyNode
	for  i:=0; i< left -1 ;i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0;i< right-left;i++ {
        next := cur.Next
        cur.Next = next.Next
        next.Next = pre.Next
        pre.Next = next
	} 
	return dummyNode.Next
}
