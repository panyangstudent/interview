package main

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

/*
传统递归方法
原始链表
node1-> node2 -> node3 -> node4 ->node5 -> node6-> nil

我们假设一部分链表已经反转

node1 -> node2 -> node3 <- node4 <-node5 <- node6

此时我们位于node3，我们需要让node4指向node3，所以会有
node3->next->next = node3
node3->next = nil
 */

func recursionReverseList(head *ListNode) *ListNode {
	if head.Next == nil || head == nil {
		return head
	}
	newHead := recursionReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
    return newHead
}