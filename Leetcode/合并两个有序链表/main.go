package main

type ListNode struct {
	Val int
	Next *ListNode
}
// 递归思想
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}
	var res *ListNode//定义结果节点
	if l1.Val < l2.value {
		res = l1 
		res.Next = mergeTwoLists(l1.Next, l2)
	} else {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	}
	return res
}
