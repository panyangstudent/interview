package main 
import (
	"fmt"
)
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 进阶：你能尝试使用一趟扫描实现吗？
// 先循环到n的地方， 然后判断当前的节点是不是为nil
// 如果不为nil。 则将两个指针都向前推进，直到fest为nil
// 然后将slow的next 指向 slow的next的next上

func removeNthFromEnd(head *ListNode, n int) *ListNode {
		dummy := &ListNode{}
		dummy.Next = head
		slow, fast := dummy, dummy
		for i := 0; i < n + 1; i++ {
			fast = fast.Next;
		}
	
		for fast != nil {
			slow = slow.Next
			fast = fast.Next
		}
	
		slow.Next = slow.Next.Next;
		return dummy.Next;
}
