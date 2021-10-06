package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	res := ""
	if len(strs) <=0 {
		return res
	}
	for i:=0;i<len(strs[0]);i++ {
		flag := string(strs[0][i])
		for j:= 1;j<len(strs);j++ {
			if i >= len(strs[j]) || flag != string(strs[j][i]) {
				return  res
			}
		}
		res = res + string(flag)
	}
	return res
}
