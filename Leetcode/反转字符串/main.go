package main
// 反转字符串，使用 O(1) 的额外空间
func reverseString(s []byte)  {
	for left, right := 0, len(s) - 1; left < right ; {
		s[left] , s[right] = s[right], s[left]
		left ++ 
		right -- 
	}
}