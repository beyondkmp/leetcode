package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	s := make([]int, len(nums))
	for i := range s {
		s[i] = 1
	}
	var i, j int
	for i = 1; i < len(nums); i++ {
		for j = i; j >= 0; j-- {
			if nums[i] > nums[j] {
				s[i] = max(s[j]+1, s[i])
			}
		}
	}
	fmt.Println(s)
	var result int
	for i := 0; i < len(s); i++ {
		if result < s[i] {
			result = s[i]
		}
	}
	return result
}

func main() {
	a := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(a))
}
