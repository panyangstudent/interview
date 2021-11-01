package main

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1 
	iter := head 
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	add := n - k%n 
	if add == n {
		return head
	}
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	ret = iter.Next
	iter.Next = nil
	return ret
}