package main

import ( 
	"fmt"
)

// 回溯死抓三个要点
// 1. 选择
//		* 在这里每次最多两个选择，选左括号还是右括号，“选择”会展开一颗解的空间树
//		* 使用DFS遍历这棵树
// 2. 约束条件
//		* 即，什么情况下可以选择做括号，什么情况下可以选择右括号
//		* 利用约束做“剪枝”，即，去掉不会产生解的选项，即不会通往合法解的分支
// 3. 目标	
//		* 构建出一个用尽n对括号的合法括号串
//		* 意味着，当构建的长度达到2*n，就可以结束递归

func generateParenthesis(n int) []string {
    res := []string{}

	var dfs func(lRemain int, rRemain int, path string)
	dfs = func(lRemain int, rRemain int, path string) {
		if 2*n == len(path) {
			res = append(res, path)
			return
		}
		if lRemain > 0 {
			dfs(lRemain-1, rRemain, path+"(")
		}
		if lRemain < rRemain {
			dfs(lRemain, rRemain-1, path+")")
		}
	}
	
	dfs(n, n, "")
	return res
}