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




func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var ret  [][]int
	dfs2(root, sum, []int{}, &res)
} 
func dfs2(root *TreeNode, sum int, arr []int, ret *[][]int)  {
	if root == nil {
		return 
	}
	arr = append(arr, root.Val)
	if root.Val == sum && root.Left != nil && root.Right != nil {
		temp := make([]int, len(arr))
		copy(temp, arr)
		*ret = append(*ret, temp)
	}
	dfs2(root.Left, sum - root.Val, arr, ret)
	dfs2(root.Right, sum - root.Val, arr, ret)
	arr = arr[:len(arr)-1]
}