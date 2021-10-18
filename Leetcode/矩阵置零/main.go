package main

func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i:=0; i< len(matrix); i++ {
		for j := 0; j< len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		} 
	}
	for i:=0; i< len(matrix); i++ {
		for j := 0; j< len(matrix[i]); j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}