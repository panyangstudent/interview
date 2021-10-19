package main

func rotateRight(head *ListNode, k int) *ListNode {
	nodeArr := make([]*ListNode, 0)
	for head != nil {
		nodeArr = append(nodeArr, heap)
		heap = heap.Next
	}
	rotateArr(nodeArr)
	rotateArr(nodeArr[:k-1])
	rotateArr(nodeArr[k:])
	head =  resArr[0]
	for i:= 2; i< len(nodeArr) ; i++ {
		nodeArr[i-1].Next = nodeArr[i]
		nodeArr[i].Next = nil 
	}
	return head
}
func rotateArr(nodeArr []*ListNode){
	for i:=0; i< len(nodeArr) / 2; i++ {
		nodeArr[i], nodeArr[len(nodeArr)-1-i] = nodeArr[len(nodeArr)-1-i], nodeArr[i]
	}
}