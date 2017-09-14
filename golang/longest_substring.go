package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	dic := make(map[rune]int)
	firstNotRepeat := 0
	max := 0
	for i, v := range s {
		if n, ok := dic[v]; ok && n > firstNotRepeat {
			tmp := i - n
			if tmp > max {
				max = tmp
			}
			firstNotRepeat = n + 1
		}
		dic[v] = i
	}
	if len(s)-firstNotRepeat > max {
		max = len(s) - firstNotRepeat
	}
	return max
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
