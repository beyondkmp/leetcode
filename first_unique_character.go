package main

import "fmt"

func firstUniqChar(s string) int {
	var nums [26]int
	for _, w := range s {
		nums[w-'a'] += 1
	}
	for i, w := range s {
		if nums[w-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))

	s = "loveleetcode"
	fmt.Println(firstUniqChar(s))
}
