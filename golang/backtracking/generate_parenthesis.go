package main

import "fmt"

func generateParenthesis(n int) []string {
	var result []string
	var dfs func(left, right int, sol []byte)

	dfs = func(left, right int, sol []byte) {
		if left == n && right == n {
			result = append(result, string(sol))
		}

		if left > n || right > n || right > left {
			return
		}

		dfs(left+1, right, append(sol, '('))
		dfs(left, right+1, append(sol, ')'))

	}
	dfs(0, 0, nil)
	return result
}

func main() {
	fmt.Println(generateParenthesis(3))
}
