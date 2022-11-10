package main
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0 , len(nums) - 1)
}
func helper(nums []int, Left, Right int) *TreeNode {
	if Left > Right {
		return nil
	}
	mid := (Left + Right) / 2
	root := &TreeNode{Val : nums[mid]}
	root.Left = helper(nums, Left, mid - 1)
	root.Right = helper(nums, mid + 1, Right)
	return root
}
