package main

func detectCycle(head *ListNode) *ListNode { 
	slow , fast := head, head 
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow == fast {
			p := head 
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return  p 
		}
	}
	return nil
}