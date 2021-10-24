package main 

import (
	"fmt"
)
// 解题思路：回溯法
func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil 
	} 
	var res ,path []string
	backTracking(s, path,0, &res)

}

func backTracking(s string,path []string,startIndex int,res *[]string){
	// 终止条件
	if startIndex == len(s) && len(path) == 4 {
		tempIpString := path[0] + "." + path[1] + "." + path[2] + "." + path[3] 
		*res = append(*res,tempIpString)
	}

	for i:= startIndex; i<len(s);i++ {
		path := append(path,s[startIndex:i+1])
		if i-startIndex+1 <= 3 && len(path) <= 4 && isNormalIp(s, startIndex, i) {
			// 递归
			backTracking(s,path,i+1,res)
		} else {
			return
		}
		// 回溯
		path=path[:len(path)-1]
	}
}
func isNormalIp(s string,startIndex,end int) bool {
	checkInt,_ := strconv.Atoi(s[startIndex:end+1])
	if end-startIndex + 1 > 1 && s[startIndex] == '0' {// 对于前导为0的Ip判断
		return false
	}
	if checkInt > 255 {
		return false
	} 
	return true
}