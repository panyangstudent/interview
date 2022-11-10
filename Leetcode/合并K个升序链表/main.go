package main
/*
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

*/
type ListNode struct {
	Val int
	Next *ListNode 
}

// 分治+优先队列
func mergeKLists(lists []*ListNode) *ListNode {
	ln := len(lists)
	if ln == 0 {
		return nil
	}

	if ln == 1 {
		return lists[0]
	}

	return merge2list(mergeKLists(lists[:ln/2]), mergeKLists(lists[ln/2:]))
}

func merge2list(list1,list2 *ListNode) *ListNode{
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := &ListNode{}
	ans := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			ans.Next = list1
			list1 = list1.Next
		} else {
			ans.Next = list2
			list2 = list2.Next
		}
		ans = ans.Next
	}

	if list1 == nil {
		ans.Next = list2 
	}
	if list2 == nil {
		ans.Next = list1 
	}
	return head.Next
}
