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
	var firstNotRepeat, length int
	firstNotRepeat = -1
	for i, w := range s {
		if alphaMap[w] > firstNotRepeat {
			firstNotRepeat = alphaMap[w]
		}
		alphaMap[w] = i
		length = max(length, i-firstNotRepeat)
	}
	return length
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
