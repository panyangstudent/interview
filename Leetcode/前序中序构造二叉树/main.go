package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

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

func buildTreeN(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 {
		return nil
	}
	temp := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	i:=0
	for ;i<len(inorder);i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	temp.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	temp.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return temp
}