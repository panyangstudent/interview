package main

/*
解题思路：
1. 题目要求是无重复的字符的连续子串,所以肯定是构造一个map，来查看这个map中是否存在，map的value存该字符的下标
2. 然后foreach， 如果不存在之前的map中，就把当前的字符加入， 并且对应的最长字串长度加1
3. 如果已经存在之前的map中了， 则删除之前已存在map中的小于当前下标的字符。这里的实现比较笨，有待后续优化，理论上只需要搞一个游走下标就可以了
 */
import (
	"fmt"
)

func main()  {
	maxlen := lengthOfLongestSubstringNew("aabaab!bb")
	fmt.Println("maxlen is %v", maxlen)
}

func lengthOfLongestSubstringNew(s string) int {
	charIndex := make(map[rune]int)
	start, maxlength ,length := 0,0,0
	for i, v :=range []rune(s) {
		if lastIndex ,ok := charIndex[v]; ok && start <= lastIndex {
			start = lastIndex + 1
			length = i - start + 1
		} else {
			length += 1
		}

		charIndex[v] = i
		if length > maxlength {
			maxlength = length 
		}
	}
	return maxlength
}