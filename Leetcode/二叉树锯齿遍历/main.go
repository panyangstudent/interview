package main 

import (
	"fmt"
)

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	for level :=0; len(queue) > 0 ; level++ {
		vals := []int{}
		q := queue
		queue = nil
		for _, node := range q {
			vals = append(vals, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
		}
		// 层次遍历 + 反转
		if level %2 == 1 {
			for i, n := 0, len(vals); i < n/2; i++ {
                vals[i], vals[n-1-i] = vals[n-1-i], vals[i]
            }
		}
		ans = append(ans, vals)
	}
	return 
}