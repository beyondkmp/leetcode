package main

import "fmt"

func scoreOfParentheses(S string) int {
	var result int
	var depth uint
	for i, _ := range S {
		if S[i] == '(' {
			depth++
		} else {
			depth--
		}

		if S[i] == ')' && S[i-1] == '(' {
			result += 1 << depth
		}
	}
	return result
}

func main() {
	fmt.Println(scoreOfParentheses("(()(()))"))
}
