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
func detectCycle(head *ListNode) *ListNode { 
	nodeMap := make(map[*ListNode]bool)
	for head != nil {
		if _, ok = nodeMap[head]; ok {
			return head
		}
		nodeMap[head] = true
		head = head.Next
	}
	return nil
}