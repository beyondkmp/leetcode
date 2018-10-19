package main

import "fmt"

func generateParenthesis(n int) []string {
	var result []string
	var backstrack func(S string, left, right int)

	backstrack = func(S string, left, right int) {
		if len(S) == 2*n {
			result = append(result, S)
		}
		if left < n {
			backstrack(S+"(", left+1, right)
		}

		if right < left {
			backstrack(S+")", left, right+1)
		}
	}
	backstrack("", 0, 0)
	return result
}

func main() {
	fmt.Println(generateParenthesis(3))
}
