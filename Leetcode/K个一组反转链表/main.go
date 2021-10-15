package main
import (
	"fmt"
)


type ListNode struct {
	Value int64
	Next *ListNode 
}

func reverseKGroup(head *ListNode, k int) *ListNode {	
	stt, end := head, head
    for i:=0; i<k; i++ {
        if end == nil {
            end = head
            break
        }
        end = end.Next
    }
    if stt == end {
        return stt
    }
    pre, cur := stt, stt
    for cur != end {
        tmp := cur.Next
        cur.Next = pre
        pre = cur
        cur = tmp
    }
    stt.Next = reverseKGroup(cur, k)
    return pre
}