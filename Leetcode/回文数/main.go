package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := 0
	tempx := x
	for {
		temp = temp * 10 + x % 10
		x = x  / 10
		if x == 0 {
			break
		}
	} 
	if temp == tempx {
		return true
	}
	return false
}

