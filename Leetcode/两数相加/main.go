package main

/*
整体思路：
1. 优先写完正常的流程, 提前判断当前输入的链表是否是单节点&为0值的链表， 如果是单节点&为0值的，直接返回另一个参数
2. 构造返回的链表头， 和一个临时的node。 这个node当作游走的node来计算
3. 两个链表相加， 说白了就是两个node相加，那我们提前将node拿出来计算
4. 获得了当前两个节点的和，同时构造当前的node， 并且在构造一个临时node， 将当前的node指向临时的node
5. 在以上的正常的流程走完后，思考当前的for循环的退出条件， 当两个链表都是nil 并且没有进位的时候肯定会退出。所以推出条件加在合适的地方就好。
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val == 0 && l1.Next == nil {
		return l2
	}
	if l2.Val == 0 && l2.Next == nil {
		return l1
	}
	result := &ListNode{}
	// 需要处理的第一个节点
	node := result
	var (
		sum = 0
		carry = 0
	)
	for {
		// 分而治之， 先把当前需要叠加的数据拿出来
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 !=nil {
			sum += l2.Val
			l2 = l2.Next
		}
		// 获取到了所有的数据。再加上上次循环的进位
		sum += carry
		// 当前的node，赋值
		node.Val = sum % 10
		carry = sum / 10
		if l1 == nil && l2 == nil && carry == 0{
			break
		}
		// 构造下一个节点，并把node指向下一个节点
		tempNode := &ListNode{}
		node.Next = tempNode
		node = node.Next
		sum = 0
	}
	return result
}
