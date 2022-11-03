package main


type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{Val: 0, Next: nil}
    if l1 == nil {
        return l2
    }

    if l2== nil {
        return l1
    }
	l1Arr := make([]int, 0)
	l2Arr := make([]int, 0)
	for l1 != nil {
		l1Arr = append(l1Arr, l1.Val)
		l1 = l1.Next 
	}
	for l2 != nil {
		l2Arr = append(l2Arr, l2.Val)
		l2 = l2.Next 
	}
	l1Len,l2Len := len(l1Arr),len(l2Arr)
    flag := true
	carry := 0
	for i := 0; ;i++ {
		a,b := 0,0
		if i < l1Len {
			a = l1Arr[l1Len - i - 1]
            flag = false  
		}

		if i < l2Len {
			b = l2Arr[l2Len - i - 1]
            flag = false
		}
        if a == b  && a == 0 && flag {
			break
		}
		node := &ListNode{Val:(a+b + carry)%10,Next:nil}
		carry =  (a + b + carry) / 10
		node.Next = res.Next
		res.Next = node
        flag = true
	}
    if carry == 1 {
        node :=  &ListNode{Val: carry,Next:nil}
        node.Next = res.Next
        res.Next = node
    }
	return res.Next
}