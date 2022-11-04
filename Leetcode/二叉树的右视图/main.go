package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
    if root == nil {
        return res
    }
	tempNodes := make([]*TreeNode, 0)
	tempNodes = append(tempNodes, root)
	for len(tempNodes) > 0 {
		res  = append(res, tempNodes[len(tempNodes) -1 ].Val)
		length := len(tempNodes)
		for _, value := range tempNodes {
			if value.Left != nil {
				tempNodes = append(tempNodes, value.Left)
			}
			if value.Right != nil {
				tempNodes = append(tempNodes, value.Right)
			}
		}
		tempNodes = tempNodes[length:]
	}
	return res
}