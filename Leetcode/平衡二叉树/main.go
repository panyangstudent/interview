package main

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return hight(dfs(root.Left), dfs(root.Right)) <= 1 && isBalanced(root.Right) && isBalanced(root.Left)
}
func hight(a, b int) int {
    if a-b <0 {
        return b-a
    }
    return a-b
}
func dfs(root *TreeNode) int {
    if root == nil {
        return 0 
    }
    return max(dfs(root.Left), dfs(root.Right)) + 1
}
func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}