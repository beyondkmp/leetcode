package main

func numDistinct(s string, t string) int {
	n, m := len(s), len(t)
	if n == 0 || m == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	s1 := []rune(s)
	t1 := []rune(t)
	if s1[0] == t1[0] {
		dp[0][0] = 1
	}
	for i := 1; i < n; i++ {
		if s1[i] == t1[0] {
			dp[0][i] = dp[0][i-1] + 1
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := i; j < n; j++ {
			if s1[j] == t1[i] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
