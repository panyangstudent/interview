package main

import "fmt"

// 暴力解法思路：
// 	单个字符一定是个回文字符串，如果单个字符的两边字符也是相同的， 那这三个字符一定是回文字符串
// 	第一层循环将字符串的每个字符都作为开头的字符，
//	第二层循环，从当前字符的下一个字符开始，假设为一个回文字符串，二分法，头尾字符拿来对比，如果相同两个指针向中间游走。如果循环完成也没问题，则证明是回文字符串

// 中心扩散法思路：
// 	暴力解法虽然时间复杂度比较高，但是胜在结构清晰。因为单个字符是回文字符串，所以左右相邻两个字符如果在相等的话，那这三个字符形成了一个回文字符串
//	以此类推，在向两边扩散，如果还是相同的，则继续，同时记录。

func main()  {
	// 暴力解法
	longestPalindrome("babad")
	// 中心扩散法
	center("ssssdfdfdfd")
}

// func longestPalindrome(s string) string {
// 	if len(s) < 2 {
// 		return s
// 	}
// 	maxLongestPalindrome := s[:1]
// 	for i:=0; i < len(s)-1; i++ {
// 		for j := i+1; j < len(s); j++  {
// 			if isPalindrome(s[i:j+1]) && j-i+1 > len(maxLongestPalindrome) {
// 				maxLongestPalindrome = s[i:j+1]
// 			}
// 		}
// 	}
// 	fmt.Printf("maxLongestPalindrome %v",maxLongestPalindrome)
// 	return maxLongestPalindrome
// }

// func isPalindrome(s string) bool {
// 	for i := 0; i < len(s)/2; i++ {
// 		if s[i] != s[len(s) - i -1] {
// 			return false
// 		}
// 	}
// 	return true
// }

// 中心扩散法
func center(s string) string {
	if len(s) < 2 {
		return s
	}
	begin := 0
	maxNum := 1
    left := 0
    right :=0 
	for i :=0 ; i < len(s); i++ {
		// 字符数为奇数
		left = i - 1
		right = i+1 
		for right < len(s) && left >= 0 && s[left] == s[right] {
			if right - left + 1 > maxNum {
				begin = left
				maxNum = right - left + 1
			}
			left--
			right++
		}
		// 字符数为偶数
		left = i
		right = i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right - left + 1 > maxNum {
				begin = left
				maxNum = right - left + 1
			}
			left--
			right++ 
		}
	}
	return s[begin:begin+maxNum]
}
