package main 

import (
	"fmt"
)
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}
// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}

// 后序遍历

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}


// 二叉树的层次遍历
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	tempNodeList := make([]*TreeNode,0)
	if root == nil {
		return nil
	}
	tempNodeList = append(tempNodeList, root)
	for {
		length := len(tempNodeList)
		res1 := make([]int,0)
		for _, v := range tempNodeList {
			res1 = append(res1, v.Val)
			if v.Left != nil {
				tempNodeList = append(tempNodeList, v.Left)
			}
			if v.Right != nil {
				tempNodeList = append(tempNodeList, v.Right)
			}
	
		}
		res = append(res,res1)
		tempNodeList = tempNodeList[length:]
		if len(tempNodeList) <= 0 {
			break
		}
	}
	return res
}