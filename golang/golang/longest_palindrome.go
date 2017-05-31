/* Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.
 *
 * This is case sensitive, for example "Aa" is not considered a palindrome here.
 *
 * Note:
 * Assume the length of given string will not exceed 1,010.
 *
 * Example:
 *
 * Input:
 * "abccccdd"
 *
 * Output:
 * 7
 *
 * Explanation:
 * One longest palindrome that can be built is "dccaccd", whose length is 7.
 *  */

package main

import "fmt"

func longestPalindrome(s string) int {
	count := make(map[rune]int)
	for _, w := range s {
		count[w]++
	}
	sum := 0
	flag := 0
	for _, c := range count {
		if c%2 == 0 {
			sum += c
		} else {
			flag = 1
			sum += c - 1
		}
	}
	return sum + flag
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
	fmt.Println(longestPalindrome("ccc"))
}
