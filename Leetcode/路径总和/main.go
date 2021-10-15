package main

import (
	"fmt"
)

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}

	if hasPathSum(root.Left, targetSum - root.Val) {
		return true
	}

	if hasPathSum(root.Right, targetSum - root.Val) {
		return true
	}
	return false
}