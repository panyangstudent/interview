package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// 深度遍历
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 层次遍历
func check(root []*TreeNode) int {
	if len (root) == 0 {
		return 0
	}
	tempRoot := make([]*TreeNode, 0)
	for _, value := range root {
		if value.Left != nil {
			tempRoot = append(tempRoot,value.Left)
		}
		if value.Right != nil {
			tempRoot = append(tempRoot,value.Right)
		}
	}
	return check(tempRoot) + 1
}