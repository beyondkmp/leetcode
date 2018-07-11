package main

import "fmt"

func longestPalindrome(s string) string {
	var result string
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = s[i] == s[j] && (j-i < 3 || dp[i+1][j-1])

			if dp[i][j] && len(result) <= j-i+1 {
				result = s[i : j+1]
			}
		}
	}
	return result
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome(""))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("aaaa"))
}
