package main

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	max := 0
	for i := 0;i< m;i++ {
		for j := 0 ; j < len(grid[i]);j++ {
			if grid[i][j] == 1 {
				tempMax := check(grid, i, j)
				if max < tempMax  {
					max = tempMax
				}
			}
		}
	}
	return max
}

func check(grid [][]int, i, j int) int {
	// 终止条件
	if i < 0 || j < 0 || i >= len(grid) || j >=len(grid[i]) || grid[i][j] == 0 {
		return 0
	}
	res := 1
	grid[i][j] = 0
	res += check(grid, i-1, j) 
	res += check(grid, i+1, j) 
	res += check(grid, i, j-1) 
	res += check(grid, i, j+1) 
	return res
}