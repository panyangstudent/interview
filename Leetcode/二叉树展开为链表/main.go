package main

import (
	"fmt"
)

func flatten(root *TreeNode)  {
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left , prev.Right = nil, curr
	}
}

func preorderTraversal(root *TreeNode) []*TreeNode  {
	list := []*TreeNode{}
	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}
	return list
}
