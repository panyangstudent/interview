package main

import (
	"fmt"
)

func findCircleNum(isConnected [][]int) (ans int) {
    vis := make([]bool, len(isConnected))
    var dfs func(int)
    dfs = func(from int) {
        vis[from] = true
        for to, conn := range isConnected[from] {
            if conn == 1 && !vis[to] {
                dfs(to)
            }
        }
    }
    for i, v := range vis {
        if !v {
            ans++
            dfs(i)
        }
    }
    return
}