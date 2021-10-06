package main 

import (
	"fmt"
)


// 经典回溯法
func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])

	for i := 0; i< n ; i++ {
		for j := 0; j < m;j++ {
			if check(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func check(board [][]byte,word string, i,j,k int) bool {
	// 找到了
	if k == len(word) {
		return true
	}
	// 如果越界了，就说明没有
	if i < 0 || j < 0 || len(board) == i || j == len(board[i]) {
		return false
	} 
	if board[i][j] != word[k] {
		return false
	}
	temp := board[i][j]
	// 重置它是为了回溯往回找的时候避免重复使用，干脆，如果找到它是对的，就直接把它置为 空
	// 等结束之后再重置回来
	board[i][j] = byte('')

	// 开始上下左右探测
	if check(board, word, i-1, j, k+1) {
		return true
	}
	if check(board, word, i+1, j, k+1) {
		return true
	}
	if check(board, word, i, j-1, k+1) {
		return true
	}
	if check(board, word, i, j+1, k+1) {
		return true
	}
	board[i][j] = temp
	return false
}