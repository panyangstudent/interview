package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func kthSmallest(root *TreeNode, k int) int {

	res := presort(root, k)
	return res[k-1]
}
func presort(root *TreeNode, k int) []int{
	if root == nil {	
		return []int{}
	}
	res := make([]int, 0)
	res = append(res, presort(root.Left, k)...)
	if len(res) == k {
		return res
	}
	res = append(res, root.Val)
	if len(res) == k {
		return res
	}
	res = append(res, presort(root.Right, k)...)
	if len(res) == k {
		return res
	}
	return res
}