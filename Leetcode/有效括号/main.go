package main 
import (
	"fmt"
)
func main() {
	fmt.Println("有效括号: ",isValid("fdfsfds"))
}

func isValid(s string) bool {
	if len(s) <= 0 {
		return true
	}
	m := map[string]string{")":"(", "]":"[", "}":"{"}
	stack := make([]string, 0)
	for i := 0; i<len(s); i++ {
		if len(stack) <= 0 {
			stack = append(stack, string(s[i]))
		} else {
			if stack[len(stack) - 1] == m[string(s[i])] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, string(s[i]))
			}
		}
	}
	if len(stack) <= 0 {
		return true
	}
	return false
}