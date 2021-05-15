package main

type ListNode struct {
	Value int64
	Next *ListNode
}

func main()  {
	
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int64
	var sum int64
	var newl1 *ListNode
	var newl2 *ListNode
	var respNode *ListNode
	for {
		temp := &ListNode{}
		if l1.Next != nil {
			
			break
		}
		if l2.Next != nil {
			break
		}
		newl1 = l1
		newl2 = l2
		sum = newl1.value + newl2.value
		if sum >= 10 {
			carry = 1
			sum = sum % 10
		} else {
			carry = 0
		}
		temp.Value = sum
		
	}
	return respNode
}

