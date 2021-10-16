package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool { 
	m,n := len(matrix), len(matrix[0])
	i := sort.Search(m*n func (i int)  {
		return matrix[i/n][i%m] >= target
	})

	return i < m*n && matrix[i/n][i%m] == target
}