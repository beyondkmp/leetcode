package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	var result int
	var count int
	longest := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		} else if count > 0 {
			longest[i] = longest[i-1] + 2
			if i-longest[i] >= 0 {
				longest[i] += longest[i-longest[i]]
			}
			result = max(result, longest[i])
		}
	}
	return result
}

func main() {
	fmt.Println(longestValidParentheses("()"))
}
