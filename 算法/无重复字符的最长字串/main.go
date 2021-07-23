package main

/*
解题思路：
1. 题目要求是无重复的字符的连续子串
 */
import "fmt"

func main()  {
	maxlen := lengthOfLongestSubstring("pwwkew")
	fmt.Println("maxlen is %v", maxlen)
}
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	var (
		existChar = make(map[rune]int)
		maxLen = 0
		tempMaxLen = 0
		tempIndex = 0
	)
	for index, value := range []rune(s) {
		if i ,ok := existChar[value]; ok{
			if i <= tempIndex {

			}
			fmt.Println("maxlen is %v index : %v", value,index)

			for key, value := range existChar {
				if value <= i {
					delete(existChar, key)
					tempMaxLen--
				}
			}
			existChar[value] = index
			tempMaxLen++
		} else {
			fmt.Println("maxlen 111 is %v index : %v", value,index)
			existChar[value] = index
			tempMaxLen++
			if tempMaxLen >= maxLen {
				maxLen = tempMaxLen
			}
		}
	}
	return maxLen
}