package main

import "fmt"

func main()  {
	maxlen := lengthOfLongestSubstring("bbbbb")
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
	)
	for index, value := range []rune(s) {
		if _,ok := existChar[value]; ok{
			fmt.Println("maxlen is %v", value)
			tempMaxLen = 1
			existChar = make(map[rune]int)
			existChar[value] = index
		} else {
			fmt.Println("maxlen is %v", value)
			existChar[value] = index
			tempMaxLen++
			if tempMaxLen >= maxLen {
				maxLen = tempMaxLen
			}
		}
	}
	return maxLen
}