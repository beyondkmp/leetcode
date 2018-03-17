package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLongestSubstring(s string) int {
	alphaMap := make(map[rune]int)
	var start, lengest int
	for i, w := range s {
		if v, ok := alphaMap[w]; ok {
			lengest = max(lengest, i-start)
			if start <= v {
				start = v + 1
			}
		}
		alphaMap[w] = i
	}
	return max(lengest, len(s)-start)
}

func main() {
	s := "dvdf"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "abcabcab"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "ac"
	fmt.Println(lengthOfLongestSubstring(s))
	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}
