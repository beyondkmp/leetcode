package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	s := make([]int, len(nums))
	s[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			s[i] = s[i-1] + 1

		} else {
			s[i] = 1
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
	a := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(a))
}
