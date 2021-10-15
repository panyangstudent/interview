package main
import (
	"fmt"
)
// 解题思路 ：回溯法 + 深度遍历	
// 1. 目标是找到矩阵中的岛屿的数量，上下左右相连的1都被认为是连续岛屿
// 2. dfs方法：设目前指针指向一个岛屿中某一点(i,j),寻找包括此点的岛屿边界
//		* 从（i，j）向此点的上下左右 (i+1,j),(i-1,j),(i,j+1),(i,j-1) 做深度搜索
//		* 终止条件：
//			* （i，j）越过矩阵边界
//			* grid[i][j] == 0，代表此分支已越过岛屿边界。
//	3. 搜索岛屿的同时执行grid[i][j] = '0'，即将岛屿所有节点删除，以免之后重复搜索相同岛屿。
func numIslands(grid [][]byte) int {
	count := 0
	// 遍历整个矩阵，当遇到 grid[i][j] == '1' 时，从此点开始做深度优先搜索 dfs，岛屿数 count + 1 且在深度优先搜索中删除此岛屿。
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}
func dfs(grid [][]byte, i, j int) {
	if (i < 0 || j < 0) || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}