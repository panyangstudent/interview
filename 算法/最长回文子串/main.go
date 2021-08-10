package main

func main()  {
	// 暴力解法
	longestPalindrome("发送到发生")
	// 中心扩散法

	// Manacher算法
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	maxLongestPalindrome := s[:1]
	for i:=0; i < len(s)-1; i++ {
		for j := i+1; j < len(s); j++  {
			if isPalindrome(s[i:j+1]) && j-i+1 > len(maxLongestPalindrome) {
				maxLongestPalindrome = s[i:j+1]
			}
		}
	}
	return maxLongestPalindrome
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s) - i -1] {
			return false
		}
	}
	return true
}