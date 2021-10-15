package main


import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i:=0; i<rows; i++ {
		visited[i] = make([]bool, columns)
	}
	var (
		total = rows * columns
		order = make([]int, total)
		row, column = 0, 0
		directions = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
        directionIndex = 0
	)
	for i:=0;i<total;i++ {
		order[i] = matrix[row][column]
		visited[row][column] = true
		nextRow ,nextColumn := row + directions[directionIndex][0], column + directions[directionIndex][1] 
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns ||  visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
        column += directions[directionIndex][1]
	}
	return order
}