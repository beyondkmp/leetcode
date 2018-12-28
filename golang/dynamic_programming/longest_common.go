package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommon(a, b string) int {
	m := len(a)
	n := len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	i := m
	j := n
	index := dp[m][n]
	result := make([]byte, index)

	for i > 0 && j > 0 {
		switch {
		case a[i-1] == b[j-1]:
			result[index-1] = a[i-1]
			index--
			i--
			j--
		case dp[i-1][j] > dp[i][j-1]:
			i--
		default:
			j--
		}
	}

	fmt.Println(string(result))

	return dp[m][n]
}

func main() {
	a := "abcdab"
	b := "abeab"
	fmt.Println(longestCommon(a, b))
}
