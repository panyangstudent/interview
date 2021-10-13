package main

import (
	""
)

func buildTree(preorder []int, inorder []int) *TreeNode { 
	if len(preorder) == 0 {
		return nil 
	}
	root := &TreeNode{preorder[0], nil,nil}
	i := 0 
	for ; i<len(inorder); i++ {
		if inorder[i] == preorder[0]{
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}