package main


// 追赶法-遍历追赶-使用双指针游走
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast.Next == nil {
		return slow
	}
	return slow.Next
}