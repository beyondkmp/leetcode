package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minOfthree(a, b, c int) int {
	return min(a, min(b, c))
}

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			switch {
			case i == 0:
				dp[i][j] = j
			case j == 0:
				dp[i][j] = i
			case word1[i-1] == word2[j-1]:
				dp[i][j] = dp[i-1][j-1]
			default:
				dp[i][j] = 1 + minOfthree(
					dp[i-1][j],
					dp[i][j-1],
					dp[i-1][j-1])

			}
		}

	}
	return dp[m][n]
}

func main() {
	a := "intention"
	b := "execution"
	fmt.Println(minDistance(a, b))

}
