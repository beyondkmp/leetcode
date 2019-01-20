package main

import "fmt"

func isInterleave2(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)

	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i < m; i++ {
		if s3[i] == s1[i] {
			dp[i+1][0] = dp[i][0]
		}
	}
	for j := 0; j < n; j++ {
		if s3[j] == s2[j] {
			dp[0][j+1] = dp[0][j]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			switch {
			case s3[i+j-1] == s1[i-1] && s3[i+j-1] == s2[j-1]:
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			case s3[i+j-1] == s1[i-1]:
				dp[i][j] = dp[i-1][j]
			case s3[i+j-1] == s2[j-1]:
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	l3 := len(s3)

	fmt.Println(s1, s2, s3)
	if l1+l2 != l3 {
		return false
	}

	if l1 == 0 || l2 == 0 {
		return s1+s2 == s3
	}

	if s3[l3-1] == s1[l1-1] && s3[l3-1] == s2[l2-1] {
		return isInterleave(s1[:l1-1], s2, s3[:l3-1]) || isInterleave(s1, s2[:l2-1], s3[:l3-1])

	}
	if s3[l3-1] == s1[l1-1] {
		return isInterleave(s1[:l1-1], s2, s3[:l3-1])
	}

	if s3[l3-1] == s2[l2-1] {
		return isInterleave(s1, s2[:l2-1], s3[:l3-1])
	}
	return false

}

func main() {
	s1 := "aabc"
	s2 := "abad"
	s3 := "aabadabc"
	fmt.Println(isInterleave(s1, s2, s3))
	fmt.Println(isInterleave2(s1, s2, s3))
}
