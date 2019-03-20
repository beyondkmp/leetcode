package main

import "fmt"

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[i-j] * dp[j-1]
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(numTrees(3))
}
