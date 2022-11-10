package main

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    i := sort.Search(m*n, func(i int) bool { return matrix[i/n][i%n] >= target })
    return i < m*n && matrix[i/n][i%n] == target
}