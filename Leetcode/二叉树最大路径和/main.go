package main

import (
	"math"
)
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// 思路 ： 考虑实现一个简化的函数maxGain(node)，该函数计算二叉树中的的一个节点的最大贡献值，具体而言，就是以该节点的为根节点的子树中寻找以该节点
// 为起点的一条路径，，使得该路径上的节点值之和最大。
// 具体而言：	
//		* 空节点的最大贡献值为0
//		* 非空节点的最大贡献值等于节点值与其子节点的最大贡献值之和（对于叶节点而言，最大贡献值等于节点值）
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(* TreeNode) int
	maxGain  = func(node *TreeNode) int {
		if node == nil {
			return 0 
		}
		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于0的时候，才会选取对应的子节点
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		priceNewPath := node.Val + leftGain + rightGain

		// 更新答案
		maxSum = max(maxSum, priceNewPath)

		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}